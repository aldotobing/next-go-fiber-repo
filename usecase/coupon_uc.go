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

// CouponUC ...
type CouponUC struct {
	*ContractUC
}

// BuildBody ...
func (uc CouponUC) BuildBody(data *models.Coupon, res *viewmodel.CouponVM) {
	res.ID = data.ID
	res.StartDate = data.StartDate
	res.EndDate = data.EndDate
	res.PointConversion = data.PointConversion
	res.Name = data.Name
	res.Description = data.Description
	res.CreatedAt = data.CreatedAt
	res.UpdatedAt = data.UpdatedAt.String
	res.DeletedAt = data.DeletedAt.String
}

// FindAll ...
func (uc CouponUC) FindAll(c context.Context, parameter models.CouponParameter) (out []viewmodel.CouponVM, p viewmodel.PaginationVM, err error) {
	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.PointRuleOrderBy, models.PointRuleOrderByrByString)

	repo := repository.NewCouponRepository(uc.DB)
	data, count, err := repo.FindAll(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	p = uc.setPaginationResponse(parameter.Page, parameter.Limit, count)
	for _, datum := range data {
		var temp viewmodel.CouponVM
		uc.BuildBody(&datum, &temp)

		out = append(out, temp)
	}

	if out == nil {
		out = make([]viewmodel.CouponVM, 0)
	}

	return
}

// SelectAll ...
func (uc CouponUC) SelectAll(c context.Context, parameter models.CouponParameter) (out []viewmodel.CouponVM, err error) {
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.PointRuleOrderBy, models.PointRuleOrderByrByString)

	repo := repository.NewCouponRepository(uc.DB)
	data, err := repo.SelectAll(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	for _, datum := range data {
		var temp viewmodel.CouponVM
		uc.BuildBody(&datum, &temp)

		out = append(out, temp)
	}

	if out == nil {
		out = make([]viewmodel.CouponVM, 0)
	}

	return
}

// FindByID ...
func (uc CouponUC) FindByID(c context.Context, parameter models.CouponParameter) (out viewmodel.CouponVM, err error) {
	repo := repository.NewCouponRepository(uc.DB)
	data, err := repo.FindByID(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	uc.BuildBody(&data, &out)

	return
}

// Add ...
func (uc CouponUC) Add(c context.Context, in requests.CouponRequest) (out viewmodel.CouponVM, err error) {
	out = viewmodel.CouponVM{
		StartDate:       in.StartDate,
		EndDate:         in.EndDate,
		PointConversion: in.PointConversion,
		Name:            in.Name,
		Description:     in.Description,
	}

	repo := repository.NewCouponRepository(uc.DB)
	out.ID, err = repo.Add(c, out)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	return
}

// AddPhoto ...
func (uc CouponUC) AddPhoto(c context.Context, image *multipart.FileHeader) (out string, err error) {
	awsUc := AwsUC{ContractUC: uc.ContractUC}
	awsUc.AWSS3.Directory = "image/voucher"
	imgBannerFile, err := awsUc.Upload("image/voucher", image)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "upload_file", c.Value("requestid"))
		return
	}
	out = config.ImagePath + imgBannerFile.FilePath

	return
}

// Update ...
func (uc CouponUC) Update(c context.Context, id string, in requests.CouponRequest) (out viewmodel.CouponVM, err error) {
	out = viewmodel.CouponVM{
		ID:              id,
		StartDate:       in.StartDate,
		EndDate:         in.EndDate,
		PointConversion: in.PointConversion,
		Name:            in.Name,
		Description:     in.Description,
	}

	repo := repository.NewCouponRepository(uc.DB)
	out.ID, err = repo.Update(c, out)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	return
}

// Delete ...
func (uc CouponUC) Delete(c context.Context, in string) (err error) {
	repo := repository.NewCouponRepository(uc.DB)
	_, err = repo.Delete(c, in)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	return
}
