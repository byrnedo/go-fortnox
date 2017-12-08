# Go Fortnox Client

A client for [Fortnox's](https://www.fortnox.se) REST api.

```go
client := fortnox.NewFortnoxClient(fortnox.WithAuthOpts("token", "secret"))

order, err := client.GetOrder(docNumber)
if err != nil {
    return errors.Wrap(err, "Error getting order")
}
```

## Methods:

### Package Function

```go
// Get an access token from an authorization code 
GetAccessToken(ctx context.Context, authorizationCode string, clientSecret string, optsFuncs ...func(*AccessTokenOptions)) (string, error)

// Create new client
NewFortnoxClient(optionsFuncs ...OptionsFunc) *Client
```

### Client methods

```go
GetOrder(ctx context.Context, id string) (*OrderFull, error) {
CreateOrder(ctx context.Context, order *CreateOrder) (*OrderFull, error) {
UpdateOrder(ctx context.Context, id string, fields map[string]interface{}) (*OrderFull, error) {

ListInvoices(ctx context.Context, p *QueryParams) (*ListInvoicesResp, error) {
GetInvoice(ctx context.Context, id string) (*InvoiceFull, error) {

GetCompanySettings(ctx context.Context) (*CompanySettings, error) {

ListArticles(ctx context.Context, p *QueryParams) (*ListArticlesResp, error) {
GetArticle(ctx context.Context, id string) (*Article, error) {

ListLabels(ctx context.Context) ([]*Label, error) {
CreateLabel(ctx context.Context, name string) (*Label, error) {

```

There are quite a few endpoints that aren't implemented yet. Feel free to make an issue or pull request.

## 'ish Types (Floatish, Intish)

For some reason the fortnox api ocassionally gives back a float but sometimes a string for certain fields. 
I made these two types for dealing with those situations for unmarshalling.

## Running Tests

You need to set the `FORTNOX_AUTH_CODE`, `FORTNOX_ACCESS_TOKEN` and `FORTNOX_CLIENT_SECRET` envs when running the tests.
