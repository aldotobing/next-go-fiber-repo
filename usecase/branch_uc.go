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

// BranchUC ...
type BranchUC struct {
	*ContractUC
}

// BuildBody ...
func (uc BranchUC) BuildBody(res *models.Branch) {
}

// SelectAll ...
// func (uc BranchUC) SelectAll(c context.Context, parameter models.BranchParameter) (res []models.Branch, err error) {
// 	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(0, 0, parameter.By, parameter.Sort, models.BranchOrderBy, models.BranchOrderByrByString)

// 	cacheKey := fmt.Sprintf("branches:user:%d:page:%d:limit:%d:by:%s:sort:%s", parameter.UserID, parameter.Page, parameter.Limit, parameter.By, parameter.Sort)

// 	// Try to fetch data from cache first
// 	cachedData, cacheErr := uc.RedisClient.Get(cacheKey)
// 	if cacheErr == nil && string(cachedData) != "" {
// 		err := json.Unmarshal(cachedData, &res)
// 		if err == nil {
// 			//fmt.Println("Data loaded from cache")
// 			return res, nil
// 		}
// 	}

// 	// If not in cache or cache data was invalid, hit the database
// 	repo := repository.NewBranchRepository(uc.DB)
// 	res, err = repo.SelectAll(c, parameter)
// 	if err != nil {
// 		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
// 		return nil, err
// 	}

// 	for i := range res {
// 		uc.BuildBody(&res[i])
// 	}

// 	// Cache the fetched result
// 	jsonData, marshalErr := json.Marshal(res)
// 	if marshalErr == nil {
// 		uc.RedisClient.Set(cacheKey, jsonData, time.Minute*30)
// 	}

// 	//fmt.Println("Data loaded from database")
// 	return res, nil
// }

func (uc BranchUC) SelectAll(c context.Context, parameter models.BranchParameter) (res []models.Branch, err error) {
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(0, 0, parameter.By, parameter.Sort, models.BranchOrderBy, models.BranchOrderByrByString)

	repo := repository.NewBranchRepository(uc.DB)
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
// func (uc BranchUC) FindAll(c context.Context, parameter models.BranchParameter) (res []models.Branch, p viewmodel.PaginationVM, err error) {

// 	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.BranchOrderBy, models.BranchOrderByrByString)

// 	cacheKey := fmt.Sprintf("branches:user:%d:page:%d:limit:%d:by:%s:sort:%s", parameter.UserID, parameter.Page, parameter.Limit, parameter.By, parameter.Sort)

// 	// Try to fetch data from cache first
// 	cachedData, cacheErr := uc.RedisClient.Get(cacheKey)
// 	if cacheErr == nil && string(cachedData) != "" {
// 		err := json.Unmarshal(cachedData, &res)
// 		if err == nil {
// 			fmt.Println("Data loaded from cache")
// 			// Assuming you also want to cache the PaginationVM.
// 			// You might need to unmarshal it separately from the cachedData if you do.
// 			return res, p, nil
// 		}
// 	}

// 	repo := repository.NewBranchRepository(uc.DB)
// 	res, count, err := repo.FindAll(c, parameter)
// 	if err != nil {
// 		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
// 		return res, p, err
// 	}

// 	p = uc.setPaginationResponse(parameter.Page, parameter.Limit, count)
// 	for i := range res {
// 		uc.BuildBody(&res[i])
// 	}

// 	// Cache the fetched result
// 	jsonData, marshalErr := json.Marshal(res)
// 	if marshalErr == nil {
// 		uc.RedisClient.Set(cacheKey, jsonData, time.Minute*30) // Cache for 30 minutes
// 	}

// 	fmt.Println("Data loaded from database")
// 	return res, p, nil
// }

func (uc BranchUC) FindAll(c context.Context, parameter models.BranchParameter) (res []models.Branch, p viewmodel.PaginationVM, err error) {
	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.BranchOrderBy, models.BranchOrderByrByString)

	var count int
	repo := repository.NewBranchRepository(uc.DB)
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
func (uc BranchUC) FindByID(c context.Context, parameter models.BranchParameter) (res models.Branch, err error) {
	repo := repository.NewBranchRepository(uc.DB)
	res, err = repo.FindByID(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}
	uc.BuildBody(&res)

	return res, err
}

func (uc BranchUC) Update(c context.Context, id string, in *requests.BranchRequest) (res models.Branch, err error) {
	res = models.Branch{
		ID:         &id,
		PICPhoneNo: &in.PICPhoneNo,
		PICName:    &in.PICName,
	}

	repo := repository.NewBranchRepository(uc.DB)
	_, err = repo.Update(c, res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return
}

func (uc BranchUC) GenerateAllUser(c context.Context, id string) (res models.Branch, err error) {
	res = models.Branch{
		ID: &id,
	}

	repo := repository.NewBranchRepository(uc.DB)
	_, err = repo.GenerateAllUser(c, res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return
}
