package usecase

import (
	"context"
	"encoding/json"
	"mime/multipart"

	"nextbasis-service-v-0.1/config"
	"nextbasis-service-v-0.1/db/repository"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/functioncaller"
	"nextbasis-service-v-0.1/pkg/logruslogger"
	"nextbasis-service-v-0.1/server/requests"
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

// PointRuleUC ...
type PointRuleUC struct {
	*ContractUC
}

// BuildBody ...
func (uc PointRuleUC) BuildBody(data *models.PointRule, res *viewmodel.PointRuleVM) {
	res.ID = data.ID
	res.StartDate = data.StartDate
	res.EndDate = data.EndDate
	res.MinOrder = data.MinOrder
	res.PointConversion = data.PointConversion
	res.MonthlyMaxPoint = data.MonthlyMaxPoint
	res.CreatedAt = data.CreatedAt
	res.UpdatedAt = data.UpdatedAt.String
	res.DeletedAt = data.DeletedAt.String

	if data.Customer.Valid {
		json.Unmarshal([]byte(data.Customer.String), &res.Customers)
	}
}

// FindAll ...
func (uc PointRuleUC) FindAll(c context.Context, parameter models.PointRuleParameter) (out []viewmodel.PointRuleVM, p viewmodel.PaginationVM, err error) {
	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.PointRuleOrderBy, models.PointRuleOrderByrByString)

	repo := repository.NewPointRuleRepository(uc.DB)
	data, count, err := repo.FindAll(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	p = uc.setPaginationResponse(parameter.Page, parameter.Limit, count)
	for _, datum := range data {
		var temp viewmodel.PointRuleVM
		uc.BuildBody(&datum, &temp)

		out = append(out, temp)
	}

	if out == nil {
		out = make([]viewmodel.PointRuleVM, 0)
	}

	return
}

// SelectAll ...
func (uc PointRuleUC) SelectAll(c context.Context, parameter models.PointRuleParameter) (out []viewmodel.PointRuleVM, err error) {
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.PointRuleOrderBy, models.PointRuleOrderByrByString)

	repo := repository.NewPointRuleRepository(uc.DB)
	data, err := repo.SelectAll(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	for _, datum := range data {
		var temp viewmodel.PointRuleVM
		uc.BuildBody(&datum, &temp)

		out = append(out, temp)
	}

	if out == nil {
		out = make([]viewmodel.PointRuleVM, 0)
	}

	return
}

// FindByID ...
func (uc PointRuleUC) FindByID(c context.Context, parameter models.PointRuleParameter) (out viewmodel.PointRuleVM, err error) {
	repo := repository.NewPointRuleRepository(uc.DB)
	data, err := repo.FindByID(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	uc.BuildBody(&data, &out)

	return
}

// Add ...
func (uc PointRuleUC) Add(c context.Context, in requests.PointRuleRequest) (out viewmodel.PointRuleVM, err error) {
	var customers []viewmodel.PointRuleCustomerVM
	for _, datum := range in.Customers {
		customers = append(customers, viewmodel.PointRuleCustomerVM{
			CustomerCode: datum.CustomerCode,
		})
	}
	out = viewmodel.PointRuleVM{
		StartDate:       in.StartDate,
		EndDate:         in.EndDate,
		MinOrder:        in.MinOrder,
		PointConversion: in.PointConversion,
		MonthlyMaxPoint: in.MonthlyMaxPoint,
		Customers:       customers,
	}

	repo := repository.NewPointRuleRepository(uc.DB)
	out.ID, err = repo.Add(c, out)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	return
}

// AddPhoto ...
func (uc PointRuleUC) AddPhoto(c context.Context, image *multipart.FileHeader) (out string, err error) {
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
func (uc PointRuleUC) Update(c context.Context, id string, in requests.PointRuleRequest) (out viewmodel.PointRuleVM, err error) {
	var customers []viewmodel.PointRuleCustomerVM
	for _, datum := range in.Customers {
		customers = append(customers, viewmodel.PointRuleCustomerVM{
			CustomerCode: datum.CustomerCode,
		})
	}

	out = viewmodel.PointRuleVM{
		ID:              id,
		StartDate:       in.StartDate,
		EndDate:         in.EndDate,
		MinOrder:        in.MinOrder,
		PointConversion: in.PointConversion,
		MonthlyMaxPoint: in.MonthlyMaxPoint,
		Customers:       customers,
	}

	repo := repository.NewPointRuleRepository(uc.DB)
	out.ID, err = repo.Update(c, out)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	return
}

// Delete ...
func (uc PointRuleUC) Delete(c context.Context, in string) (err error) {
	repo := repository.NewPointRuleRepository(uc.DB)
	_, err = repo.Delete(c, in)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	return
}
