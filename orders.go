package fortnox

import (
	"context"
	"fmt"
	"net/url"
	"time"
)

// EmailInformation data type
type EmailInformation struct {
	EmailAddressBCC  string `json:"EmailAddressBCC"`
	EmailAddressCC   string `json:"EmailAddressCC"`
	EmailAddressFrom string `json:"EmailAddressFrom"`
	EmailAddressTo   string `json:"EmailAddressTo"`
	EmailBody        string `json:"EmailBody"`
	EmailSubject     string `json:"EmailSubject"`
}

// OrderShort data type
type OrderShort struct {
	URL                       string  `json:"@url"`
	Cancelled                 bool    `json:"Cancelled"`
	Currency                  string  `json:"Currency"`
	CustomerName              string  `json:"CustomerName"`
	CustomerNumber            string  `json:"CustomerNumber"`
	DeliveryDate              Date    `json:"DeliveryDate"`
	DocumentNumber            string  `json:"DocumentNumber"`
	ExternalInvoiceReference1 string  `json:"ExternalInvoiceReference1"`
	ExternalInvoiceReference2 string  `json:"ExternalInvoiceReference2"`
	OrderDate                 Date    `json:"OrderDate"`
	Project                   string  `json:"Project"`
	Total                     float64 `json:"Total"`
}

// OrderRow data type
type OrderRow struct {
	AccountNumber          int      `json:"AccountNumber"`
	ArticleNumber          string   `json:"ArticleNumber"`
	ContributionPercent    Floatish `json:"ContributionPercent,omitempty"`
	ContributionValue      Floatish `json:"ContributionValue,omitempty"`
	CostCenter             string   `json:"CostCenter"`
	DeliveredQuantity      string   `json:"DeliveredQuantity"`
	Description            string   `json:"Description"`
	Discount               int      `json:"Discount"`
	DiscountType           string   `json:"DiscountType"`
	HouseWork              bool     `json:"HouseWork"`
	HouseWorkHoursToReport int      `json:"HouseWorkHoursToReport"`
	HouseWorkType          string   `json:"HouseWorkType"`
	OrderedQuantity        string   `json:"OrderedQuantity"`
	Price                  float64  `json:"Price"`
	Project                string   `json:"Project"`
	Total                  float64  `json:"Total"`
	Unit                   string   `json:"Unit"`
	VAT                    float64  `json:"VAT"`
}

// CreateOrderRow payload for order rows when creating new order. Pointers since most fields are not required.
type CreateOrderRow struct {
	AccountNumber          *int64   `json:"AccountNumber"`
	ArticleNumber          *string  `json:"ArticleNumber,omitempty"`
	CostCenter             *string  `json:"CostCenter"`
	DeliveredQuantity      *string  `json:"DeliveredQuantity,omitempty"`
	Description            *string  `json:"Description,omitempty"`
	Discount               *int64   `json:"Discount,omitempty"`
	DiscountType           *string  `json:"DiscountType,omitempty"`
	HouseWork              *bool    `json:"HouseWork,omitempty"`
	HouseWorkHoursToReport *int64   `json:"HouseWorkHoursToReport,omitempty"`
	HouseWorkType          *string  `json:"HouseWorkType,omitempty"`
	OrderedQuantity        *string  `json:"OrderedQuantity,omitempty"`
	Price                  *float64 `json:"Price,omitempty"`
	Project                *string  `json:"Project,omitempty"`
	Unit                   *string  `json:"Unit,omitempty"`
	VAT                    *float64 `json:"VAT,omitempty"`
}

