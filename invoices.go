package fortnox

import (
	"context"
	"fmt"
)

// InvoiceRow data type
type InvoiceRow OrderRow

// InvoiceShort data type
type InvoiceShort struct {
	URL                       string   `json:"@url"`
	Balance                   float64  `json:"Balance"`
	Booked                    bool     `json:"Booked"`
	Cancelled                 bool     `json:"Cancelled"`
	Currency                  string   `json:"Currency"`
	CurrencyRate              Floatish `json:"CurrencyRate"`
	CurrencyUnit              float64  `json:"CurrencyUnit"`
	CustomerName              string   `json:"CustomerName"`
	CustomerNumber            string   `json:"CustomerNumber"`
	DocumentNumber            Intish   `json:"DocumentNumber"`
	DueDate                   Date     `json:"DueDate"`
	ExternalInvoiceReference1 string   `json:"ExternalInvoiceReference1"`
	ExternalInvoiceReference2 string   `json:"ExternalInvoiceReference2"`
	InvoiceDate               Date     `json:"InvoiceDate"`
	NoxFinans                 bool     `json:"NoxFinans"`
	OCR                       string   `json:"OCR"`
	Project                   string   `json:"Project"`
	Sent                      bool     `json:"Sent"`
	TermsOfPayment            Intish   `json:"TermsOfPayment"`
	Total                     float64  `json:"Total"`
	WayOfDelivery             string   `json:"WayOfDelivery"`
}

// EDIInformation data type
type EDIInformation struct {
	EDIGlobalLocationNumber         string `json:"EDIGlobalLocationNumber"`
	EDIGlobalLocationNumberDelivery string `json:"EDIGlobalLocationNumberDelivery"`
	EDIInvoiceExtra1                string `json:"EDIInvoiceExtra1"`
	EDIInvoiceExtra2                string `json:"EDIInvoiceExtra2"`
	EDIOurElectronicReference       string `json:"EDIOurElectronicReference"`
	EDIYourElectronicReference      string `json:"EDIYourElectronicReference"`
}

// InvoiceFull data type
type InvoiceFull struct {
	URL                       string           `json:"@url"`
	URLTaxReductionList       string           `json:"@urlTaxReductionList"`
	Address1                  string           `json:"Address1"`
	Address2                  string           `json:"Address2"`
	AdministrationFee         float64          `json:"AdministrationFee"`
	AdministrationFeeVAT      float64          `json:"AdministrationFeeVAT"`
	Balance                   float64          `json:"Balance"`
	BasisTaxReduction         float64          `json:"BasisTaxReduction"`
	Booked                    bool             `json:"Booked"`
	Cancelled                 bool             `json:"Cancelled"`
	City                      string           `json:"City"`
	Comments                  string           `json:"Comments"`
	ContractReference         int              `json:"ContractReference"`
	ContributionPercent       Floatish         `json:"ContributionPercent"`
	ContributionValue         Floatish         `json:"ContributionValue"`
	CostCenter                string           `json:"CostCenter"`
	Country                   string           `json:"Country"`
	Credit                    string           `json:"Credit"`
	CreditInvoiceReference    string           `json:"CreditInvoiceReference"`
	Currency                  string           `json:"Currency"`
	CurrencyRate              Floatish         `json:"CurrencyRate"`
	CurrencyUnit              float64          `json:"CurrencyUnit"`
	CustomerName              string           `json:"CustomerName"`
	CustomerNumber            string           `json:"CustomerNumber"`
	DeliveryAddress1          string           `json:"DeliveryAddress1"`
	DeliveryAddress2          string           `json:"DeliveryAddress2"`
	DeliveryCity              string           `json:"DeliveryCity"`
	DeliveryCountry           string           `json:"DeliveryCountry"`
	DeliveryDate              Date             `json:"DeliveryDate"`
	DeliveryName              string           `json:"DeliveryName"`
	DeliveryZipCode           string           `json:"DeliveryZipCode"`
	DocumentNumber            Intish           `json:"DocumentNumber"`
	DueDate                   Date             `json:"DueDate"`
	EDIInformation            EDIInformation   `json:"EDIInformation"`
	EUQuarterlyReport         bool             `json:"EUQuarterlyReport"`
	EmailInformation          EmailInformation `json:"EmailInformation"`
	ExternalInvoiceReference1 string           `json:"ExternalInvoiceReference1"`
	ExternalInvoiceReference2 string           `json:"ExternalInvoiceReference2"`
	Freight                   float64          `json:"Freight"`
	FreightVAT                float64          `json:"FreightVAT"`
	Gross                     float64          `json:"Gross"`
	HouseWork                 bool             `json:"HouseWork"`
	InvoiceDate               Date             `json:"InvoiceDate"`
	InvoicePeriodEnd          Date             `json:"InvoicePeriodEnd"`
	InvoicePeriodStart        Date             `json:"InvoicePeriodStart"`
	InvoiceReference          Intish           `json:"InvoiceReference"`
	InvoiceRows               []InvoiceRow     `json:"InvoiceRows"`
	InvoiceType               string           `json:"InvoiceType"`
	Labels                    []Label          `json:"Labels"`
	Language                  string           `json:"Language"`
	LastRemindDate            Date             `json:"LastRemindDate"`
	Net                       float64          `json:"Net"`
	NotCompleted              bool             `json:"NotCompleted"`
	NoxFinans                 bool             `json:"NoxFinans"`
	OCR                       string           `json:"OCR"`
	OfferReference            Intish           `json:"OfferReference"`
	OrderReference            Intish           `json:"OrderReference"`
	OrganisationNumber        string           `json:"OrganisationNumber"`
	OurReference              string           `json:"OurReference"`
	PaymentWay                string           `json:"PaymentWay"`
	Phone1                    string           `json:"Phone1"`
	Phone2                    string           `json:"Phone2"`
	PriceList                 string           `json:"PriceList"`
	PrintTemplate             string           `json:"PrintTemplate"`
	Project                   string           `json:"Project"`
	Remarks                   string           `json:"Remarks"`
	Reminders                 int              `json:"Reminders"`
	RoundOff                  float64          `json:"RoundOff"`
	Sent                      bool             `json:"Sent"`
	TaxReduction              int              `json:"TaxReduction"`
	TermsOfDelivery           string           `json:"TermsOfDelivery"`
	TermsOfPayment            Intish           `json:"TermsOfPayment"`
	Total                     float64          `json:"Total"`
	TotalToPay                float64          `json:"TotalToPay"`
	TotalVAT                  float64          `json:"TotalVAT"`
	VATIncluded               bool             `json:"VATIncluded"`
	VoucherNumber             int              `json:"VoucherNumber"`
	VoucherSeries             string           `json:"VoucherSeries"`
	VoucherYear               int              `json:"VoucherYear"`
	WayOfDelivery             string           `json:"WayOfDelivery"`
	YourOrderNumber           string           `json:"YourOrderNumber"`
	YourReference             string           `json:"YourReference"`
	ZipCode                   string           `json:"ZipCode"`
}

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

// InvoiceResp Response for single invoice
type InvoiceResp struct {
	Invoice InvoiceFull `json:"Invoice"`
}

// GetInvoice gets one invoice
func (c *Client) GetInvoice(ctx context.Context, id int) (*InvoiceFull, error) {

	resp := &InvoiceResp{}

	err := c.request(ctx, "GET", fmt.Sprintf("invoices/%d", id), nil, nil, resp)
	if err != nil {
		return nil, err
	}

	return &resp.Invoice, nil
}
