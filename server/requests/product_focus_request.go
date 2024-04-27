package requests

// ProductFocusRequest ...
type ProductFocusRequest struct {
	ItemCodes []string `json:"item_codes"`
	BranchIDs []string `json:"branch_ids"`
}