// CreateOrder payload for creating orders
type CreateOrder struct {
	AdministrationFee         *float64          `json:"AdministrationFee,omitempty"`
	Address1                  *string           `json:"Address1,omitempty"`
	Address2                  *string           `json:"Address2,omitempty"`
	City                      *string           `json:"City,omitempty"`
	Comments                  *string           `json:"Comments,omitempty"`
	CopyRemarks               *bool             `json:"CopyRemarks,omitempty"`
	Country                   *string           `json:"Country,omitempty"`
	CostCenter                *string           `json:"CostCenter,omitempty"`
	Currency                  *string           `json:"Currency,omitempty"`
	CurrencyRate              *float64          `json:"CurrencyRate,omitempty"`
	CurrencyUnit              *float64          `json:"CurrencyUnit,omitempty"`
	CustomerName              *string           `json:"CustomerName,omitempty"`
	CustomerNumber            *string           `json:"CustomerNumber,omitempty"`
	DeliveryAddress1          *string           `json:"DeliveryAddress1,omitempty"`
	DeliveryAddress2          *string           `json:"DeliveryAddress2,omitempty"`
	DeliveryCity              *string           `json:"DeliveryCity,omitempty"`
	DeliveryCountry           *string           `json:"DeliveryCountry,omitempty"`
	DeliveryDate              *string           `json:"DeliveryDate,omitempty"`
	DeliveryName              *string           `json:"DeliveryName,omitempty"`
	DeliveryZipCode           *string           `json:"DeliveryZipCode,omitempty"`
	DocumentNumber            *Intish           `json:"DocumentNumber,omitempty"`
	EmailInformation          *EmailInformation `json:"EmailInformation,omitempty"`
	ExternalInvoiceReference1 *string           `json:"ExternalInvoiceReference1,omitempty"`
	ExternalInvoiceReference2 *string           `json:"ExternalInvoiceReference2,omitempty"`
	Freight                   *float64          `json:"Freight,omitempty"`
	Language                  *string           `json:"Language,omitempty"`
	Labels                    []*Label          `json:"Labels,omitempty"`
	NotCompleted              *bool             `json:"NotCompleted,omitempty"`
	OrderDate                 *string           `json:"OrderDate,omitempty"`
	OrderRows                 []*CreateOrderRow `json:"OrderRows,omitempty"`
	OurReference              *string           `json:"OurReference,omitempty"`
	Phone1                    *string           `json:"Phone1,omitempty"`
	Phone2                    *string           `json:"Phone2,omitempty"`
	PriceList                 *string           `json:"PriceList,omitempty"`
	PrintTemplate             *string           `json:"PrintTemplate,omitempty"`
	Project                   *string           `json:"Project,omitempty"`
	Remarks                   *string           `json:"Remarks,omitempty"`
	TermsOfDelivery           *string           `json:"TermsOfDelivery,omitempty"`
	TermsOfPayment            *StringIsh        `json:"TermsOfPayment,omitempty"`
	VATIncluded               *bool             `json:"VATIncluded,omitempty"`
	WayOfDelivery             *string           `json:"WayOfDelivery,omitempty"`
	YourReference             *string           `json:"YourReference,omitempty"`
	YourOrderNumber           *string           `json:"YourOrderNumber,omitempty"`
	ZipCode                   *string           `json:"ZipCode,omitempty"`
}

// UpdateOrder payload for updating orders
type UpdateOrder CreateOrder

