package usecase

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"mime/multipart"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis/v7"
	"nextbasis-service-v-0.1/db/repository"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/functioncaller"
	"nextbasis-service-v-0.1/pkg/logruslogger"
	"nextbasis-service-v-0.1/server/requests"
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

// WebCustomerUC ...
type WebCustomerUC struct {
	*ContractUC
	CustomerLogUC
}

// BuildBody ...
func (uc WebCustomerUC) BuildBody(data *models.WebCustomer, res *viewmodel.CustomerVM, birthdateFull bool) {
	if !data.IsDataComplete {
		res.CustomerProfileStatus = models.CustomerProfileStatusIncomplete
	} else {
		res.CustomerProfileStatus = models.CustomerProfileStatusComplete
	}

	res.ID = data.ID.String
	res.Code = data.Code.String
	res.CustomerName = data.CustomerName.String

	var profilePictureURL string
	if data.CustomerProfilePicture.Valid {
		profilePictureURL = models.CustomerImagePath + data.CustomerProfilePicture.String
	}
	res.CustomerProfilePicture = profilePictureURL

	res.CustomerActiveStatus = data.CustomerActiveStatus.String

	if data.CustomerBirthDate.Valid && birthdateFull {
		birthDate, _ := time.Parse("2006-01-02T15:04:05Z", data.CustomerBirthDate.String)
		birthDateString := birthDate.Format("02 January 2006")
		res.CustomerBirthDate = birthDateString
	} else if data.CustomerBirthDate.Valid && !birthdateFull {
		birthDate, _ := time.Parse("2006-01-02T15:04:05Z", data.CustomerBirthDate.String)
		birthDateString := birthDate.Format("2006-01-02")
		res.CustomerBirthDate = birthDateString
	}

	res.CustomerReligion = data.CustomerReligion.String
	res.CustomerLatitude = data.CustomerLatitude.String
	res.CustomerLongitude = data.CustomerLongitude.String
	res.CustomerBranchCode = data.CustomerBranchCode.String
	res.CustomerBranchName = data.CustomerBranchName.String
	res.CustomerBranchArea = data.CustomerBranchArea.String
	res.CustomerBranchAddress = data.CustomerAddress.String
	res.CustomerBranchLat = data.CustomerBranchLat.String
	res.CustomerBranchLng = data.CustomerBranchLng.String
	res.CustomerBranchPicPhoneNo = data.CustomerBranchPicPhoneNo.String
	res.CustomerBranchPicName = data.CustomerBranchPicName.String
	res.CustomerRegionCode = data.CustomerRegionCode.String
	res.CustomerRegionName = data.CustomerRegionName.String
	res.CustomerRegionGroup = data.CustomerRegionGroup.String
	res.CustomerEmail = data.CustomerEmail.String
	res.CustomerCpName = data.CustomerCpName.String
	res.CustomerAddress = data.CustomerAddress.String
	res.CustomerPostalCode = data.CustomerPostalCode.String
	res.CustomerProvinceID = data.CustomerProvinceID.String
	res.CustomerProvinceName = data.CustomerProvinceName.String
	res.CustomerCityID = data.CustomerCityID.String
	res.CustomerCityName = data.CustomerCityName.String
	res.CustomerDistrictID = data.CustomerDistrictID.String
	res.CustomerDistrictName = data.CustomerDistrictName.String
	res.CustomerSubdistrictID = data.CustomerSubdistrictID.String
	res.CustomerSubdistrictName = data.CustomerSubdistrictName.String
	res.CustomerSalesmanCode = data.CustomerSalesmanCode.String
	res.CustomerSalesmanName = data.CustomerSalesmanName.String
	res.CustomerSalesmanPhone = data.CustomerSalesmanPhone.String
	res.CustomerSalesCycle = data.CustomerSalesCycle.String
	res.CustomerTypeId = data.CustomerTypeId.String
	res.CustomerTypeName = data.CustomerTypeName.String
	res.CustomerPhone = data.CustomerPhone.String
	res.CustomerPoint = data.CustomerPoint.String
	res.GiftName = data.GiftName.String
	res.Loyalty = data.Loyalty.String
	res.VisitDay = data.VisitDay.String
	res.CustomerTaxCalcMethod = data.CustomerTaxCalcMethod.String
	res.CustomerBranchID = data.CustomerBranchID.String
	res.CustomerSalesmanID = data.CustomerSalesmanID.String
	res.CustomerNik = data.CustomerNik.String

	var photoktpURL string
	if data.CustomerPhotoKtp.Valid {
		photoktpURL = models.CustomerImagePath + data.CustomerPhotoKtp.String
	}
	res.CustomerPhotoKtp = photoktpURL

	res.CustomerLevelID = int(data.CustomerLevelID.Int64)
	res.CustomerLevel = data.CustomerLevel.String
	res.CustomerUserID = data.CustomerUserID.String
	res.CustomerUserName = data.CustomerUserName.String
	res.CustomerGender = data.CustomerGender.String
	res.ModifiedBy = data.ModifiedBy.String
	res.ModifiedDate = data.ModifiedDate.String

	res.CustomerPriceListID = data.CustomerPriceListID.String
	res.CustomerPriceListName = data.CustomerPriceListName.String
	res.CustomerShowInApp = data.ShowInApp.String

	res.CustomerStatusInstall = true
	if data.CustomerUserToken.Valid {
		res.CustomerStatusInstall = false
	} else {
		res.CustomerFCMToken = data.CustomerUserToken.String
	}

	res.SalesmanTypeCode = data.SalesmanTypeCode.String
	res.SalesmanTypeName = data.SalesmanTypeName.String
}

