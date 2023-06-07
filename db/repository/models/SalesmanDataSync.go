package models

// SalesmanDataSync ...
type SalesmanDataSync struct {
	ID                *string `json:"customer_id"`
	PartnerID         *string `json:"partner_id"`
	Code              *string `json:"saleman_code"`
	Name              *string `json:"salesman_name"`
	SalesmanType      *string `json:"salesman_type"`
	EffectiveSalesman *string `json:"effective_salesman"`
	PhoneNo           *string `json:"phone_no"`
	Address           *string `json:"salesman_address"`
	BranchID          *string `json:"branch_id"`
}

// SalesmanDataSyncParameter ...
type SalesmanDataSyncParameter struct {
	ID        string `json:"id_customer"`
	Code      string `json:"customer_code"`
	DateParam string `json:"date_param"`
	MysmOnly  string `json:"mysm_only"`
	Name      string `json:"name"`
	Search    string `json:"search"`
	Page      int    `json:"page"`
	Offset    int    `json:"offset"`
	Limit     int    `json:"limit"`
	By        string `json:"by"`
	Sort      string `json:"sort"`
}

var (
	// SalesmanDataSyncOrderBy ...
	SalesmanDataSyncOrderBy = []string{"c.id", "p._name", "c.created_date"}
	// SalesmanDataSyncOrderByrByString ...
	SalesmanDataSyncOrderByrByString = []string{
		"p._name",
	}

	// SalesmanDataSyncSelectStatement ...
	SalesmanDataSyncSelectStatement = `select s.id,p.id as partner_id, p.code as saleman_code,p._name as salesman_name, 
	p.address as salesman_address,p.phone_no as phone_no,
	s.salesman_type_id as salesman_type,s.effective_salesman as effective_salesman,
	b.id as branch_id 
	from salesman s join branch b on b.id = s.branch_id join partner p on p.id =s.partner_id
	`

	// SalesmanDataSyncWhereStatement ...
	SalesmanDataSyncWhereStatement = `WHERE p._name IS not NULL`
)
