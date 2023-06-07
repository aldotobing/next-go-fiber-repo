package requests

// SalesmanDataSyncRequest ...
type SalesmanDataSyncRequest struct {
	Code              *string `json:"saleman_code"`
	Name              *string `json:"salesman_name"`
	SalesmanType      *string `json:"salesman_type"`
	EffectiveSalesman *string `json:"effective_salesman"`
	PhoneNo           *string `json:"phone_no"`
	BranchID          *string `json:"branch_id"`
	Address           *string `json:"salesman_address"`
}
