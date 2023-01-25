package models

// WebCustomerTypeEligiblePromo ...
type WebCustomerTypeEligiblePromo struct {
	ID               *string `json:"id"`
	PromoID          *string `json:"promo_id"`
	Code             *string `json:"promo_code"`
	PromoName        *string `json:"promo_name"`
	CustomerTypeId   *string `json:"customer_type_id"`
	CustomerTypeName *string `json:"customer_type_name"`
}

// WebCustomerTypeEligiblePromoParameter ...
type WebCustomerTypeEligiblePromoParameter struct {
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
	WebCustomerTypeEligiblePromoOrderBy = []string{"pr.id", "pr._name", "pr.created_date"}
	// WebCustomerTypeEligiblePromoOrderByrByString ...
	WebCustomerTypeEligiblePromoOrderByrByString = []string{
		"pr._name",
	}

	// WebCustomerTypeEligiblePromoSelectStatement ...

	WebCustomerTypeEligiblePromoSelectStatement = `
	SELECT 
		CTEP.ID as customer_type_eligible_promo_id,
		PR.ID AS promo_id, 
		PR.CODE AS promo_code, 
		PR._NAME AS promo_name,
		CT.ID AS customer_type_id,
		CT._NAME AS customer_type_name
	FROM customer_type_eligible_promo CTEP
	LEFT JOIN customer_type ct ON ct.id = CTEP.customer_type_id
	LEFT JOIN promo PR ON PR.ID = CTEP.promo_id
	`
	// WebCustomerTypeEligiblePromoWhereStatement ...
	WebCustomerTypeEligiblePromoWhereStatement = ` 
	WHERE CTEP.ID IS NOT NULL 
	`
)
