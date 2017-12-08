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
GetAccessToken(authorizationCode string, clientSecret string, optsFuncs ...func(*AccessTokenOptions)) (string, error)
```

### Client methods

```go
ListOrders(p *QueryParams) (*ListOrdersResp, error) {
GetOrder(id string) (*OrderFull, error) {
CreateOrder(order *CreateOrder) (*OrderFull, error) {
UpdateOrder(id string, fields map[string]interface{}) (*OrderFull, error) {
ListInvoices(p *QueryParams) (*ListInvoicesResp, error) {
GetInvoice(id string) (*InvoiceFull, error) {
GetCompanySettings() (*CompanySettings, error) {
ListArticles(p *QueryParams) (*ListArticlesResp, error) {
GetArticle(id string) (*Article, error) {
ListLabels() ([]*Label, error) {
CreateLabel(name string) (*Label, error) {
```

There are quite a few endpoints that aren't implemented yet. Feel free to make an issue or pull request.

## 'ish Types (Floatish, Intish)

For some reason the fortnox api ocassionally gives back a float but sometimes a string for certain fields. 
I made these two types for dealing with those situations for unmarshalling.

## Running Tests

You need to set the `FORTNOX_AUTH_CODE`, `FORTNOX_ACCESS_TOKEN` and `FORTNOX_CLIENT_SECRET` envs when running the tests.
