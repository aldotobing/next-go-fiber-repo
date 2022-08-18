package usecase

import (
	"nextbasis-service-v-0.1/pkg/file"
	"nextbasis-service-v-0.1/pkg/logruslogger"
)

// FileDownloadUC ...
type FileDownloadUC struct {
	*ContractUC
}

// Download download and save file to random filename
func (uc FileDownloadUC) Download(url string) (res string, err error) {
	ctx := "FileDownloadUC.Download"

	uploadPath := uc.ContractUC.EnvConfig["FILE_STATIC_FILE"]
	res, err = file.Download(url, uploadPath)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "download", uc.ReqID)
		return res, err
	}

	return res, err
}

// DownloadImage download and save file to parameter filename
func (uc FileDownloadUC) DownloadImage(url, name string) (res, contentType string, err error) {
	ctx := "FileDownloadUC.DownloadImage"

	uploadPath := uc.ContractUC.EnvConfig["FILE_STATIC_FILE"]
	res, contentType, err = file.DownloadImage(url, uploadPath, name)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "download", uc.ReqID)
		return res, contentType, err
	}

	return res, contentType, err
}
