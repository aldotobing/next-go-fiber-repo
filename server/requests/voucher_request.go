package requests

// VoucherRequest ...
type VoucherRequest struct {
	Code              string `json:"code" validate:"required"`
	Name              string `json:"name" validate:"required"`
	StartDate         string `json:"start_date" validate:"required"`
	EndDate           string `json:"end_date" validate:"required"`
	ImageURL          string `json:"image_url"`
	VoucherCategoryID string `json:"voucher_category_id"`
	CashValue         string `json:"cash_value"`
}
