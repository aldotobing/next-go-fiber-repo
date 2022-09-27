package models

// PriceListVersionSync ...
type PriceListVersionSync struct {
	ID            *string `json:"price_list_version_id"`
	PriceListID   *string `json:"price_list_id"`
	PriceListCode *string `json:"price_list_code"`
	StartDate     *string `json:"price_list_version_strat_date"`
	EndDate       *string `json:"price_list_version_end_date"`
	Description   *string `json:"price_list_version_description"`
	CreatedDate   *string `json:"created_date"`
	ModifiedDate  *string `json:"modified_date"`
}

// PriceListVersionSyncParameter ...
type PriceListVersionSyncParameter struct {
	ID            string `json:"price_list_version_id"`
	DateParam     string `json:"date"`
	PriceListId   string `json:"price_list_id"`
	PriceListCode string `json:"price_list_code"`
	Description   string `json:"desc"`
	Search        string `json:"search"`
	Page          int    `json:"page"`
	Offset        int    `json:"offset"`
	Limit         int    `json:"limit"`
	By            string `json:"by"`
	Sort          string `json:"sort"`
	ExceptId      string `json:"except_id"`
}

var (
	PriceListVersionSyncSelectStatement = `
	SELECT 
    def.id::varchar as price_list_version_id
    , pl.id::varchar as price_list_id
    , to_char(def.start_date ,'YYYY-MM-DD') as  price_list_version_strat_date
    , to_char(def.end_date ,'YYYY-MM-DD') as price_list_version_end_date
    , def.description as price_list_version_description
    , to_char(def.modified_date,'YYYY-MM-DD HH24:mi:ss') as modified_date
    , to_char(def.created_date,'YYYY-MM-DD HH24:mi:ss') as created_date
  from price_list_version def
  join price_list pl on def.price_list_id = pl.id
	`

	// PriceListVersionSyncWhereStatement ...
	PriceListVersionSyncWhereStatement = ` WHERE def.created_date IS not NULL `
)