// SelectAll ...
func (uc WebCustomerUC) SelectAll(c context.Context, parameter models.WebCustomerParameter) (res []viewmodel.CustomerVM, err error) {
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(0, 0, parameter.By, parameter.Sort, models.WebCustomerOrderBy, models.WebCustomerOrderByrByString)

	repo := repository.NewWebCustomerRepository(uc.DB)
	data, err := repo.SelectAll(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	for i := range data {
		var temp viewmodel.CustomerVM

		uc.BuildBody(&data[i], &temp, true)
		res = append(res, temp)
	}

	return res, err
}

// FindAll ...
// func (uc WebCustomerUC) FindAll(c context.Context, parameter models.WebCustomerParameter) (res []viewmodel.CustomerVM, p viewmodel.PaginationVM, err error) {
// 	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.WebCustomerOrderBy, models.WebCustomerOrderByrByString)

// 	var count int
// 	repo := repository.NewWebCustomerRepository(uc.DB)
// 	data, count, err := repo.FindAll(c, parameter)
// 	if err != nil {
// 		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
// 		return res, p, err
// 	}

// 	p = uc.setPaginationResponse(parameter.Page, parameter.Limit, count)
// 	for i := range data {
// 		var temp viewmodel.CustomerVM
// 		uc.BuildBody(&data[i], &temp, false)
// 		res = append(res, temp)
// 	}

// 	return res, p, err
// }

func (uc WebCustomerUC) FindAll(c context.Context, parameter models.WebCustomerParameter) ([]viewmodel.CustomerVM, viewmodel.PaginationVM, error) {
	var response viewmodel.PaginatedResponse

	cacheKey := fmt.Sprintf("customer:admin_user_id:%s:page:%d:search:%s:branch_id:%s:phone_number:%s:show_in_app:%s:by:%s:sort:&%s:customer_type:%s",
		parameter.UserId, parameter.Page, parameter.Search, parameter.BranchId, parameter.PhoneNumber, parameter.ShowInApp, parameter.By, parameter.Sort, parameter.CustomerTypeId)

	// Try getting data from cache
	cachedData, err := uc.RedisClient.Get(cacheKey)
	if err == nil && string(cachedData) != "" {
		err := json.Unmarshal(cachedData, &response)
		if err == nil {
			return response.Data.ListCustomer, response.Meta, nil
		}
	}

	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.WebCustomerOrderBy, models.WebCustomerOrderByrByString)

	repo := repository.NewWebCustomerRepository(uc.DB)
	data, count, err := repo.FindAll(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return nil, viewmodel.PaginationVM{}, err
	}

	p := uc.setPaginationResponse(parameter.Page, parameter.Limit, count)
	for _, d := range data {
		var temp viewmodel.CustomerVM
		uc.BuildBody(&d, &temp, true)
		response.Data.ListCustomer = append(response.Data.ListCustomer, temp)
	}

	response.Meta = p

	// Cache the entire response (don't cache empty result)
	if len(response.Data.ListCustomer) > 0 {
		jsonData, err := json.Marshal(response)
		if err == nil {
			uc.RedisClient.Set(cacheKey, jsonData, time.Minute*30) // Cache for 30 minutes only if there's data
		}
	}

	return response.Data.ListCustomer, response.Meta, nil
}

// func (uc WebCustomerUC) FindAll(c context.Context, parameter models.WebCustomerParameter) (res []viewmodel.CustomerVM, p viewmodel.PaginationVM, err error) {
//     parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.WebCustomerOrderBy, models.WebCustomerOrderByrByString)

//     // Prepare a cache key
//     cacheKey := fmt.Sprintf("customers_page_%d", parameter.Page)

