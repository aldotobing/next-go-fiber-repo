package requests

// PriceListSyncRequest ...
type PriceListSyncRequest struct {
	ID                string `json:"price_list_id"`
	Code              string `json:"price_list_code"`
	Name              string `json:"price_list_name"`
	PriceListPrint    string `json:"price_list_print"`
	PriceListBranchID string `json:"price_list_branch_id"`
	CreatedDate       string `json:"created_date"`
	ModifiedDate      string `json:"modified_date"`
}
