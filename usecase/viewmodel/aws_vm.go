package viewmodel

//AwsVM ....
type AwsVM struct {
	Name     string `json:"name"`
	FileName string `json:"file_name"`
	SPath    string `json:"s_path"`
	Size     int64  `json:"size"`
	Type     string `json:"type"`
}