//     // Try to get the data from cache
//     cachedData, cacheErr := uc.RedisClient.Client.Get(cacheKey).Result()

//     if cacheErr == nil && cachedData != "" {
//         // Cache hit - unmarshal JSON data to struct and return
//         err := json.Unmarshal([]byte(cachedData), &res)
//         if err != nil {
//             return nil, p, err
//         }
//     } else {
//         // Cache miss - fetch data from DB
//         var count int
//         repo := repository.NewWebCustomerRepository(uc.DB)
//         data, count, err := repo.FindAll(c, parameter)
//         if err != nil {
//             logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
//             return res, p, err
//         }

//         p = uc.setPaginationResponse(parameter.Page, parameter.Limit, count)
//         for i := range data {
//             var temp viewmodel.CustomerVM
//             uc.BuildBody(&data[i], &temp, false)
//             res = append(res, temp)
//         }

//         // Cache data
//         jsonData, err := json.Marshal(res)
//         if err != nil {
//             return nil, p, err
//         }
//         uc.RedisClient.Client.Set(cacheKey, string(jsonData), time.Hour)
//     }

//     return res, p, err
// }

// func (uc WebCustomerUC) SelectAll(c context.Context, parameter models.WebCustomerParameter) (res []viewmodel.CustomerVM, err error) {
// 	// Define cache key
// 	cacheKey := "web_customers:SelectAll:" + parameter.By + ":" + parameter.Sort

// 	// Try to fetch the data from Redis first
// 	err = uc.RedisClient.Client.Get(cacheKey).Scan(&res)
// 	if err == nil {
// 		// Data was found in Redis, return it
// 		return res, nil
// 	} else if err != nil && err != redis.Nil {
// 		// An error occurred that wasn't just "key doesn't exist"
// 		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "redis_get", uc.ReqID)
// 		return res, err
// 	}

// 	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(0, 0, parameter.By, parameter.Sort, models.WebCustomerOrderBy, models.WebCustomerOrderByrByString)

// 	repo := repository.NewWebCustomerRepository(uc.DB)
// 	data, err := repo.SelectAll(c, parameter)
// 	if err != nil {
// 		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
// 		return res, err
// 	}

// 	for i := range data {
// 		var temp viewmodel.CustomerVM
// 		uc.BuildBody(&data[i], &temp, true)
// 		res = append(res, temp)
// 	}

// 	// Cache the data in Redis, ignore the result
// 	err = uc.StoreToRedis(cacheKey, res)
// 	if err != nil {
// 		logruslogger.Log(logruslogger.WarnLevel, err.Error(), "StoreToRedis", "redis_set", uc.ReqID)
// 	}

// 	return res, err
// }

// FindAll ...
// func (uc WebCustomerUC) FindAll(c context.Context, parameter models.WebCustomerParameter) (res []viewmodel.CustomerVM, p viewmodel.PaginationVM, err error) {
// 	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.WebCustomerOrderBy, models.WebCustomerOrderByrByString)

// 	// Redis integration
// 	cacheKey := "web_customers_findAll:" + strconv.Itoa(parameter.Page) + ":" + strconv.Itoa(parameter.Limit)

// 	val, err := uc.RedisClient.Client.Get(cacheKey).Result()
// 	if err == nil {
// 		// If cache exists
// 		err = json.Unmarshal([]byte(val), &res)
// 		if err == nil {
// 			return res, p, nil
// 		}
// 	}

// 	var count int
// 	repo := repository.NewWebCustomerRepository(uc.DB)
// 	data, count, err := repo.FindAll(c, parameter)
// 	if err != nil {
// 		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
// 		return res, p, err
// 	}

// 	p = uc.setPaginationResponse(parameter.Page, parameter.Limit, count)
// 	for i := range data {
// 		var temp viewmodel.CustomerVM
// 		uc.BuildBody(&data[i], &temp, false)
// 		res = append(res, temp)
// 	}

// 	// Save result into Redis
// 	jsonData, errRedis := json.Marshal(res)
// 	if errRedis != nil {
// 		logruslogger.Log(logruslogger.WarnLevel, errRedis.Error(), functioncaller.PrintFuncName(), "json_marshal", uc.ReqID)
// 	} else {
// 		errRedis = uc.RedisClient.Client.Set(cacheKey, jsonData, time.Hour).Err()
// 		if errRedis != nil {
// 			logruslogger.Log(logruslogger.WarnLevel, errRedis.Error(), functioncaller.PrintFuncName(), "redis_set", uc.ReqID)
// 		}
// 	}

// 	return res, p, err
// }

