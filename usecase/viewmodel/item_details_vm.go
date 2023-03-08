package viewmodel

type ItemDetailsVM struct {
	ID                      *string `json:"item_id"`
	Code                    *string `json:"item_code"`
	Name                    *string `json:"item_name"`
	Description             *string `json:"item_description"`
	ItemDetailsCategoryId   *string `json:"item_category_id"`
	ItemDetailsCategoryName *string `json:"item_category_name"`
	ItemDetailsPicture      *string `json:"item_picture"`
	Uom                     []Uom   `json:"uom"`
	PriceListVersionId      *string `json:"price_list_version_id"`
}

type Uom struct {
	ID               *string `json:"uom_id"`
	Name             *string `json:"uom_name"`
	Conversion       *string `json:"uom_line_conversion"`
	ItemDetailsPrice *string `json:"item_price"`
}
