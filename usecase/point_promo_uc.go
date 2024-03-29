package usecase

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"nextbasis-service-v-0.1/db/repository"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/functioncaller"
	"nextbasis-service-v-0.1/pkg/logruslogger"
	"nextbasis-service-v-0.1/server/requests"
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

// PointPromoUC ...
type PointPromoUC struct {
	*ContractUC
}

// BuildBody ...
func (uc PointPromoUC) BuildBody(data *models.PointPromo, res *viewmodel.PointPromoVM) {
	res.ID = data.ID
	startDate, _ := time.Parse("2006-01-02T15:04:05.999999999Z", data.StartDate)
	res.StartDate = startDate.Format("2006-01-02")
	endDate, _ := time.Parse("2006-01-02T15:04:05.999999999Z", data.EndDate)
	res.EndDate = endDate.Format("2006-01-02")
	res.CreatedAt = data.CreatedAt
	res.UpdatedAt = data.UpdatedAt.String
	res.DeletedAt = data.DeletedAt.String
	res.Multiplicator = data.Multiplicator
	res.PointConversion = data.PointConversion.String
	res.QuantityConversion = data.QuantityConversion.String
	res.PromoType = data.PromoType.String

	_ = json.Unmarshal([]byte(data.Strata.String), &res.Strata)

	if res.Strata == nil {
		res.Strata = make([]viewmodel.PointPromoStrataVM, 0)
	}

	var items []viewmodel.PointPromoItemVM
	additional := strings.Split(data.Items, "|")
	if len(additional) > 0 && additional[0] != "" {
		// Find Lowest Price and lowest conversion
		for _, addDatum := range additional {
			perAddDatum := strings.Split(addDatum, "#sep#")
			items = append(items, viewmodel.PointPromoItemVM{
				ID:       perAddDatum[0],
				ItemName: perAddDatum[1],
			})
		}
	}
	if items == nil {
		items = make([]viewmodel.PointPromoItemVM, 0)
	}
	res.Items = items
}

// FindAll ...
func (uc PointPromoUC) FindAll(c context.Context, parameter models.PointPromoParameter) (out []viewmodel.PointPromoVM, p viewmodel.PaginationVM, err error) {
	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.PointPromoOrderBy, models.PointPromoOrderByrByString)

	repo := repository.NewPointPromoRepository(uc.DB)
	data, count, err := repo.FindAll(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	p = uc.setPaginationResponse(parameter.Page, parameter.Limit, count)

	for _, datum := range data {
		var temp viewmodel.PointPromoVM
		uc.BuildBody(&datum, &temp)

		out = append(out, temp)
	}

	if out == nil {
		out = make([]viewmodel.PointPromoVM, 0)
	}

	return
}

// SelectAll ...
func (uc PointPromoUC) SelectAll(c context.Context, parameter models.PointPromoParameter) (out []viewmodel.PointPromoVM, err error) {
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.PointPromoOrderBy, models.PointPromoOrderByrByString)

	repo := repository.NewPointPromoRepository(uc.DB)
	data, err := repo.SelectAll(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	for _, datum := range data {
		var temp viewmodel.PointPromoVM
		uc.BuildBody(&datum, &temp)

		out = append(out, temp)
	}

	if out == nil {
		out = make([]viewmodel.PointPromoVM, 0)
	}

	return
}

// FindByID ...
func (uc PointPromoUC) FindByID(c context.Context, parameter models.PointPromoParameter) (out viewmodel.PointPromoVM, err error) {
	repo := repository.NewPointPromoRepository(uc.DB)
	data, err := repo.FindByID(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	uc.BuildBody(&data, &out)

	return
}

// Add ...
func (uc PointPromoUC) Add(c context.Context, in requests.PointPromoRequest) (out viewmodel.PointPromoVM, err error) {
	var strata []viewmodel.PointPromoStrataVM
	for _, datum := range in.Strata {
		strata = append(strata, viewmodel.PointPromoStrataVM(datum))
	}

	var items []viewmodel.PointPromoItemVM
	for _, datum := range in.Items {
		items = append(items, viewmodel.PointPromoItemVM{
			ID:        datum,
			ItemName:  "",
			CreatedAt: "",
			UpdatedAt: "",
			DeletedAt: "",
		})
	}
	out = viewmodel.PointPromoVM{
		StartDate:          in.StartDate,
		EndDate:            in.EndDate,
		Multiplicator:      in.Multiplicator,
		PointConversion:    in.PointConversion,
		QuantityConversion: in.QuantityConversion,
		PromoType:          in.PromoType,
		Strata:             strata,
		Items:              items,
	}

	repo := repository.NewPointPromoRepository(uc.DB)
	out.ID, err = repo.Add(c, out)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query add", c.Value("requestid"))
		return
	}

	err = PointPromoItemUC{uc.ContractUC}.AddBulk(c, out.ID, in.Items)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "add point_promo_item", c.Value("requestid"))
		return
	}

	return
}

// Update ...
func (uc PointPromoUC) Update(c context.Context, id string, in requests.PointPromoRequest) (out viewmodel.PointPromoVM, err error) {
	var strata []viewmodel.PointPromoStrataVM
	for _, datum := range in.Strata {
		strata = append(strata, viewmodel.PointPromoStrataVM(datum))
	}
	var items []viewmodel.PointPromoItemVM
	for _, datum := range in.Items {
		items = append(items, viewmodel.PointPromoItemVM{
			ID:        datum,
			ItemName:  "",
			CreatedAt: "",
			UpdatedAt: "",
			DeletedAt: "",
		})
	}
	out = viewmodel.PointPromoVM{
		ID:                 id,
		StartDate:          in.StartDate,
		EndDate:            in.EndDate,
		Multiplicator:      in.Multiplicator,
		PointConversion:    in.PointConversion,
		QuantityConversion: in.QuantityConversion,
		PromoType:          in.PromoType,
		Strata:             strata,
		Items:              items,
	}

	repo := repository.NewPointPromoRepository(uc.DB)
	out.ID, err = repo.Update(c, out)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	err = PointPromoItemUC{uc.ContractUC}.Delete(c, out.ID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "add point_promo_item", c.Value("requestid"))
		return
	}

	err = PointPromoItemUC{uc.ContractUC}.AddBulk(c, out.ID, in.Items)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "add point_promo_item", c.Value("requestid"))
		return
	}

	return
}

// Delete ...
func (uc PointPromoUC) Delete(c context.Context, in string) (err error) {
	repo := repository.NewPointPromoRepository(uc.DB)
	_, err = repo.Delete(c, in)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	err = PointPromoItemUC{uc.ContractUC}.Delete(c, in)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "add point_promo_item", c.Value("requestid"))
		return
	}

	return
}
