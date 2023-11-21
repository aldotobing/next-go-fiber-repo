package usecase

import (
	"context"
	"time"

	"nextbasis-service-v-0.1/db/repository"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/functioncaller"
	"nextbasis-service-v-0.1/pkg/logruslogger"
	"nextbasis-service-v-0.1/server/requests"
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

// ShoppingCartUC ...
type ShoppingCartUC struct {
	*ContractUC
}

// BuildBody ...
func (uc ShoppingCartUC) BuildBody(res *models.ShoppingCart) {
}

func (uc ShoppingCartUC) BuildHroupedBody(res *models.GroupedShoppingCart) {
}

func (uc ShoppingCartUC) BuildBonusBody(res *models.ShoppingCartItemBonus) {
}

// SelectAll ...
func (uc ShoppingCartUC) SelectAll(c context.Context, parameter models.ShoppingCartParameter) (res []models.ShoppingCart, err error) {
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(0, 0, parameter.By, parameter.Sort, models.ShoppingCartOrderBy, models.ShoppingCartOrderByrByString)

	repo := repository.NewShoppingCartRepository(uc.DB)
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
func (uc ShoppingCartUC) FindAll(c context.Context, parameter models.ShoppingCartParameter) (res []models.ShoppingCart, p viewmodel.PaginationVM, err error) {
	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.ShoppingCartOrderBy, models.ShoppingCartOrderByrByString)

	var count int
	repo := repository.NewShoppingCartRepository(uc.DB)
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
func (uc ShoppingCartUC) FindByID(c context.Context, parameter models.ShoppingCartParameter) (res models.ShoppingCart, err error) {
	repo := repository.NewShoppingCartRepository(uc.DB)
	res, err = repo.FindByID(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}
	uc.BuildBody(&res)

	return res, err
}

// Add ...
func (uc ShoppingCartUC) Add(c context.Context, data *requests.ShoppingCartRequest) (res models.ShoppingCart, err error) {

	repo := repository.NewShoppingCartRepository(uc.DB)
	now := time.Now().UTC()
	strnow := now.Format(time.RFC3339)
	res = models.ShoppingCart{
		CustomerID: &data.CustomerID,
		ItemID:     &data.ItemID,
		UomID:      &data.UomID,
		Price:      &data.Price,
		CreatedBy:  &data.CustomerID,
		CreatedAt:  &strnow,
		Qty:        &data.Qty,
		StockQty:   &data.StockQty,
		TotalPrice: &data.TotalPrice,
		OldPriceID: &data.OldPriceID,
	}
	res.ID, err = repo.Add(c, &res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return res, err
}

// Edit ,...
func (uc ShoppingCartUC) Edit(c context.Context, id string, data *requests.ShoppingCartRequest) (res models.ShoppingCart, err error) {
	repo := repository.NewShoppingCartRepository(uc.DB)
	now := time.Now().UTC()
	strnow := now.Format(time.RFC3339)
	res = models.ShoppingCart{
		ID:         &id,
		CustomerID: &data.CustomerID,
		ItemID:     &data.ItemID,
		UomID:      &data.UomID,
		Price:      &data.Price,
		ModifiedAt: &strnow,
		ModifiedBy: &data.CustomerID,
		TotalPrice: &data.TotalPrice,
	}

	res.ID, err = repo.Edit(c, &res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return res, err
}

// Delete ...
func (uc ShoppingCartUC) Delete(c context.Context, id string) (res viewmodel.ShoppingCartVM, err error) {
	// now := time.Now().UTC()
	// repo := repository.NewShoppingCartRepository(uc.DB)
	// res.ID, err = repo.Delete(c, id)
	// if err != nil {
	// 	logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
	// 	return res, err
	// }

	return res, err

}

func (uc ShoppingCartUC) MultipleEdit(c context.Context, data *[]requests.ShoppingCartRequest) (res []models.ShoppingCart, err error) {

	repo := repository.NewShoppingCartRepository(uc.DB)

	var listobjectData []models.ShoppingCart

	for _, input := range *data {

		now := time.Now().UTC()
		strnow := now.Format(time.RFC3339)
		ShoppingcartOject := models.ShoppingCart{
			ID:         &input.ID,
			CustomerID: &input.CustomerID,
			ItemID:     &input.ItemID,
			UomID:      &input.UomID,
			ModifiedBy: &input.CustomerID,
			ModifiedAt: &strnow,
			Qty:        &input.Qty,
			StockQty:   &input.StockQty,
			Price:      &input.Price,
			// TotalPrice: &input.TotalPrice,
		}
		ShoppingcartOject.ID, err = repo.Edit(c, &ShoppingcartOject)
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
			return res, err
		}
		var objectData models.ShoppingCart
		objectData = ShoppingcartOject
		listobjectData = append(listobjectData, objectData)
	}

	res = listobjectData

	return res, err
}

func (uc ShoppingCartUC) MultipleEditByCartID(c context.Context, data *[]requests.ShoppingCartRequest) (res []models.ShoppingCart, err error) {

	repo := repository.NewShoppingCartRepository(uc.DB)

	var listobjectData []models.ShoppingCart

	for _, input := range *data {

		now := time.Now().UTC()
		strnow := now.Format(time.RFC3339)
		ShoppingcartOject := models.ShoppingCart{
			ID:         &input.ID,
			CustomerID: &input.CustomerID,
			ItemID:     &input.ItemID,
			UomID:      &input.UomID,
			ModifiedBy: &input.CustomerID,
			ModifiedAt: &strnow,
			Qty:        &input.Qty,
			StockQty:   &input.StockQty,
			Price:      &input.Price,
			// TotalPrice: &input.TotalPrice,
		}
		ShoppingcartOject.ID, err = repo.EditQuantity(c, &ShoppingcartOject)
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
			return res, err
		}
		var objectData models.ShoppingCart
		objectData = ShoppingcartOject
		listobjectData = append(listobjectData, objectData)
	}

	res = listobjectData

	return res, err
}

func (uc ShoppingCartUC) MultipleDelete(c context.Context, data *[]requests.ShoppingCartRequest) (res []models.ShoppingCart, err error) {

	repo := repository.NewShoppingCartRepository(uc.DB)

	var listobjectData []models.ShoppingCart

	for _, input := range *data {

		ShoppingcartOject := models.ShoppingCart{
			ID: &input.ID,
		}
		ShoppingcartOject.ID, err = repo.Delete(c, input.ID)
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
			return res, err
		}
		var objectData models.ShoppingCart
		objectData = ShoppingcartOject
		listobjectData = append(listobjectData, objectData)
	}

	res = listobjectData

	return res, err
}

func (uc ShoppingCartUC) SelectAllForGroup(c context.Context, parameter models.ShoppingCartParameter) (res []models.GroupedShoppingCart, err error) {
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(0, 0, parameter.By, parameter.Sort, models.ShoppingCartOrderBy, models.ShoppingCartOrderByrByString)

	repo := repository.NewShoppingCartRepository(uc.DB)
	res, err = repo.SelectAllForGroup(c, parameter)

	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	for i := range res {
		uc.BuildHroupedBody(&res[i])
	}

	return res, err
}

func (uc ShoppingCartUC) SelectAllBonus(c context.Context, parameter models.ShoppingCartParameter) (res []models.ShoppingCartItemBonus, err error) {

	repo := repository.NewShoppingCartRepository(uc.DB)
	res, err = repo.SelectAllBonus(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	for i := range res {
		uc.BuildBonusBody(&res[i])
	}

	return res, err
}
