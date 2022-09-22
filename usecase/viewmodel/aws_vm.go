package viewmodel

//AwsVM ....
type AwsVM struct {
	FileName string `json:"file_name"`
	FilePath string `json:"file_path"`
	FileSize int64  `json:"size"`
	FileType string `json:"type"`
}
