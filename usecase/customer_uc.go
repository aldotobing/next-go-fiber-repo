package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"strings"
	"time"

	"nextbasis-service-v-0.1/db/repository"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/functioncaller"
	"nextbasis-service-v-0.1/pkg/logruslogger"
	"nextbasis-service-v-0.1/server/requests"
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

// CustomerUC ...
type CustomerUC struct {
	*ContractUC
}

// BuildBody ...
func (uc CustomerUC) BuildBody(res *models.Customer) {
}

const CustomerCacheKey = "customer:"

// SelectAll ...
// func (uc CustomerUC) SelectAll(c context.Context, parameter models.CustomerParameter) (res []models.Customer, err error) {
// 	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(0, 0, parameter.By, parameter.Sort, models.CustomerOrderBy, models.CustomerOrderByrByString)

// 	repo := repository.NewCustomerRepository(uc.DB)
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

func (uc CustomerUC) SelectAll(c context.Context, parameter models.CustomerParameter) (res []models.Customer, err error) {

	// Define the cache key
	cacheKey := CustomerCacheKey + parameter.By + ":" + parameter.Sort + ":" + parameter.BranchID + ":" +
		parameter.CustomerTypeId + ":" + parameter.RegionID + ":" + parameter.RegionGroupID + ":" + parameter.CustomerLevelId +
		":" + parameter.CustomerCodes + ":" + parameter.CustomerReligion

	// Try to get data from Redis cache first
	err = uc.RedisClient.GetFromRedis(cacheKey, &res)
	if err == nil {
		fmt.Println("from redis : ", cacheKey)
		// If the data exists in the cache, return it
		return res, nil
	}

	// If the data does not exist in the cache, fetch it from the DB
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(0, 0, parameter.By, parameter.Sort, models.CustomerOrderBy, models.CustomerOrderByrByString)

	repo := repository.NewCustomerRepository(uc.DB)
	res, err = repo.SelectAll(c, parameter)

	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	for i := range res {
		uc.BuildBody(&res[i])
	}

	// Cache the result in Redis with an expiration time of 1 hour
	uc.RedisClient.StoreToRedistWithExpired(cacheKey, res, "1h")

	return res, err
}

// FindAll ...
func (uc CustomerUC) FindAll(c context.Context, parameter models.CustomerParameter) (res []models.Customer, p viewmodel.PaginationVM, err error) {
	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.CustomerOrderBy, models.CustomerOrderByrByString)

	var count int
	repo := repository.NewCustomerRepository(uc.DB)
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
func (uc CustomerUC) FindByID(c context.Context, parameter models.CustomerParameter) (res models.Customer, err error) {

	repo := repository.NewCustomerRepository(uc.DB)
	res, err = repo.FindByID(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}
	uc.BuildBody(&res)

	return res, err
}

// func (uc CustomerUC) FindByID(c context.Context, parameter models.CustomerParameter) (res models.Customer, err error) {
// 	// Define cache key
// 	cacheKey := CustomerCacheKey + parameter.ID

// 	// Try to fetch the data from Redis first
// 	val, err := uc.RedisClient.Client.Get(cacheKey).Result()
// 	if err != nil && err != redis.Nil {
// 		// If there is an error other than "key does not exist", return error
// 		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "redis_get", uc.ReqID)
// 		return res, err
// 	}
// 	if err == nil {
// 		// If cache exists
// 		err = json.Unmarshal([]byte(val), &res)
// 		if err != nil {
// 			// If there is an error in unmarshaling, log and proceed to fetch data from repository
// 			logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "json_unmarshal", uc.ReqID)
// 		} else {
// 			// If unmarshaling successful, return result
// 			return res, nil
// 		}
// 	}

// 	repo := repository.NewCustomerRepository(uc.DB)
// 	res, err = repo.FindByID(c, parameter)
// 	if err != nil {
// 		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
// 		return res, err
// 	}
// 	uc.BuildBody(&res)

// 	// Save result into Redis
// 	jsonData, err := json.Marshal(res)
// 	if err != nil {
// 		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "json_marshal", uc.ReqID)
// 	} else {
// 		err = uc.RedisClient.Client.Set(cacheKey, jsonData, time.Hour).Err()
// 		if err != nil {
// 			logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "redis_set", uc.ReqID)
// 		}
// 	}

// 	return res, err
// }