// FindByID ...
func (uc WebCustomerUC) FindByID(c context.Context, parameter models.WebCustomerParameter) (res viewmodel.CustomerVM, err error) {
	// Redis integration
	cacheKey := CustomerCacheKey + parameter.ID

	val, err := uc.RedisClient.Client.Get(cacheKey).Result()

	if err != nil {
		if errors.Is(err, redis.Nil) {
			// If cache does not exist, fetch from the repository
			repo := repository.NewWebCustomerRepository(uc.DB)
			datum, err := repo.FindByID(c, parameter)
			if err != nil {
				logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
				return res, err
			}

			uc.BuildBody(&datum, &res, false)

			// Save result into Redis
			jsonData, err := json.Marshal(res)
			if err != nil {
				logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "json_marshal", uc.ReqID)
				return res, err // return here if error occurred
			}
			err = uc.RedisClient.Client.Set(cacheKey, jsonData, time.Hour).Err()
			if err != nil {
				logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "redis_set", uc.ReqID)
				return res, err // return here if error occurred
			}

			return res, nil
		} else {
			// If there is an error other than "key does not exist", log and return error
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "redis_get", uc.ReqID)
			return res, err
		}
	} else {
		err = json.Unmarshal([]byte(val), &res)
		if err != nil {
			// If there is an error in unmarshaling, log it
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "json_unmarshal", uc.ReqID)
		}
		return res, nil
	}
}

// FindByIDNoCache ...
func (uc WebCustomerUC) FindByIDNoCache(c context.Context, parameter models.WebCustomerParameter) (res viewmodel.CustomerVM, err error) {
	// Fetch data from DB
	repo := repository.NewWebCustomerRepository(uc.DB)
	datum, err := repo.FindByID(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	// Transform data
	uc.BuildBody(&datum, &res, false)

	return res, nil
}

func (uc WebCustomerUC) Edit(c context.Context, id string, data *requests.WebCustomerRequest, imgProfile, imgKtp *multipart.FileHeader) (res viewmodel.CustomerVM, err error) {

	// Invalidate the cache before update
	cacheKey := CustomerCacheKey + id
	err = uc.RedisClient.Client.Del(cacheKey).Err()
	if err != nil {
		// Log error
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "redis_del", uc.ReqID)
		return res, err
	}

	// currentObjectUc, err := uc.FindByID(c, models.MpBankParameter{ID: id})
	currentObjectUc, err := uc.FindByIDNoCache(c, models.WebCustomerParameter{ID: id})
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "invalid id", uc.ReqID)
		return
	}

	// if currentObjectUc.CustomerPhone != nil && *currentObjectUc.CustomerPhone != data.CustomerPhone {
	// 	checkerPhoneNumberData, _ := uc.SelectAll(c, models.WebCustomerParameter{
	// 		PhoneNumber: data.CustomerPhone,
	// 		By:          "c.created_date",
	// 	})
	// 	if len(checkerPhoneNumberData) > 0 {
	// 		err = errors.New("Duplicate Phone Number")
	// 		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "duplicate phone number", uc.ReqID)
	// 		return
	// 	}
	// }

	// If the phone number has changed, check if it's unique
	if currentObjectUc.CustomerPhone != data.CustomerPhone {
		checkerPhoneNumberData, _ := uc.SelectAll(c, models.WebCustomerParameter{
			PhoneNumber: data.CustomerPhone,
			By:          "c.created_date",
		})

		// If phone number is not unique, return error
		if len(checkerPhoneNumberData) > 0 {
			err = errors.New("Duplicate phone number")
			return res, err
		}
	}

	ctx := "FileUC.Upload"
	awsUc := AwsUC{ContractUC: uc.ContractUC}

	var strImgprofile = ""
	if currentObjectUc.CustomerProfilePicture != "" {
		strImgprofile = strings.ReplaceAll(currentObjectUc.CustomerProfilePicture, models.CustomerImagePath, "")
	}
	if imgProfile != nil {
		awsUc.AWSS3.Directory = "image/customer"
		if &strImgprofile != nil && strings.Trim(strImgprofile, " ") != "" {
			_, err = awsUc.Delete(awsUc.AWSS3.Directory, strImgprofile)
			if err != nil {
				logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "s3", uc.ReqID)
			}
		}

		imgBannerFile, err := awsUc.Upload(awsUc.AWSS3.Directory, imgProfile)
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "upload_file", c.Value("requestid"))
			return res, err
		}
		strImgprofile = imgBannerFile.FileName
	}

	var stringImageKTP string
	if currentObjectUc.CustomerPhotoKtp != "" {
		stringImageKTP = strings.ReplaceAll(currentObjectUc.CustomerPhotoKtp, models.CustomerImagePath, "")
	}
	if imgKtp != nil {
		awsUc.AWSS3.Directory = "image/customer"
		if &stringImageKTP != nil && strings.Trim(stringImageKTP, " ") != "" {
			_, err = awsUc.Delete(awsUc.AWSS3.Directory, stringImageKTP)
			if err != nil {
				logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "s3", uc.ReqID)
			}
		}

		imgBannerFile, err := awsUc.Upload(awsUc.AWSS3.Directory, imgKtp)
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "upload_file", c.Value("requestid"))
			return res, err
		}
		stringImageKTP = imgBannerFile.FileName
	}

	repo := repository.NewWebCustomerRepository(uc.DB)
	// now := time.Now().UTC()
	// strnow := now.Format(time.RFC3339)

	birthDate, _ := time.Parse("2006-01-02", data.CustomerBirthDate)
	data.CustomerBirthDate = birthDate.Format("2006-01-02")

	if data.CustomerShowInApp == "" {
		data.CustomerShowInApp = currentObjectUc.CustomerShowInApp
	}
	in := models.WebCustomer{
		ID:                     sql.NullString{String: id},
		Code:                   sql.NullString{String: data.Code},
		CustomerName:           sql.NullString{String: data.CustomerName},
		CustomerAddress:        sql.NullString{String: data.CustomerAddress},
		CustomerPhone:          sql.NullString{String: data.CustomerPhone},
		CustomerEmail:          sql.NullString{String: data.CustomerEmail},
		CustomerCpName:         sql.NullString{String: data.CustomerCpName},
		CustomerProfilePicture: sql.NullString{String: strImgprofile},
		CustomerTaxCalcMethod:  sql.NullString{String: data.CustomerTaxCalcMethod},
		CustomerActiveStatus:   sql.NullString{String: data.CustomerActiveStatus},
		CustomerSalesmanID:     sql.NullString{String: data.CustomerSalesmanID},
		CustomerBranchID:       sql.NullString{String: data.CustomerBranchID},
		CustomerNik:            sql.NullString{String: data.CustomerNik},
		CustomerUserID:         sql.NullString{String: data.CustomerUserID},
		CustomerReligion:       sql.NullString{String: data.CustomerReligion},
		CustomerLevelID:        sql.NullInt64{Int64: int64(data.CustomerLevelID)},
		CustomerGender:         sql.NullString{String: data.CustomerGender},
		CustomerBirthDate:      sql.NullString{String: data.CustomerBirthDate},
		CustomerPhotoKtp:       sql.NullString{String: stringImageKTP},
		UserID:                 sql.NullInt64{Int64: int64(data.UserID)},
		ShowInApp:              sql.NullString{String: data.CustomerShowInApp},
	}

	in.ID.String, err = repo.Edit(c, in)
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
	res, err = uc.FindByID(c, models.WebCustomerParameter{ID: in.ID.String})
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	err = CustomerLogUC{ContractUC: uc.ContractUC}.Add(c, currentObjectUc, res, id, data.UserID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "add_logs", uc.ReqID)
		return res, err
	}

	return res, err
}

