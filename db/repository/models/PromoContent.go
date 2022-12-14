package models

// PromoContent ...
type PromoContent struct {
	ID               *string `json:"promo_id"`
	Code             *string `json:"promo_code"`
	PromoName        *string `json:"promo_name"`
	PromoDescription *string `json:"promo_description"`
	PromoUrlBanner   *string `json:"promo_url_banner"`
	StartDate        *string `json:"start_date"`
	EndDate          *string `json:"end_date"`
	Active           *string `json:"active"`
}

// PromoContentParameter ...
type PromoContentParameter struct {
	ID               string `json:"promo_id"`
	Code             string `json:"promo_code"`
	PromoName        string `json:"promo_name"`
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
	// PromoContentOrderBy ...
	PromoContentOrderBy = []string{"pc.id", "pc._name", "pc.created_date"}
	// PromoContentOrderByrByString ...
	PromoContentOrderByrByString = []string{
		"pc._name",
	}

	// PromoContentSelectStatement ...

	PromoContentSelectStatement = `
	SELECT 
		PC.ID AS PROMO_ID,
		PC.CODE AS PROMO_CODE,
		PC._NAME AS PROMO_NAME,
		PC.DESCRIPTION AS PROMO_DESCRIPTION,
		(concat('https://sidomuncul.s3.ap-southeast-3.amazonaws.com/',PC.URL_BANNER)) AS PROMO_URL_BANNER,
		PC.START_DATE AS PROMO_START_DATE,
		PC.END_DATE AS PROMO_END_DATE,
		PC.ACTIVE AS ACTIVE 
	FROM PROMO PC
	`
	// PromoContentWhereStatement ...
	PromoContentWhereStatement = ` 
	WHERE PC.ID IS NOT NULL AND PC.ACTIVE = 1
	`
)
