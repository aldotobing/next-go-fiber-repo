package requests

// ItemPriceSyncRequest ...
type ItemPriceSyncRequest struct {
	ID                   string `json:"item_price_id"`
	PriceListVersionID   string `json:"price_list_version_id"`
	ItemId               string `json:"item_id"`
	UomId                string `json:"uom_id"`
	Price                string `json:"price"`
	CreatedDate          string `json:"created_date"`
	ModifiedDate         string `json:"modified_date"`
	PriceListVersionCode string `json:"price_list_version_desc"`
	PriceListId          string `json:"price_list_id"`
	PriceListCode        string `json:"price_list_code"`
}
