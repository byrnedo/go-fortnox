# Go Fortnox Client

[![Go Report Card](https://goreportcard.com/badge/github.com/byrnedo/go-fortnox)](https://goreportcard.com/report/github.com/byrnedo/go-fortnox) [![GoDoc](https://godoc.org/github.com/byrnedo/go-fortnox?status.svg)](https://godoc.org/github.com/byrnedo/go-fortnox)

A client for [Fortnox's](https://www.fortnox.se) REST api.

```go
client := fortnox.NewFortnoxClient(fortnox.WithAuthOpts("token", "secret"))

order, err := client.GetOrder(docNumber)
if err != nil {
    return errors.Wrap(err, "Error getting order")
}
```

There are quite a few endpoints that aren't implemented yet. Feel free to make an issue or pull request.

## 'ish Types (Floatish, Intish)

For some reason the fortnox api ocassionally gives back a float but sometimes a string for certain fields. 
I made these two types for dealing with those situations for unmarshalling.

## Running Tests

You need to set the `FORTNOX_AUTH_CODE`, `FORTNOX_ACCESS_TOKEN` and `FORTNOX_CLIENT_SECRET` envs when running the tests.
