package usecase

import (
	"context"
	"strconv"
	"strings"

	"nextbasis-service-v-0.1/db/repository"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/functioncaller"
	"nextbasis-service-v-0.1/pkg/logruslogger"
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

// ItemProductFocusUC ...
type ItemProductFocusUC struct {
	*ContractUC
}

// BuildBody ...
func (uc ItemProductFocusUC) BuildBody(res *models.ItemProductFocus) {
}

// SelectAll ...
func (uc ItemProductFocusUC) SelectAll(c context.Context, parameter models.ItemProductFocusParameter) (res []models.ItemProductFocus, err error) {
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(0, 0, parameter.By, parameter.Sort, models.ItemProductFocusOrderBy, models.ItemProductFocusOrderByrByString)

	repo := repository.NewItemProductFocusRepository(uc.DB)
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
func (uc ItemProductFocusUC) FindAll(c context.Context, parameter models.ItemProductFocusParameter) (res []models.ItemProductFocus, p viewmodel.PaginationVM, err error) {
	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.ItemProductFocusOrderBy, models.ItemProductFocusOrderByrByString)

	var count int
	repo := repository.NewItemProductFocusRepository(uc.DB)
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
func (uc ItemProductFocusUC) FindByID(c context.Context, parameter models.ItemProductFocusParameter) (res models.ItemProductFocus, err error) {
	repo := repository.NewItemProductFocusRepository(uc.DB)
	res, err = repo.FindByID(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}
	uc.BuildBody(&res)

	return res, err
}

// SelectAllV2 ...
func (uc ItemProductFocusUC) SelectAllV2(c context.Context, parameter models.ItemProductFocusParameter) (res []viewmodel.ItemVM, err error) {
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(0, 0, parameter.By, parameter.Sort, models.ItemProductFocusOrderBy, models.ItemProductFocusOrderByrByString)

	customerUC := CustomerUC{ContractUC: uc.ContractUC}
	userData, err := customerUC.FindByID(c, models.CustomerParameter{
		ID: parameter.CustomerID,
	})
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "find_customer_by_id", c.Value("requestid"))
		return res, err
	}

	repo := repository.NewItemProductFocusRepository(uc.DB)
	data, err := repo.SelectAllV2(c, parameter, *userData.CustomerBranchID, *userData.CustomerTypeId, *userData.CustomerPriceListID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	for i := range data {
		additional := strings.Split(*data[i].AdditionalData, "|")

		var uoms []viewmodel.Uom
		if len(additional) > 0 && additional[0] != "" {
			// Find Lowest Price and lowest conversion
			var lowestPrice, lowestConversion float64
			for _, addDatum := range additional {
				perAddDatum := strings.Split(addDatum, "#sep#")
				price, _ := strconv.ParseFloat(perAddDatum[3], 64)
				conversion, _ := strconv.ParseFloat(perAddDatum[2], 64)
				if price < lowestPrice || lowestPrice == 0 {
					lowestPrice = price
					lowestConversion = conversion
				}
			}

			multiplyData := strings.Split(*data[i].MultiplyData, "|")
			if len(multiplyData) > 0 && multiplyData[0] != "" {
				basePrice := lowestPrice / lowestConversion
				for _, multiplyDatum := range multiplyData {
					perMultiDatum := strings.Split(multiplyDatum, "#sep#")

					if perMultiDatum[3] == "1" {
						conversion, _ := strconv.ParseFloat(perMultiDatum[2], 64)
						price := strconv.FormatFloat(basePrice*conversion, 'f', 2, 64)

						uoms = append(uoms, viewmodel.Uom{
							ID:               &perMultiDatum[0],
							Name:             &perMultiDatum[1],
							Conversion:       &perMultiDatum[2],
							ItemDetailsPrice: &price,
						})
					}
				}
			}
		}

		if len(uoms) > 0 {
			itemCategoryData := strings.Split(*data[i].ItemCategory, "#sep#")
			res = append(res, viewmodel.ItemVM{
				ID:               data[i].ID,
				Code:             data[i].Code,
				Name:             data[i].Name,
				Description:      data[i].Description,
				ItemCategoryId:   &itemCategoryData[0],
				ItemCategoryName: &itemCategoryData[1],
				ItemPicture:      data[i].ItemPicture,
				Uom:              uoms,
			})
		}
	}

	return res, err
}