// Edit ,...
func (uc CustomerUC) Edit(c context.Context, id string, data *requests.CustomerRequest, imgProfile *multipart.FileHeader, imgKtp *multipart.FileHeader) (res models.Customer, err error) {

	// Invalidate the cache before update
	cacheKey := CustomerCacheKey + id
	err = uc.RedisClient.Client.Del(cacheKey).Err()
	if err != nil {
		// Log error
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "redis_del", uc.ReqID)
		return res, err
	}

	currentObjectUc, err := uc.FindByID(c, models.CustomerParameter{ID: id})
	ctx := "FileUC.Upload"
	awsUc := AwsUC{ContractUC: uc.ContractUC}

	var strImgprofile = ""
	var strImgktp = ""
	if currentObjectUc.CustomerPhotoKtp != nil && *currentObjectUc.CustomerPhotoKtp != "" {
		strImgktp = strings.ReplaceAll(*currentObjectUc.CustomerPhotoKtp, models.CustomerImagePath, "")
	}
	if currentObjectUc.CustomerProfilePicture != nil && *currentObjectUc.CustomerProfilePicture != "" {
		strImgprofile = strings.ReplaceAll(*currentObjectUc.CustomerProfilePicture, models.CustomerImagePath, "")
	}

	if imgProfile != nil {

		if &strImgprofile != nil && strings.Trim(strImgprofile, " ") != "" {
			_, err = awsUc.Delete("image/customer", strImgprofile)
			if err != nil {
				logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "s3", uc.ReqID)
			}
		}

		imgProfileFile, err := awsUc.Upload("image/customer", imgProfile)
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "upload_file", c.Value("requestid"))
			return res, err
		}
		strImgprofile = imgProfileFile.FileName

	}

	if imgKtp != nil {

		if &strImgktp != nil && strings.Trim(strImgktp, " ") != "" {
			_, err = awsUc.Delete("image/customer", strImgktp)
			if err != nil {
				logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "s3", uc.ReqID)
			}
		}

		imgKtpFile, err := awsUc.Upload("image/customer", imgKtp)
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "upload_file", c.Value("requestid"))
			return res, err
		}
		strImgktp = imgKtpFile.FileName

	}

	repo := repository.NewCustomerRepository(uc.DB)

	res = models.Customer{
		ID:                     &id,
		Code:                   &data.Code,
		CustomerName:           &data.CustomerName,
		CustomerAddress:        &data.CustomerAddress,
		CustomerPhone:          &data.CustomerPhone,
		CustomerEmail:          &data.CustomerEmail,
		CustomerCpName:         &data.CustomerCpName,
		CustomerProfilePicture: &strImgprofile,
		CustomerNik:            &data.CustomerNik,
		CustomerReligion:       &data.CustomerReligion,
		CustomerPhotoKtp:       &strImgktp,
		CustomerBirthDate:      &data.CustomerBirthDate,
	}

	res.ID, err = repo.Edit(c, &res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	birthDate, _ := time.Parse("2006-01-02T15:04:05Z", *currentObjectUc.CustomerBirthDate)
	birthDateString := birthDate.Format("2006-01-02")
	currentObjectUc.CustomerBirthDate = &birthDateString
	err = CustomerLogUC{ContractUC: uc.ContractUC}.Add(c, currentObjectUc, res, id, 0)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "add_logs", uc.ReqID)
		return res, err
	}
	// Invalidate the cache before update
	err = uc.RedisClient.Client.Del(cacheKey).Err()
	if err != nil {
		// Log error
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "redis_del", uc.ReqID)
		return res, err
	}

	// Refresh the data from the repository
	refreshedRes, err := repo.FindByID(c, models.CustomerParameter{ID: *res.ID})
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return refreshedRes, err
}

func (uc CustomerUC) EditAddress(c context.Context, id string, data *requests.CustomerRequest) (res models.Customer, err error) {
	repo := repository.NewCustomerRepository(uc.DB)

	res = models.Customer{
		ID:                    &id,
		CustomerName:          &data.CustomerName,
		CustomerAddress:       &data.CustomerAddress,
		CustomerProvinceID:    &data.CustomerProvinceID,
		CustomerCityID:        &data.CustomerCityID,
		CustomerDistrictID:    &data.CustomerDistrictID,
		CustomerSubdistrictID: &data.CustomerSubdistrictID,
		CustomerPostalCode:    &data.CustomerPostalCode,
	}

	res.ID, err = repo.EditAddress(c, &res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	// Redis integration
	cacheKey := CustomerCacheKey + id

	// Update the cache with new data
	jsonData, err := json.Marshal(res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "json_marshal", uc.ReqID)
		return res, err
	}
	err = uc.RedisClient.Client.Set(cacheKey, jsonData, time.Hour).Err()
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "redis_set", uc.ReqID)
		return res, err
	}

	return res, err
}

