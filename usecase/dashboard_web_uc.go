package usecase

import (
	"context"
	"strconv"

	"github.com/leekchan/accounting"
	"nextbasis-service-v-0.1/db/repository"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/functioncaller"
	"nextbasis-service-v-0.1/pkg/logruslogger"
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

// DashboardWebUC ...
type DashboardWebUC struct {
	*ContractUC
}

// BuildBody ...
func (uc DashboardWebUC) BuildBody(res *models.DashboardWeb) {
}

func (uc DashboardWebUC) BuildRegionDetailBody(res *models.DashboardWebRegionDetail) {
}

func (uc DashboardWebUC) BuildBranchDetailCustomerBody(res *models.DashboardWebBranchDetail) {
}

// FindByID ...
func (uc DashboardWebUC) GetData(c context.Context, parameter models.DashboardWebParameter) (res []models.DashboardWeb, err error) {
	repo := repository.NewDashboardWebRepository(uc.DB)
	res, err = repo.GetData(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	for i := range res {
		uc.BuildBody(&res[i])
	}

	return res, err
}

// FindByID ...
func (uc DashboardWebUC) GetRegionDetailData(c context.Context, parameter models.DashboardWebRegionParameter) (res []models.DashboardWebRegionDetail, err error) {
	repo := repository.NewDashboardWebRepository(uc.DB)
	res, err = repo.GetRegionDetailData(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	for i := range res {
		uc.BuildRegionDetailBody(&res[i])
	}

	return res, err
}

func (uc DashboardWebUC) GetBranchDetailCustomerData(c context.Context, parameter models.DashboardWebBranchParameter) (res []models.DashboardWebBranchDetail, p viewmodel.PaginationVM, err error) {
	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.DashboardWebBranchDetailOrderBy, models.DashboardWebBranchDetailOrderByrByString)

	var count int
	repo := repository.NewDashboardWebRepository(uc.DB)
	res, count, err = repo.GetBranchDetailCustomerData(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, p, err
	}

	p = uc.setPaginationResponse(parameter.Page, parameter.Limit, count)
	for i := range res {
		uc.BuildBranchDetailCustomerBody(&res[i])
	}

	return res, p, err
}

func (uc DashboardWebUC) GetAllBranchDetailCustomerData(c context.Context, parameter models.DashboardWebBranchParameter) (res []models.DashboardWebBranchDetail, p viewmodel.PaginationVM, err error) {
	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.DashboardWebBranchDetailOrderBy, models.DashboardWebBranchDetailOrderByrByString)

	var count int
	repo := repository.NewDashboardWebRepository(uc.DB)
	res, err = repo.GetAllBranchDetailCustomerData(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, p, err
	}

	p = uc.setPaginationResponse(parameter.Page, parameter.Limit, count)
	for i := range res {
		uc.BuildBranchDetailCustomerBody(&res[i])
	}

	return res, p, err
}

func (uc DashboardWebUC) GetAllDetailCustomerDataWithUserID(c context.Context, parameter models.DashboardWebBranchParameter) (res []models.DashboardWebBranchDetail, p viewmodel.PaginationVM, err error) {
	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.DashboardWebBranchDetailOrderBy, models.DashboardWebBranchDetailOrderByrByString)

	var count int
	repo := repository.NewDashboardWebRepository(uc.DB)
	res, err = repo.GetAllDetailCustomerDataWithUserID(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, p, err
	}

	p = uc.setPaginationResponse(parameter.Page, parameter.Limit, count)

	return res, p, err
}

func (uc DashboardWebUC) GetOmzetValue(c context.Context, parameter models.DashboardWebBranchParameter) (res []viewmodel.OmzetValueVM, err error) {
	regionUC := WebRegionAreaUC{ContractUC: uc.ContractUC}
	regionData, err := regionUC.SelectAllGroupByRegion(c)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "get_region_data", c.Value("requestid"))
		return res, err
	}

	// initiate region national
	var regionIDNational, regionNameNational string
	regionIDNational = "0"
	regionNameNational = "Nasional"
	res = append(res, viewmodel.OmzetValueVM{
		RegionID:   &regionIDNational,
		RegionName: &regionNameNational,
	})

	//Append the rest of region
	for i := range regionData {
		res = append(res, viewmodel.OmzetValueVM{
			RegionID:   regionData[i].GroupID,
			RegionName: regionData[i].GroupName,
		})
	}

	repo := repository.NewDashboardWebRepository(uc.DB)
	omzetData, err := repo.GetOmzetValue(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	acOmzet := accounting.Accounting{Symbol: "Rp. ", Thousand: ".", Decimal: ","}
	acQuantity := accounting.Accounting{Thousand: "."}
	for i := range res {
		var totalAmount, totalQuantity float64
		for j := range omzetData {
			if i == 0 {
				amount, _ := strconv.ParseFloat(omzetData[j].TotalNettAmount, 64)
				totalAmount += amount
				amount, _ = strconv.ParseFloat(omzetData[j].TotalQuantity, 64)
				totalQuantity += amount
			} else {
				if omzetData[j].RegionID.String == *res[i].RegionID {
					amount, _ := strconv.ParseFloat(omzetData[j].TotalNettAmount, 64)
					totalAmount += amount
					amount, _ = strconv.ParseFloat(omzetData[j].TotalQuantity, 64)
					totalQuantity += amount

					break
				}
			}
		}

		resTotalOmzet := acOmzet.FormatMoney(totalAmount)
		res[i].TotalOmzet = &resTotalOmzet
		resTotalQuantity := acQuantity.FormatMoney(totalQuantity)
		res[i].TotalQuantity = &resTotalQuantity
	}

	return res, err
}
