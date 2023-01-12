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

// WebItemUC ...
type WebItemUC struct {
	*ContractUC
}

// BuildBody ...
func (uc WebItemUC) BuildBody(res *models.WebItem) {
}

// SelectAll ...
func (uc WebItemUC) SelectAll(c context.Context, parameter models.WebItemParameter) (res []models.WebItem, err error) {
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(0, 0, parameter.By, parameter.Sort, models.ItemOrderBy, models.ItemOrderByrByString)

	repo := repository.NewWebItemRepository(uc.DB)
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
func (uc WebItemUC) FindAll(c context.Context, parameter models.WebItemParameter) (res []models.WebItem, p viewmodel.PaginationVM, err error) {
	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.ItemOrderBy, models.ItemOrderByrByString)

	var count int
	repo := repository.NewWebItemRepository(uc.DB)
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
func (uc WebItemUC) FindByID(c context.Context, parameter models.WebItemParameter) (res models.WebItem, err error) {
	repo := repository.NewWebItemRepository(uc.DB)
	res, err = repo.FindByID(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}
	uc.BuildBody(&res)

	return res, err
}

// Edit ...
func (uc WebItemUC) Edit(c context.Context, id string, data *requests.WebItemRequest, itemImage *multipart.FileHeader) (res models.WebItem, err error) {

	// currentObjectUc, err := uc.FindByID(c, models.MpBankParameter{ID: id})
	currentObjectUc, err := uc.FindByID(c, models.WebItemParameter{ID: id})
	ctx := "FileUC.Upload"
	awsUc := AwsUC{ContractUC: uc.ContractUC}

	var strImg = ""

	if currentObjectUc.ItemPicture != nil && *currentObjectUc.ItemPicture != "" {
		strImg = strings.ReplaceAll(*currentObjectUc.ItemPicture, models.ItemImagePath, "")
	}

	if itemImage != nil {
		if &strImg != nil && strings.Trim(strImg, " ") != "" {
			_, err = awsUc.Delete("image/item", strImg)
			if err != nil {
				logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "s3", uc.ReqID)
			}
		}

		awsUc.AWSS3.Directory = "image/item"
		imgFile, err := awsUc.Upload("image/item", itemImage)
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "upload_file", c.Value("requestid"))
			return res, err
		}
		strImg = imgFile.FileName

	}
	repo := repository.NewWebItemRepository(uc.DB)
	// now := time.Now().UTC()
	// strnow := now.Format(time.RFC3339)
	res = models.WebItem{
		ID:          &id,
		Code:        &data.Code,
		Name:        &data.Name,
		ItemPicture: &strImg,
	}

	res.ID, err = repo.Edit(c, &res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return res, err
}
