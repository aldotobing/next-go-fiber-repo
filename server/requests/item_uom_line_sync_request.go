package requests

// ItemUomLineSyncRequest ...
type ItemUomLineSyncRequest struct {
	ItemCode      string `json:"item_code"`
	UomCode       string `json:"uom_code"`
	UomConversion string `json:"uom_conversion"`
}
