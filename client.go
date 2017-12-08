package fortnox

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"net/url"
	"time"
)

const (
	mimeJson = "application/json"

	TIME_FORMAT = "2006-01-02 15:04"
	DATE_FORMAT = "2006-01-02"
	API_URL     = "https://api.fortnox.se/3/"
)

var (
	// Error return in the case of 401
	ErrUnauthorized = errors.New("Unauthorized")

	defaultTimeout  = time.Duration(20 * time.Second)
	defaultHeaders  = map[string]string{
		"Accept":       mimeJson,
		"Content-Type": mimeJson,
	}
)

type AccessTokenOptions struct {
	BaseUrl    string
	HttpClient *http.Client
}

func GetAccessToken(ctx context.Context, authorizationCode string, clientSecret string, optsFuncs ...func(*AccessTokenOptions)) (string, error) {

	opts := &AccessTokenOptions{
		BaseUrl:    API_URL,
		HttpClient: &http.Client{Timeout: defaultTimeout},
	}
	for _, f := range optsFuncs {
		f(opts)
	}

	headers := map[string]string{
		"Authorization-Code": authorizationCode,
		"Client-Secret":      clientSecret,
	}

	result := &struct {
		Authorization struct {
			AccessToken string `json:"AccessToken"`
		} `json:"Authorization"`
	}{}

	if err := request(opts.HttpClient, ctx, headers, "GET", opts.BaseUrl, nil, nil, result); err != nil {
		return "", err
	}

	return result.Authorization.AccessToken, nil
}

type ClientOptions struct {
	// Users's access token (obtained by user when they add our integration)
	AccessToken string
	// Client's integration secret
	ClientSecret string
	Accepts      string
	ContentType  string
	BaseUrl      string
	SkipVerify   bool
	HttpClient   *http.Client
}

// Query param info from docs
//:
// SEARCH NAME	DESCRIPTION	EXAMPLE VALUE
// lastmodified	Retrieves all records since the provided timestamp.	2014-03-10 12:30
// financialyear	Selects what financial year that should be used	5
// financialyeardate	Selects by date, what financial year that should be used	2014-03-10
// fromdate	Defines a selection based on a start date.
// Only available for invoices, orders, offers and vouchers.	2014-03-10
// todate	Defines a selection based on an end date.
// Only available for invoices, orders, offers and vouchers	 2014-03-10
type QueryParams struct {
	LastModified      time.Time
	FinancialYear     int
	FinancialYearDate string
	FromDate          string
	ToDate            string
	Page              int
	Limit             int
	Offset            int
	Extra             map[string][]string
}

func (this *QueryParams) toValues() url.Values {

	ret := make(url.Values)
	if !this.LastModified.IsZero() {
		ret["lastmodified"] = []string{this.LastModified.Format(TIME_FORMAT)}
	}
	if this.FinancialYear > 0 {
		ret["financialyear"] = []string{fmt.Sprintf("%d", this.FinancialYear)}
	}
	if len(this.FinancialYearDate) > 0 {
		ret["financialyeardate"] = []string{this.FinancialYearDate}
	}
	if len(this.FromDate) > 0 {
		ret["fromdate"] = []string{this.FromDate}
	}
	if len(this.ToDate) > 0 {
		ret["todate"] = []string{this.ToDate}
	}
	if this.Limit > 0 {
		ret["limit"] = []string{fmt.Sprintf("%d", this.Limit)}
	}
	if this.Offset > 0 {
		ret["offset"] = []string{fmt.Sprintf("%d", this.Offset)}
	}
	if this.Page > 0 {
		ret["page"] = []string{fmt.Sprintf("%d", this.Page)}
	}
	for k, vs := range this.Extra {
		ret[k] = vs
	}
	return ret
}

type FilterParamFunc func(*QueryParams)

// Client for taklking to fnox with
//
type Client struct {
	clientOptions *ClientOptions
}

type OptionsFunc func(o *ClientOptions)

func WithAuthOpts(token, secret string) OptionsFunc {
	return func(o *ClientOptions) {
		o.AccessToken = token
		o.ClientSecret = secret
	}
}

func WithURLOpts(url string) OptionsFunc {
	return func(o *ClientOptions) {
		o.BaseUrl = url
	}
}

func NewFortnoxClient(optionsFuncs ...OptionsFunc) *Client {

	c := &http.Client{Timeout: defaultTimeout}

	o := &ClientOptions{
		Accepts:     mimeJson,
		ContentType: mimeJson,
		BaseUrl:     API_URL,
		HttpClient:  c,
	}
	for _, f := range optionsFuncs {
		f(o)
	}

	return &Client{
		clientOptions: o,
	}
}

func (c *Client) makeUrl(section string) (*url.URL, error) {
	u, err := url.Parse(c.clientOptions.BaseUrl)
	if err != nil {
		return nil, err
	}
	u2, err := url.Parse(section)
	if err != nil {
		return nil, err
	}
	return u.ResolveReference(u2), nil
}

type ListOrdersResp struct {
	Orders          []*OrderShort    `json:"Orders"`
	MetaInformation *MetaInformation `json:"MetaInformation"`
}

