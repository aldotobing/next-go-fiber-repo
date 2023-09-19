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

// VoucherUC ...
type VoucherUC struct {
	*ContractUC
}

// BuildBody ...
func (uc VoucherUC) BuildBody(data *models.Voucher, res *viewmodel.VoucherVM) {
	res.ID = data.ID
	res.Code = data.Code
	res.Name = data.Name
	res.StartDate = data.StartDate
	res.EndDate = data.EndDate
	res.ImageURL = data.ImageURL
	res.VoucherCategoryID = data.VoucherCategoryID
	res.CashValue = data.CashValue
	res.Description = data.Description.String
	res.TermAndCondition = data.TermAndCondition.String
	res.CreatedAt = data.CreatedAt
	res.UpdatedAt = data.UpdatedAt.String
	res.DeletedAt = data.DeletedAt.String
}

// FindAll ...
func (uc VoucherUC) FindAll(c context.Context, parameter models.VoucherParameter) (out []viewmodel.VoucherVM, p viewmodel.PaginationVM, err error) {
	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.VoucherOrderBy, models.VoucherOrderByrByString)

	repo := repository.NewVoucherRepository(uc.DB)
	data, count, err := repo.FindAll(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	p = uc.setPaginationResponse(parameter.Page, parameter.Limit, count)
	for _, datum := range data {
		var temp viewmodel.VoucherVM
		uc.BuildBody(&datum, &temp)

		out = append(out, temp)
	}

	if out == nil {
		out = make([]viewmodel.VoucherVM, 0)
	}

	return
}

// SelectAll ...
func (uc VoucherUC) SelectAll(c context.Context, parameter models.VoucherParameter) (out []viewmodel.VoucherVM, err error) {
	repo := repository.NewVoucherRepository(uc.DB)
	data, err := repo.SelectAll(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	for _, datum := range data {
		var temp viewmodel.VoucherVM
		uc.BuildBody(&datum, &temp)

		out = append(out, temp)
	}

	if out == nil {
		out = make([]viewmodel.VoucherVM, 0)
	}

	return
}

// FindByID ...
func (uc VoucherUC) FindByID(c context.Context, parameter models.VoucherParameter) (out viewmodel.VoucherVM, err error) {
	repo := repository.NewVoucherRepository(uc.DB)
	data, err := repo.FindByID(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	uc.BuildBody(&data, &out)

	return
}

// Add ...
func (uc VoucherUC) Add(c context.Context, in requests.VoucherRequest) (out viewmodel.VoucherVM, err error) {
	out = viewmodel.VoucherVM{
		Code:              in.Code,
		Name:              in.Name,
		StartDate:         in.StartDate,
		EndDate:           in.EndDate,
		ImageURL:          in.ImageURL,
		VoucherCategoryID: in.VoucherCategoryID,
		CashValue:         in.CashValue,
		Description:       in.Description,
		TermAndCondition:  in.TermAndCondition,
	}

	repo := repository.NewVoucherRepository(uc.DB)
	out.ID, err = repo.Add(c, out)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	return
}

// AddPhoto ...
func (uc VoucherUC) AddPhoto(c context.Context, image *multipart.FileHeader) (out string, err error) {
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
func (uc VoucherUC) Update(c context.Context, id string, in requests.VoucherRequest) (out viewmodel.VoucherVM, err error) {
	out = viewmodel.VoucherVM{
		ID:                id,
		Code:              in.Code,
		Name:              in.Name,
		StartDate:         in.StartDate,
		EndDate:           in.EndDate,
		ImageURL:          in.ImageURL,
		VoucherCategoryID: in.VoucherCategoryID,
		CashValue:         in.CashValue,
		Description:       in.Description,
		TermAndCondition:  in.TermAndCondition,
	}

	repo := repository.NewVoucherRepository(uc.DB)
	out.ID, err = repo.Update(c, out)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	return
}

// Delete ...
func (uc VoucherUC) Delete(c context.Context, in string) (err error) {
	repo := repository.NewVoucherRepository(uc.DB)
	_, err = repo.Delete(c, in)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	return
}
