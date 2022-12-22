package usecase

import (
	"mime/multipart"

	"nextbasis-service-v-0.1/pkg/logruslogger"
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

// AwsUC ...
type AwsUC struct {
	*ContractUC
}

// Upload upload file to aws server
func (uc AwsUC) Upload(filePath string, f *multipart.FileHeader) (res viewmodel.AwsVM, err error) {
	ctx := "AwsUC.Upload"
	// uc.ContractUC.EnvConfig["AWS_BUCKET_NAME"]
	uc.AWSS3.Directory = filePath
	s3path, fileName, err := uc.AWSS3.UploadManager(f)
	//data, err := AwsModel.Upload(defaultBucket, filePath, f)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "upload", uc.ReqID)
		return res, err
	}

	res.FileSize = f.Size
	res.FilePath = s3path
	res.FileName = fileName

	return res, err
}

func (uc AwsUC) Delete(filePath, file_name string) (res viewmodel.AwsVM, err error) {
	ctx := "AwsUC.Delete"

	FileToRemove := filePath + "/" + file_name
	s3path, fileName, err := uc.AWSS3.DeleteManager(FileToRemove)

	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "upload", uc.ReqID)
		return res, err
	}

	res.FilePath = s3path
	res.FileName = fileName

	return res, err
}
