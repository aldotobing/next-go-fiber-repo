package viewmodel

type DashboardByGroupID struct {
	RegionID                 *string `json:"region_id_detail"`
	RegionName               *string `json:"region_name_detail"`
	TotalVisitUser           *string `json:"total_visit_user_detail"`
	TotalRepeatUser          *string `json:"total_repeat_order_user_detail"`
	TotalOrderUser           *string `json:"total_order_user_detail"`
	TotalInvoice             *string `json:"total_invoice_user_detail"`
	TotalRegisteredUser      *string `json:"total_registered_user_detail"`
	CustomerCountRepeatOrder *string `json:"customer_count_repeat_order_detail"`
	TotalActiveOutlet        *string `json:"total_active_outlet_detail"`
	TotalOutlet              *string `json:"total_outlet"`
	TotalCompleteCustomer    *string `json:"total_complete_customer"`
}

type DashboardBranchByUserID struct {
	BranchID              *string `json:"branch_id"`
	BranchName            *string `json:"branch_name"`
	BranchCode            *string `json:"branch_code"`
	RegionName            *string `json:"region_name"`
	RegionGroupName       *string `json:"region_group_name"`
	TotalRepeatUser       *string `json:"total_repeat_order_user_detail"`
	TotalRepeatToko       *string `json:"customer_count_repeat_order_detail"`
	TotalOrderUser        *string `json:"total_order_user_detail"`
	TotalInvoice          *string `json:"total_invoice_user_detail"`
	TotalCheckin          *string `json:"total_visit_user_detail"`
	TotalAktifOutlet      *string `json:"total_active_outlet_detail"`
	TotalOutlet           *string `json:"total_outlet"`
	TotalOutletAll        *string `json:"total_outlet_all"`
	TotalRegisteredUser   *string `json:"total_registered_user_detail"`
	TotalCompleteCustomer *string `json:"total_complete_customer"`
}

type DashboardCustomerByUserID struct {
	CustomerID             *string `json:"customer_id"`
	CustomerName           *string `json:"customer_name"`
	CustomerCode           *string `json:"customer_code"`
	BranchName             *string `json:"branch_name"`
	BranchCode             *string `json:"branch_code"`
	RegionName             *string `json:"region_name"`
	RegionGroupName        *string `json:"region_group_name"`
	CustomerLevelName      *string `json:"customer_level_name"`
	CustomerTypeName       *string `json:"customer_type_name"`
	CustomerCityName       *string `json:"cutomer_city_name"`
	TotalRepeatUser        *string `json:"total_repeat_order_user"`
	TotalOrderUser         *string `json:"total_order_user"`
	TotalInvoice           *string `json:"total_invoice_user"`
	TotalCheckin           *string `json:"total_checkin_user"`
	TotalAktifOutlet       *string `json:"total_aktif_outlet"`
	StatusCompleteCustomer *string `json:"status_complete_customer"`
}

type OmzetValueVM struct {
	RegionGroupID   *string `json:"region_group_id"`
	RegionGroupName *string `json:"region_group_name"`
	TotalQuantity   *string `json:"total_quantity"`
	TotalOmzet      *string `json:"total_omzet"`
}

type OmzetValueByRegionVM struct {
	TotalOmzet    *string            `json:"total_omzet"`
	TotalQuantity *string            `json:"total_quantity"`
	Area          []OmzetValueAreaVM `json:"area"`
}

type OmzetValueAreaVM struct {
	RegionID        *string `json:"region_id"`
	RegionName      *string `json:"region_name"`
	Quantity        *string `json:"quantity"`
	Omzet           *string `json:"omzet"`
	RegionGroupID   *string `json:"region_group_id"`
	RegionGroupName *string `json:"region_group_name"`
}

type OmzetValueByBranchVM struct {
	TotalOmzet    *string              `json:"total_omzet"`
	TotalQuantity *string              `json:"total_quantity"`
	Branches      []OmzetValueBranchVM `json:"branches"`
}
type OmzetValueBranchVM struct {
	RegionID        *string `json:"region_id"`
	RegionName      *string `json:"region_name"`
	RegionGroupID   *string `json:"region_group_id"`
	RegionGroupName *string `json:"region_group_name"`
	BranchID        *string `json:"branch_id"`
	BranchName      *string `json:"branch_name"`
	BranchCode      *string `json:"branch_code"`
	Quantity        *string `json:"quantity"`
	Omzet           *string `json:"omzet"`
	ActiveCustomer  *string `json:"active_customer"`
}

type OmzetValueByCustomerVM struct {
	TotalOmzet    *string                `json:"total_omzet"`
	TotalQuantity *string                `json:"total_quantity"`
	Customers     []OmzetValueCustomerVM `json:"customers"`
}

type OmzetValueCustomerVM struct {
	RegionGroupName *string `json:"region_group_name"`
	RegionName      *string `json:"region_name"`
	BranchName      *string `json:"branch_name"`
	BranchCode      *string `json:"branch_code"`
	CustomerID      *string `json:"customer_id"`
	CustomerCode    *string `json:"customer_code"`
	CustomerName    *string `json:"customer_name"`
	CustomerType    *string `json:"customer_type"`
	ProvinceName    *string `json:"customer_province_name"`
	CityName        *string `json:"customer_city_name"`
	CustomerLevel   *string `json:"customer_level"`
	Quantity        *string `json:"quantity"`
	Omzet           *string `json:"omzet"`
}

type OmzetValueByItemVM struct {
	TotalOmzet    *string            `json:"total_omzet"`
	TotalQuantity *string            `json:"total_quantity"`
	Customers     []OmzetValueItemVM `json:"customers"`
}

type OmzetValueItemVM struct {
	ItemID   *string `json:"item_id"`
	ItemName *string `json:"item_name"`
	Quantity *string `json:"quantity"`
	Omzet    *string `json:"omzet"`
}

type DashboardTrackingInvoiceVM struct {
	RegionGroupName             string `json:"region_group_name"`
	RegionName                  string `json:"region_name"`
	BranchName                  string `json:"branch_name"`
	BranchArea                  string `json:"branch_area"`
	BranchCode                  string `json:"branch_code"`
	CustomerName                string `json:"customer_name"`
	CustomerCode                string `json:"customer_code"`
	CustomerLevelName           string `json:"customer_level_name"`
	SalesOrderDocumentNumber    string `json:"sales_order_document_number"`
	CustomerOrderDocumentNumber string `json:"customer_order_document_number"`
	InvoiceID                   string `json:"invoice_id"`
	InvoiceNumber               string `json:"invoice_number"`
	InvoiceDate                 string `json:"invoice_date"`
	AcceptedDate                string `json:"accepted_date"`
	ProcessedDate               string `json:"processed_date"`
	ConfimationDate             string `json:"confirmation_date"`
	DueDate                     string `json:"due_date"`
	PaidOffDate                 string `json:"paid_off_date"`
	SourceTransaction           string `json:"source_transaction"`
}
