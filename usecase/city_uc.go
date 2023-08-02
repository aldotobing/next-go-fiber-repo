package usecase

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
	"nextbasis-service-v-0.1/db/repository"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/functioncaller"
	"nextbasis-service-v-0.1/pkg/logruslogger"
	"nextbasis-service-v-0.1/server/requests"
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

// CityUC ...
type CityUC struct {
	*ContractUC
}

// BuildBody ...
func (uc CityUC) BuildBody(res *models.City) {
}

// SelectAll ...
// func (uc CityUC) SelectAll(c context.Context, parameter models.CityParameter) (res []models.City, err error) {
// 	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(0, 0, parameter.By, parameter.Sort, models.CityOrderBy, models.CityOrderByrByString)

// 	repo := repository.NewCityRepository(uc.DB)
// 	res, err = repo.SelectAll(c, parameter)

// 	if err != nil {
// 		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
// 		return res, err
// 	}

// 	for i := range res {
// 		uc.BuildBody(&res[i])
// 	}

// 	return res, err
// }

func (uc CityUC) SelectAll(c context.Context, parameter models.CityParameter) (res []models.City, err error) {
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(0, 0, parameter.By, parameter.Sort, models.CityOrderBy, models.CityOrderByrByString)

	// Redis integration
	cacheKey := "cities:" + parameter.IDs // cacheKey can be different based on your needs

	val, err := uc.RedisClient.Client.Get(cacheKey).Result()
	if err == nil {
		// If cache exists
		err = json.Unmarshal([]byte(val), &res)
		if err == nil {
			return res, nil
		}
	}

	repo := repository.NewCityRepository(uc.DB)
	res, err = repo.SelectAll(c, parameter)

	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	// Save result into Redis
	jsonData, err := json.Marshal(res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "json_marshal", uc.ReqID)
	} else {
		err = uc.RedisClient.Client.Set(cacheKey, jsonData, time.Hour).Err()
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "redis_set", uc.ReqID)
		}
	}

	for i := range res {
		uc.BuildBody(&res[i])
	}

	return res, err
}

// FindAll ...
func (uc CityUC) FindAll(c context.Context, parameter models.CityParameter) (res []models.City, p viewmodel.PaginationVM, err error) {
	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.CityOrderBy, models.CityOrderByrByString)

	var count int
	repo := repository.NewCityRepository(uc.DB)
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
// func (uc CityUC) FindByID(c context.Context, parameter models.CityParameter) (res models.City, err error) {
// 	repo := repository.NewCityRepository(uc.DB)
// 	res, err = repo.FindByID(c, parameter)
// 	if err != nil {
// 		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
// 		return res, err
// 	}
// 	uc.BuildBody(&res)

// 	return res, err
// }

func (uc CityUC) FindByID(c context.Context, parameter models.CityParameter) (res models.City, err error) {
	// Try to fetch the data from Redis first
	cacheKey := "city:" + parameter.ID // Directly use parameter.ID
	err = uc.RedisClient.Client.Get(cacheKey).Scan(&res)
	if err == nil {
		// Data was found in Redis, return it
		return res, nil
	} else if err != redis.Nil {
		// An error occurred that wasn't just "key doesn't exist"
		return res, err
	}

	// Data was not found in Redis, fetch from DB instead
	repo := repository.NewCityRepository(uc.DB)
	res, err = repo.FindByID(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}
	uc.BuildBody(&res)

	// Cache the data in Redis, ignore the result
	uc.StoreToRedis(cacheKey, res)

	return res, err
}

// Add ...
func (uc CityUC) Add(c context.Context, data *requests.CityRequest) (res models.City, err error) {

	repo := repository.NewCityRepository(uc.DB)
	// now := time.Now().UTC()
	// strnow := now.Format(time.RFC3339)
	res = models.City{
		Name: &data.Name,
	}
	res.ID, err = repo.Add(c, &res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return res, err
}

// Edit ,...
func (uc CityUC) Edit(c context.Context, id string, data *requests.CityRequest) (res models.City, err error) {
	repo := repository.NewCityRepository(uc.DB)
	// now := time.Now().UTC()
	// strnow := now.Format(time.RFC3339)
	res = models.City{
		ID: &id,
		// ProvinceID: &data.ProvinceID,
		Name: &data.Name,
		// Longitude:  &data.Lat,
		// Latitude:   &data.Long,
		// CreatedAt:  &strnow,
		// UpdatedAt:  &strnow,
		// CreatedBy:  &data.CreatedBy,
		// UpdatedBy:  &data.UpdatedBy,
	}

	res.ID, err = repo.Edit(c, &res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}
	uc.RedisClient.Client.Del("city:" + id)

	return res, err
}

// Delete ...
func (uc CityUC) Delete(c context.Context, id string) (res viewmodel.CityVM, err error) {
	now := time.Now().UTC()
	repo := repository.NewCityRepository(uc.DB)
	res.ID, err = repo.Delete(c, id, now)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return res, err

}

func (uc CityUC) AddAll(c context.Context, data *[]requests.MpCityDataBreakDownRequest) (res []models.MpCityDataBreakDown, err error) {

	repo := repository.NewCityRepository(uc.DB)

	for _, input := range *data {

		// lat := fmt.Sprintf("%f", input.LatCity)
		// longi := fmt.Sprintf("%f", input.LongCity)
		cityOject := models.MpCityDataBreakDown{
			Name:       &input.Name,
			ProvinceID: &input.ProvinceID,
			OldID:      &input.OldID,
			LatCity:    &input.LatCity,
			LongCity:   &input.LongCity,
		}
		cityOject.ID, err = repo.AddDataBreakDown(c, &cityOject)
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
			return res, err
		}
		res = append(res, cityOject)
	}

	return res, err
}
