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
	s3path, fileName, err := uc.AWSS3.UploadManager(f)
	//data, err := AwsModel.Upload(defaultBucket, filePath, f)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "upload", uc.ReqID)
		return res, err
	}
	// res.Name = data
	res.Size = f.Size
	res.SPath = s3path
	res.FileName = fileName

	return res, err
}

// // UploadOs upload file to min.io server
// func (uc AwsUC) UploadOs(filePath, fileName, fileLocalPath, contentType string) (res string, err error) {
// 	ctx := "AwsUC.UploadOs"

// 	defaultBucket := uc.ContractUC.EnvConfig["Aws_DEFAULT_BUCKET"]
// 	AwsModel := Aws.NewAwsModel(uc.Aws)
// 	res, err = AwsModel.UploadOs(defaultBucket, filePath, fileName, fileLocalPath, contentType)
// 	if err != nil {
// 		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "upload", uc.ReqID)
// 		return res, err
// 	}

// 	err = os.Remove(fileLocalPath)
// 	if err != nil {
// 		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "delete_file", uc.ReqID)
// 		return res, err
// 	}

// 	return res, err
// }

// // GetURL get file url by object name
// func (uc AwsUC) GetURL(objectName string) (res string, err error) {
// 	ctx := "AwsUC.GetURL"

// 	defaultBucket := uc.ContractUC.EnvConfig["Aws_DEFAULT_BUCKET"]
// 	AwsModel := Aws.NewAwsModel(uc.Aws)
// 	if objectName == "" {
// 		logruslogger.Log(logruslogger.WarnLevel, "", ctx, "empty_parameter", uc.ReqID)
// 		return res, err
// 	}

// 	res, err = AwsModel.GetFile(defaultBucket, objectName)
// 	if err != nil {
// 		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "get_file_url", uc.ReqID)
// 		return res, err
// 	}

// 	res = strings.Replace(res, "http://"+uc.ContractUC.EnvConfig["Aws_ENDPOINT"], uc.ContractUC.EnvConfig["Aws_BASE_URL"], 1)

// 	return res, err
// }

// // GetURLNoErr get file url by object name wo err response
// func (uc AwsUC) GetURLNoErr(objectName string) (res string) {
// 	err := uc.RedisClient.GetFromRedis("Aws"+objectName, &res)
// 	if err == nil {
// 		return res
// 	}

// 	res, _ = uc.GetURL(objectName)

// 	uc.StoreToRedisExp("Aws"+objectName, res, "15m")

// 	return res
// }

// // Delete delete object
// func (uc AwsUC) Delete(objectName string) (err error) {
// 	ctx := "AwsUC.Delete"

// 	defaultBucket := uc.ContractUC.EnvConfig["Aws_DEFAULT_BUCKET"]
// 	AwsModel := Aws.NewAwsModel(uc.Aws)
// 	err = AwsModel.Delete(defaultBucket, objectName)
// 	if err != nil {
// 		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "delete", uc.ReqID)
// 		return err
// 	}

// 	return err
// }
