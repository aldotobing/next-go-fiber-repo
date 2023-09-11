package viewmodel

// VoucherVM ....
type VoucherVM struct {
	ID                string `json:"id"`
	Code              string `json:"code"`
	Name              string `json:"name"`
	StartDate         string `json:"start_date"`
	EndDate           string `json:"end_date"`
	ImageURL          string `json:"image_url"`
	VoucherCategoryID string `json:"voucher_category_id"`
	CashValue         string `json:"cash_value"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
	DeletedAt         string `json:"deleted_at"`
}