func (uc WebCustomerUC) EditBulk(c context.Context, data requests.WebCustomerBulkRequest) (err error) {
	repo := repository.NewWebCustomerRepository(uc.DB)

	err = repo.EditBulk(c, data)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	return
}

// func (uc WebCustomerUC) Edit(c context.Context, id string, data *requests.WebCustomerRequest, imgProfile, imgKtp *multipart.FileHeader) (res models.WebCustomer, err error) {

// 	// currentObjectUc, err := uc.FindByID(c, models.MpBankParameter{ID: id})
// 	currentObjectUc, err := uc.FindByID(c, models.WebCustomerParameter{ID: id})
// 	if err != nil {
// 		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "invalid id", uc.ReqID)
// 		return
// 	}

// 	if currentObjectUc.CustomerPhone != nil && *currentObjectUc.CustomerPhone != data.CustomerPhone {
// 		checkerPhoneNumberData, _ := uc.SelectAll(c, models.WebCustomerParameter{
// 			PhoneNumber: data.CustomerPhone,
// 			By:          "c.created_date",
// 		})
// 		if len(checkerPhoneNumberData) > 0 {
// 			err = errors.New("Duplicate Phone Number")
// 			logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "duplicate phone number", uc.ReqID)
// 			return
// 		}
// 	}

// 	ctx := "FileUC.Upload"
// 	awsUc := AwsUC{ContractUC: uc.ContractUC}

