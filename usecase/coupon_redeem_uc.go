package usecase

import (
	"context"
	"errors"
	"mime/multipart"
	"strconv"
	"time"

	"nextbasis-service-v-0.1/config"
	"nextbasis-service-v-0.1/db/repository"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/helper"
	"nextbasis-service-v-0.1/pkg/functioncaller"
	"nextbasis-service-v-0.1/pkg/logruslogger"
	"nextbasis-service-v-0.1/server/requests"
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

// CouponRedeemUC ...
type CouponRedeemUC struct {
	*ContractUC
}

// BuildBody ...
func (uc CouponRedeemUC) BuildBody(data *models.CouponRedeem, res *viewmodel.CouponRedeemVM) {
	res.ID = data.ID
	res.CouponID = data.CouponID
	res.CouponName = data.CouponName
	res.CouponDescription = data.CouponDescription
	res.CouponPointConversion = data.CouponPointConversion
	res.CouponPhotoURL = data.CouponPhotoURL.String
	res.CustomerID = data.CustomerID
	res.CustomerName = data.CustomerName
	res.Redeem = data.Redeem
	res.RedeemAt = data.RedeemAt.String
	res.RedeemedToDocumentNo = data.RedeemedToDocumentNo.String
	res.CreatedAt = data.CreatedAt
	res.UpdatedAt = data.UpdatedAt.String
	res.DeletedAt = data.DeletedAt.String
	res.ExpiredAt = data.ExpiredAt.String
	res.CouponCode = data.CouponCode.String
	res.InvoiceNo = data.InvoiceNo.String
}

// FindAll ...
func (uc CouponRedeemUC) FindAll(c context.Context, parameter models.CouponRedeemParameter) (out []viewmodel.CouponRedeemVM, p viewmodel.PaginationVM, err error) {
	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.PointRuleOrderBy, models.PointRuleOrderByrByString)

	repo := repository.NewCouponRedeemRepository(uc.DB)
	data, count, err := repo.FindAll(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	p = uc.setPaginationResponse(parameter.Page, parameter.Limit, count)
	for _, datum := range data {
		var temp viewmodel.CouponRedeemVM
		uc.BuildBody(&datum, &temp)

		out = append(out, temp)
	}

	if out == nil {
		out = make([]viewmodel.CouponRedeemVM, 0)
	}

	return
}

// SelectAll ...
func (uc CouponRedeemUC) SelectAll(c context.Context, parameter models.CouponRedeemParameter) (out []viewmodel.CouponRedeemVM, err error) {
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.PointRuleOrderBy, models.PointRuleOrderByrByString)

	repo := repository.NewCouponRedeemRepository(uc.DB)
	data, err := repo.SelectAll(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	for _, datum := range data {
		var temp viewmodel.CouponRedeemVM
		uc.BuildBody(&datum, &temp)

		out = append(out, temp)
	}

	if out == nil {
		out = make([]viewmodel.CouponRedeemVM, 0)
	}

	return
}

// FindByID ...
func (uc CouponRedeemUC) FindByID(c context.Context, parameter models.CouponRedeemParameter) (out viewmodel.CouponRedeemVM, err error) {
	repo := repository.NewCouponRedeemRepository(uc.DB)
	data, err := repo.FindByID(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	uc.BuildBody(&data, &out)

	return
}

// Add ...
func (uc CouponRedeemUC) Add(c context.Context, in requests.CouponRedeemRequest) (out viewmodel.CouponRedeemVM, err error) {
	couponData, err := CouponUC{ContractUC: uc.ContractUC}.FindByID(c, models.CouponParameter{ID: in.CouponID})
	if err != nil {
		err = errors.New("coupon not found")
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	pointUC := PointUC{ContractUC: uc.ContractUC}
	customerPoint, err := pointUC.GetBalance(c, models.PointParameter{
		CustomerID: in.CustomerID,
		Renewal:    "1",
	})
	if err != nil {
		err = errors.New("customer point not found")
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	point, _ := strconv.ParseFloat(customerPoint.Balance, 64)
	couponConversion, _ := strconv.ParseFloat(couponData.PointConversion, 64)
	if point < couponConversion {
		err = errors.New("there are insufficient point on your account")
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "insufficient_point", c.Value("requestid"))
		return
	}

	customerData, _ := WebCustomerUC{ContractUC: uc.ContractUC}.FindByID(c, models.WebCustomerParameter{ID: in.CustomerID})
	now := time.Now()
	out = viewmodel.CouponRedeemVM{
		CouponID:   in.CouponID,
		CustomerID: in.CustomerID,
		ExpiredAt:  helper.GetExpiredWithInterval(time.Now(), couponData.Interval),
		CouponCode: customerData.CustomerBranchCode + strconv.Itoa(int(now.Month())) + strconv.Itoa(now.Year()) + helper.StringWithCharset(6),
	}

	repo := repository.NewCouponRedeemRepository(uc.DB)
	out.ID, err = repo.Add(c, out)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	pointUC.AddWithdraw(c, requests.PointRequest{
		Point:      couponData.PointConversion,
		CustomerID: in.CustomerID,
	})

	return
}

// Redeem ...
func (uc CouponRedeemUC) Redeem(c context.Context, in models.CouponRedeemParameter) (out viewmodel.CouponRedeemVM, err error) {
	out = viewmodel.CouponRedeemVM{
		ID:                   in.ID,
		Redeem:               "1",
		RedeemedToDocumentNo: in.RedeemedToDocumentNo,
	}

	repo := repository.NewCouponRedeemRepository(uc.DB)
	out.ID, err = repo.Redeem(c, out)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	return
}

// Revert ...
func (uc CouponRedeemUC) Revert(c context.Context, invoiceNo string) (out viewmodel.CouponRedeemVM, err error) {
	out = viewmodel.CouponRedeemVM{
		Redeem:               "0",
		RedeemedToDocumentNo: invoiceNo,
	}

	repo := repository.NewCouponRedeemRepository(uc.DB)
	out.ID, err = repo.Revert(c, out)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	return
}

// AddPhoto ...
func (uc CouponRedeemUC) AddPhoto(c context.Context, image *multipart.FileHeader) (out string, err error) {
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

// SelectReport ...
func (uc CouponRedeemUC) SelectReport(c context.Context, parameter models.CouponRedeemParameter) (out []viewmodel.CouponRedeemReportVM, err error) {
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.PointRuleOrderBy, models.PointRuleOrderByrByString)

	repo := repository.NewCouponRedeemRepository(uc.DB)
	data, err := repo.SelectReport(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	for _, datum := range data {
		out = append(out, viewmodel.CouponRedeemReportVM{
			ID:                    datum.ID,
			CouponID:              datum.CouponID,
			CustomerID:            datum.CustomerID,
			Redeem:                datum.Redeem,
			RedeemAt:              datum.RedeemAt.String,
			RedeemedToDocumentNo:  datum.RedeemedToDocumentNo.String,
			CreatedAt:             datum.CreatedAt,
			UpdatedAt:             datum.UpdatedAt.String,
			DeletedAt:             datum.DeletedAt.String,
			ExpiredAt:             datum.ExpiredAt.String,
			CouponName:            datum.CouponName,
			CouponDescription:     datum.CouponDescription,
			CouponPointConversion: datum.CouponPointConversion,
			CustomerName:          datum.CustomerName,
			CustomerCode:          datum.CustomerCode,
			BranchName:            datum.BranchName,
			BranchCode:            datum.BranchCode,
			RegionName:            datum.RegionName,
			RegionGroupName:       datum.RegionGroupName,
			CustomerLevelName:     datum.CustomerLevelName.String,
			CouponCode:            datum.CouponCode.String,
			InvoiceNo:             datum.InvoiceNo.String,
			SalesOrderDocumentNo:  datum.SalesOrderDocumentNo.String,
		})
	}

	if out == nil {
		out = make([]viewmodel.CouponRedeemReportVM, 0)
	}

	return
}
