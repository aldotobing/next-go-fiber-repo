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

// PromoItemLineUC ...
type PromoItemLineUC struct {
	*ContractUC
}

// BuildBody ...
func (uc PromoItemLineUC) BuildBody(res *models.PromoItemLine) {
}

// SelectAll ...
func (uc PromoItemLineUC) SelectAll(c context.Context, parameter models.PromoItemLineParameter) (res []viewmodel.PromoItemLineVM, err error) {
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(0, 0, parameter.By, parameter.Sort, models.PromoItemLineOrderBy, models.PromoItemLineOrderByrByString)

	repo := repository.NewPromoItemLineRepository(uc.DB)
	data, err := repo.SelectAll(c, parameter)

	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	var ids, priceListVersionID string
	for i := range data {
		if ids == "" {
			ids += *data[i].ItemID
		} else {
			ids += "," + *data[i].ItemID
		}
		priceListVersionID = *data[i].PriceListVersionID
	}
	latestPrice, _ := ItemUC{ContractUC: uc.ContractUC}.SelectAllV2(c, models.ItemParameter{IDs: ids, PriceListVersionId: priceListVersionID, By: "def.created_date"}, true)
	doubleChecker := make(map[string]string)
	for i := range data {
		if doubleChecker[*data[i].ID] == "" {
			doubleChecker[*data[i].ID] = "done"
			var finalPrice string
			for j := range latestPrice {
				if *latestPrice[j].ID == *data[i].ItemID {
					for k := range latestPrice[j].Uom {
						if *latestPrice[j].Uom[k].Conversion == *data[i].UomLineConversion {
							finalPrice = *latestPrice[j].Uom[k].ItemDetailsPrice
						}
					}
					break
				}
			}

			res = append(res, viewmodel.PromoItemLineVM{
				ID:                 data[i].ID,
				ItemID:             data[i].ItemID,
				UomLineConversion:  data[i].UomLineConversion,
				PromoID:            data[i].PromoID,
				PromoLineID:        data[i].PromoLineID,
				PromoName:          data[i].PromoName,
				ItemCode:           data[i].ItemCode,
				ItemName:           data[i].ItemName,
				ItemDescription:    data[i].ItemDescription,
				ItemCategoryID:     data[i].ItemCategoryID,
				ItemCategoryName:   data[i].ItemCategoryName,
				ItemPicture:        data[i].ItemPicture,
				Qty:                data[i].Qty,
				UomID:              data[i].UomID,
				UomName:            data[i].UomID,
				ItemPrice:          &finalPrice,
				PriceListVersionID: data[i].PriceListVersionID,
				GlobalMaxQty:       data[i].GlobalMaxQty,
				CustomerMaxQty:     data[i].CustomerMaxQty,
				DiscPercent:        data[i].DiscPercent,
				DiscAmount:         data[i].DiscAmount,
				MinValue:           data[i].MinValue,
				MinQty:             data[i].MinQty,
				Description:        data[i].Description,
				Multiply:           data[i].Multiply,
				MinQtyUomID:        data[i].MinQtyUomID,
				PromoType:          data[i].PromoType,
				Strata:             data[i].Strata,
				StartDate:          data[i].StartDate,
				EndDate:            data[i].EndDate,
			})
		}
	}

	return res, err
}

// FindAll ...
func (uc PromoItemLineUC) FindAll(c context.Context, parameter models.PromoItemLineParameter) (res []models.PromoItemLine, p viewmodel.PaginationVM, err error) {
	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.PromoItemLineOrderBy, models.PromoItemLineOrderByrByString)

	var count int
	repo := repository.NewPromoItemLineRepository(uc.DB)
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
func (uc PromoItemLineUC) FindByID(c context.Context, parameter models.PromoItemLineParameter) (res models.PromoItemLine, err error) {
	repo := repository.NewPromoItemLineRepository(uc.DB)
	res, err = repo.FindByID(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}
	uc.BuildBody(&res)

	return res, err
}

func (uc PromoItemLineUC) Add(c context.Context, data *requests.PromoItemLineRequest) (res models.PromoItemLineBreakDown, err error) {

	repo := repository.NewPromoItemLineRepository(uc.DB)
	// now := time.Now().UTC()
	// strNow := now.Format(time.RFC3339)
	res = models.PromoItemLineBreakDown{
		PromoLineID: &data.PromoLineID,
		ItemID:      &data.ItemID,
		UomID:       &data.UomID,
		Qty:         &data.Qty,
	}
	res.ID, err = repo.Add(c, &res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return res, err
}
