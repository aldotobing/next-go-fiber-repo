package models

// WebBranchEligiblePromo ...
type WebBranchEligiblePromo struct {
	ID                *string `json:"id"`
	PromoID           *string `json:"promo_id"`
	Code              *string `json:"promo_code"`
	PromoName         *string `json:"promo_name"`
	CustomerLevelId   *string `json:"customer_level_id"`
	CustomerLevelName *string `json:"customer_level_name"`
}

// WebBranchEligiblePromoParameter ...
type WebBranchEligiblePromoParameter struct {
	ID               string `json:"id"`
	PromoID          string `json:"promo_id"`
	Code             string `json:"promo_code"`
	PromoName        string `json:"promo_name"`
	BranchId         string `json:"branch_id"`
	BranchName       string `json:"branch_name"`
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
	WebBranchEligiblePromoOrderBy = []string{"pr.id", "pr._name", "pr.created_date"}
	// WebCustomerTypeEligiblePromoOrderByrByString ...
	WebBranchEligiblePromoOrderByrByString = []string{
		"pr._name",
	}

	// WebCustomerTypeEligiblePromoSelectStatement ...

	WebBranchEligiblePromoSelectStatement = `
	SELECT 
		BEP.ID as customer_type_eligible_promo_id,
		PR.ID AS promo_id, 
		PR.CODE AS promo_code, 
		PR._NAME AS promo_name,
		b.ID AS branch_id,
		b._NAME AS branch_name
	FROM branch_eligible_promo BEP
	LEFT JOIN branch b ON b.id = BEP.branch_id
	LEFT JOIN promo PR ON PR.ID = BEP.promo_id
	`
	// WebCustomerTypeEligiblePromoWhereStatement ...
	WebBranchEligiblePromoWhereStatement = ` 
	WHERE BEP.ID IS NOT NULL 
	`
)
