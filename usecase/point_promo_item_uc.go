package usecase

import (
	"context"
	"nextbasis-service-v-0.1/db/repository"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/functioncaller"
	"nextbasis-service-v-0.1/pkg/logruslogger"
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

// PointPromoItemUC ...
type PointPromoItemUC struct {
	*ContractUC
}

// BuildBody ...
func (uc PointPromoItemUC) BuildBody(data *models.PointPromo, res *viewmodel.PointPromoVM) {
}

// AddBulk ...
func (uc PointPromoItemUC) AddBulk(c context.Context, pointPromoID string, items []string) (err error) {
	repo := repository.NewPointPromoItemRepository(uc.DB)
	err = repo.AddBulk(c, pointPromoID, items)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	return
}

// Delete ...
func (uc PointPromoItemUC) Delete(c context.Context, pointPromoID string) (err error) {
	repo := repository.NewPointPromoItemRepository(uc.DB)
	err = repo.Delete(c, pointPromoID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	return
}
