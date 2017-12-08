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
GetAccessToken(baseUrl string, authorizationCode string, clientSecret string, cFuncs ...HttpClientFunc) (string, error)
```

### Client methods

```go
GetOrders(p *QueryParams) ([]*OrderShort, *MetaInformation, error) {
GetOrder(id string) (*OrderFull, error) {
CreateOrder(order *CreateOrder) (*OrderFull, error) {
PutOrder(id string, fields map[string]interface{}) (*OrderFull, error) {
GetInvoices(p *QueryParams) ([]*InvoiceShort, *MetaInformation, error) {
GetInvoice(id string) (*InvoiceFull, error) {
GetCompanySettings() (*CompanySettings, error) {
GetArticles(p *QueryParams) ([]*Article, *MetaInformation, error) {
GetArticle(id string) (*Article, error) {
GetLabels() ([]*Label, error) {
CreateLabel(name string) (*Label, error) {
```

There are quite a few endpoints that aren't implemented yet. Feel free to make an issue or pull request.

## 'ish Types (Floatish, Intish)

For some reason the fortnox api ocassionally gives back a float but sometimes a string for certain fields. 
I made these two types for dealing with those situations for unmarshalling.

## Running Tests

You need to set the `FORTNOX_AUTH_CODE`, `FORTNOX_ACCESS_TOKEN` and `FORTNOX_CLIENT_SECRET` envs when running the tests.
