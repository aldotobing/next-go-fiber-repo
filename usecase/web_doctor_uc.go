package usecase

import (
	"context"
	"mime/multipart"
	"strings"

	"nextbasis-service-v-0.1/db/repository"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/functioncaller"
	"nextbasis-service-v-0.1/pkg/logruslogger"
	"nextbasis-service-v-0.1/server/requests"
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

// WebDoctorUC ...
type WebDoctorUC struct {
	*ContractUC
}

// BuildBody ...
func (uc WebDoctorUC) BuildBody(res *models.WebDoctor) {
}

// SelectAll ...
func (uc WebDoctorUC) SelectAll(c context.Context, parameter models.WebDoctorParameter) (res []models.WebDoctor, err error) {
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(0, 0, parameter.By, parameter.Sort, models.WebDoctorOrderBy, models.WebDoctorOrderByrByString)

	repo := repository.NewWebDoctorRepository(uc.DB)
	res, err = repo.SelectAll(c, parameter)

	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	for i := range res {
		uc.BuildBody(&res[i])
	}

	return res, err
}

// FindAll ...
func (uc WebDoctorUC) FindAll(c context.Context, parameter models.WebDoctorParameter) (res []models.WebDoctor, p viewmodel.PaginationVM, err error) {
	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.WebDoctorOrderBy, models.WebDoctorOrderByrByString)

	var count int
	repo := repository.NewWebDoctorRepository(uc.DB)
	res, count, err = repo.FindAll(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, p, err
	}

	p = uc.setPaginationResponse(parameter.Page, parameter.Limit, count)
	for i := range res {
		uc.BuildBody(&res[i])
	}

	return res, p, err
}

// FindByID ...
func (uc WebDoctorUC) FindByID(c context.Context, parameter models.WebDoctorParameter) (res models.WebDoctor, err error) {

	repo := repository.NewWebDoctorRepository(uc.DB)
	res, err = repo.FindByID(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}
	uc.BuildBody(&res)

	return res, err
}

func (uc WebDoctorUC) Edit(c context.Context, id string, data *requests.WebDoctorRequest, imgProfile *multipart.FileHeader) (res models.WebDoctor, err error) {

	// currentObjectUc, err := uc.FindByID(c, models.MpBankParameter{ID: id})
	currentObjectUc, err := uc.FindByID(c, models.WebDoctorParameter{ID: id})
	ctx := "FileUC.Upload"
	awsUc := AwsUC{ContractUC: uc.ContractUC}

	var strImgprofile = ""

	if currentObjectUc.DoctorProfilePicture != nil && *currentObjectUc.DoctorProfilePicture != "" {
		strImgprofile = strings.ReplaceAll(*currentObjectUc.DoctorProfilePicture, models.CustomerImagePath, "")
	}

	if imgProfile != nil {
		if &strImgprofile != nil && strings.Trim(strImgprofile, " ") != "" {
			_, err = awsUc.Delete("image/doctor", strImgprofile)
			if err != nil {
				logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "s3", uc.ReqID)
			}
		}

		awsUc.AWSS3.Directory = "image/customer"
		imgBannerFile, err := awsUc.Upload("image/customer", imgProfile)
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "upload_file", c.Value("requestid"))
			return res, err
		}
		strImgprofile = imgBannerFile.FilePath

	}
	repo := repository.NewWebDoctorRepository(uc.DB)
	// now := time.Now().UTC()
	// strnow := now.Format(time.RFC3339)
	res = models.WebDoctor{
		ID:                   &id,
		Code:                 &data.Code,
		DoctorName:           &data.DoctorName,
		DoctorAddress:        &data.DoctorAddress,
		DoctorPhone:          &data.DoctorPhone,
		DoctorUserID:         &data.DoctorUserID,
		DoctorProfilePicture: &strImgprofile,
	}

	res.ID, err = repo.Edit(c, &res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return res, err
}

func (uc WebDoctorUC) Add(c context.Context, data *requests.WebDoctorRequest, imgProfile *multipart.FileHeader) (res models.WebDoctor, err error) {

	// currentObjectUc, err := uc.FindByID(c, models.MpBankParameter{ID: id})
	ctx := "FileUC.Upload"
	awsUc := AwsUC{ContractUC: uc.ContractUC}

	var strImgprofile = ""

	if imgProfile != nil {

		awsUc.AWSS3.Directory = "image/customer"
		imgBannerFile, err := awsUc.Upload("image/customer", imgProfile)
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "upload_file", c.Value("requestid"))
			return res, err
		}
		strImgprofile = imgBannerFile.FilePath

	}
	repo := repository.NewWebDoctorRepository(uc.DB)
	// now := time.Now().UTC()
	// strnow := now.Format(time.RFC3339)
	res = models.WebDoctor{
		Code:                 &data.Code,
		DoctorName:           &data.DoctorName,
		DoctorAddress:        &data.DoctorAddress,
		DoctorPhone:          &data.DoctorPhone,
		DoctorUserID:         &data.DoctorUserID,
		DoctorProfilePicture: &strImgprofile,
	}

	res.ID, err = repo.Add(c, &res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return res, err
}
