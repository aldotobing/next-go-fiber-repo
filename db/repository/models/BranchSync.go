package models

// Branch ...
type BranchSync struct {
	ID                               *string `json:"branch_id"`
	Code                             *string `json:"branch_code"`
	Name                             *string `json:"branch_name"`
	RegionCode                       *string `json:"region_code"`
	GroupCode                        *string `json:"branch_group_code"`
	CityCode                         *string `json:"branch_city_code"`
	LocationCode                     *string `json:"location_code"`
	SOCheckCreditLimit               *string `json:"sales_order_check_credit_limit"`
	DefalutTrackingHistoryRoadMarker *string `json:"default_tracking_history_show_road_marker"`
	DefaultDistanceLimit             *string `json:"default_distance_limit"`
	WorkingStart                     *string `json:"working_hours_start"`
	WorkingEnd                       *string `json:"working_hours_end"`
	Latitude                         *string `json:"latitude"`
	Longitude                        *string `json:"longitude"`
	IdealSalesQty                    *string `json:"ideal_salesman_qty"`
	ManagerID                        *string `json:"manager_id"`
	MinVisitAmount                   *string `json:"minimum_visit_amount"`
	MinOmzetAmount                   *string `json:"min_omzet_amount"`
	DeaultTaxCalcMethod              *string `json:"default_tax_calc"`
	CreatedDate                      *string `json:"created_date"`
	ModifiedDate                     *string `json:"modified_date"`
}

// BranchParameter ...
type BranchSyncParameter struct {
	ID                 string `json:"branch_id"`
	Code               string `json:"branch_code"`
	Name               string `json:"branch_name"`
	DateParam          string `json:"date"`
	BranchCategoryId   string `json:"branch_category_id"`
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
	BranchSyncSelectStatement = `
	select 
	def.id as branch_id,def._name as branch_name,def.branch_code as branch_code,
	r.code as region_code, bg.code as branch_group_code, c.code as branch_city_code,
	wl.code as location_code, def.sales_order_check_credit_limit::varchar, def.default_tracking_history_show_road_marker::varchar,
	def.default_distance_limit::varchar, to_char(def.working_hours_start,'HH24:MI:SS'),to_char(def.working_hours_end,'HH24:MI:SS'),
	def.latitude::varchar, def.longitude::varchar, def.ideal_salesman_qty::varchar,def.manager_id::varchar,
	def.minimum_visit_amount::varchar,def.min_omzet_amount::varchar,def.default_tax_calc
	from branch def
	left join region r on r.id = def.region_id
	left join branch_group bg on bg.id = def.branch_group_id
	left join city c on c.id = def.city_id
	left join warehouse_location wl on wl.id = def.default_location_id
	`

	// BranchWhereStatement ...
	BranchSyncWhereStatement = ` WHERE def.created_date IS not NULL `
)
