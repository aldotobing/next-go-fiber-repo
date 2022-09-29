package requests

// ItemSyncRequest ...
type ItemSyncRequest struct {
	Code           string `json:"item_code"`
	Name           string `json:"item_name"`
	Picture        string `json:"item_picture" `
	ItemCategoryID string `json:"item_category_id"`
	ItemActive     int    `json:"item_active"`
	CreatedDate    int    `json:"created_date"`
	ModifiedDate   int    `json:"modified_date"`
}
