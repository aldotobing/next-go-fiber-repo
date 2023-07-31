package requests

// BranchRequest ...
type BranchRequest struct {
	PICPhoneNo string `json:"pic_phone_no" validate:"required"`
}