// 	var strImgprofile = ""
// 	if currentObjectUc.CustomerProfilePicture != nil && *currentObjectUc.CustomerProfilePicture != "" {
// 		strImgprofile = strings.ReplaceAll(*currentObjectUc.CustomerProfilePicture, models.CustomerImagePath, "")
// 	}
// 	if imgProfile != nil {
// 		awsUc.AWSS3.Directory = "image/customer"
// 		if &strImgprofile != nil && strings.Trim(strImgprofile, " ") != "" {
// 			_, err = awsUc.Delete(awsUc.AWSS3.Directory, strImgprofile)
// 			if err != nil {
// 				logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "s3", uc.ReqID)
// 			}
// 		}

// 		imgBannerFile, err := awsUc.Upload(awsUc.AWSS3.Directory, imgProfile)
// 		if err != nil {
// 			logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "upload_file", c.Value("requestid"))
// 			return res, err
// 		}
// 		strImgprofile = imgBannerFile.FileName
// 	}

// 	var stringImageKTP string
// 	if currentObjectUc.CustomerPhotoKtp != nil && *currentObjectUc.CustomerPhotoKtp != "" {
// 		stringImageKTP = strings.ReplaceAll(*currentObjectUc.CustomerPhotoKtp, models.CustomerImagePath, "")
// 	}
// 	if imgKtp != nil {
// 		awsUc.AWSS3.Directory = "image/customer"
// 		if &stringImageKTP != nil && strings.Trim(stringImageKTP, " ") != "" {
// 			_, err = awsUc.Delete(awsUc.AWSS3.Directory, stringImageKTP)
// 			if err != nil {
// 				logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "s3", uc.ReqID)
// 			}
// 		}

// 		imgBannerFile, err := awsUc.Upload(awsUc.AWSS3.Directory, imgKtp)
// 		if err != nil {
// 			logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "upload_file", c.Value("requestid"))
// 			return res, err
// 		}
// 		stringImageKTP = imgBannerFile.FileName
// 	}

// 	repo := repository.NewWebCustomerRepository(uc.DB)
// 	// now := time.Now().UTC()
// 	// strnow := now.Format(time.RFC3339)

// 	birthDate, _ := time.Parse("2006-01-02", data.CustomerBirthDate)
// 	data.CustomerBirthDate = birthDate.Format("2006-01-02")

// 	res = models.WebCustomer{
// 		ID:                     &id,
// 		Code:                   &data.Code,
// 		CustomerName:           &data.CustomerName,
// 		CustomerAddress:        &data.CustomerAddress,
// 		CustomerPhone:          &data.CustomerPhone,
// 		CustomerEmail:          &data.CustomerEmail,
// 		CustomerCpName:         &data.CustomerCpName,
// 		CustomerProfilePicture: &strImgprofile,
// 		CustomerTaxCalcMethod:  &data.CustomerTaxCalcMethod,
// 		CustomerActiveStatus:   &data.CustomerActiveStatus,
// 		CustomerSalesmanID:     &data.CustomerSalesmanID,
// 		CustomerBranchID:       &data.CustomerBranchID,
// 		CustomerNik:            &data.CustomerNik,
// 		CustomerUserID:         &data.CustomerUserID,
// 		CustomerReligion:       &data.CustomerReligion,
// 		CustomerLevelID:        &data.CustomerLevelID,
// 		CustomerGender:         &data.CustomerGender,
// 		CustomerBirthDate:      &data.CustomerBirthDate,
// 		CustomerPhotoKtp:       &stringImageKTP,
// 		UserID:                 &data.UserID,
// 	}

// 	res.ID, err = repo.Edit(c, &res)
// 	if err != nil {
// 		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
// 		return res, err
// 	}

// 	return res, err
// }

