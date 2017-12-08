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
	mimeJSON = "application/json"
	// TimeFormat is the format that fortnox expects
	TimeFormat = "2006-01-02 15:04"
	// DateFormat that fortnox expects
	DateFormat = "2006-01-02"

	// DefaultURL is the default api url
	DefaultURL = "https://api.fortnox.se/3/"
)

var (
	// ErrUnauthorized is the error returned in the case of 401
	ErrUnauthorized = errors.New("Unauthorized")

	defaultTimeout = time.Duration(20 * time.Second)
	defaultHeaders = map[string]string{
		"Accept":       mimeJSON,
		"Content-Type": mimeJSON,
	}
)

//AccessTokenOptions are options used when creating access tokens
type AccessTokenOptions struct {
	BaseURL    string
	HTTPClient *http.Client
}

// GetAccessToken from an auth code for a client. Careful, do this only once per auth code
func GetAccessToken(ctx context.Context, authorizationCode string, clientSecret string, optsFuncs ...func(*AccessTokenOptions)) (string, error) {

	opts := &AccessTokenOptions{
		BaseURL:    DefaultURL,
		HTTPClient: &http.Client{Timeout: defaultTimeout},
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

	if err := request(ctx, opts.HTTPClient, headers, "GET", opts.BaseURL, nil, nil, result); err != nil {
		return "", err
	}

	return result.Authorization.AccessToken, nil
}

// ClientOptions for creating the main client
type ClientOptions struct {
	// Users's access token (obtained by user when they add our integration)
	AccessToken string
	// Client's integration secret
	ClientSecret string
	Accepts      string
	ContentType  string
	BaseURL      string
	SkipVerify   bool
	HTTPClient   *http.Client
}

// QueryParams when searching orders/invoices
// From fortnox docs:
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

func (p *QueryParams) toValues() url.Values {

	ret := make(url.Values)
	if !p.LastModified.IsZero() {
		ret["lastmodified"] = []string{p.LastModified.Format(TimeFormat)}
	}
	if p.FinancialYear > 0 {
		ret["financialyear"] = []string{fmt.Sprintf("%d", p.FinancialYear)}
	}
	if len(p.FinancialYearDate) > 0 {
		ret["financialyeardate"] = []string{p.FinancialYearDate}
	}
	if len(p.FromDate) > 0 {
		ret["fromdate"] = []string{p.FromDate}
	}
	if len(p.ToDate) > 0 {
		ret["todate"] = []string{p.ToDate}
	}
	if p.Limit > 0 {
		ret["limit"] = []string{fmt.Sprintf("%d", p.Limit)}
	}
	if p.Offset > 0 {
		ret["offset"] = []string{fmt.Sprintf("%d", p.Offset)}
	}
	if p.Page > 0 {
		ret["page"] = []string{fmt.Sprintf("%d", p.Page)}
	}
	for k, vs := range p.Extra {
		ret[k] = vs
	}
	return ret
}

// Client for fortnox api calls
type Client struct {
	clientOptions *ClientOptions
}

// OptionsFunc sig for customising options
type OptionsFunc func(o *ClientOptions)

// WithAuthOpts helper for adding auth
func WithAuthOpts(token, secret string) OptionsFunc {
	return func(o *ClientOptions) {
		o.AccessToken = token
		o.ClientSecret = secret
	}
}

// WithURLOpts helper for changing base url
func WithURLOpts(url string) OptionsFunc {
	return func(o *ClientOptions) {
		o.BaseURL = url
	}
}

// NewFortnoxClient creates a new client
func NewFortnoxClient(optionsFuncs ...OptionsFunc) *Client {

	c := &http.Client{Timeout: defaultTimeout}

	o := &ClientOptions{
		Accepts:     mimeJSON,
		ContentType: mimeJSON,
		BaseURL:     DefaultURL,
		HTTPClient:  c,
	}
	for _, f := range optionsFuncs {
		f(o)
	}

	return &Client{
		clientOptions: o,
	}
}

func (c *Client) makeURL(section string) (*url.URL, error) {
	u, err := url.Parse(c.clientOptions.BaseURL)
	if err != nil {
		return nil, err
	}
	u2, err := url.Parse(section)
	if err != nil {
		return nil, err
	}
	return u.ResolveReference(u2), nil
}

// ListOrdersResp Response when listing orders
type ListOrdersResp struct {
	Orders          []*OrderShort    `json:"Orders"`
	MetaInformation *MetaInformation `json:"MetaInformation"`
}

// ListOrders or search orders
func (c *Client) ListOrders(ctx context.Context, p *QueryParams) (*ListOrdersResp, error) {

	resp := &ListOrdersResp{}

	err := c.request(ctx, "GET", "orders", nil, p, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetOrder gets one order by id
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

// CreateOrder creates an order
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

// UpdateOrder updates an order
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

// ListInvoicesResp is the response for listing invoices
type ListInvoicesResp struct {
	Invoices        []*InvoiceShort  `json:"Invoices"`
	MetaInformation *MetaInformation `json:"MetaInformation"`
}

// ListInvoices lists invoices
func (c *Client) ListInvoices(ctx context.Context, p *QueryParams) (*ListInvoicesResp, error) {
	resp := &ListInvoicesResp{}

	err := c.request(ctx, "GET", "invoices", nil, p, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetInvoice gets one invoice
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

// GetCompanySettings fetches company info
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

// ListArticlesResp is the response for ListArticles
type ListArticlesResp struct {
	Articles        []*Article       `json:"Articles"`
	MetaInformation *MetaInformation `json:"MetaInformation"`
}

// ListArticles lists or searches articles
func (c *Client) ListArticles(ctx context.Context, p *QueryParams) (*ListArticlesResp, error) {
	resp := &ListArticlesResp{}

	err := c.request(ctx, "GET", "articles", nil, p, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetArticle gets one article
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

// ListLabels lists labels
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

// CreateLabelReq is the request used for creating labels
type CreateLabelReq struct {
	Label struct {
		Description string `json:"Description"`
	} `json:"Label"`
}

// CreateLabel creates a label
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
	u, err := c.makeURL(resource)
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

	return request(ctx, c.clientOptions.HTTPClient, headers, method, u.String(), bodyBuffer, p, result)

}

func request(ctx context.Context, client *http.Client, headers map[string]string, method, url string, data io.Reader, p *QueryParams, result interface{}) error {

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
		return fmt.Errorf("%d: %s", errMsg.ErrorInformation.Code, errMsg.ErrorInformation.Message)
	}

}
