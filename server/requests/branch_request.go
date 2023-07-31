package requests

// BranchRequest ...
type BranchRequest struct {
	PICPhoneNo string `json:"pic_phone_no" validate:"required"`
	PICName    string `json:"pic_name" validate:"required"`
}
