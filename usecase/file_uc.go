package usecase

import (
	"context"
	"mime/multipart"
	"time"

	"nextbasis-service-v-0.1/db/repository"
	"nextbasis-service-v-0.1/pkg/logruslogger"
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

// FileUC ...
type FileUC struct {
	*ContractUC
}

// FindByID ...
func (uc FileUC) FindByID(id string) (res viewmodel.FileVM, err error) {
	ctx := "FileUC.FindByID"

	repo := repository.NewFileRepository(uc.DB)
	data, err := repo.FindByID(id)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "find_id", uc.ReqID)
		return res, err
	}

	minioUc := MinioUC{ContractUC: uc.ContractUC}
	tempURL, err := minioUc.GetURL(data.URL.String)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "get_url", uc.ReqID)
		return res, err
	}

	res = viewmodel.FileVM{
		ID:         data.ID,
		Type:       data.Type.String,
		URL:        data.URL.String,
		TempURL:    tempURL,
		UserUpload: data.UserUpload.String,
		CreatedAt:  data.CreatedAt,
		UpdatedAt:  data.UpdatedAt,
		DeletedAt:  data.DeletedAt.String,
	}

	return res, err
}

// FindOneUnassigned check if image is unassigned
func (uc FileUC) FindOneUnassigned(id, types, userUpload string) (res viewmodel.FileVM, err error) {
	ctx := "FileUC.FindOneUnassigned"

	repo := repository.NewFileRepository(uc.DB)
	data, err := repo.FindUnassignedByID(id, types, userUpload)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "find_id", uc.ReqID)
		return res, err
	}

	minioUc := MinioUC{ContractUC: uc.ContractUC}
	tempURL, err := minioUc.GetURL(data.URL.String)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "get_url", uc.ReqID)
		return res, err
	}

	res = viewmodel.FileVM{
		ID:         data.ID,
		Type:       data.Type.String,
		URL:        data.URL.String,
		TempURL:    tempURL,
		UserUpload: data.UserUpload.String,
		CreatedAt:  data.CreatedAt,
		UpdatedAt:  data.UpdatedAt,
		DeletedAt:  data.DeletedAt.String,
	}

	return res, err
}

// FindByType check if image is exist
func (uc FileUC) FindByType(id, types string) (res viewmodel.FileVM, err error) {
	ctx := "FileUC.FindByType"

	repo := repository.NewFileRepository(uc.DB)
	data, err := repo.FindByType(id, types)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "find_id", uc.ReqID)
		return res, err
	}

	res = viewmodel.FileVM{
		ID:         data.ID,
		Type:       data.Type.String,
		URL:        data.URL.String,
		UserUpload: data.UserUpload.String,
		CreatedAt:  data.CreatedAt,
		UpdatedAt:  data.UpdatedAt,
		DeletedAt:  data.DeletedAt.String,
	}

	return res, err
}

// Create ...
func (uc FileUC) Create(types, url, userUpload string) (res viewmodel.FileVM, err error) {
	ctx := "FileUC.Create"

	// Delete all unused files first
	// if !str.Contains(models.FileMultipleUploadWhitelist, types) {
	// 	err = uc.DeleteAllUnused(userUpload, types)
	// 	if err != nil {
	// 		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "delete_unused", uc.ReqID)
	// 		return res, err
	// 	}
	// }

	now := time.Now().UTC()
	res = viewmodel.FileVM{
		Type:       types,
		URL:        url,
		UserUpload: userUpload,
		CreatedAt:  now.Format(time.RFC3339),
		UpdatedAt:  now.Format(time.RFC3339),
	}
	// repo := repository.NewFileRepository(uc.DB)
	// res.ID, err = repo.Store(&res, now)
	// if err != nil {
	// 	logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "create", uc.ReqID)
	// 	return res, err
	// }

	// Get temp url
	minioUc := MinioUC{ContractUC: uc.ContractUC}
	res.TempURL, err = minioUc.GetURL(res.URL)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "get_url", uc.ReqID)
		return res, err
	}

	return res, err
}

// UpdateIsUsed ...
func (uc FileUC) UpdateIsUsed(id string, isUsed bool) (res viewmodel.FileVM, err error) {
	ctx := "FileUC.UpdateIsUsed"

	now := time.Now().UTC()
	repo := repository.NewFileRepository(uc.DB)
	res.ID, err = repo.UpdateIsUsed(id, isUsed, now)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "delete", uc.ReqID)
		return res, err
	}

	return res, err
}

// Delete ...
func (uc FileUC) Delete(id string) (res viewmodel.FileVM, err error) {
	ctx := "FileUC.Delete"

	now := time.Now().UTC()
	repo := repository.NewFileRepository(uc.DB)
	res.ID, err = repo.Destroy(id, now)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "delete", uc.ReqID)
		return res, err
	}

	return res, err
}

// DeleteAllUnused ...
func (uc FileUC) DeleteAllUnused(userUpload, types string) (err error) {
	ctx := "FileUC.DeleteAllUnused"
	repo := repository.NewFileRepository(uc.DB)
	unusedFile, err := repo.FindAllUnassignedByUserID(userUpload, types)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "find_all_unused", uc.ReqID)
		return err
	}

	minioUc := MinioUC{ContractUC: uc.ContractUC}
	now := time.Now().UTC()
	for _, r := range unusedFile {
		err = minioUc.Delete(r.URL.String)
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "s3", uc.ReqID)
		}

		_, err = repo.Destroy(r.ID, now)
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "user_file", uc.ReqID)
		}
	}
	err = nil

	return err
}

// Upload ...
func (uc FileUC) Upload(c context.Context, types string, file *multipart.FileHeader) (res viewmodel.FileVM, err error) {
	ctx := "FileUC.Upload"

	minioUc := MinioUC{ContractUC: uc.ContractUC}
	fileKey, err := minioUc.Upload(types+"/"+c.Value("user_id").(string), file)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "upload_file", c.Value("requestid"))
		return res, err
	}

	res, err = uc.Create(types, fileKey.Name, c.Value("user_id").(string))
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, "", ctx, "create", c.Value("requestid"))
		return res, err
	}

	return res, err
}

func (uc FileUC) UploadTes(c context.Context, types string, file *multipart.FileHeader) (res viewmodel.FileVM, err error) {
	ctx := "FileUC.Upload"

	minioUc := MinioUC{ContractUC: uc.ContractUC}
	fileKey, err := minioUc.Upload(types+"/"+c.Value("user_id").(string), file)
	if err != nil {

		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "upload_file", c.Value("requestid"))
		return res, err
	}

	res, err = uc.Create(types, fileKey.Name, c.Value("user_id").(string))
	if err != nil {

		logruslogger.Log(logruslogger.WarnLevel, "", ctx, "create", c.Value("requestid"))
		return res, err
	}

	return res, err
}
