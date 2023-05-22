package requests

// WebPromoRequest ...
type WebPromoRequest struct {
	ID                 string `json:"promo_id"`
	Code               string `json:"promo_code"`
	PromoName          string `json:"promo_name"`
	PromoDescription   string `json:"promo_description"`
	PromoUrlBanner     string `json:"promo_url_banner"`
	StartDate          string `json:"start_date"`
	EndDate            string `json:"end_date"`
	ShowInApp          int    `json:"show_in_app"`
	CustomerTypeIdList string `json:"customer_type_id_list"`
	RegionIDList       string `json:"region_area_id_list"`
	Active             int    `json:"active"`
}
