package requests

// WebPromoRequest ...
type WebItemRequest struct {
	ID              string `json:"item_id"`
	Code            string `json:"item_code"`
	Name            string `json:"item_name"`
	ItemPicture     string `json:"item_picture"`
	ItemCategoryId  string `json:"item_category_id"`
	ItemHide        string `json:"item_hide"`
	ItemActive      string `json:"item_active"`
	ItemDescription string `json:"item_description"`
}
