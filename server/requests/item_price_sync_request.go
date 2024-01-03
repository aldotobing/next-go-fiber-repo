package requests

// ItemPriceSyncRequest ...
type ItemPriceSyncRequest struct {
	ID                   string `json:"item_price_id"`
	PriceListVersionID   string `json:"price_list_version_id"`
	ItemCode             string `json:"item_code"`
	UomCode              string `json:"uom_code"`
	Price                string `json:"price"`
	CreatedDate          string `json:"created_date"`
	ModifiedDate         string `json:"modified_date"`
	PriceListVersionCode string `json:"price_list_version_desc"`
	PriceListId          string `json:"price_list_id"`
	PriceListCode        string `json:"price_list_code"`
}