func (c *Client) ListOrders(ctx context.Context, p *QueryParams) (*ListOrdersResp, error) {

	resp := &ListOrdersResp{}

	err := c.request(ctx, "GET", "orders", nil, p, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) GetOrder(ctx context.Context, id string) (*OrderFull, error) {

	resp := &struct {
		Order *OrderFull `json:"Order"`
	}{}
	err := c.request(ctx, "GET", "orders/"+id, nil, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp.Order, nil
}

func (c *Client) CreateOrder(ctx context.Context, order *CreateOrder) (*OrderFull, error) {
	orderResp := &struct {
		Order *OrderFull `json:"Order"`
	}{}
	err := c.request(ctx, "POST", "orders/", &struct {
		Order *CreateOrder `json:"Order"`
	}{
		Order: order,
	}, nil, orderResp)
	if err != nil {
		return nil, err
	}

	return orderResp.Order, nil
}

func (c *Client) UpdateOrder(ctx context.Context, id string, fields map[string]interface{}) (*OrderFull, error) {

	resp := &struct {
		Order *OrderFull `json:"Order"`
	}{}
	err := c.request(ctx, "PUT", "orders/"+id, &struct {
		Order map[string]interface{} `json:"Order"`
	}{
		Order: fields,
	}, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp.Order, nil
}

/*
 "MetaInformation": {
        "@CurrentPage": 1,
        "@TotalPages": 1,
        "@TotalResources": 32
    }
*/

type ListInvoicesResp struct {
	Invoices        []*InvoiceShort  `json:"Invoices"`
	MetaInformation *MetaInformation `json:"MetaInformation"`
}

func (c *Client) ListInvoices(ctx context.Context, p *QueryParams) (*ListInvoicesResp, error) {
	resp := &ListInvoicesResp{}

	err := c.request(ctx, "GET", "invoices", nil, p, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) GetInvoice(ctx context.Context, id string) (*InvoiceFull, error) {

	resp := &struct {
		Invoice *InvoiceFull `json:"Invoice"`
	}{}
	err := c.request(ctx, "GET", "invoices/"+id, nil, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp.Invoice, nil
}

func (c *Client) GetCompanySettings(ctx context.Context) (*CompanySettings, error) {

	resp := &struct {
		CompanySettings *CompanySettings `json:"CompanySettings"`
	}{}
	err := c.request(ctx, "GET", "settings/company", nil, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp.CompanySettings, nil
}

type ListArticlesResp struct {
	Articles        []*Article       `json:"Articles"`
	MetaInformation *MetaInformation `json:"MetaInformation"`
}

func (c *Client) ListArticles(ctx context.Context, p *QueryParams) (*ListArticlesResp, error) {
	resp := &ListArticlesResp{}

	err := c.request(ctx, "GET", "articles", nil, p, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) GetArticle(ctx context.Context, id string) (*Article, error) {

	resp := &struct {
		Article *Article `json:"Article"`
	}{}
	err := c.request(ctx, "GET", "articles/"+id, nil, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp.Article, nil
}

func (c *Client) ListLabels(ctx context.Context) ([]*Label, error) {
	resp := &struct {
		Labels []*Label `json:"Labels"`
	}{}

	err := c.request(ctx, "GET", "labels", nil, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp.Labels, nil
}

type CreateLabelReq struct {
	Label struct {
		Description string `json:"Description"`
	} `json:"Label"`
}

func (c *Client) CreateLabel(ctx context.Context, name string) (*Label, error) {

	resp := &struct {
		Label *Label `json:"Label"`
	}{}
	req := CreateLabelReq{}
	req.Label.Description = name
	err := c.request(ctx, "POST", "labels", &req, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp.Label, nil
}

func (c *Client) request(ctx context.Context, method, resource string, body interface{}, p *QueryParams, result interface{}) error {
	u, err := c.makeUrl(resource)
	if err != nil {
		return err
	}

	if p != nil {
		u.RawQuery = p.toValues().Encode()
	}

	headers := map[string]string{
		"Access-Token":  c.clientOptions.AccessToken,
		"Client-Secret": c.clientOptions.ClientSecret,
	}

	bodyBuffer := new(bytes.Buffer)
	json.NewEncoder(bodyBuffer).Encode(body)

	return request(c.clientOptions.HttpClient, ctx, headers, method, u.String(), bodyBuffer, p, result)

}

func request(client *http.Client, ctx context.Context, headers map[string]string, method, url string, data io.Reader, p *QueryParams, result interface{}) error {

	req, err := http.NewRequest(method, url, data)
	if err != nil {
		return errors.Wrap(err, "error creating request")
	}

	req.WithContext(ctx)

	for k, v := range defaultHeaders {
		req.Header.Add(k, v)
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return errors.Wrap(err, "error sending request")
	}

	defer resp.Body.Close()

	switch resp.StatusCode {
	case 401:
		return ErrUnauthorized
	case 200, 201:

		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			return errors.Wrap(err, "failed to decode json")
		}

		return nil

	default:

		errMsg := &struct {
			ErrorInformation *ErrorMessage
		}{
			&ErrorMessage{},
		}
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			return errors.Wrap(err, "failed to decode json")
		}
		return errors.New(fmt.Sprintf("%d: %s", errMsg.ErrorInformation.Code, errMsg.ErrorInformation.Message))
	}

}
