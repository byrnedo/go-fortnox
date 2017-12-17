package fortnox

import (
	"context"
	"net/url"
	"fmt"
)

// A Customer is the payload in the responses from the customer endpoint
type Customer struct {
	URL            string `json:"@url"`
	Active         bool   `json:"Active"`
	Address1       string `json:"Address1"`
	Address2       string `json:"Address2"`
	City           string `json:"City"`
	Comments       string `json:"Comments"`
	CostCenter     string `json:"CostCenter"`
	Country        string `json:"Country"`
	CountryCode    string `json:"CountryCode"`
	Currency       string `json:"Currency"`
	CustomerNumber string `json:"CustomerNumber"`
	DefaultDeliveryTypes struct {
		Invoice string `json:"Invoice"`
		Offer   string `json:"Offer"`
		Order   string `json:"Order"`
	} `json:"DefaultDeliveryTypes"`
	DefaultTemplates struct {
		CashInvoice string `json:"CashInvoice"`
		Invoice     string `json:"Invoice"`
		Offer       string `json:"Offer"`
		Order       string `json:"Order"`
	} `json:"DefaultTemplates"`
	DeliveryAddress1         string  `json:"DeliveryAddress1"`
	DeliveryAddress2         string  `json:"DeliveryAddress2"`
	DeliveryCity             string  `json:"DeliveryCity"`
	DeliveryCountry          string  `json:"DeliveryCountry"`
	DeliveryCountryCode      string  `json:"DeliveryCountryCode"`
	DeliveryFax              string  `json:"DeliveryFax"`
	DeliveryName             string  `json:"DeliveryName"`
	DeliveryPhone1           string  `json:"DeliveryPhone1"`
	DeliveryPhone2           string  `json:"DeliveryPhone2"`
	DeliveryZipCode          string  `json:"DeliveryZipCode"`
	Email                    string  `json:"Email"`
	EmailInvoice             string  `json:"EmailInvoice"`
	EmailInvoiceBCC          string  `json:"EmailInvoiceBCC"`
	EmailInvoiceCC           string  `json:"EmailInvoiceCC"`
	EmailOffer               string  `json:"EmailOffer"`
	EmailOfferBCC            string  `json:"EmailOfferBCC"`
	EmailOfferCC             string  `json:"EmailOfferCC"`
	EmailOrder               string  `json:"EmailOrder"`
	EmailOrderBCC            string  `json:"EmailOrderBCC"`
	EmailOrderCC             string  `json:"EmailOrderCC"`
	Fax                      string  `json:"Fax"`
	GLN                      string  `json:"GLN"`
	GLNDelivery              string  `json:"GLNDelivery"`
	InvoiceAdministrationFee float64 `json:"InvoiceAdministrationFee"`
	InvoiceDiscount          float64 `json:"InvoiceDiscount"`
	InvoiceFreight           float64 `json:"InvoiceFreight"`
	InvoiceRemark            string  `json:"InvoiceRemark"`
	Name                     string  `json:"Name"`
	OrganisationNumber       string  `json:"OrganisationNumber"`
	OurReference             string  `json:"OurReference"`
	Phone1                   string  `json:"Phone1"`
	Phone2                   string  `json:"Phone2"`
	PriceList                string  `json:"PriceList"`
	Project                  string  `json:"Project"`
	SalesAccount             Intish  `json:"SalesAccount"`
	ShowPriceVATIncluded     bool    `json:"ShowPriceVATIncluded"`
	TermsOfDelivery          string  `json:"TermsOfDelivery"`
	TermsOfPayment           string  `json:"TermsOfPayment"`
	Type                     string  `json:"Type"`
	VATNumber                string  `json:"VATNumber"`
	VATType                  string  `json:"VATType"`
	VisitingAddress          string  `json:"VisitingAddress"`
	VisitingCity             string  `json:"VisitingCity"`
	VisitingCountry          string  `json:"VisitingCountry"`
	VisitingCountryCode      string  `json:"VisitingCountryCode"`
	VisitingZipCode          string  `json:"VisitingZipCode"`
	WWW                      string  `json:"WWW"`
	WayOfDelivery            string  `json:"WayOfDelivery"`
	YourReference            string  `json:"YourReference"`
	ZipCode                  string  `json:"ZipCode"`
}

