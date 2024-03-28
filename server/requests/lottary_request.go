package requests

// LottaryRequest ...
type LottaryRequest struct {
	CustomerCode string `json:"customer_code"`
	Jumlah       int    `json:"jumlah"`
}

type LottaryRequestHeader struct {
	Year    string           `json:"_year"`
	Quartal string           `json:"_quartal"`
	Detail  []LottaryRequest `json:"detail"`
}
