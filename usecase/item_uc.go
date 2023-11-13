package usecase

import (
	"context"
	"strconv"
	"strings"
	"time"

	"nextbasis-service-v-0.1/db/repository"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/functioncaller"
	"nextbasis-service-v-0.1/pkg/logruslogger"
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

// ItemUC ...
type ItemUC struct {
	*ContractUC
}

// BuildBody ...
func (uc ItemUC) BuildBody(res *models.Item) {
}

// SelectAll ...
func (uc ItemUC) SelectAll(c context.Context, parameter models.ItemParameter) (res []models.Item, err error) {
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(0, 0, parameter.By, parameter.Sort, models.ItemOrderBy, models.ItemOrderByrByString)

	repo := repository.NewItemRepository(uc.DB)
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

// SelectAllV2 ...
func (uc ItemUC) SelectAllV2(c context.Context, parameter models.ItemParameter, allParam bool) (res []viewmodel.ItemVM, err error) {
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(0, 0, parameter.By, parameter.Sort, models.ItemOrderBy, models.ItemOrderByrByString)

	repo := repository.NewItemRepository(uc.DB)
	data, err := repo.SelectAllV2(c, parameter, allParam)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	for i := range data {
		additional := strings.Split(*data[i].AdditionalData, "|")

		var uoms []viewmodel.Uom
		var lowestVisibleUOM string
		var lowestVisibleConversion float64
		if len(additional) > 0 && additional[0] != "" {
			// Find Lowest Price and lowest conversion
			var lowestPrice, lowestConversion float64
			var newestModifiedDate time.Time
			for _, addDatum := range additional {
				perAddDatum := strings.Split(addDatum, "#sep#")
				price, _ := strconv.ParseFloat(perAddDatum[4], 64)
				conversion, _ := strconv.ParseFloat(perAddDatum[2], 64)

				dbUpdatedDate, errParse := time.Parse("2006-01-02 15:04:05.999999", perAddDatum[3])
				if (newestModifiedDate.Before(dbUpdatedDate) || errParse != nil) || lowestPrice == 0 {
					lowestPrice = price
					lowestConversion = conversion
					newestModifiedDate = dbUpdatedDate
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

						if lowestVisibleUOM == "" || conversion < lowestVisibleConversion {
							lowestVisibleUOM = perMultiDatum[1]
							lowestVisibleConversion = conversion
						}

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
			res = append(res, viewmodel.ItemVM{
				ID:               data[i].ID,
				Name:             data[i].Name,
				Code:             data[i].Code,
				Description:      data[i].Description,
				ItemCategoryId:   data[i].ItemCategoryId,
				ItemCategoryName: data[i].ItemCategoryName,
				ItemPicture:      data[i].ItemPicture,
				Uom:              uoms,
				LowestUom:        lowestVisibleUOM,
			})
		}
	}

	return
}

// FindAll ...
func (uc ItemUC) FindAll(c context.Context, parameter models.ItemParameter) (res []models.Item, p viewmodel.PaginationVM, err error) {
	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.ItemOrderBy, models.ItemOrderByrByString)

	var count int
	repo := repository.NewItemRepository(uc.DB)
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
func (uc ItemUC) FindByID(c context.Context, parameter models.ItemParameter) (res models.Item, err error) {
	repo := repository.NewItemRepository(uc.DB)
	res, err = repo.FindByID(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}
	uc.BuildBody(&res)

	return res, err
}
