package usecase

import (
	"context"
	"fmt"
	"strconv"

	"nextbasis-service-v-0.1/db/repository"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/functioncaller"
	"nextbasis-service-v-0.1/pkg/logruslogger"
	"nextbasis-service-v-0.1/server/requests"
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

// ItemOldPriceUC ...
type ItemOldPriceUC struct {
	*ContractUC
}

// BuildBody ...
func (uc ItemOldPriceUC) BuildBody(in *models.ItemOldPrice, out *viewmodel.ItemOldPriceVM) {
	qty, _ := strconv.ParseFloat(in.Quantity, 64)

	out.ID = in.ID
	out.CustomerID = in.CustomerID
	out.CustomerCode = in.CustomerCode
	out.CustomerName = in.CustomerName
	out.ItemID = in.ItemID
	out.ItemCode = in.ItemCode
	out.ItemName = in.ItemName
	out.ItemPicture = in.ItemPicture.String
	out.UomID = in.UomID
	out.UomName = in.UomName
	out.StartDate = in.StartDate
	out.EndDate = in.EndDate
	out.Quantity = int(qty)
	out.PriceListID = in.PriceListID
	out.SellPrice = in.SellPrice
	out.PreservedQty = in.PreservedQty
	out.InvoiceQty = in.InvoiceQty
	out.CreatedAt = in.CreatedAt.String
	out.UpdatedAt = in.UpdatedAt.String
	out.DeletedAt = in.DeletedAt.String
}

// SelectAll ...
func (uc ItemOldPriceUC) SelectAll(c context.Context, parameter models.ItemOldPriceParameter) (res []viewmodel.ItemOldPriceVM, err error) {
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(0, 0, parameter.By, parameter.Sort, models.ItemOldPriceOrderBy, models.ItemOldPriceOrderByrByString)

	repo := repository.NewItemOldPriceRepository(uc.DB)
	data, err := repo.SelectAll(c, parameter)

	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	for i := range data {
		var temp viewmodel.ItemOldPriceVM
		uc.BuildBody(&data[i], &temp)

		res = append(res, temp)
	}

	if res == nil {
		res = make([]viewmodel.ItemOldPriceVM, 0)
	}

	return res, err
}

// FindAll ...
func (uc ItemOldPriceUC) FindAll(c context.Context, parameter models.ItemOldPriceParameter) (res []viewmodel.ItemOldPriceVM, p viewmodel.PaginationVM, err error) {
	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.ItemOldPriceOrderBy, models.ItemOldPriceOrderByrByString)

	repo := repository.NewItemOldPriceRepository(uc.DB)
	data, count, err := repo.FindAll(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, p, err
	}

	p = uc.setPaginationResponse(parameter.Page, parameter.Limit, count)
	for i := range data {
		var temp viewmodel.ItemOldPriceVM
		uc.BuildBody(&data[i], &temp)
		res = append(res, temp)
	}

	if res == nil {
		res = make([]viewmodel.ItemOldPriceVM, 0)
	}

	return res, p, err
}

// FindByID ...
func (uc ItemOldPriceUC) FindByID(c context.Context, parameter models.ItemOldPriceParameter) (res viewmodel.ItemOldPriceVM, err error) {
	repo := repository.NewItemOldPriceRepository(uc.DB)
	data, err := repo.FindByID(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}
	uc.BuildBody(&data, &res)

	return res, err
}