// A CreateCustomer is the payload when creating customer
type CreateCustomer struct {
	Active         *bool   `json:"Active,omitempty"`
	Address1       *string `json:"Address1,omitempty"`
	Address2       *string `json:"Address2,omitempty"`
	City           *string `json:"City,omitempty"`
	Comments       *string `json:"Comments,omitempty"`
	CostCenter     *string `json:"CostCenter,omitempty"`
	CountryCode    *string `json:"CountryCode,omitempty"`
	Currency       *string `json:"Currency,omitempty"`
	CustomerNumber *string  `json:"CustomerNumber,omitempty"`
	DefaultDeliveryTypes *struct {
		Invoice *string `json:"Invoice,omitempty"`
		Offer   *string `json:"Offer,omitempty"`
		Order   *string `json:"Order,omitempty"`
	} `json:"DefaultDeliveryTypes,omitempty"`
	DefaultTemplates *struct {
		CashInvoice *string `json:"CashInvoice,omitempty"`
		Invoice     *string `json:"Invoice,omitempty"`
		Offer       *string `json:"Offer,omitempty"`
		Order       *string `json:"Order,omitempty"`
	} `json:"DefaultTemplates,omitempty"`
	DeliveryAddress1         *string  `json:"DeliveryAddress1,omitempty"`
	DeliveryAddress2         *string  `json:"DeliveryAddress2,omitempty"`
	DeliveryCity             *string  `json:"DeliveryCity,omitempty"`
	DeliveryCountryCode      *string  `json:"DeliveryCountryCode,omitempty"`
	DeliveryFax              *string  `json:"DeliveryFax,omitempty"`
	DeliveryName             *string  `json:"DeliveryName,omitempty"`
	DeliveryPhone1           *string  `json:"DeliveryPhone1,omitempty"`
	DeliveryPhone2           *string  `json:"DeliveryPhone2,omitempty"`
	DeliveryZipCode          *string  `json:"DeliveryZipCode,omitempty"`
	Email                    *string  `json:"Email,omitempty"`
	EmailInvoice             *string  `json:"EmailInvoice,omitempty"`
	EmailInvoiceBCC          *string  `json:"EmailInvoiceBCC,omitempty"`
	EmailInvoiceCC           *string  `json:"EmailInvoiceCC,omitempty"`
	EmailOffer               *string  `json:"EmailOffer,omitempty"`
	EmailOfferBCC            *string  `json:"EmailOfferBCC,omitempty"`
	EmailOfferCC             *string  `json:"EmailOfferCC,omitempty"`
	EmailOrder               *string  `json:"EmailOrder,omitempty"`
	EmailOrderBCC            *string  `json:"EmailOrderBCC,omitempty"`
	EmailOrderCC             *string  `json:"EmailOrderCC,omitempty"`
	Fax                      *string  `json:"Fax,omitempty"`
	GLN                      *string  `json:"GLN,omitempty"`
	GLNDelivery              *string  `json:"GLNDelivery,omitempty"`
	InvoiceAdministrationFee *float64 `json:"InvoiceAdministrationFee,omitempty"`
	InvoiceDiscount          *float64 `json:"InvoiceDiscount,omitempty"`
	InvoiceFreight           *float64 `json:"InvoiceFreight,omitempty"`
	InvoiceRemark            *string  `json:"InvoiceRemark,omitempty"`
	Name                     *string  `json:"Name,omitempty"`
	OrganisationNumber       *string  `json:"OrganisationNumber,omitempty"`
	OurReference             *string  `json:"OurReference,omitempty"`
	Phone1                   *string  `json:"Phone1,omitempty"`
	Phone2                   *string  `json:"Phone2,omitempty"`
	PriceList                *string  `json:"PriceList,omitempty"`
	Project                  *string  `json:"Project,omitempty"`
	SalesAccount             *Intish  `json:"SalesAccount,omitempty"`
	ShowPriceVATIncluded     *bool    `json:"ShowPriceVATIncluded,omitempty"`
	TermsOfDelivery          *string  `json:"TermsOfDelivery,omitempty"`
	TermsOfPayment           *string  `json:"TermsOfPayment,omitempty"`
	Type                     *string  `json:"Type,omitempty"`
	VATNumber                *string  `json:"VATNumber,omitempty"`
	VATType                  *string  `json:"VATType,omitempty"`
	VisitingAddress          *string  `json:"VisitingAddress,omitempty"`
	VisitingCity             *string  `json:"VisitingCity,omitempty"`
	VisitingCountryCode      *string  `json:"VisitingCountryCode,omitempty"`
	VisitingZipCode          *string  `json:"VisitingZipCode,omitempty"`
	WWW                      *string  `json:"WWW,omitempty"`
	WayOfDelivery            *string  `json:"WayOfDelivery,omitempty"`
	YourReference            *string  `json:"YourReference,omitempty"`
	ZipCode                  *string  `json:"ZipCode,omitempty"`
}

