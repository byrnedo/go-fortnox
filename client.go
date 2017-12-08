package fortnox

import (
	"crypto/tls"
	"fmt"
	"github.com/byrnedo/apibase/utils"
	"github.com/pkg/errors"
	"net/http"
	"net/url"
	"time"
)

var (
	ErrUnauthorized = errors.New("Unauthorized")
)

const (
	mime_json = "application/json"

	TIME_FORMAT = "2006-01-02 15:04"
	DATE_FORMAT = "2006-01-02"
)

type HttpClientFunc func(c *http.Client)

func GetAccessToken(baseUrl string, authorizationCode string, clientSecret string, cFuncs ...HttpClientFunc) (string, error) {

	restClient := utils.NewRestClient(func(c *http.Client) {
		c.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
		}

		for _, f := range cFuncs {
			f(c)
		}
	})

	restClient.Headers = map[string]string{
		"Authorization-Code": authorizationCode,
		"Client-Secret":      clientSecret,
		"Accept":             mime_json,
		"Content-Type":       mime_json,
	}

	r, err := restClient.Get(baseUrl)
	if err != nil {
		return "", err
	}

	result := &struct {
		Authorization struct {
			AccessToken string `json:"AccessToken"`
		} `json:"Authorization"`
	}{}
	switch r.Response.StatusCode {
	case 401:
		return "", ErrUnauthorized
	case 200, 201:
		err := r.AsJson(result)
		if err != nil {
			return "", errors.Wrap(err, "Failed to decode json from ["+string(r.GetBody())+"]")
		}

		return result.Authorization.AccessToken, nil

	default:

		errMsg := &struct {
			ErrorInformation *ErrorMessage
		}{
			&ErrorMessage{},
		}
		if e := r.AsJson(errMsg); e != nil {
			return "", errors.Wrap(e, "Failed to decode json from ["+string(r.GetBody())+"]")
		}
		return "", errors.New(fmt.Sprintf("%d: %s", errMsg.ErrorInformation.Code, errMsg.ErrorInformation.Message))
	}

}

type FortnoxClientOptions struct {
	// Users's access token (obtained by user when they add our integration)
	AccessToken string
	// Client's integration secret
	ClientSecret string
	Accepts      string
	ContentType  string
	BaseUrl      string
	SkipVerify   bool
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
type FortnoxClient struct {
	restClient *utils.RestClient
	*FortnoxClientOptions
}

type OptionsFunc func(o *FortnoxClientOptions)

func WithAuthOpts(token, secret string) OptionsFunc {
	return func(o *FortnoxClientOptions) {
		o.AccessToken = token
		o.ClientSecret = secret
	}
}

func WithURLOpts(url string) OptionsFunc {
	return func(o *FortnoxClientOptions) {
		o.BaseUrl = url
	}
}

func NewFortnoxClient(optionsFuncs ...OptionsFunc) *FortnoxClient {

	o := &FortnoxClientOptions{
		Accepts:     mime_json,
		ContentType: mime_json,
	}
	for _, f := range optionsFuncs {
		f(o)
	}

	return &FortnoxClient{
		restClient: utils.NewRestClient(func(c *http.Client) {
			c.Transport = &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: o.SkipVerify},
			}
		}),
		FortnoxClientOptions: o,
	}
}

func (this *FortnoxClient) makeUrl(section string) (*url.URL, error) {
	u, err := url.Parse(this.BaseUrl)
	if err != nil {
		return nil, err
	}
	u2, err := url.Parse(section)
	if err != nil {
		return nil, err
	}
	return u.ResolveReference(u2), nil
}

func (this *FortnoxClient) GetOrders(p *QueryParams) ([]*OrderShort, *MetaInformation, error) {

	resp := &struct {
		Orders          []*OrderShort    `json:"Orders"`
		MetaInformation *MetaInformation `json:"MetaInformation"`
	}{}

	err := this.request("GET", "orders", nil, p, resp)
	if err != nil {
		return nil, nil, err
	}
	return resp.Orders, resp.MetaInformation, nil
}