func (uc WebCustomerUC) Add(c context.Context, data *requests.WebCustomerRequest, imgProfile *multipart.FileHeader) (res viewmodel.CustomerVM, err error) {

	// currentObjectUc, err := uc.FindByID(c, models.MpBankParameter{ID: id})
	ctx := "FileUC.Upload"
	awsUc := AwsUC{ContractUC: uc.ContractUC}

	checkerPhoneNumberData, _ := uc.SelectAll(c, models.WebCustomerParameter{
		PhoneNumber: data.CustomerPhone,
		By:          "c.created_date",
	})
	if len(checkerPhoneNumberData) > 0 {
		err = errors.New("Duplicate Phone Number")
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "duplicate phone number", uc.ReqID)
		return
	}

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

	birthDate, _ := time.Parse("2006-01-02", data.CustomerBirthDate)
	data.CustomerBirthDate = birthDate.Format("2006-01-02")

	repo := repository.NewWebCustomerRepository(uc.DB)
	// now := time.Now().UTC()
	// strnow := now.Format(time.RFC3339)
	in := models.WebCustomer{
		Code:                   sql.NullString{String: data.Code},
		CustomerName:           sql.NullString{String: data.CustomerName},
		CustomerAddress:        sql.NullString{String: data.CustomerAddress},
		CustomerPhone:          sql.NullString{String: data.CustomerPhone},
		CustomerEmail:          sql.NullString{String: data.CustomerEmail},
		CustomerCpName:         sql.NullString{String: data.CustomerCpName},
		CustomerProfilePicture: sql.NullString{String: strImgprofile},
		CustomerTaxCalcMethod:  sql.NullString{String: data.CustomerTaxCalcMethod},
		CustomerActiveStatus:   sql.NullString{String: data.CustomerActiveStatus},
		CustomerSalesmanID:     sql.NullString{String: data.CustomerSalesmanID},
		CustomerBranchID:       sql.NullString{String: data.CustomerBranchID},
		CustomerUserID:         sql.NullString{String: data.CustomerUserID},
		CustomerReligion:       sql.NullString{String: data.CustomerReligion},
		CustomerLevelID:        sql.NullInt64{Int64: int64(data.CustomerLevelID)},
		CustomerGender:         sql.NullString{String: data.CustomerGender},
		CustomerBirthDate:      sql.NullString{String: data.CustomerBirthDate},
	}

	in.ID.String, err = repo.Add(c, in)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query_insert", c.Value("requestid"))
		return res, err
	}

	res, _ = uc.FindByID(c, models.WebCustomerParameter{ID: in.ID.String})

	return res, err
}

