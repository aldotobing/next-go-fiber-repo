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

var (
	amountDefultValue    = "Rp. 0"
	quantityDefaultValue = "0"
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

// GetDataByGroupID ...
func (uc DashboardWebUC) GetDataByGroupID(c context.Context, parameter models.DashboardWebParameter) (res []viewmodel.DashboardByGroupID, err error) {
	repo := repository.NewDashboardWebRepository(uc.DB)
	data, err := repo.GetDataByGroupID(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	for i := range data {
		res = append(res, viewmodel.DashboardByGroupID{
			RegionID:                 data[i].RegionID,
			RegionName:               data[i].RegionName,
			TotalVisitUser:           data[i].TotalVisitUser,
			TotalRepeatUser:          data[i].TotalRepeatUser,
			TotalOrderUser:           data[i].TotalOrderUser,
			TotalInvoice:             data[i].TotalInvoice,
			TotalRegisteredUser:      data[i].TotalRegisteredUser,
			CustomerCountRepeatOrder: data[i].CustomerCountRepeatOrder,
			TotalActiveOutlet:        data[i].TotalActiveOutlet,
			TotalOutlet:              data[i].TotalOutlet,
			TotalCompleteCustomer:    data[i].TotalCompleteCustomer,
		})
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

func (uc DashboardWebUC) GetUserByRegionDetailData(c context.Context, parameter models.DashboardWebRegionParameter) (res []models.DashboardWebBranchDetail, err error) {
	repo := repository.NewDashboardWebRepository(uc.DB)
	res, err = repo.GetUserByRegionDetailData(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
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

func (uc DashboardWebUC) GetAllReportBranchDetailCustomerData(c context.Context, parameter models.DashboardWebBranchParameter) (res []models.DashboardWebBranchDetail, err error) {
	_, _, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.DashboardWebBranchDetailOrderBy, models.DashboardWebBranchDetailOrderByrByString)

	repo := repository.NewDashboardWebRepository(uc.DB)
	res, err = repo.GetAllReportBranchDetailCustomerData(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return res, err
}

func (uc DashboardWebUC) GetAllBranchDataWithUserID(c context.Context, parameter models.DashboardWebBranchParameter) (res []viewmodel.DashboardBranchByUserID, err error) {
	repo := repository.NewDashboardWebRepository(uc.DB)
	data, err := repo.GetAllBranchDataWithUserID(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	for i := range data {
		res = append(res, viewmodel.DashboardBranchByUserID{
			BranchID:              data[i].CustomerBranchID,
			BranchName:            data[i].CustomerBranchName,
			BranchCode:            data[i].CustomerBranchCode,
			RegionName:            data[i].CustomerRegionName,
			RegionGroupName:       data[i].CustomerRegionGroupName,
			TotalRepeatUser:       data[i].TotalRepeatUser,
			TotalRepeatToko:       data[i].TotalRepeatToko,
			TotalOrderUser:        data[i].TotalOrderUser,
			TotalInvoice:          data[i].TotalInvoice,
			TotalCheckin:          data[i].TotalCheckin,
			TotalAktifOutlet:      data[i].TotalAktifOutlet,
			TotalOutlet:           data[i].TotalOutlet,
			TotalOutletAll:        data[i].TotalOutletAll,
			TotalRegisteredUser:   data[i].TotalRegisteredUser,
			TotalCompleteCustomer: data[i].CompleteCustomer,
		})
	}

	if res == nil {
		res = make([]viewmodel.DashboardBranchByUserID, 0)
	}

	return res, err
}

func (uc DashboardWebUC) GetAllDetailCustomerDataWithUserID(c context.Context, parameter models.DashboardWebBranchParameter) (res []viewmodel.DashboardCustomerByUserID, err error) {
	repo := repository.NewDashboardWebRepository(uc.DB)
	data, err := repo.GetAllDetailCustomerDataWithUserID(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	for i := range data {
		res = append(res, viewmodel.DashboardCustomerByUserID{
			CustomerID:             data[i].CustomerID,
			CustomerName:           data[i].CustomerName,
			CustomerCode:           data[i].CustomerCode,
			BranchName:             data[i].CustomerBranchName,
			BranchCode:             data[i].CustomerBranchCode,
			RegionName:             data[i].CustomerRegionName,
			RegionGroupName:        data[i].CustomerRegionGroupName,
			CustomerTypeName:       data[i].CustomerTypeName,
			CustomerLevelName:      data[i].CustomerLevelName,
			CustomerCityName:       data[i].CustomerCityName,
			TotalRepeatUser:        data[i].TotalRepeatUser,
			TotalOrderUser:         data[i].TotalOrderUser,
			TotalInvoice:           data[i].TotalInvoice,
			TotalCheckin:           data[i].TotalCheckin,
			TotalAktifOutlet:       data[i].TotalAktifOutlet,
			StatusCompleteCustomer: data[i].CompleteCustomer,
		})
	}

	if res == nil {
		res = make([]viewmodel.DashboardCustomerByUserID, 0)
	}

	return res, err
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
		RegionGroupID:   &regionIDNational,
		RegionGroupName: &regionNameNational,
	})

	//Append the rest of region
	for i := range regionData {
		res = append(res, viewmodel.OmzetValueVM{
			RegionGroupID:   regionData[i].GroupID,
			RegionGroupName: regionData[i].GroupName,
			TotalQuantity:   &quantityDefaultValue,
			TotalOmzet:      &amountDefultValue,
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
				if omzetData[j].RegionGroupID.String == *res[i].RegionGroupID {
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

func (uc DashboardWebUC) GetOmzetValueByRegionGroupID(c context.Context, parameter models.DashboardWebBranchParameter, regionGroupID string) (res viewmodel.OmzetValueByRegionVM, err error) {
	regionUC := WebRegionAreaUC{ContractUC: uc.ContractUC}
	regionData, err := regionUC.SelectByGroupID(c, regionGroupID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "get_region_data", c.Value("requestid"))
		return res, err
	}

	var areas []viewmodel.OmzetValueAreaVM
	for i := range regionData {
		areas = append(areas, viewmodel.OmzetValueAreaVM{
			RegionID:        regionData[i].ID,
			RegionName:      regionData[i].Name,
			Quantity:        &quantityDefaultValue,
			Omzet:           &amountDefultValue,
			RegionGroupID:   regionData[i].GroupID,
			RegionGroupName: regionData[i].GroupName,
		})
	}

	repo := repository.NewDashboardWebRepository(uc.DB)
	omzetData, err := repo.GetOmzetValueByGroupID(c, parameter, regionGroupID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	var grandTotalQuantity, grandTotalOmzet float64
	acOmzet := accounting.Accounting{Symbol: "Rp. ", Thousand: ".", Decimal: ","}
	acQuantity := accounting.Accounting{Thousand: "."}

	for j := range omzetData {
		for i := range areas {
			if omzetData[j].RegionID.String == *areas[i].RegionID {
				amount, _ := strconv.ParseFloat(omzetData[j].TotalNettAmount, 64)
				resTotalOmzet := acOmzet.FormatMoney(amount)
				areas[i].Omzet = &resTotalOmzet
				grandTotalOmzet += amount

				amount, _ = strconv.ParseFloat(omzetData[j].TotalQuantity, 64)
				resTotalQuantity := acQuantity.FormatMoney(amount)
				areas[i].Quantity = &resTotalQuantity
				grandTotalQuantity += amount

				break
			} else if omzetData[j].RegionID.Valid == false {
				amount, _ := strconv.ParseFloat(omzetData[j].TotalNettAmount, 64)
				grandTotalOmzet += amount

				amount, _ = strconv.ParseFloat(omzetData[j].TotalQuantity, 64)
				grandTotalQuantity += amount
			}
		}
	}

	grandTotalOmzetString := acOmzet.FormatMoney(grandTotalOmzet)
	grandTotalQuantityString := acQuantity.FormatMoney(grandTotalQuantity)

	if areas == nil {
		areas = make([]viewmodel.OmzetValueAreaVM, 0)
	}

	res = viewmodel.OmzetValueByRegionVM{
		TotalOmzet:    &grandTotalOmzetString,
		TotalQuantity: &grandTotalQuantityString,
		Area:          areas,
	}

	return res, err
}

func (uc DashboardWebUC) GetOmzetValueByRegionID(c context.Context, parameter models.DashboardWebBranchParameter, regionID string) (res viewmodel.OmzetValueByBranchVM, err error) {
	branchUC := BranchUC{ContractUC: uc.ContractUC}
	branchData, err := branchUC.SelectAll(c, models.BranchParameter{
		RegionID: regionID,
		By:       "def.id",
	})
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "get_branch_data", c.Value("requestid"))
		return res, err
	}

	var branches []viewmodel.OmzetValueBranchVM
	for i := range branchData {
		branches = append(branches, viewmodel.OmzetValueBranchVM{
			BranchID:        branchData[i].ID,
			BranchName:      branchData[i].Name,
			BranchCode:      branchData[i].Code,
			RegionID:        branchData[i].RegionID,
			RegionName:      branchData[i].RegionName,
			RegionGroupID:   branchData[i].RegionGroupID,
			RegionGroupName: branchData[i].RegionGroupName,
			Quantity:        &quantityDefaultValue,
			Omzet:           &amountDefultValue,
		})
	}

	repo := repository.NewDashboardWebRepository(uc.DB)
	omzetData, err := repo.GetOmzetValueByRegionID(c, parameter, regionID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	var grandTotalQuantity, grandTotalOmzet float64

	acOmzet := accounting.Accounting{Symbol: "Rp. ", Thousand: ".", Decimal: ","}
	acQuantity := accounting.Accounting{Thousand: "."}

	for j := range omzetData {
		for i := range branches {
			if omzetData[j].BranchID.String == *branches[i].BranchID {
				amount, _ := strconv.ParseFloat(omzetData[j].TotalNettAmount, 64)
				resTotalOmzet := acOmzet.FormatMoney(amount)
				branches[i].Omzet = &resTotalOmzet
				grandTotalOmzet += amount

				amount, _ = strconv.ParseFloat(omzetData[j].TotalQuantity, 64)
				resTotalQuantity := acQuantity.FormatMoney(amount)
				branches[i].Quantity = &resTotalQuantity
				grandTotalQuantity += amount

				branches[i].ActiveCustomer = &omzetData[j].TotalActiveCustomer

				break
			}
		}
	}

	grandTotalOmzetString := acOmzet.FormatMoney(grandTotalOmzet)
	grandTotalQuantityString := acQuantity.FormatMoney(grandTotalQuantity)

	if branches == nil {
		branches = make([]viewmodel.OmzetValueBranchVM, 0)
	}

	res = viewmodel.OmzetValueByBranchVM{
		TotalOmzet:    &grandTotalOmzetString,
		TotalQuantity: &grandTotalQuantityString,
		Branches:      branches,
	}

	return res, err
}

func (uc DashboardWebUC) GetOmzetValueByBranchID(c context.Context, parameter models.DashboardWebBranchParameter, branchID string) (res viewmodel.OmzetValueByCustomerVM, err error) {
	repo := repository.NewDashboardWebRepository(uc.DB)
	omzetData, err := repo.GetOmzetValueByBranchID(c, parameter, branchID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	var grandTotalQuantity, grandTotalOmzet float64

	acOmzet := accounting.Accounting{Symbol: "Rp. ", Thousand: ".", Decimal: ","}
	acQuantity := accounting.Accounting{Thousand: "."}

	var customers []viewmodel.OmzetValueCustomerVM
	for i := range omzetData {
		amount, _ := strconv.ParseFloat(omzetData[i].TotalNettAmount, 64)
		grandTotalOmzet += amount
		resTotalOmzet := acOmzet.FormatMoney(amount)

		amount, _ = strconv.ParseFloat(omzetData[i].TotalQuantity, 64)
		grandTotalQuantity += amount
		resTotalQuantity := acQuantity.FormatMoney(amount)

		customers = append(customers, viewmodel.OmzetValueCustomerVM{
			RegionGroupName: omzetData[i].RegionGroupName,
			RegionName:      omzetData[i].RegionName,
			BranchName:      omzetData[i].BranchName,
			BranchCode:      omzetData[i].BranchCode,
			CustomerID:      omzetData[i].CustomerID,
			CustomerCode:    omzetData[i].CustomerCode,
			CustomerName:    omzetData[i].CustomerName,
			CustomerType:    omzetData[i].CustomerType,
			ProvinceName:    omzetData[i].ProvinceName,
			CityName:        omzetData[i].CityName,
			CustomerLevel:   omzetData[i].CustomerLevel,
			Quantity:        &resTotalQuantity,
			Omzet:           &resTotalOmzet,
		})
	}

	grandTotalOmzetString := acOmzet.FormatMoney(grandTotalOmzet)
	grandTotalQuantityString := acQuantity.FormatMoney(grandTotalQuantity)

	if customers == nil {
		customers = make([]viewmodel.OmzetValueCustomerVM, 0)
	}

	res = viewmodel.OmzetValueByCustomerVM{
		TotalOmzet:    &grandTotalOmzetString,
		TotalQuantity: &grandTotalQuantityString,
		Customers:     customers,
	}

	return res, err
}

func (uc DashboardWebUC) GetOmzetValueByCustomerID(c context.Context, parameter models.DashboardWebBranchParameter, customerID string) (res viewmodel.OmzetValueByItemVM, err error) {
	repo := repository.NewDashboardWebRepository(uc.DB)
	omzetData, err := repo.GetOmzetValueByCustomerID(c, parameter, customerID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	var grandTotalQuantity, grandTotalOmzet float64

	acOmzet := accounting.Accounting{Symbol: "Rp. ", Thousand: ".", Decimal: ","}
	acQuantity := accounting.Accounting{Thousand: "."}

	var customers []viewmodel.OmzetValueItemVM
	for i := range omzetData {
		amount, _ := strconv.ParseFloat(omzetData[i].TotalNettAmount, 64)
		grandTotalOmzet += amount
		resTotalOmzet := acOmzet.FormatMoney(amount)

		amount, _ = strconv.ParseFloat(omzetData[i].TotalQuantity, 64)
		grandTotalQuantity += amount
		resTotalQuantity := acQuantity.FormatMoney(amount)

		customers = append(customers, viewmodel.OmzetValueItemVM{
			ItemID:   &omzetData[i].ItemID.String,
			ItemName: &omzetData[i].ItemName.String,
			Quantity: &resTotalQuantity,
			Omzet:    &resTotalOmzet,
		})
	}

	grandTotalOmzetString := acOmzet.FormatMoney(grandTotalOmzet)
	grandTotalQuantityString := acQuantity.FormatMoney(grandTotalQuantity)

	if customers == nil {
		customers = make([]viewmodel.OmzetValueItemVM, 0)
	}

	res = viewmodel.OmzetValueByItemVM{
		TotalOmzet:    &grandTotalOmzetString,
		TotalQuantity: &grandTotalQuantityString,
		Customers:     customers,
	}

	return res, err
}

func (uc DashboardWebUC) GetTrackingInvoiceData(c context.Context, parameter models.DashboardWebBranchParameter) (res []viewmodel.DashboardCustomerByUserID, err error) {
	repo := repository.NewDashboardWebRepository(uc.DB)
	data, err := repo.GetAllDetailCustomerDataWithUserID(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	for i := range data {
		res = append(res, viewmodel.DashboardCustomerByUserID{
			CustomerID:        data[i].CustomerID,
			CustomerName:      data[i].CustomerName,
			CustomerCode:      data[i].CustomerCode,
			BranchName:        data[i].CustomerBranchName,
			BranchCode:        data[i].CustomerBranchCode,
			RegionName:        data[i].CustomerRegionName,
			RegionGroupName:   data[i].CustomerRegionGroupName,
			CustomerTypeName:  data[i].CustomerTypeName,
			CustomerLevelName: data[i].CustomerLevelName,
			CustomerCityName:  data[i].CustomerCityName,
			TotalRepeatUser:   data[i].TotalRepeatUser,
			TotalOrderUser:    data[i].TotalOrderUser,
			TotalInvoice:      data[i].TotalInvoice,
			TotalCheckin:      data[i].TotalCheckin,
			TotalAktifOutlet:  data[i].TotalAktifOutlet,
		})
	}

	if res == nil {
		res = make([]viewmodel.DashboardCustomerByUserID, 0)
	}

	return res, err
}
