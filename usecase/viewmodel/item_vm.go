package viewmodel

// ItemVM ...
type ItemVM struct {
	ID                 *string `json:"item_id"`
	UOMLineID          *string `json:"item_uom_line_id"`
	Code               *string `json:"item_code"`
	Name               *string `json:"item_name"`
	Description        *string `json:"item_description"`
	ItemCategoryId     *string `json:"item_category_id"`
	ItemCategoryName   *string `json:"item_category_name"`
	ItemPicture        *string `json:"item_picture"`
	PriceListVersionId *string `json:"price_list_version_id"`
	LowestUom          string  `json:"lowest_uom"`
	Uom                []Uom   `json:"item_uom"`
}
