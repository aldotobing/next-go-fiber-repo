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

// VideoPromoteUC ...
type VideoPromoteUC struct {
	*ContractUC
}

// BuildBody ...
func (uc VideoPromoteUC) BuildBody(res *models.VideoPromote) {
}

// SelectAll ...
func (uc VideoPromoteUC) SelectAll(c context.Context, parameter models.VideoPromoteParameter) (res []models.VideoPromote, err error) {
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(0, 0, parameter.By, parameter.Sort, models.VideoPromoteOrderBy, models.VideoPromoteOrderByrByString)

	repo := repository.NewVideoPromoteRepository(uc.DB)
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
func (uc VideoPromoteUC) FindAll(c context.Context, parameter models.VideoPromoteParameter) (res []models.VideoPromote, p viewmodel.PaginationVM, err error) {
	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.VideoPromoteOrderBy, models.VideoPromoteOrderByrByString)

	var count int
	repo := repository.NewVideoPromoteRepository(uc.DB)
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

// Add ...
func (uc VideoPromoteUC) Add(c context.Context, data *requests.VideoPromoteRequest) (res models.VideoPromote, err error) {

	repo := repository.NewVideoPromoteRepository(uc.DB)
	// now := time.Now().UTC()
	// strnow := now.Format(time.RFC3339)
	res = models.VideoPromote{
		Title:       &data.Title,
		Description: &data.Description,
		StartDate:   &data.StartDate,
		EndDate:     &data.EndDate,
		Active:      &data.Active,
		Url:         &data.Url,
	}
	res.ID, err = repo.Add(c, &res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return res, err
}
