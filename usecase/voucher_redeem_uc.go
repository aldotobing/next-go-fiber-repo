package usecase

import (
	"context"

	"nextbasis-service-v-0.1/db/repository"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/functioncaller"
	"nextbasis-service-v-0.1/pkg/logruslogger"
	"nextbasis-service-v-0.1/server/requests"
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

// VoucherRedeemUC ...
type VoucherRedeemUC struct {
	*ContractUC
}

// BuildBody ...
func (uc VoucherRedeemUC) BuildBody(data *models.VoucherRedeem, res *viewmodel.VoucherRedeemVM) {
	res.ID = data.ID
	res.CustomerCode = data.CustomerCode
	res.Redeemed = data.Redeemed
	res.RedeemedAt = data.RedeemedAt.String
	res.RedeemedToDocumentNo = data.RedeemedToDocNo.String
	res.CreatedAt = data.CreatedAt
	res.UpdatedAt = data.UpdatedAt.String
	res.DeletedAt = data.DeletedAt.String

	res.VoucherID = data.VoucherID
	res.VoucherName = data.VoucherName
	res.VoucherCashValue = data.VoucherCashValue
	res.VoucherDescription = data.VoucherDescription.String
	res.VoucherImageURL = data.VoucherImageURL
	res.VoucherStartDate = data.VoucherStartDate
	res.VoucherEndDate = data.VoucherEndDate
}

// FindAll ...
func (uc VoucherRedeemUC) FindAll(c context.Context, parameter models.VoucherRedeemParameter) (out []viewmodel.VoucherRedeemVM, p viewmodel.PaginationVM, err error) {
	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.VoucherRedeemOrderBy, models.VoucherRedeemOrderByrByString)

	repo := repository.NewVoucherRedeemRepository(uc.DB)
	data, count, err := repo.FindAll(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	p = uc.setPaginationResponse(parameter.Page, parameter.Limit, count)
	for _, datum := range data {
		var temp viewmodel.VoucherRedeemVM
		uc.BuildBody(&datum, &temp)

		out = append(out, temp)
	}

	if out == nil {
		out = make([]viewmodel.VoucherRedeemVM, 0)
	}

	return
}

// SelectAll ...
func (uc VoucherRedeemUC) SelectAll(c context.Context, parameter models.VoucherRedeemParameter) (out []viewmodel.VoucherRedeemVM, err error) {
	repo := repository.NewVoucherRedeemRepository(uc.DB)
	data, err := repo.SelectAll(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	for _, datum := range data {
		var temp viewmodel.VoucherRedeemVM
		uc.BuildBody(&datum, &temp)

		out = append(out, temp)
	}

	if out == nil {
		out = make([]viewmodel.VoucherRedeemVM, 0)
	}

	return
}

// FindByID ...
func (uc VoucherRedeemUC) FindByID(c context.Context, parameter models.VoucherRedeemParameter) (out viewmodel.VoucherRedeemVM, err error) {
	repo := repository.NewVoucherRedeemRepository(uc.DB)
	data, err := repo.FindByID(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	uc.BuildBody(&data, &out)

	return
}

// FindByDocumentNo ...
func (uc VoucherRedeemUC) FindByDocumentNo(c context.Context, parameter models.VoucherRedeemParameter) (out viewmodel.VoucherRedeemVM, err error) {
	repo := repository.NewVoucherRedeemRepository(uc.DB)
	data, err := repo.FindByDocumentNo(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	uc.BuildBody(&data, &out)

	return
}

// Add ...
func (uc VoucherRedeemUC) Add(c context.Context, in requests.VoucherRedeemRequest) (out viewmodel.VoucherRedeemVM, err error) {
	out = viewmodel.VoucherRedeemVM{
		CustomerCode: in.CustomerCode,
		VoucherID:    in.VoucherID,
	}

	broadcastRepo := repository.NewVoucherRedeemRepository(uc.DB)
	out.ID, err = broadcastRepo.Add(c, out)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	return
}

// AddBulk ...
func (uc VoucherRedeemUC) AddBulk(c context.Context, in requests.VoucherRedeemBulkRequest) (out []viewmodel.VoucherRedeemVM, err error) {
	for _, datum := range in.VouchersRedeem {
		out = append(out, viewmodel.VoucherRedeemVM{
			CustomerCode: datum.CustomerCode,
			VoucherID:    in.VoucherID,
		})
	}

	broadcastRepo := repository.NewVoucherRedeemRepository(uc.DB)
	err = broadcastRepo.AddBulk(c, out)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	return
}

// Update ...
func (uc VoucherRedeemUC) Update(c context.Context, id string, in requests.VoucherRedeemRequest) (out viewmodel.VoucherRedeemVM, err error) {
	out = viewmodel.VoucherRedeemVM{
		ID:           id,
		CustomerCode: in.CustomerCode,
		VoucherID:    in.VoucherID,
	}

	broadcastRepo := repository.NewVoucherRedeemRepository(uc.DB)
	out.ID, err = broadcastRepo.Update(c, out)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	return
}

// Redeem ...
func (uc VoucherRedeemUC) Redeem(c context.Context, id string, in requests.VoucherRedeemRequest) (out viewmodel.VoucherRedeemVM, err error) {
	out = viewmodel.VoucherRedeemVM{
		ID:                   id,
		RedeemedToDocumentNo: in.RedeemToDocumentNo,
	}

	broadcastRepo := repository.NewVoucherRedeemRepository(uc.DB)
	out.ID, err = broadcastRepo.Redeem(c, out)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	return
}

// Delete ...
func (uc VoucherRedeemUC) Delete(c context.Context, in string) (err error) {
	repo := repository.NewVoucherRedeemRepository(uc.DB)
	_, err = repo.Delete(c, in)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	return
}

// Paid ...
func (uc VoucherRedeemUC) PaidRedeem(c context.Context, in viewmodel.VoucherRedeemVM) (out viewmodel.VoucherRedeemVM, err error) {

	broadcastRepo := repository.NewVoucherRedeemRepository(uc.DB)
	out.ID, err = broadcastRepo.PaidRedeem(c, in)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	return
}
