package models

// WebPromo ...
type WebPromo struct {
	ID                  *string                          `json:"promo_id"`
	Code                *string                          `json:"promo_code"`
	PromoName           *string                          `json:"promo_name"`
	PromoDescription    *string                          `json:"promo_description"`
	PromoUrlBanner      *string                          `json:"promo_url_banner"`
	StartDate           *string                          `json:"start_date"`
	EndDate             *string                          `json:"end_date"`
	Active              *int                             `json:"active"`
	ShowInApp           *int                             `json:"show_in_app"`
	CustomerTypeIdList  *string                          `json:"customer_type_id_list"`
	CustomerTypeList    *[]WebCustomerTypeEligiblePromo  `json:"customer_type_list"`
	RegionAreaIdList    *string                          `json:"region_area_id_list"`
	RegionAreaList      *[]WebRegionAreaEligiblePromo    `json:"region_area_list"`
	CustomerLevelIdList *string                          `json:"customer_Level_id_list"`
	CustomerLevelList   *[]WebCustomerLevelEligiblePromo `json:"customer_Level_list"`
}

// WebPromoParameter ...
type WebPromoParameter struct {
	ID               string `json:"promo_id"`
	Code             string `json:"promo_code"`
	PromoName        string `json:"promo_name"`
	CustomerTypeId   string `json:"customer_type_id"`
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
	// WebPromoOrderBy ...
	WebPromoOrderBy = []string{"pc.id", "pc._name", "pc.created_date"}
	// WebPromoOrderByrByString ...
	WebPromoOrderByrByString = []string{
		"pc._name",
	}

	// WebPromoSelectStatement ...

	WebPromoSelectStatement = `
	SELECT 
		PC.ID AS PROMO_ID,
		PC.CODE AS PROMO_CODE,
		PC._NAME AS PROMO_NAME,
		PC.DESCRIPTION AS PROMO_DESCRIPTION,
		(concat('` + PromoImagePath + `',PC.URL_BANNER)) AS PROMO_URL_BANNER,
		PC.START_DATE AS PROMO_START_DATE,
		PC.END_DATE AS PROMO_END_DATE,
		PC.ACTIVE AS ACTIVE,
		PC.show_in_app 
	FROM PROMO PC
	`
	// WebPromoWhereStatement ...
	WebPromoWhereStatement = ` 
	WHERE PC.ID IS NOT NULL AND PC.ACTIVE = 1
	`
)
