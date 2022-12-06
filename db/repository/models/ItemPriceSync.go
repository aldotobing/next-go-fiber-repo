package models

// ItemPriceSync ...
type ItemPriceSync struct {
	ID                   *string `json:"item_price_id"`
	PriceListId          *string `json:"price_list_id"`
	PriceListCode        *string `json:"price_list_code"`
	PriceListVersionID   *string `json:"price_list_version_id"`
	PriceListVersionCode *string `json:"price_list_version_desc"`
	ItemId               *string `json:"item_id"`
	UomId                *string `json:"uom_id"`
	Price                *string `json:"price"`
	CreatedDate          *string `json:"created_date"`
	ModifiedDate         *string `json:"modified_date"`
}

// ItemPriceSyncParameter ...
type ItemPriceSyncParameter struct {
	ID                   string `json:"item_price_id"`
	DateParam            string `json:"date"`
	PriceListVersionId   string `json:"price_list_version_id"`
	PriceListVersionCode string `json:"price_list_version_desc"`
	PriceListCode        string `json:"price_list_code"`
	ItemId               string `json:"item_id"`
	UomId                string `json:"uom_id"`
	Search               string `json:"search"`
	Page                 int    `json:"page"`
	Offset               int    `json:"offset"`
	Limit                int    `json:"limit"`
	By                   string `json:"by"`
	Sort                 string `json:"sort"`
	ExceptId             string `json:"except_id"`
}

var (
	ItemPriceSyncSelectStatement = `
	SELECT 
  	pl.id::varchar as price_list_id,pl.code as price_list_code,
    def.id::varchar as item_price_id
    , plv.id::varchar as price_list_version_id
    , plv.description as price_list_version_desc
    , def.item_id::varchar as item_id
    , def.uom_id::varchar as uom_id
    , def.price::varchar as price
    , to_char(def.created_date,'YYYY-MM-DD HH24:mi:ss')  as created_date
    , to_char(def.modified_date,'YYYY-MM-DD HH24:mi:ss')  as modified_date
  from item_price def
  join item i on i.id = def.item_id
  join uom u on u.id = def.uom_id
  join price_list_version plv on def.price_list_version_id = plv.id
  join price_list pl on pl.id = plv.price_list_id
	`

	// ItemPriceSyncWhereStatement ...
	ItemPriceSyncWhereStatement = ` WHERE def.created_date IS not NULL `
)
