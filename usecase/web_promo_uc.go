package usecase

import (
	"context"
	"mime/multipart"

	"nextbasis-service-v-0.1/db/repository"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/functioncaller"
	"nextbasis-service-v-0.1/pkg/logruslogger"
	"nextbasis-service-v-0.1/server/requests"
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

// WebPromoUC ...
type WebPromoUC struct {
	*ContractUC
}

// BuildBody ...
func (uc WebPromoUC) BuildBody(res *models.WebPromo) {
}

// SelectAll ...
func (uc WebPromoUC) SelectAll(c context.Context, parameter models.WebPromoParameter) (res []models.WebPromo, err error) {
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(0, 0, parameter.By, parameter.Sort, models.WebPromoOrderBy, models.WebPromoOrderByrByString)

	repo := repository.NewWebPromoRepository(uc.DB)
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
func (uc WebPromoUC) FindAll(c context.Context, parameter models.WebPromoParameter) (res []models.WebPromo, p viewmodel.PaginationVM, err error) {
	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.WebPromoOrderBy, models.WebPromoOrderByrByString)

	var count int
	repo := repository.NewWebPromoRepository(uc.DB)
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

// Add ...
func (uc WebPromoUC) Add(c context.Context, data *requests.WebPromoRequest, imgBanner *multipart.FileHeader) (res models.WebPromo, err error) {

	ctx := "FileUC.Upload"
	awsUc := AwsUC{ContractUC: uc.ContractUC}
	var strImgBanner = ""
	if imgBanner != nil {
		imgProfileFile, err := awsUc.Upload("image/promo", imgBanner)
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "upload_file", c.Value("requestid"))
			return res, err
		}
		strImgBanner = imgProfileFile.FileName
	}

	repo := repository.NewWebPromoRepository(uc.DB)
	// now := time.Now().UTC()
	// strNow := now.Format(time.RFC3339)
	res = models.WebPromo{

		Code:             &data.Code,
		PromoName:        &data.PromoName,
		PromoDescription: &data.PromoDescription,
		PromoUrlBanner:   &strImgBanner,
		StartDate:        &data.StartDate,
		EndDate:          &data.EndDate,
		ShowInApp:        &data.ShowInApp,
	}
	res.ID, err = repo.Add(c, &res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return res, err
}

// Delete ...
func (uc WebPromoUC) Delete(c context.Context, id string) (res viewmodel.CommonDeletedObjectVM, err error) {
	repo := repository.NewWebPromoRepository(uc.DB)
	res.ID, err = repo.Delete(c, id)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return res, err

}
