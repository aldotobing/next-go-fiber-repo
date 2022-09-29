package models

// Item ...
type ItemSync struct {
	ID             *string `json:"item_id"`
	Code           *string `json:"item_code"`
	Name           *string `json:"item_name"`
	ItemPicture    *string `json:"item_picture"`
	Description    *string `json:"item_description"`
	ItemCategoryId *string `json:"item_category_id"`
	ItemActive     *string `json:"item_active"`
	ItemParentID   *string `json:"item_parent_id"`
	HaveVariant    *string `json:"item_have_variant"`
	ItemAlias      *string `json:"item_alias"`
	Keterangan     *string `json:"item_keterangan"`
	UrlVideo       *string `json:"item_url_video"`
	CreatedDate    *string `json:"created_date"`
	ModifiedDate   *string `json:"modified_date"`
}

// ItemParameter ...
type ItemSyncParameter struct {
	ID                 string `json:"item_id"`
	Code               string `json:"item_code"`
	Name               string `json:"item_name"`
	DateParam          string `json:"date"`
	ItemCategoryId     string `json:"item_category_id"`
	PriceListVersionId string `json:"price_list_version_id"`
	Search             string `json:"search"`
	Page               int    `json:"page"`
	Offset             int    `json:"offset"`
	Limit              int    `json:"limit"`
	By                 string `json:"by"`
	Sort               string `json:"sort"`
	ExceptId           string `json:"except_id"`
}

var (
	ItemSyncSelectStatement = `
	select def.id, def.code, def.code
	from item def
	`

	// ItemWhereStatement ...
	ItemSyncWhereStatement = ` WHERE def.created_date IS not NULL `
)