// ReportSelect ...
func (uc WebCustomerUC) ReportSelect(c context.Context, parameter models.WebCustomerReportParameter) (res []viewmodel.CustomerVM, err error) {
	repo := repository.NewWebCustomerRepository(uc.DB)
	data, err := repo.ReportSelect(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	var cityIDs, proviceIDs, districtIDs, subDistrictids, salesmanIDs, customerLevelIDs, customerTypeIDs []string
	cityIDChecker := make(map[string]string)
	proviceIDChecker := make(map[string]string)
	districtIDChecker := make(map[string]string)
	subdistrictIDChecker := make(map[string]string)
	salesmanIDChecker := make(map[string]string)
	customerLevelIDChecker := make(map[int]string)
	customerTypeIDChecker := make(map[string]string)
	for i := range data {
		if data[i].CustomerCityID.Valid && cityIDChecker[data[i].CustomerCityID.String] == "" {
			cityIDChecker[data[i].CustomerCityID.String] = "done"
			cityIDs = append(cityIDs, data[i].CustomerCityID.String)
		}
		if data[i].CustomerProvinceID.Valid && proviceIDChecker[data[i].CustomerProvinceID.String] == "" {
			proviceIDChecker[data[i].CustomerProvinceID.String] = "done"
			proviceIDs = append(proviceIDs, data[i].CustomerProvinceID.String)
		}
		if data[i].CustomerDistrictID.Valid && districtIDChecker[data[i].CustomerDistrictID.String] == "" {
			districtIDChecker[data[i].CustomerDistrictID.String] = "done"
			districtIDs = append(districtIDs, data[i].CustomerDistrictID.String)
		}
		if data[i].CustomerSubdistrictID.Valid && subdistrictIDChecker[data[i].CustomerSubdistrictID.String] == "" {
			subdistrictIDChecker[data[i].CustomerSubdistrictID.String] = "done"
			subDistrictids = append(subDistrictids, data[i].CustomerSubdistrictID.String)
		}
		if data[i].CustomerSalesmanID.Valid && salesmanIDChecker[data[i].CustomerSalesmanID.String] == "" {
			salesmanIDChecker[data[i].CustomerSalesmanID.String] = "done"
			salesmanIDs = append(salesmanIDs, data[i].CustomerSalesmanID.String)
		}
		if data[i].CustomerLevelID.Valid && customerLevelIDChecker[int(data[i].CustomerLevelID.Int64)] == "" {
			customerLevelIDChecker[int(data[i].CustomerLevelID.Int64)] = "done"
			customerLevelIDs = append(customerLevelIDs, strconv.Itoa(int(data[i].CustomerLevelID.Int64)))
		}
		if data[i].CustomerTypeId.Valid && customerTypeIDChecker[data[i].CustomerTypeId.String] == "" {
			customerTypeIDChecker[data[i].CustomerTypeId.String] = "done"
			customerTypeIDs = append(customerTypeIDs, data[i].CustomerTypeId.String)
		}
	}

	cityUC := CityUC{ContractUC: uc.ContractUC}
	cityData, err := cityUC.SelectAll(c, models.CityParameter{IDs: strings.Join(cityIDs, ","), By: "def.id"})
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "find_city_data", c.Value("requestid"))
		return res, err
	}

	proviceUC := ProvinceUC{ContractUC: uc.ContractUC}
	proviceData, err := proviceUC.SelectAll(c, models.ProvinceParameter{IDs: strings.Join(proviceIDs, ","), By: "def.id"})
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "find_provice_data", c.Value("requestid"))
		return res, err
	}

	districtUC := DistrictUC{ContractUC: uc.ContractUC}
	districtData, err := districtUC.SelectAll(c, models.DistrictParameter{IDs: strings.Join(districtIDs, ","), By: "def.id"})
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "find_district_data", c.Value("requestid"))
		return res, err
	}

	subDistrictUC := SubDistrictUC{ContractUC: uc.ContractUC}
	subDistrictData, err := subDistrictUC.SelectAll(c, models.SubDistrictParameter{IDs: strings.Join(subDistrictids, ","), By: "def.id"})
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "find_subdistrict_data", c.Value("requestid"))
		return res, err
	}

	salesmanUC := SalesmanUC{ContractUC: uc.ContractUC}
	salesmanData, err := salesmanUC.SelectAll(c, models.SalesmanParameter{IDs: strings.Join(salesmanIDs, ","), By: "def.id"})
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "find_salesman_data", c.Value("requestid"))
		return res, err
	}

	customerLevelUC := CustomerLevelUC{ContractUC: uc.ContractUC}
	customerLevelData, err := customerLevelUC.FindAll(c, models.CustomerLevelParameter{IDs: strings.Join(customerLevelIDs, ","), By: "def.id"})
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "find_salesman_data", c.Value("requestid"))
		return res, err
	}

	customerTypeUC := CustomerTypeUC{ContractUC: uc.ContractUC}
	customerTypeData, err := customerTypeUC.SelectAll(c, models.CustomerTypeParameter{IDs: strings.Join(customerTypeIDs, ","), By: "def.id"})
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "find_salesman_data", c.Value("requestid"))
		return res, err
	}

	for i := range data {
		var temp viewmodel.CustomerVM

		for j := range cityData {
			if data[i].CustomerCityID.Valid && data[i].CustomerCityID.String == *cityData[j].ID {
				data[i].CustomerCityName.String = *cityData[j].Name
				break
			}
		}

		for j := range proviceData {
			if data[i].CustomerProvinceID.Valid && data[i].CustomerProvinceID.String == proviceData[j].ID {
				data[i].CustomerProvinceName.String = *proviceData[j].Name
				break
			}
		}

		for j := range districtData {
			if data[i].CustomerDistrictID.Valid && data[i].CustomerDistrictID.String == districtData[j].ID {
				data[i].CustomerDistrictName.String = districtData[j].Name
				break
			}
		}

		for j := range subDistrictData {
			if data[i].CustomerSubdistrictID.Valid && data[i].CustomerSubdistrictID.String == subDistrictData[j].ID {
				data[i].CustomerSubdistrictName.String = subDistrictData[j].Name
				break
			}
		}

		for j := range salesmanData {
			if data[i].CustomerSalesmanID.Valid && data[i].CustomerSalesmanID.String == *salesmanData[j].ID {
				data[i].CustomerSalesmanName.String = *salesmanData[j].Name
				data[i].CustomerSalesmanPhone.String = *salesmanData[j].PhoneNo
				data[i].CustomerSalesmanCode.String = *salesmanData[j].Code
				break
			}
		}

		for j := range customerLevelData {
			if data[i].CustomerLevelID.Valid && strconv.Itoa(int(data[i].CustomerLevelID.Int64)) == customerLevelData[j].ID {
				data[i].CustomerLevel.String = customerLevelData[j].Name
				break
			}
		}

		for j := range customerTypeData {
			if data[i].CustomerTypeId.Valid && data[i].CustomerTypeId.String == *customerTypeData[j].ID {
				data[i].CustomerTypeName.String = *customerTypeData[j].Name
				break
			}
		}

		if parameter.CustomerProfileStatus == "0" {
			if !data[i].IsDataComplete {
				uc.BuildBody(&data[i], &temp, true)
				res = append(res, temp)
			}
		} else if parameter.CustomerProfileStatus == "1" {
			if data[i].IsDataComplete {
				uc.BuildBody(&data[i], &temp, true)
				res = append(res, temp)
			}
		} else {
			uc.BuildBody(&data[i], &temp, true)
			res = append(res, temp)
		}
	}

	return res, err
}
