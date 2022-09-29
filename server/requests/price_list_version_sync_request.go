package requests

// PriceListVersionSyncRequest ...
type PriceListVersionSyncRequest struct {
	ID            string `json:"price_list_version_id"`
	PriceListID   string `json:"price_list_id"`
	PriceListCode string `json:"price_list_code"`
	StartDate     string `json:"price_list_version_strat_date"`
	EndDate       string `json:"price_list_version_end_date"`
	Description   string `json:"price_list_version_description"`
	CreatedDate   string `json:"created_date"`
	ModifiedDate  string `json:"modified_date"`
}
