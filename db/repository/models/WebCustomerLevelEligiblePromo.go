package models

// WebCustomerLevelEligiblePromo ...
type WebCustomerLevelEligiblePromo struct {
	ID                *string `json:"id"`
	PromoID           *string `json:"promo_id"`
	Code              *string `json:"promo_code"`
	PromoName         *string `json:"promo_name"`
	CustomerLevelId   *string `json:"customer_level_id"`
	CustomerLevelName *string `json:"customer_level_name"`
}

// WebCustomerLevelEligiblePromoParameter ...
type WebCustomerLevelEligiblePromoParameter struct {
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
	// WebCustomerTypeEligiblePromoOrderBy ...
	WebCustomerLevelEligiblePromoOrderBy = []string{"pr.id", "pr._name", "pr.created_date"}
	// WebCustomerTypeEligiblePromoOrderByrByString ...
	WebCustomerLevelEligiblePromoOrderByrByString = []string{
		"pr._name",
	}

	// WebCustomerTypeEligiblePromoSelectStatement ...

	WebCustomerLevelEligiblePromoSelectStatement = `
	SELECT 
		CTEP.ID as customer_type_eligible_promo_id,
		PR.ID AS promo_id, 
		PR.CODE AS promo_code, 
		PR._NAME AS promo_name,
		CT.ID AS customer_level_id,
		CT._NAME AS customer_level_name
	FROM customer_level_eligible_promo CTEP
	LEFT JOIN customer_level ct ON ct.id = CTEP.customer_level_id
	LEFT JOIN promo PR ON PR.ID = CTEP.promo_id
	`
	// WebCustomerTypeEligiblePromoWhereStatement ...
	WebCustomerLevelEligiblePromoWhereStatement = ` 
	WHERE CTEP.ID IS NOT NULL 
	`
)