// ItemDetailFindByID ...
func (uc ItemOldPriceUC) ItemDetailFindByID(c context.Context, parameter models.ItemOldPriceParameter) (res []viewmodel.ItemVM, err error) {
	repo := repository.NewItemOldPriceRepository(uc.DB)
	data, err := repo.FindByID(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	var resData viewmodel.ItemOldPriceVM
	uc.BuildBody(&data, &resData)

	res, err = ItemUC{ContractUC: uc.ContractUC}.SelectAllV2(c, models.ItemParameter{
		ID:          resData.ItemID,
		By:          "def.id",
		PriceListId: resData.PriceListID,
	}, false, false)

	itemUomLineData, err := WebItemUomLineUC{ContractUC: uc.ContractUC}.SelectAll(c, models.WebItemUomLineParameter{
		ItemID: resData.ItemID,
		By:     "def.id",
		Sort:   "asc",
	})

	var basePrice float64
	sellPriceFloat, _ := strconv.ParseFloat(resData.SellPrice, 64)
	for i := range itemUomLineData {
		if *itemUomLineData[i].ItemUomID == resData.UomID {
			conversion, _ := strconv.ParseFloat(*itemUomLineData[i].ItemUomConversion, 64)
			basePrice = sellPriceFloat / float64(conversion)

			break
		}
	}

	//convert to old price
	for i := range res {
		for j := range res[i].Uom {
			conversion, _ := strconv.ParseFloat(*res[i].Uom[j].Conversion, 64)

			price := strconv.FormatFloat(basePrice*conversion, 'f', 2, 64)

			preservedQtyFloat, _ := strconv.ParseFloat(resData.PreservedQty, 64)
			InvoiceQtyFloat, _ := strconv.ParseFloat(resData.InvoiceQty, 64)

			limitQuantity := strconv.FormatFloat(
				(float64(resData.Quantity)/conversion)-(preservedQtyFloat/conversion)-(InvoiceQtyFloat/conversion),
				'f', 2, 64)

			res[i].Uom[j].ItemDetailsPrice = &price
			res[i].Uom[j].LimitQuantity = &limitQuantity
		}
	}
	return res, err
}

// Add ...
func (uc ItemOldPriceUC) Add(c context.Context, in requests.ItemOldPriceBulkRequest) (res []viewmodel.ItemOldPriceVM, err error) {
	var customerCodes, itemCodes string
	for _, datum := range in.OldPrice {
		if customerCodes != "" {
			customerCodes += `,'` + datum.CustomerCode + `'`
		} else {
			customerCodes += `'` + datum.CustomerCode + `'`
		}

		if itemCodes != "" {
			itemCodes += `,'` + datum.ItemCode + `'`
		} else {
			itemCodes += `'` + datum.ItemCode + `'`
		}
	}
	customerData, err := WebCustomerUC{ContractUC: uc.ContractUC}.FindByCodes(c, models.WebCustomerParameter{Code: customerCodes})
	for _, datum := range in.OldPrice {
		var customerID, customerPriceListID string
		for _, customerDatum := range customerData {
			if customerDatum.Code == datum.CustomerCode {
				customerID = customerDatum.ID
				customerPriceListID = customerDatum.CustomerPriceListID
				break
			}
		}

		itemData, _ := ItemUC{ContractUC: uc.ContractUC}.SelectAllV2(c, models.ItemParameter{
			Code:        datum.ItemCode,
			By:          "def.id",
			PriceListId: customerPriceListID,
		}, false, true)

		var itemID, uomID string
		var itemPrice, conversion float64
		for _, itemDatum := range itemData {
			if *itemDatum.Code == datum.ItemCode {
				itemID = *itemDatum.ID

				for i, itemPriceDatum := range itemDatum.Uom {
					if i == 0 {
						itemPrice, _ = strconv.ParseFloat(*itemPriceDatum.ItemDetailsPrice, 64)
						conversion, _ = strconv.ParseFloat(*itemPriceDatum.Conversion, 64)
						uomID = *itemPriceDatum.ID
					} else {
						itemPriceTemp, _ := strconv.ParseFloat(*itemPriceDatum.ItemDetailsPrice, 64)
						conversionTemp, _ := strconv.ParseFloat(*itemPriceDatum.Conversion, 64)
						if conversionTemp <= conversion {
							itemPrice = itemPriceTemp
							conversion = conversionTemp
							uomID = *itemPriceDatum.ID
						}
					}
				}
				break
			}
		}

		res = append(res, viewmodel.ItemOldPriceVM{
			CustomerID:   customerID,
			ItemID:       itemID,
			Quantity:     datum.Quantity,
			PreservedQty: "0",
			InvoiceQty:   "0",
			PriceListID:  customerPriceListID,
			SellPrice:    fmt.Sprintf("%.2f", itemPrice),
			UomID:        uomID,
			StartDate:    datum.StartDate,
			EndDate:      datum.EndDate,
		})
	}

	repo := repository.NewItemOldPriceRepository(uc.DB)
	err = repo.Add(c, res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "add_query", c.Value("requestid"))
		return res, err
	}

	return res, err
}

// Update ...
func (uc ItemOldPriceUC) Update(c context.Context, id string, in requests.ItemOldPriceRequest) (out viewmodel.ItemOldPriceVM, err error) {
	out = viewmodel.ItemOldPriceVM{
		ID:        id,
		StartDate: in.StartDate,
		EndDate:   in.EndDate,
		Quantity:  in.Quantity,
	}

	repo := repository.NewItemOldPriceRepository(uc.DB)
	out.ID, err = repo.Update(c, out)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	return
}

// UpdatePreservedQuantity ...
func (uc ItemOldPriceUC) UpdatePreservedQuantity(c context.Context, id string, quantity string) (out viewmodel.ItemOldPriceVM, err error) {
	out = viewmodel.ItemOldPriceVM{
		ID:           id,
		PreservedQty: quantity,
	}

	repo := repository.NewItemOldPriceRepository(uc.DB)
	out.ID, err = repo.UpdatePreservedQuantity(c, out)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	return
}

// Delete ...
func (uc ItemOldPriceUC) Delete(c context.Context, in string) (err error) {
	repo := repository.NewItemOldPriceRepository(uc.DB)
	_, err = repo.Delete(c, in)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	return
}
