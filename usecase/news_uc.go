package usecase

import (
	"context"
	"mime/multipart"

	"nextbasis-service-v-0.1/config"
	"nextbasis-service-v-0.1/db/repository"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/functioncaller"
	"nextbasis-service-v-0.1/pkg/logruslogger"
	"nextbasis-service-v-0.1/server/requests"
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

// NewsUC ...
type NewsUC struct {
	*ContractUC
}

// BuildBody ...
func (uc NewsUC) BuildBody(res *models.News) {
}

// SelectAll ...
func (uc NewsUC) SelectAll(c context.Context, parameter models.NewsParameter) (res []models.News, err error) {
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(0, 0, parameter.By, parameter.Sort, models.NewsOrderBy, models.NewsOrderByrByString)

	repo := repository.NewNewsRepository(uc.DB)
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
func (uc NewsUC) FindAll(c context.Context, parameter models.NewsParameter) (res []models.News, p viewmodel.PaginationVM, err error) {
	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.NewsOrderBy, models.NewsOrderByrByString)

	var count int
	repo := repository.NewNewsRepository(uc.DB)
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
func (uc NewsUC) Add(c context.Context, data *requests.NewsRequest) (res models.News, err error) {

	repo := repository.NewNewsRepository(uc.DB)
	// now := time.Now().UTC()
	// strnow := now.Format(time.RFC3339)
	res = models.News{
		Title:       &data.Title,
		Description: &data.Description,
		StartDate:   &data.StartDate,
		EndDate:     &data.EndDate,
		Active:      &data.Active,
		ImageUrl:    &data.ImageUrl,
	}
	res.ID, err = repo.Add(c, &res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return res, err
}

// AddBulk ...
func (uc NewsUC) AddBulk(c context.Context, data requests.NewsBulkRequest) (err error) {

	repo := repository.NewNewsRepository(uc.DB)
	// now := time.Now().UTC()
	// strnow := now.Format(time.RFC3339)
	var in []models.News
	for i := range data.News {
		in = append(in, models.News{
			Title:       &data.News[i].Title,
			Description: &data.News[i].Description,
			StartDate:   &data.News[i].StartDate,
			EndDate:     &data.News[i].EndDate,
			Active:      &data.News[i].Active,
			ImageUrl:    &data.News[i].ImageUrl,
		})
	}

	err = repo.AddBulk(c, in)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	return
}

// AddPhoto ...
func (uc NewsUC) AddPhoto(c context.Context, image *multipart.FileHeader) (out string, err error) {
	awsUc := AwsUC{ContractUC: uc.ContractUC}
	awsUc.AWSS3.Directory = "image/news"
	imgBannerFile, err := awsUc.Upload("image/news", image)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "upload_file", c.Value("requestid"))
		return
	}
	out = config.ImagePath + imgBannerFile.FilePath

	return
}

// Delete ...
func (uc NewsUC) Delete(c context.Context, id string) (res viewmodel.CommonDeletedObjectVM, err error) {
	repo := repository.NewNewsRepository(uc.DB)
	res.ID, err = repo.Delete(c, id)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return res, err

}

func (uc NewsUC) Edit(c context.Context, id string, data *requests.NewsRequest) (res models.News, err error) {
	repo := repository.NewNewsRepository(uc.DB)
	// now := time.Now().UTC()
	// strnow := now.Format(time.RFC3339)
	res = models.News{
		ID:          &id,
		Title:       &data.Title,
		Description: &data.Description,
		StartDate:   &data.StartDate,
		EndDate:     &data.EndDate,
		Active:      &data.Active,
		ImageUrl:    &data.ImageUrl,
	}

	res.ID, err = repo.Edit(c, &res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return res, err
}