// UpdateCustomer data type
type UpdateCustomer CreateCustomer

// ListCustomersResp is the response for ListCustomers
type ListCustomersResp struct {
	Customers       []*Customer      `json:"Customers"`
	MetaInformation *MetaInformation `json:"MetaInformation"`
}

// A CustomerQueryParams is used for querying customers using ListCustomers
type CustomerQueryParams struct {
	City               string
	CustomerNumber     string
	Email              string
	GLN                string
	GLNDelivery        string
	Name               string
	OrganisationNumber string
	Phone1             string
	ZipCode            string
	Page               int
	Limit              int
	Offset             int
	Extra              map[string][]string
}

func (p *CustomerQueryParams) toValues() url.Values {
	ret := make(url.Values)

	if len(p.City) > 0 {
		ret["city"] = []string{p.City}
	}
	if len(p.CustomerNumber) > 0 {
		ret["customernumber"] = []string{p.CustomerNumber}
	}
	if len(p.Email) > 0 {
		ret["email"] = []string{p.Email}
	}

	if len(p.GLN) > 0 {
		ret["gln"] = []string{p.GLN}
	}
	if len(p.GLNDelivery) > 0 {
		ret["glndelivery"] = []string{p.GLNDelivery}
	}
	if len(p.Name) > 0 {
		ret["name"] = []string{p.Name}
	}
	if len(p.OrganisationNumber) > 0 {
		ret["organisationnumber"] = []string{p.OrganisationNumber}
	}
	if len(p.Phone1) > 0 {
		ret["phone1"] = []string{p.Phone1}
	}
	if len(p.ZipCode) > 0 {
		ret["zipcode"] = []string{p.ZipCode}
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

// ListCustomers lists or searches customers
func (c *Client) ListCustomers(ctx context.Context, p *CustomerQueryParams) (*ListCustomersResp, error) {
	resp := &ListCustomersResp{}

	var vals url.Values
	if p != nil {
		vals = p.toValues()
	}
	err := c.request(ctx, "GET", "customers", nil, vals, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// CustomerResp Response for single customer
type CustomerResp struct {
	Customer Customer `json:"Customer"`
}

// GetCustomer gets one customer
func (c *Client) GetCustomer(ctx context.Context, custNum string) (*Customer, error) {

	resp := &CustomerResp{}

	err := c.request(ctx, "GET", "customers/"+custNum, nil, nil, resp)
	if err != nil {
		return nil, err
	}

	return &resp.Customer, nil
}

// CreateCustomer creates an order
func (c *Client) CreateCustomer(ctx context.Context, customer *CreateCustomer) (*Customer, error) {
	resp := &CustomerResp{}
	err := c.request(ctx, "POST", "customers/", &struct {
		Customer *CreateCustomer `json:"Customer"`
	}{
		Customer: customer,
	}, nil, resp)
	if err != nil {
		return nil, err
	}

	return &resp.Customer, nil
}

// UpdateCustomer updates an order
func (c *Client) UpdateCustomer(ctx context.Context, custNum string, customer *UpdateCustomer) (*Customer, error) {
	resp := &CustomerResp{}
	err := c.request(ctx, "PUT", "customers/"+custNum, &struct {
		Customer *UpdateCustomer `json:"Customer"`
	}{
		Customer: customer,
	}, nil, resp)
	if err != nil {
		return nil, err
	}

	return &resp.Customer, nil
}

// DeleteCustomer deletes one customer
func (c *Client) DeleteCustomer(ctx context.Context, custNum string) error {
	return c.deleteResource(ctx, "customers/"+custNum)
}
