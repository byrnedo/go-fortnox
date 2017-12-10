package fortnox

import (
	"encoding/json"
	"fmt"
)

// Floatish type to allow unmarshalling from either string or float
type Floatish struct {
	Value float64
}

func unmarshalIsh(data []byte, receiver interface{}) error {
	if len(data) == 0 {
		return nil
	}
	if data[0] == '"' {
		data = data[1:]
		data = data[:len(data)-1]
	}

	if len(data) < 1 {
		return nil
	}
	return json.Unmarshal(data, receiver)
}

// UnmarshalJSON to allow unmarshalling from either string or float
func (f *Floatish) UnmarshalJSON(data []byte) error {
	return unmarshalIsh(data, &f.Value)
}

// MarshalJSON to allow marshalling of underlying float
func (f *Floatish) MarshalJSON() ([]byte, error) {
	return json.Marshal(f.Value)
}

// Floatish type to allow unmarshalling from either string or int
type Intish struct {
	Value int
}

// UnmarshalJSON to allow unmarshalling from either string or int
func (f *Intish) UnmarshalJSON(data []byte) error {
	return unmarshalIsh(data, &f.Value)
}

// MarshalJSON to allow marshalling of underlying int
func (f *Intish) MarshalJSON() ([]byte, error) {
	return json.Marshal(f.Value)
}

// Date simple fortnox date holder
type Date struct {
	Year  int
	Month int
	Date  int
}

// String representation of fnox date
func (d *Date) String() string {
	return fmt.Sprintf("%04d-%02d-%02d", d.Year, d.Month, d.Date)
}

// UnmarshalJSON of fnox date
func (d *Date) UnmarshalJSON(data []byte) error {

	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	if len(v) != 10 {
		return nil
	}

	if _, err := fmt.Sscanf(v,"%d-%d-%d", &d.Year, &d.Month, &d.Date); err != nil {
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

// Label data type
type Label struct {
	ID          int    `json:"Id"`
	Description string `json:"Description,omitempty"`
}

// CompanySettings data type
type CompanySettings struct {
	Address            string `json:"Address"`
	BG                 string `json:"BG"`
	BIC                string `json:"BIC"`
	BranchCode         string `json:"BranchCode"`
	City               string `json:"City"`
	ContactFirstName   string `json:"ContactFirstName"`
	ContactLastName    string `json:"ContactLastName"`
	Country            string `json:"Country"`
	CountryCode        string `json:"CountryCode"`
	DatabaseNumber     Intish `json:"DatabaseNumber"`
	Domicile           string `json:"Domicile"`
	Email              string `json:"Email"`
	Fax                string `json:"Fax"`
	IBAN               string `json:"IBAN"`
	Name               string `json:"Name"`
	OrganizationNumber string `json:"OrganizationNumber"`
	PG                 string `json:"PG"`
	Phone1             string `json:"Phone1"`
	Phone2             string `json:"Phone2"`
	TaxEnabled         bool   `json:"TaxEnabled"`
	VATNumber          string `json:"VATNumber"`
	VisitAddress       string `json:"VisitAddress"`
	VisitCity          string `json:"VisitCity"`
	VisitCountry       string `json:"VisitCountry"`
	VisitCountryCode   string `json:"VisitCountryCode"`
	VisitName          string `json:"VisitName"`
	VisitZipCode       string `json:"VisitZipCode"`
	WWW                string `json:"WWW"`
	ZipCode            string `json:"ZipCode"`
}

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

// InvoiceRow data type
type InvoiceRow OrderRow

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
	DocumentNumber            *string           `json:"DocumentNumber,omitempty"`
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
	TermsOfPayment            *Intish           `json:"TermsOfPayment,omitempty"`
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
	DocumentNumber            string           `json:"DocumentNumber"`
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
	TermsOfPayment            Intish           `json:"TermsOfPayment"`
	Total                     float64          `json:"Total"`
	TotalToPay                float64          `json:"TotalToPay"`
	TotalVat                  float64          `json:"TotalVat,omitempty"`
	VATIncluded               bool             `json:"VATIncluded"`
	WayOfDelivery             string           `json:"WayOfDelivery"`
	YourReference             string           `json:"YourReference"`
	YourOrderNumber           string           `json:"YourOrderNumber"`
	ZipCode                   string           `json:"ZipCode"`
}

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
	DocumentNumber            string   `json:"DocumentNumber"`
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
	DocumentNumber            string           `json:"DocumentNumber"`
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

// Article data type
type Article struct {
	URL                       string   `json:"@url"`
	ArticleNumber             string   `json:"ArticleNumber"`
	Bulky                     bool     `json:"Bulky"`
	ConstructionAccount       int      `json:"ConstructionAccount"`
	Depth                     int      `json:"Depth"`
	Description               string   `json:"Description"`
	DisposableQuantity        Floatish `json:"DisposableQuantity"`
	EAN                       string   `json:"EAN"`
	EUAccount                 int      `json:"EUAccount"`
	EUVATAccount              int      `json:"EUVATAccount"`
	Expired                   bool     `json:"Expired"`
	ExportAccount             int      `json:"ExportAccount"`
	Height                    int      `json:"Height"`
	Housework                 bool     `json:"Housework"`
	HouseworkType             string   `json:"HouseworkType"`
	Manufacturer              string   `json:"Manufacturer"`
	ManufacturerArticleNumber string   `json:"ManufacturerArticleNumber"`
	Note                      string   `json:"Note"`
	PurchaseAccount           int      `json:"PurchaseAccount"`
	PurchasePrice             Floatish `json:"PurchasePrice"`
	QuantityInStock           Floatish `json:"QuantityInStock"`
	ReservedQuantity          Floatish `json:"ReservedQuantity"`
	SalesAccount              int      `json:"SalesAccount"`
	SalesPrice                Floatish `json:"SalesPrice"`
	StockGoods                bool     `json:"StockGoods"`
	StockPlace                string   `json:"StockPlace"`
	StockValue                Floatish `json:"StockValue"`
	StockWarning              Floatish `json:"StockWarning"`
	SupplierName              string   `json:"SupplierName"`
	SupplierNumber            string   `json:"SupplierNumber"`
	Type                      string   `json:"Type"`
	Unit                      string   `json:"Unit"`
	VAT                       Floatish `json:"VAT"`
	WebshopArticle            bool     `json:"WebshopArticle"`
	Weight                    int      `json:"Weight"`
	Width                     int      `json:"Width"`
}

