package fortnox

import (
	"context"
	"gopkg.in/jarcoal/httpmock.v1"
	"math/rand"
	"os"
	"testing"
	"time"
)

var (
	accessToken = os.Getenv("FORTNOX_ACCESS_TOKEN")
	secret      = os.Getenv("FORTNOX_CLIENT_SECRET")
)

func init() {
	rand.Seed(time.Now().UnixNano())
	if accessToken == "" {
		panic("must give FORTNOX_ACCESS_TOKEN env")
	}
	if secret == "" {
		panic("must give FORTNOX_CLIENT_SECRET env")
	}
}

func addTestOpts() []OptionsFunc {
	return []OptionsFunc{WithAuthOpts(accessToken, secret)}
}

func TestGetAccessToken(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", DefaultURL,
		httpmock.NewStringResponder(200, `{"Authorization": {"AccessToken": "test"}}`))

	token, err := GetAccessToken(context.Background(), "test", secret, func(opts *AccessTokenOptions) {
		httpmock.ActivateNonDefault(opts.HTTPClient)
	})
	if err != nil {
		t.Fatal(err)
	}
	if token == "" {
		t.Fatal("Token empty")
	}
}

func TestNewFortnoxClient(t *testing.T) {

	c := NewClient(WithAuthOpts("token", "secret"), WithURLOpts("url"))

	if c.clientOptions.BaseURL != "url" {
		t.Fatal("Incorrect url")
	}

	if c.clientOptions.AccessToken != "token" {
		t.Fatal("Incorrect token")
	}

	if c.clientOptions.ClientSecret != "secret" {
		t.Fatal("Incorrect secret")
	}
}

func TestGetOrders(t *testing.T) {
	c := NewClient(addTestOpts()...)

	r, err := c.ListOrders(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if r.MetaInformation == nil {
		t.Fatal("Meta was nil")
	}

	if r.Orders == nil {
		t.Fatal("Response was nil")
	}
	//pretty.Print(r)
}

func TestGetOrder(t *testing.T) {
	c := NewClient(addTestOpts()...)
	for i := 1; i < 10; i++ {
		_, err := c.GetOrder(context.Background(), i)
		if err != nil {
			t.Fatal(err)
		}
	}

}

func TestGetInvoices(t *testing.T) {
	c := NewClient(addTestOpts()...)

	r, err := c.ListInvoices(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if r.MetaInformation == nil {
		t.Fatal("Meta was nil")
	}

	if r.Invoices == nil {
		t.Fatal("Response was nil")
	}
	//pretty.Print(r)
}

func TestGetInvoice(t *testing.T) {
	c := NewClient(addTestOpts()...)
	for i := 1; i < 10; i++ {
		r, err := c.GetInvoice(context.Background(), i)
		if err != nil {
			t.Fatal(err)
		}
		if r == nil {
			t.Fatal("Response was nil")
		}
	}

}

func TestClient_GetCompanySettings(t *testing.T) {
	c := NewClient(addTestOpts()...)
	r, err := c.GetCompanySettings(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	if r == nil {
		t.Fatal("Response was nil")
	}

}

func TestClient_GetArticles(t *testing.T) {

	c := NewClient(addTestOpts()...)

	r, err := c.ListArticles(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if r.MetaInformation == nil {
		t.Fatal("Meta was nil")
	}

	if r.Articles == nil {
		t.Fatal("Response was nil")
	}

}

func TestClient_GetArticle(t *testing.T) {

	c := NewClient(addTestOpts()...)

	r, err := c.GetArticle(context.Background(), "10")
	if err != nil {
		t.Fatal(err)
	}
	if r == nil {
		t.Fatal("Response was nil")
	}

}

func TestClient_GetLabels(t *testing.T) {

	c := NewClient(addTestOpts()...)

	r, err := c.ListLabels(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	if r == nil {
		t.Fatal("Response was nil")
	}

}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func TestClient_CreateLabel(t *testing.T) {

	c := NewClient(addTestOpts()...)
	name := "test" + RandStringBytes(4)
	t.Log(name)
	r, err := c.CreateLabel(context.Background(), name)
	if err != nil {
		t.Fatal(err)
	}

	if r == nil {
		t.Fatal("Response was nil")
	}
}

func TestClient_UpdateLabel(t *testing.T) {
	c := NewClient(addTestOpts()...)
	name := "test" + RandStringBytes(4)
	t.Log(name)
	r, err := c.CreateLabel(context.Background(), name)
	if err != nil {
		t.Fatal(err)
	}
	r2, err := c.UpdateLabel(context.Background(), r.ID, name+"update")
	if err != nil {
		t.Fatal(err)
	}

	if r2 == nil {
		t.Fatal("Response was nil")
	}

}

func TestClient_DeleteLabel(t *testing.T) {
	c := NewClient(addTestOpts()...)
	name := "test" + RandStringBytes(4)
	t.Log(name)
	r, err := c.CreateLabel(context.Background(), name)
	if err != nil {
		t.Fatal(err)
	}
	err = c.DeleteLabel(context.Background(), r.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_CreateOrder(t *testing.T) {

	var (
		c    = NewClient(addTestOpts()...)
		one  = "1"
		desc = "Desc Text"
	)

	order := &CreateOrder{
		CustomerNumber: &one,
		OrderRows: []*CreateOrderRow{
			{Description: &desc},
		},
	}
	r, err := c.CreateOrder(context.Background(), order)
	if err != nil {
		t.Fatal(err)
	}

	if r == nil {
		t.Fatal("Response was nil")
	}

	if len(r.OrderRows) != 1 {
		t.Fatalf("unexpected number of order rows, expected 1, got %d", len(r.OrderRows))
	}

	row := r.OrderRows[0]
	checkTextRow(row, desc, t)

}

func TestClient_UpdateOrder(t *testing.T) {

	var (
		c    = NewClient(addTestOpts()...)
		one  = "1"
		desc = "Desc Text"
		gbg  = "Gothenburg"
	)

	order := &CreateOrder{
		CustomerNumber: &one,
		OrderRows: []*CreateOrderRow{
			{Description: &desc},
		},
	}
	r, err := c.CreateOrder(context.Background(), order)
	if err != nil {
		t.Fatal(err)
	}

	if r == nil {
		t.Fatal("Response was nil")
	}

	desc2 := "Desc Text 2"

	update := &UpdateOrder{
		CustomerNumber: &one,
		DeliveryCity:   &gbg,
		OrderRows: []*CreateOrderRow{
			{Description: &desc},
			{Description: &desc2},
		},
	}

	r, err = c.UpdateOrder(context.Background(), r.DocumentNumber.Value, update)
	if err != nil {
		t.Fatal(err)
	}

	if r == nil {
		t.Fatal("Response was nil")
	}

	if len(r.OrderRows) != 2 {
		t.Fatalf("unexpected number of order rows, expected 2, got %d", len(r.OrderRows))
	}

	checkTextRow(r.OrderRows[0], desc, t)
	checkTextRow(r.OrderRows[1], desc2, t)

}

func checkTextRow(row OrderRow, desc string, t *testing.T) {
	if row.Description != desc {
		t.Fatalf("unexpected description: %s", row.Description)
	}
	// if no article
	if row.AccountNumber != 0 {
		t.Fatalf("unexpected account number: %d", row.AccountNumber)
	}

	if row.CostCenter != "" {
		t.Fatalf("unexpected cost center: %s", row.CostCenter)
	}

}
