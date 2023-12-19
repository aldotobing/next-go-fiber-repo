package models

// Item ...
type ItemUomLineSync struct {
	ID            *string `json:"item_uom_line_id"`
	ItemCode      *string `json:"item_code"`
	ItemName      *string `json:"item_name"`
	UomCode       *string `json:"uom_code"`
	UomName       *string `json:"uom_name"`
	UomConversion *string `json:"uom_conversion"`
}

// ItemParameter ...
type ItemUomLineSyncParameter struct {
	ID                 string `json:"item_id"`
	ItemCode           string `json:"item_code"`
	UomCode            string `json:"uom_code"`
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
	ItemUomLineSyncSelectStatement = `
	select def.id as def_id, i.code as item_code,i._name as item_name,uo.code as uom_code, uo._name as uom_name,
	def.conversion as uom_conversion from item_uom_line def 
	join item i on i.id = def.item_id join uom uo on uo.id =  def.uom_id 
	`

	// ItemWhereStatement ...
	ItemUomLineSyncWhereStatement = ` WHERE def.created_date IS not NULL `
)