func (this *FortnoxClient) GetOrder(id string) (*OrderFull, error) {

	resp := &struct {
		Order *OrderFull `json:"Order"`
	}{}
	err := this.request("GET", "orders/"+id, nil, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp.Order, nil
}

func (this *FortnoxClient) CreateOrder(order *CreateOrder) (*OrderFull, error) {
	orderResp := &struct {
		Order *OrderFull `json:"Order"`
	}{}
	err := this.request("POST", "orders/", &struct {
		Order *CreateOrder `json:"Order"`
	}{
		Order: order,
	}, nil, orderResp)
	if err != nil {
		return nil, err
	}

	return orderResp.Order, nil
}

func (this *FortnoxClient) PutOrder(id string, fields map[string]interface{}) (*OrderFull, error) {

	resp := &struct {
		Order *OrderFull `json:"Order"`
	}{}
	err := this.request("PUT", "orders/"+id, &struct {
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
func (this *FortnoxClient) GetInvoices(p *QueryParams) ([]*InvoiceShort, *MetaInformation, error) {

	resp := &struct {
		Invoices        []*InvoiceShort  `json:"Invoices"`
		MetaInformation *MetaInformation `json:"MetaInformation"`
	}{}

	err := this.request("GET", "invoices", nil, p, resp)
	if err != nil {
		return nil, nil, err
	}
	return resp.Invoices, resp.MetaInformation, nil
}

func (this *FortnoxClient) GetInvoice(id string) (*InvoiceFull, error) {

	resp := &struct {
		Invoice *InvoiceFull `json:"Invoice"`
	}{}
	err := this.request("GET", "invoices/"+id, nil, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp.Invoice, nil
}

func (this *FortnoxClient) GetCompanySettings() (*CompanySettings, error) {

	resp := &struct {
		CompanySettings *CompanySettings `json:"CompanySettings"`
	}{}
	err := this.request("GET", "settings/company", nil, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp.CompanySettings, nil
}

func (this *FortnoxClient) GetArticles(p *QueryParams) ([]*Article, *MetaInformation, error) {

	resp := &struct {
		Articles        []*Article       `json:"Articles"`
		MetaInformation *MetaInformation `json:"MetaInformation"`
	}{}

	err := this.request("GET", "articles", nil, p, resp)
	if err != nil {
		return nil, nil, err
	}
	return resp.Articles, resp.MetaInformation, nil
}

func (this *FortnoxClient) GetArticle(id string) (*Article, error) {

	resp := &struct {
		Article *Article `json:"Article"`
	}{}
	err := this.request("GET", "articles/"+id, nil, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp.Article, nil
}

func (this *FortnoxClient) GetLabels() ([]*Label, error) {
	resp := &struct {
		Labels []*Label `json:"Labels"`
	}{}

	err := this.request("GET", "labels", nil, nil, resp)
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

func (this *FortnoxClient) CreateLabel(name string) (*Label, error) {

	resp := &struct {
		Label *Label `json:"Label"`
	}{}
	req := CreateLabelReq{}
	req.Label.Description = name
	err := this.request("POST", "labels", &req, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp.Label, nil
}

func (this *FortnoxClient) request(method, resource string, data interface{}, p *QueryParams, result interface{}) error {
	u, err := this.makeUrl(resource)
	if err != nil {
		return err
	}

	if p != nil {
		u.RawQuery = p.toValues().Encode()
	}

	this.restClient.Headers = map[string]string{
		"Access-Token":  this.AccessToken,
		"Client-Secret": this.ClientSecret,
		"Accept":        this.Accepts,
		"Content-Type":  this.ContentType,
	}

	r, err := this.restClient.DoJson(method, u.String(), data)
	if err != nil {
		return err
	}

	switch r.Response.StatusCode {
	case 401:
		return ErrUnauthorized
	case 200, 201:
		if result != nil {
			if e := r.AsJson(result); e != nil {
				return errors.Wrap(e, "Failed to decode json from ["+string(r.GetBody())+"]")
			}
		}
		return nil
	default:

		errMsg := &struct {
			ErrorInformation *ErrorMessage
		}{
			&ErrorMessage{},
		}
		if e := r.AsJson(errMsg); e != nil {
			return errors.Wrap(e, "Failed to decode json from ["+string(r.GetBody())+"]")
		}
		return errors.New(fmt.Sprintf("%d: %s", errMsg.ErrorInformation.Code, errMsg.ErrorInformation.Message))

	}

}
