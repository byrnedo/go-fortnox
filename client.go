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
	mimeJSON  = "application/json"
	userAgent = "go-fortnox/client.go (godoc.org/github.com/byrnedo/go-fortnox)"

	// TimeFormat is the format that fortnox expects
	TimeFormat = "2006-01-02 15:04"
	// DateFormat that fortnox expects
	DateFormat = "2006-01-02"

	// DefaultURL is the default api url
	DefaultURL = "https://api.fortnox.se/3/"
)

var (
	defaultTimeout = time.Duration(20 * time.Second)
	defaultHeaders = map[string]string{
		"Accept":       mimeJSON,
		"Content-Type": mimeJSON,
		"User-Agent":   userAgent,
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

	if err := request(ctx, opts.HTTPClient, headers, "GET", opts.BaseURL, nil, result); err != nil {
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
	BaseURL      string
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

// NewClient creates a new client
func NewClient(optionsFuncs ...OptionsFunc) *Client {

	c := &http.Client{Timeout: defaultTimeout}

	o := &ClientOptions{
		BaseURL:    DefaultURL,
		HTTPClient: c,
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

func (c *Client) deleteResource(ctx context.Context, resource string) error {

	err := c.request(ctx, "DELETE", resource, nil, nil, nil)
	if err != nil {
		return err
	}
	return nil
}

// MetaInformation for responses
type MetaInformation struct {
	CurrentPage    int `json:"@CurrentPage"`
	TotalPages     int `json:"@TotalPages"`
	TotalResources int `json:"@TotalResources"`
}

// ErrorMessage response type
type ErrorMessage struct {
	Error   int    `json:"Error"`
	Message string `json:"Message"`
	Code    int    `json:"Code"`
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

	return request(ctx, c.clientOptions.HTTPClient, headers, method, u.String(), bodyBuffer, result)

}

// ErrorResp error response from fnox
type ErrorResp struct {
	ErrorInformation ErrorMessage
}

// FnoxError Our internal remote error holder
type FnoxError struct {
	HttpStatus int
	Code       int
	Message    string
}

// Error pretty print error
func (f FnoxError) Error() string {
	return fmt.Sprintf("%d - %s", f.Code, f.Message)
}

func request(ctx context.Context, client *http.Client, headers map[string]string, method, url string, data io.Reader, result interface{}) error {

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
	case 200, 201:

		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			return errors.Wrap(err, "failed to decode json")
		}

		return nil

	default:
		errMsg := &ErrorResp{}
		if err := json.NewDecoder(resp.Body).Decode(&errMsg); err != nil {
			return errors.Wrap(err, "failed to decode json")
		}
		return FnoxError{HttpStatus: resp.StatusCode, Code: errMsg.ErrorInformation.Code, Message: errMsg.ErrorInformation.Message}
	}

}
