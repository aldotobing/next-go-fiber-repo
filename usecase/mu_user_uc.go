package usecase

import (
	"context"

	"nextbasis-service-v-0.1/db/repository"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/functioncaller"
	"nextbasis-service-v-0.1/pkg/logruslogger"
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

type MuUserUC struct {
	*ContractUC
}

// BuildBody ...
func (uc MuUserUC) BuildBody(res *models.MuUser) {
}

func (uc MuUserUC) GenerateToken(c context.Context, id string) (res viewmodel.JwtVM, err error) {

	payload := map[string]interface{}{
		"user_id": id,
	}
	jwtUc := JwtUC{ContractUC: uc.ContractUC}
	err = jwtUc.GenerateToken(c, payload, &res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "generate_token")
		return res, err
	}
	return res, err
}

func (uc MuUserUC) FindByRefferalCode(c context.Context, parameter models.MuUserParameter) (res models.MuUser, err error) {
	repo := repository.NewMuUserRepository(uc.DB)
	res, err = repo.FindByRefferalCode(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}
	uc.BuildBody(&res)

	return res, err
}

func (uc MuUserUC) FindByEmail(c context.Context, parameter models.MuUserParameter) (res models.MuUser, err error) {
	repo := repository.NewMuUserRepository(uc.DB)
	res, err = repo.FindByEmail(c, parameter)
	if err != nil {

		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}
	uc.BuildBody(&res)

	return res, err
}

func (uc MuUserUC) FindByID(c context.Context, parameter models.MuUserParameter) (res models.MuUser, err error) {
	repo := repository.NewMuUserRepository(uc.DB)
	res, err = repo.FindByID(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	uc.BuildBody(&res)
	return res, err
}

func (uc MuUserUC) FindAll(c context.Context, parameter models.MuUserParameter) (res []models.MuUser, p viewmodel.PaginationVM, err error) {
	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.MuUserOrderBy, models.MuUserOrderByrByString)

	var count int
	repo := repository.NewMuUserRepository(uc.DB)
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