// OrderFull data type
type OrderFull struct {
	URL                       string           `json:"@url"`
	URLTaxReductionList       string           `json:"@urlTaxReductionList"`
	AdministrationFee         float64          `json:"AdministrationFee"`
	AdministrationFeeVAT      float64          `json:"AdministrationFeeVAT,omitempty"`
	Address1                  string           `json:"Address1"`
	Address2                  string           `json:"Address2"`
	BasisTaxReduction         float64          `json:"BasisTaxReduction,omitempty"`
	Cancelled                 bool             `json:"Cancelled,omitempty"`
	City                      string           `json:"City"`
	Comments                  string           `json:"Comments"`
	ContributionPercent       Floatish         `json:"ContributionPercent,omitempty"`
	ContributionValue         Floatish         `json:"ContributionValue,omitempty"`
	CopyRemarks               bool             `json:"CopyRemarks"`
	Country                   string           `json:"Country"`
	CostCenter                string           `json:"CostCenter"`
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
	EmailInformation          EmailInformation `json:"EmailInformation"`
	ExternalInvoiceReference1 string           `json:"ExternalInvoiceReference1"`
	ExternalInvoiceReference2 string           `json:"ExternalInvoiceReference2"`
	Freight                   float64          `json:"Freight"`
	FreightVAT                float64          `json:"FreightVAT"`
	Gross                     float64          `json:"Gross"`
	HouseWork                 bool             `json:"HouseWork"`
	InvoiceReference          Intish           `json:"InvoiceReference"`
	Language                  string           `json:"Language"`
	Labels                    []Label          `json:"Labels"`
	Net                       float64          `json:"Net"`
	NotCompleted              bool             `json:"NotCompleted"`
	OfferReference            Intish           `json:"OfferReference"`
	OrderDate                 Date             `json:"OrderDate"`
	OrderRows                 []OrderRow       `json:"OrderRows"`
	OrganisationNumber        string           `json:"OrganisationNumber"`
	OurReference              string           `json:"OurReference"`
	Phone1                    string           `json:"Phone1"`
	Phone2                    string           `json:"Phone2"`
	PriceList                 string           `json:"PriceList"`
	PrintTemplate             string           `json:"PrintTemplate"`
	Project                   string           `json:"Project"`
	Remarks                   string           `json:"Remarks"`
	RoundOff                  float64          `json:"RoundOff"`
	Sent                      bool             `json:"Sent"`
	TaxReduction              float64          `json:"TaxReduction"`
	TermsOfDelivery           string           `json:"TermsOfDelivery"`
	TermsOfPayment            StringIsh        `json:"TermsOfPayment"`
	Total                     float64          `json:"Total"`
	TotalToPay                float64          `json:"TotalToPay"`
	TotalVat                  float64          `json:"TotalVat,omitempty"`
	VATIncluded               bool             `json:"VATIncluded"`
	WayOfDelivery             string           `json:"WayOfDelivery"`
	YourReference             string           `json:"YourReference"`
	YourOrderNumber           string           `json:"YourOrderNumber"`
	ZipCode                   string           `json:"ZipCode"`
}

// OrderQueryParams when searching orders/invoices
// From fortnox docs:
// SEARCH NAME	DESCRIPTION	EXAMPLE VALUE
// lastmodified	Retrieves all records since the provided timestamp.	2014-03-10 12:30
// financialyear	Selects what financial year that should be used	5
// financialyeardate	Selects by date, what financial year that should be used	2014-03-10
// fromdate	Defines a selection based on a start date.
// Only available for invoices, orders, offers and vouchers.	2014-03-10
// todate	Defines a selection based on an end date.
// Only available for invoices, orders, offers and vouchers	 2014-03-10
type OrderQueryParams struct {
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

func (p OrderQueryParams) toValues() url.Values {

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

// ListOrdersResp Response when listing orders
type ListOrdersResp struct {
	Orders          []*OrderShort    `json:"Orders"`
	MetaInformation *MetaInformation `json:"MetaInformation"`
}

// ListOrders or search orders
func (c *Client) ListOrders(ctx context.Context, p *OrderQueryParams) (*ListOrdersResp, error) {

	resp := &ListOrdersResp{}

	var vals url.Values
	if p != nil {
		vals = p.toValues()
	}

	err := c.request(ctx, "GET", "orders", nil, vals, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// An OrderResp is the json response for singular order resources
type OrderResp struct {
	Order OrderFull `json:"Order"`
}

// GetOrder gets one order by id
func (c *Client) GetOrder(ctx context.Context, id int) (*OrderFull, error) {

	resp := &OrderResp{}
	err := c.request(ctx, "GET", fmt.Sprintf("orders/%d", id), nil, nil, resp)
	if err != nil {
		return nil, err
	}

	return &resp.Order, nil
}

// CreateOrder creates an order
func (c *Client) CreateOrder(ctx context.Context, order *CreateOrder) (*OrderFull, error) {
	orderResp := &OrderResp{}
	err := c.request(ctx, "POST", "orders/", &struct {
		Order *CreateOrder `json:"Order"`
	}{
		Order: order,
	}, nil, orderResp)
	if err != nil {
		return nil, err
	}

	return &orderResp.Order, nil
}

// UpdateOrder updates an order
func (c *Client) UpdateOrder(ctx context.Context, id int, order *UpdateOrder) (*OrderFull, error) {

	resp := &OrderResp{}
	err := c.request(ctx, "PUT", fmt.Sprintf("orders/%d", id), &struct {
		Order *UpdateOrder `json:"Order"`
	}{
		Order: order,
	}, nil, resp)
	if err != nil {
		return nil, err
	}

	return &resp.Order, nil
}