func (uc CustomerUC) BackendEdit(c context.Context, id string, data *requests.CustomerRequest, imgProfile *multipart.FileHeader) (res models.Customer, err error) {

	// Invalidate the cache before update
	cacheKey := CustomerCacheKey + id
	err = uc.RedisClient.Client.Del(cacheKey).Err()
	if err != nil {
		// Log error
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "redis_del", uc.ReqID)
		return res, err
	}

	// currentObjectUc, err := uc.FindByID(c, models.MpBankParameter{ID: id})
	currentObjectUc, err := uc.FindByID(c, models.CustomerParameter{ID: id})
	ctx := "FileUC.Upload"
	awsUc := AwsUC{ContractUC: uc.ContractUC}

	var strImgprofile = ""

	if currentObjectUc.CustomerProfilePicture != nil && *currentObjectUc.CustomerProfilePicture != "" {
		strImgprofile = strings.ReplaceAll(*currentObjectUc.CustomerProfilePicture, models.CustomerImagePath, "")
	}

	if imgProfile != nil {
		if &strImgprofile != nil && strings.Trim(strImgprofile, " ") != "" {
			_, err = awsUc.Delete("image/customer", strImgprofile)
			if err != nil {
				logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "s3", uc.ReqID)
			}
		}

		awsUc.AWSS3.Directory = "image/customer"
		imgBannerFile, err := awsUc.Upload("image/customer", imgProfile)
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "upload_file", c.Value("requestid"))
			return res, err
		}
		strImgprofile = imgBannerFile.FilePath

	}
	repo := repository.NewCustomerRepository(uc.DB)

	res = models.Customer{
		ID:                     &id,
		Code:                   &data.Code,
		CustomerName:           &data.CustomerName,
		CustomerAddress:        &data.CustomerAddress,
		CustomerPhone:          &data.CustomerPhone,
		CustomerEmail:          &data.CustomerEmail,
		CustomerCpName:         &data.CustomerCpName,
		CustomerProfilePicture: &strImgprofile,
		CustomerTaxCalcMethod:  &data.CustomerTaxCalcMethod,
		CustomerActiveStatus:   &data.CustomerActiveStatus,
		CustomerSalesmanID:     &data.CustomerSalesmanID,
		CustomerBranchID:       &data.CustomerBranchID,
		CustomerNik:            &data.CustomerNik,
	}

	res.ID, err = repo.BackendEdit(c, &res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}
	// Invalidate the cache before refresh
	err = uc.RedisClient.Client.Del(cacheKey).Err()
	if err != nil {
		// Log error
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "redis_del", uc.ReqID)
		return res, err
	}

	// Refresh the data from the repository
	refreshedRes, err := repo.FindByID(c, models.CustomerParameter{ID: *res.ID})
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return refreshedRes, err
}

func (uc CustomerUC) BackendAdd(c context.Context, data *requests.CustomerRequest, imgProfile *multipart.FileHeader) (res models.Customer, err error) {

	// currentObjectUc, err := uc.FindByID(c, models.MpBankParameter{ID: id})
	ctx := "FileUC.Upload"
	awsUc := AwsUC{ContractUC: uc.ContractUC}

	var strImgprofile = ""

	if imgProfile != nil {

		awsUc.AWSS3.Directory = "image/customer"
		imgBannerFile, err := awsUc.Upload("image/customer", imgProfile)
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "upload_file", c.Value("requestid"))
			return res, err
		}
		strImgprofile = imgBannerFile.FilePath

	}
	repo := repository.NewCustomerRepository(uc.DB)
	// now := time.Now().UTC()
	// strnow := now.Format(time.RFC3339)
	res = models.Customer{
		Code:                   &data.Code,
		CustomerName:           &data.CustomerName,
		CustomerAddress:        &data.CustomerAddress,
		CustomerPhone:          &data.CustomerPhone,
		CustomerEmail:          &data.CustomerEmail,
		CustomerCpName:         &data.CustomerCpName,
		CustomerProfilePicture: &strImgprofile,
		CustomerTaxCalcMethod:  &data.CustomerTaxCalcMethod,
		CustomerActiveStatus:   &data.CustomerActiveStatus,
		CustomerSalesmanID:     &data.CustomerSalesmanID,
		CustomerBranchID:       &data.CustomerBranchID,
	}

	res.ID, err = repo.BackendAdd(c, &res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return res, err
}
