package usecase

import (
	"context"
	"fmt"

	"nextbasis-service-v-0.1/db/repository"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/firestore"
	"nextbasis-service-v-0.1/pkg/functioncaller"
	"nextbasis-service-v-0.1/pkg/logruslogger"
)

type FireBaseUIDUC struct {
	*ContractUC
}

// BuildBody ...
func (uc FireBaseUIDUC) BuildBody(res *models.FireStoreUser) {
}

func (uc FireBaseUIDUC) SyncData(c context.Context) (res []models.FireStoreUser, err error) {

	firestoreModel := firestore.NewFireStoreModel(uc.Firestore)

	res, err = firestoreModel.GetData(c, "users")

	for _, data := range res {
		userrepo := repository.NewUserAccountRepository(uc.DB)
		_, err := userrepo.FIreStoreIDSync(c, &models.UserAccount{ID: data.ID, FireStoreUID: &data.UID})
		if err == nil {
			firestoreModel.UpdateData(c, data, "users")
		}
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	return res, err
}

func (uc FireBaseUIDUC) SelectAll(c context.Context, parameter models.FireStoreUserParameter) (res []models.FireStoreUser, err error) {
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(0, 0, parameter.By, parameter.Sort, models.FireStoreUserOrderBy, models.FireStoreUserOrderByrByString)

	repo := repository.NewFireStoreUserRepository(uc.DB)
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
