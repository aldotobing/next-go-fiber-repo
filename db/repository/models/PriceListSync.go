package models

// PriceListSync ...
type PriceListSync struct {
	ID                *string `json:"price_list_id"`
	Code              *string `json:"price_list_code"`
	Name              *string `json:"price_list_name"`
	PriceListPrint    *string `json:"price_list_print"`
	PriceListBranchID *string `json:"price_list_branch_id"`
	CreatedDate       *string `json:"created_date"`
	ModifiedDate      *string `json:"modified_date"`
}

// PriceListSyncParameter ...
type PriceListSyncParameter struct {
	ID                 string `json:"price_list_id"`
	Code               string `json:"price_list_code"`
	Name               string `json:"price_list_name"`
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
	PriceListSyncSelectStatement = `
	SELECT 
    def.id::varchar as price_list_id
    , def.code as price_list_code
    , def._name as price_list_name
	, def.print_price_list::varchar as price_list_print
	, def.branch_id::varchar as price_list_branch_id
    , to_char(def.created_date,'YYYY-MM-DD HH24:mi:ss') as price_list_created_date
    , to_char(def.modified_date,'YYYY-MM-DD HH24:mi:ss') as price_list_modified_date
  	from price_list def
	`

	// PriceListSyncWhereStatement ...
	PriceListSyncWhereStatement = ` WHERE def.created_date IS not NULL `
)
