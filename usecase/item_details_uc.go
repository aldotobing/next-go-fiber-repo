package usecase

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"nextbasis-service-v-0.1/db/repository"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/functioncaller"
	"nextbasis-service-v-0.1/pkg/logruslogger"
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

// ItemDetailsUC ...
type ItemDetailsUC struct {
	*ContractUC
}

// BuildBody ...
func (uc ItemDetailsUC) BuildBody(res *models.ItemDetails) {
}

// SelectAll ...
func (uc ItemDetailsUC) SelectAll(c context.Context, parameter models.ItemDetailsParameter) (res []models.ItemDetails, err error) {
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(0, 0, parameter.By, parameter.Sort, models.ItemDetailsOrderBy, models.ItemDetailsOrderByrByString)

	repo := repository.NewItemDetailsRepository(uc.DB)
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
func (uc ItemDetailsUC) FindAll(c context.Context, parameter models.ItemDetailsParameter) (res []models.ItemDetails, p viewmodel.PaginationVM, err error) {
	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.ItemDetailsOrderBy, models.ItemDetailsOrderByrByString)

	var count int
	repo := repository.NewItemDetailsRepository(uc.DB)
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
func (uc ItemDetailsUC) FindByID(c context.Context, parameter models.ItemDetailsParameter) (res models.ItemDetails, err error) {
	repo := repository.NewItemDetailsRepository(uc.DB)
	res, err = repo.FindByID(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return res, err
}

// FindByIDV2 ...
func (uc ItemDetailsUC) FindByIDV2(c context.Context, parameter models.ItemDetailsParameter) (res viewmodel.ItemDetailsVM, err error) {
	repo := repository.NewItemDetailsRepository(uc.DB)

	data, err := repo.FindByIDs(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	// Find Lowest Price and lowest conversion
	var lowestPrice, lowestConversion float64
	var lowestUOMName, updatedData string
	for _, datum := range data {
		price, _ := strconv.ParseFloat(*datum.ItemDetailsPrice, 64)
		conversion, _ := strconv.ParseFloat(*datum.UomLineConversion, 64)
		if price < lowestPrice || lowestPrice == 0 {
			lowestPrice = price
			lowestConversion = conversion
			lowestUOMName = *datum.UomName

			updatedData = datum.ItemPriceCreatedAT.String
			fmt.Println("append:", datum.ItemPriceCreatedAT.String)
		}

		dbUpdatedData, _ := time.Parse("2006-01-02T15:04:05.999999Z", datum.ItemPriceCreatedAT.String)
		updatedDataTime, errParse := time.Parse("2006-01-02T15:04:05.999999Z", updatedData)
		if lowestConversion == conversion && (updatedDataTime.Before(dbUpdatedData) || errParse != nil) {
			lowestPrice = price
			lowestConversion = conversion
			lowestUOMName = *datum.UomName

			updatedData = datum.ItemPriceCreatedAT.String
		}
		parameter.PriceListVersionId = *datum.PriceListVersionId
	}

	data, err = repo.FindByIDV2(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	basePrice := lowestPrice / lowestConversion
	var uoms []viewmodel.Uom
	for _, datum := range data {
		if *datum.Visibility == "1" {
			conversion, _ := strconv.ParseFloat(*datum.UomLineConversion, 64)
			price := strconv.FormatFloat(basePrice*conversion, 'f', 2, 64)

			uoms = append(uoms, viewmodel.Uom{
				ID:               datum.UomID,
				Name:             datum.UomName,
				Conversion:       datum.UomLineConversion,
				ItemDetailsPrice: &price,
			})
		}
	}

	if len(uoms) == 0 {
		err = errors.New("uom not available")
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uom_checker", c.Value("requestid"))
		return
	}

	res = viewmodel.ItemDetailsVM{
		ID:                      data[0].ID,
		Code:                    data[0].Code,
		Name:                    data[0].Name,
		Description:             data[0].Description,
		ItemDetailsCategoryId:   data[0].ItemDetailsCategoryId,
		ItemDetailsCategoryName: data[0].ItemDetailsCategoryName,
		ItemDetailsPicture:      data[0].ItemDetailsPicture,
		LowestUOMName:           &lowestUOMName,
		Uom:                     uoms,
		PriceListVersionId:      &parameter.PriceListVersionId,
	}

	return res, err
}
