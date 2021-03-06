package fortnox

import (
	"context"
	"fmt"
	"net/url"
)

// Article data type
type Article struct {
	URL                       string   `json:"@url"`
	Active                    bool     `json:"Active"`
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

// CreateArticle data type
type CreateArticle struct {
	ArticleNumber             *string   `json:"ArticleNumber,omitempty"`
	Active                    *bool     `json:"Active,omitempty"`
	Bulky                     *bool     `json:"Bulky,omitempty"`
	ConstructionAccount       *int      `json:"ConstructionAccount,omitempty"`
	Depth                     *int      `json:"Depth,omitempty"`
	Description               *string   `json:"Description,omitempty"`
	EAN                       *string   `json:"EAN,omitempty"`
	EUAccount                 *int      `json:"EUAccount,omitempty"`
	EUVATAccount              *int      `json:"EUVATAccount,omitempty"`
	Expired                   *bool     `json:"Expired,omitempty"`
	ExportAccount             *int      `json:"ExportAccount,omitempty"`
	Height                    *int      `json:"Height,omitempty"`
	Housework                 *bool     `json:"Housework,omitempty"`
	HouseworkType             *string   `json:"HouseworkType,omitempty"`
	Manufacturer              *string   `json:"Manufacturer,omitempty"`
	ManufacturerArticleNumber *string   `json:"ManufacturerArticleNumber,omitempty"`
	Note                      *string   `json:"Note,omitempty"`
	PurchaseAccount           *int      `json:"PurchaseAccount,omitempty"`
	PurchasePrice             *Floatish `json:"PurchasePrice,omitempty"`
	QuantityInStock           *Floatish `json:"QuantityInStock,omitempty"`
	SalesAccount              *int      `json:"SalesAccount,omitempty"`
	StockGoods                *bool     `json:"StockGoods,omitempty"`
	StockPlace                *string   `json:"StockPlace,omitempty"`
	StockWarning              *Floatish `json:"StockWarning,omitempty"`
	SupplierNumber            *string   `json:"SupplierNumber,omitempty"`
	Type                      *string   `json:"Type,omitempty"`
	Unit                      *string   `json:"Unit,omitempty"`
	VAT                       *Floatish `json:"VAT,omitempty"`
	WebshopArticle            *bool     `json:"WebshopArticle,omitempty"`
	Weight                    *int      `json:"Weight,omitempty"`
	Width                     *int      `json:"Width,omitempty"`
}

// UpdateArticle data type
type UpdateArticle CreateArticle

// ListArticlesResp is the response for ListArticles
type ListArticlesResp struct {
	Articles        []*Article       `json:"Articles"`
	MetaInformation *MetaInformation `json:"MetaInformation"`
}

// ArticleQueryParams are used to query articles
type ArticleQueryParams struct {
	ArticleNumber             string
	Description               string
	EAN                       string
	Manufacturer              string
	ManufacturerArticleNumber string
	SupplierName              string
	Page                      int
	Limit                     int
	Offset                    int
	Extra                     map[string][]string
}

func (p ArticleQueryParams) toValues() url.Values {

	ret := make(url.Values)

	if len(p.ArticleNumber) > 0 {
		ret["articlenumber"] = []string{p.ArticleNumber}
	}
	if len(p.Description) > 0 {
		ret["description"] = []string{p.Description}
	}
	if len(p.EAN) > 0 {
		ret["ean"] = []string{p.EAN}
	}
	if len(p.Manufacturer) > 0 {
		ret["manufacturer"] = []string{p.Manufacturer}
	}
	if len(p.ManufacturerArticleNumber) > 0 {
		ret["manufacturerarticlenumber"] = []string{p.ManufacturerArticleNumber}
	}
	if len(p.SupplierName) > 0 {
		ret["suppliername"] = []string{p.SupplierName}
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

// ListArticles lists or searches articles
func (c *Client) ListArticles(ctx context.Context, p *ArticleQueryParams) (*ListArticlesResp, error) {
	resp := &ListArticlesResp{}

	var vals url.Values
	if p != nil {
		vals = p.toValues()
	}

	err := c.request(ctx, "GET", "articles", nil, vals, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ArticleResp Response for single article
type ArticleResp struct {
	Article Article `json:"Article"`
}

// GetArticle gets one article
func (c *Client) GetArticle(ctx context.Context, artNum string) (*Article, error) {

	resp := &ArticleResp{}

	err := c.request(ctx, "GET", "articles/"+artNum, nil, nil, resp)
	if err != nil {
		return nil, err
	}

	return &resp.Article, nil
}

// CreateArticle creates an order
func (c *Client) CreateArticle(ctx context.Context, article *CreateArticle) (*Article, error) {
	resp := &ArticleResp{}
	err := c.request(ctx, "POST", "articles/", &struct {
		Article *CreateArticle `json:"Article"`
	}{
		Article: article,
	}, nil, resp)
	if err != nil {
		return nil, err
	}

	return &resp.Article, nil
}

// UpdateArticle updates an order
func (c *Client) UpdateArticle(ctx context.Context, artNum string, article *UpdateArticle) (*Article, error) {
	resp := &ArticleResp{}
	err := c.request(ctx, "PUT", "articles/"+artNum, &struct {
		Article *UpdateArticle `json:"Article"`
	}{
		Article: article,
	}, nil, resp)
	if err != nil {
		return nil, err
	}

	return &resp.Article, nil
}

// DeleteArticle deletes one article
func (c *Client) DeleteArticle(ctx context.Context, artNum string) error {
	return c.deleteResource(ctx, "articles/"+artNum)
}
