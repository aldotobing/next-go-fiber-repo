package usecase

import (
	"mime/multipart"
	"os"
	"strings"

	"nextbasis-service-v-0.1/pkg/logruslogger"
	"nextbasis-service-v-0.1/pkg/minio"
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

// MinioUC ...
type MinioUC struct {
	*ContractUC
}

// Upload upload file to min.io server
func (uc MinioUC) Upload(filePath string, f *multipart.FileHeader) (res viewmodel.MinioVM, err error) {
	ctx := "MinioUC.Upload"

	defaultBucket := uc.ContractUC.EnvConfig["MINIO_DEFAULT_BUCKET"]
	minioModel := minio.NewMinioModel(uc.Minio)
	data, err := minioModel.Upload(defaultBucket, filePath, f)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "upload", uc.ReqID)
		return res, err
	}
	res.Name = data
	res.Size = f.Size

	return res, err
}

// UploadOs upload file to min.io server
func (uc MinioUC) UploadOs(filePath, fileName, fileLocalPath, contentType string) (res string, err error) {
	ctx := "MinioUC.UploadOs"

	defaultBucket := uc.ContractUC.EnvConfig["MINIO_DEFAULT_BUCKET"]
	minioModel := minio.NewMinioModel(uc.Minio)
	res, err = minioModel.UploadOs(defaultBucket, filePath, fileName, fileLocalPath, contentType)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "upload", uc.ReqID)
		return res, err
	}

	err = os.Remove(fileLocalPath)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "delete_file", uc.ReqID)
		return res, err
	}

	return res, err
}

// GetURL get file url by object name
func (uc MinioUC) GetURL(objectName string) (res string, err error) {
	ctx := "MinioUC.GetURL"

	defaultBucket := uc.ContractUC.EnvConfig["MINIO_DEFAULT_BUCKET"]
	minioModel := minio.NewMinioModel(uc.Minio)
	if objectName == "" {
		logruslogger.Log(logruslogger.WarnLevel, "", ctx, "empty_parameter", uc.ReqID)
		return res, err
	}

	res, err = minioModel.GetFile(defaultBucket, objectName)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "get_file_url", uc.ReqID)
		return res, err
	}

	res = strings.Replace(res, "http://"+uc.ContractUC.EnvConfig["MINIO_ENDPOINT"], uc.ContractUC.EnvConfig["MINIO_BASE_URL"], 1)

	return res, err
}

// GetURLNoErr get file url by object name wo err response
func (uc MinioUC) GetURLNoErr(objectName string) (res string) {
	err := uc.RedisClient.GetFromRedis("minio"+objectName, &res)
	if err == nil {
		return res
	}

	res, _ = uc.GetURL(objectName)

	uc.StoreToRedisExp("minio"+objectName, res, "15m")

	return res
}

// Delete delete object
func (uc MinioUC) Delete(objectName string) (err error) {
	ctx := "MinioUC.Delete"

	defaultBucket := uc.ContractUC.EnvConfig["MINIO_DEFAULT_BUCKET"]
	minioModel := minio.NewMinioModel(uc.Minio)
	err = minioModel.Delete(defaultBucket, objectName)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "delete", uc.ReqID)
		return err
	}

	return err
}
