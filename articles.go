package fortnox

import "context"

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

// ListArticlesResp is the response for ListArticles
type ListArticlesResp struct {
	Articles        []*Article       `json:"Articles"`
	MetaInformation *MetaInformation `json:"MetaInformation"`
}

// ListArticles lists or searches articles
func (c *Client) ListArticles(ctx context.Context, p *QueryParams) (*ListArticlesResp, error) {
	resp := &ListArticlesResp{}

	err := c.request(ctx, "GET", "articles", nil, p, resp)
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
func (c *Client) GetArticle(ctx context.Context, id string) (*Article, error) {

	resp := &ArticleResp{}

	err := c.request(ctx, "GET", "articles/"+id, nil, nil, resp)
	if err != nil {
		return nil, err
	}

	return &resp.Article, nil
}
