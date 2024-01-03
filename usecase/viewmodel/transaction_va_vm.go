package viewmodel

// VaBilInfo Vm ....
type VaBillInfoVM struct {
	BillInfo1     string          `json:"billInfo1"`
	BillInfo2     string          `json:"billInfo2"`
	BillInfo3     string          `json:"billInfo3"`
	Currency      string          `json:"currency"`
	VabillDetails VaBillDetailsVM `json:"billDetails"`
	Status        VaBillStatus    `json:"status"`
}

type PaymentVaBillInfoVM struct {
	BillInfo1 string       `json:"billInfo1"`
	BillInfo2 string       `json:"billInfo2"`
	Status    VaBillStatus `json:"status"`
}

type VaBillDetailVM struct {
	BillCode      string  `json:"billCode"`
	BillName      string  `json:"billName"`
	BillShortName string  `json:"billShortName"`
	BillAmount    string  `json:"billAmount"`
	Reference1    *string `json:"reference1"`
	Reference2    *string `json:"reference2"`
	Reference3    *string `json:"reference3"`
}

type VaBillDetailsVM struct {
	BillDetail []VaBillDetailVM `json:"BillDetail"`
}

type VaBillStatus struct {
	IsError           string `json:"isError"`
	ErrorCode         string `json:"errorCode"`
	StatusDescription string `json:"statusDescription"`
}
