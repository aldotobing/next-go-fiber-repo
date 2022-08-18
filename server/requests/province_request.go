package requests

// ProvinceRequest ...
type ProvinceRequest struct {
	Code      string `json:"code_province" validate:"required"`
	Name      string `json:"name_province" validate:"required"`
	IdNation  string `json:"id_nation"`
	CreatedBy int    `json:"created_by_province"`
	UpdatedBy int    `json:"updated_by_province"`
	DeletedBy int    `json:"deleted_by_province"`
}

type MpProvinceDataBreakDownRequest struct {
	Name     *string `json:"name"`
	Code     *string `json:"code"`
	OldID    *int    `json:"id"`
	NationID *int    `json:"negara_id"`
}
