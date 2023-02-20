package models

// WebRegionAreaEligiblePromo ...
type WebRegionAreaEligiblePromo struct {
	ID         *string `json:"id"`
	PromoID    *string `json:"promo_id"`
	Code       *string `json:"promo_code"`
	PromoName  *string `json:"promo_name"`
	RegionID   *string `json:"region_id"`
	RegionName *string `json:"region_name"`
}

// WebRegionAreaEligiblePromoParameter ...
type WebRegionAreaEligiblePromoParameter struct {
	ID               string `json:"id"`
	PromoID          string `json:"promo_id"`
	Code             string `json:"promo_code"`
	PromoName        string `json:"promo_name"`
	CustomerTypeId   string `json:"customer_type_id"`
	CustomerTypeName string `json:"customer_type_name"`
	PromoDescription string `json:"promo_description"`
	PromoUrlBanner   string `json:"promo_url_banner"`
	StartDate        string `json:"start_date"`
	EndDate          string `json:"end_date"`
	Search           string `json:"search"`
	Page             int    `json:"page"`
	Offset           int    `json:"offset"`
	Limit            int    `json:"limit"`
	By               string `json:"by"`
	Sort             string `json:"sort"`
}

var (
	// WebRegionAreaEligiblePromoOrderBy ...
	WebRegionAreaEligiblePromoOrderBy = []string{"pr.id", "pr._name", "pr.created_date"}
	// WebRegionAreaEligiblePromoOrderByrByString ...
	WebRegionAreaEligiblePromoOrderByrByString = []string{
		"pr._name",
	}

	// WebRegionAreaEligiblePromoSelectStatement ...

	WebRegionAreaEligiblePromoSelectStatement = `
	SELECT 
		RAEP.ID as region_area_eligible_promo_id,
		PR.ID AS promo_id, 
		PR.CODE AS promo_code, 
		PR._NAME AS promo_name,
		r.ID AS region_id,
		r._NAME AS region_name
	FROM region_area_eligible_promo RAEP
	LEFT JOIN region r ON r.id = RAEP.region_id
	LEFT JOIN promo PR ON PR.ID = RAEP.promo_id
	`
	// WebRegionAreaEligiblePromoWhereStatement ...
	WebRegionAreaEligiblePromoWhereStatement = ` 
	WHERE RAEP.ID IS NOT NULL 
	`
)
