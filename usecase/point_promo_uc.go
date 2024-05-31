package usecase

import (
	"context"
	"encoding/json"
	"mime/multipart"
	"strconv"
	"strings"
	"time"

	"nextbasis-service-v-0.1/config"
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
				ID:         perAddDatum[0],
				ItemName:   perAddDatum[1],
				Image:      models.ItemImagePath + perAddDatum[2],
				UomID:      perAddDatum[3],
				UomName:    perAddDatum[4],
				Convertion: perAddDatum[5],
				Quantity:   perAddDatum[6],
			})
		}
	}
	if items == nil {
		items = make([]viewmodel.PointPromoItemVM, 0)
	}
	res.Items = items
	res.Image = data.Image.String
	res.Title = data.Title.String
	res.Description = data.Description.String
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

// EligiblePoint ...
func (uc PointPromoUC) EligiblePoint(c context.Context, cartList string) (out string, err error) {
	var cartListQuery string
	cartListSplit := strings.Split(cartList, ",")
	if len(cartListSplit) > 0 {
		for _, datum := range cartListSplit {
			if cartListQuery != "" {
				cartListQuery += `,'` + datum + `'`
			} else {
				cartListQuery += `'` + datum + `'`
			}
		}
	}

	cart, err := ShoppingCartUC{ContractUC: uc.ContractUC}.SelectAll(c, models.ShoppingCartParameter{
		ListID: cartListQuery,
		Sort:   "asc",
		By:     "def.id",
	})
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "select_shopping_cart", c.Value("requestid"))
		return
	}

	pointPromo, err := uc.SelectAll(c, models.PointPromoParameter{
		Now:  true,
		Sort: "asc",
		By:   "def.id",
	})
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "select_point_promo", c.Value("requestid"))
		return
	}

	var pointEligible float64
	for _, pointPromoData := range pointPromo {
		switch pointPromoData.PromoType {
		case models.PromoTypePoint:
			var multiplicator bool
			for _, itemPromo := range pointPromoData.Items {
				for _, itemCart := range cart {
					if multiplicator {
						continue
					}
					if itemCart.ItemID != nil && itemPromo.ID == *itemCart.ItemID {
						cartStock, _ := strconv.ParseFloat(*itemCart.StockQty, 64)
						itemCartTotalQty := cartStock

						itemPromoConvertion, _ := strconv.ParseFloat(itemPromo.Convertion, 64)
						itemPromoQty, _ := strconv.ParseFloat(itemPromo.Quantity, 64)
						itemPromoTotalQty := itemPromoConvertion * itemPromoQty

						eligibleMultiply := itemCartTotalQty / itemPromoTotalQty
						if eligibleMultiply >= 1 {
							if !pointPromoData.Multiplicator {
								eligibleMultiply = 1
								multiplicator = true
							}
							pointConvertion, _ := strconv.ParseFloat(pointPromoData.PointConversion, 64)
							pointGet := pointConvertion * float64(int(eligibleMultiply))
							pointEligible += pointGet
						}

					}
				}
			}
		case models.PromoTypeStrata:
			var totalPrice float64
			for _, itemPromo := range pointPromoData.Items {
				for _, itemCart := range cart {
					if itemCart.ItemID != nil && itemPromo.ID == *itemCart.ItemID {
						price, _ := strconv.ParseFloat(*itemCart.Price, 64)
						qty, _ := strconv.ParseFloat(*itemCart.Qty, 64)
						totalPrice += price * qty
					}
				}
			}

			var flag bool
			for x, strata := range pointPromoData.Strata {
				from, _ := strconv.ParseFloat(strata.From, 64)
				to, _ := strconv.ParseFloat(strata.To, 64)
				if totalPrice >= from && totalPrice <= to {
					getPoint, _ := strconv.ParseFloat(strata.Point, 64)

					pointEligible += getPoint
					flag = true
				} else if len(pointPromoData.Strata)-1 == x && !flag && totalPrice > to {
					getPoint, _ := strconv.ParseFloat(strata.Point, 64)

					pointEligible += getPoint
				}
			}
		case models.PromoTypeStrataTotal:
			for _, itemPromo := range pointPromoData.Items {
				var totalItem float64

				for _, itemCart := range cart {
					if itemCart.ItemID != nil && itemPromo.ID == *itemCart.ItemID {
						stockQty, _ := strconv.ParseFloat(*itemCart.StockQty, 64)
						totalItem += stockQty
					}
				}

				var flag bool
				for x, strata := range pointPromoData.Strata {
					stockQty, _ := strconv.ParseFloat(strata.StockQty, 64)
					from, _ := strconv.ParseFloat(strata.From, 64)
					to, _ := strconv.ParseFloat(strata.To, 64)
					if totalItem >= from*stockQty && totalItem <= to*stockQty {
						getPoint, _ := strconv.ParseFloat(strata.Point, 64)

						pointEligible += getPoint
						flag = true
					} else if len(pointPromoData.Strata)-1 == x && !flag && totalItem > to*stockQty {
						getPoint, _ := strconv.ParseFloat(strata.Point, 64)

						pointEligible += getPoint
					}
				}
			}

		case models.PromoTypeStrataPerUOM:
			var totalItem float64
			for _, itemPromo := range pointPromoData.Items {
				for _, itemCart := range cart {
					if itemCart.ItemID != nil && itemPromo.ID == *itemCart.ItemID {
						stockQty, _ := strconv.ParseFloat(*itemCart.StockQty, 64)
						totalItem += stockQty
					}
				}
			}
			var flag bool
			for x, strata := range pointPromoData.Strata {
				stockQty, _ := strconv.ParseFloat(strata.StockQty, 64)
				from, _ := strconv.ParseFloat(strata.From, 64)
				to, _ := strconv.ParseFloat(strata.To, 64)
				if totalItem >= from*stockQty && totalItem <= to*stockQty {
					getPoint, _ := strconv.ParseFloat(strata.Point, 64)

					pointEligible += getPoint * float64(int((totalItem / stockQty)))
					flag = true
				} else if len(pointPromoData.Strata)-1 == x && !flag && totalItem > to*stockQty {
					getPoint, _ := strconv.ParseFloat(strata.Point, 64)

					pointEligible += getPoint * float64(int((totalItem / stockQty)))
				}
			}
		}
	}

	out = strconv.FormatFloat(pointEligible, 'f', 0, 64)

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
			ID:         datum.ItemID,
			ItemName:   "",
			UomID:      datum.UomID,
			UomName:    datum.UomName,
			Convertion: datum.Convertion,
			Quantity:   datum.Quantity,
			CreatedAt:  "",
			UpdatedAt:  "",
			DeletedAt:  "",
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
		Image:              in.Image,
		Title:              in.Title,
		Description:        in.Description,
	}

	repo := repository.NewPointPromoRepository(uc.DB)
	out.ID, err = repo.Add(c, out)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query add", c.Value("requestid"))
		return
	}

	err = PointPromoItemUC{uc.ContractUC}.AddBulk(c, out.ID, out.Items)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "add_point_promo_item", c.Value("requestid"))
		return
	}

	return
}

// AddPhoto ...
func (uc PointPromoUC) AddPhoto(c context.Context, image *multipart.FileHeader) (out string, err error) {
	awsUc := AwsUC{ContractUC: uc.ContractUC}
	awsUc.AWSS3.Directory = "image/point_promo"
	imgBannerFile, err := awsUc.Upload("image/point_promo", image)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "upload_file", c.Value("requestid"))
		return
	}
	out = config.ImagePath + imgBannerFile.FilePath

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
			ID:         datum.ItemID,
			ItemName:   "",
			UomID:      datum.UomID,
			UomName:    datum.UomName,
			Convertion: datum.Convertion,
			Quantity:   datum.Quantity,
			CreatedAt:  "",
			UpdatedAt:  "",
			DeletedAt:  "",
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
		Image:              in.Image,
		Title:              in.Title,
		Description:        in.Description,
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

	err = PointPromoItemUC{uc.ContractUC}.AddBulk(c, out.ID, out.Items)
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
