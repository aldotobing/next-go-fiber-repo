package usecase

import (
	"context"
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

// WebCustomerUC ...
type WebCustomerUC struct {
	*ContractUC
}

// BuildBody ...
func (uc WebCustomerUC) BuildBody(data *models.WebCustomer, res *viewmodel.CustomerVM) {
	if data.CustomerProfilePicture == nil ||
		data.CustomerName == nil ||
		data.CustomerBranchName == nil ||
		data.CustomerBranchCode == nil ||
		data.CustomerPhone == nil ||
		data.CustomerBranchPicPhoneNo == nil ||
		data.CustomerReligion == nil ||
		data.CustomerBirthDate == nil ||
		data.CustomerNik == nil ||
		data.CustomerPhotoKtp == nil ||
		data.CustomerProfilePicture == nil {
		res.CustomerProfileStatus = &models.CustomerProfileStatusIncomplete
	} else {
		res.CustomerProfileStatus = &models.CustomerProfileStatusComplete
	}

	res.ID = data.ID
	res.Code = data.Code
	res.CustomerName = data.CustomerName
	res.CustomerProfilePicture = data.CustomerProfilePicture
	res.CustomerActiveStatus = data.CustomerActiveStatus
	res.CustomerBirthDate = data.CustomerBirthDate
	res.CustomerReligion = data.CustomerReligion
	res.CustomerLatitude = data.CustomerLatitude
	res.CustomerLongitude = data.CustomerLongitude
	res.CustomerBranchCode = data.CustomerBranchCode
	res.CustomerBranchName = data.CustomerBranchName
	res.CustomerBranchArea = data.CustomerBranchArea
	res.CustomerBranchAddress = data.CustomerAddress
	res.CustomerBranchLat = data.CustomerBranchLat
	res.CustomerBranchLng = data.CustomerBranchLng
	res.CustomerBranchPicPhoneNo = data.CustomerBranchPicPhoneNo
	res.CustomerRegionCode = data.CustomerRegionCode
	res.CustomerRegionName = data.CustomerRegionName
	res.CustomerRegionGroup = data.CustomerRegionGroup
	res.CustomerEmail = data.CustomerEmail
	res.CustomerCpName = data.CustomerCpName
	res.CustomerAddress = data.CustomerAddress
	res.CustomerPostalCode = data.CustomerPostalCode
	res.CustomerProvinceID = data.CustomerProvinceID
	res.CustomerProvinceName = data.CustomerProvinceName
	res.CustomerCityID = data.CustomerCityID
	res.CustomerCityName = data.CustomerCityName
	res.CustomerDistrictID = data.CustomerDistrictID
	res.CustomerDistrictName = data.CustomerDistrictName
	res.CustomerSubdistrictID = data.CustomerSubdistrictID
	res.CustomerSubdistrictName = data.CustomerSubdistrictName
	res.CustomerSalesmanCode = data.CustomerSalesmanCode
	res.CustomerSalesmanName = data.CustomerSalesmanName
	res.CustomerSalesmanPhone = data.CustomerSalesmanPhone
	res.CustomerSalesCycle = data.CustomerSalesCycle
	res.CustomerTypeId = data.CustomerTypeId
	res.CustomerTypeName = data.CustomerTypeName
	res.CustomerPhone = data.CustomerPhone
	res.CustomerPoint = data.CustomerPoint
	res.GiftName = data.GiftName
	res.Loyalty = data.Loyalty
	res.VisitDay = data.VisitDay
	res.CustomerTaxCalcMethod = data.CustomerTaxCalcMethod
	res.CustomerBranchID = data.CustomerBranchID
	res.CustomerSalesmanID = data.CustomerSalesmanID
	res.CustomerNik = data.CustomerNik
	res.CustomerPhotoKtp = data.CustomerPhotoKtp
	res.CustomerLevelID = data.CustomerLevelID
	res.CustomerLevel = data.CustomerLevel
	res.CustomerUserID = data.CustomerUserID
	res.CustomerUserName = data.CustomerUserName
	res.CustomerGender = data.CustomerGender
}

// SelectAll ...
func (uc WebCustomerUC) SelectAll(c context.Context, parameter models.WebCustomerParameter) (res []models.WebCustomer, err error) {
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(0, 0, parameter.By, parameter.Sort, models.WebCustomerOrderBy, models.WebCustomerOrderByrByString)

	repo := repository.NewWebCustomerRepository(uc.DB)
	res, err = repo.SelectAll(c, parameter)

	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return res, err
}

// FindAll ...
func (uc WebCustomerUC) FindAll(c context.Context, parameter models.WebCustomerParameter) (res []viewmodel.CustomerVM, p viewmodel.PaginationVM, err error) {
	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.WebCustomerOrderBy, models.WebCustomerOrderByrByString)

	var count int
	repo := repository.NewWebCustomerRepository(uc.DB)
	data, count, err := repo.FindAll(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, p, err
	}

	p = uc.setPaginationResponse(parameter.Page, parameter.Limit, count)
	for i := range data {
		var temp viewmodel.CustomerVM
		uc.BuildBody(&data[i], &temp)
		res = append(res, temp)
	}

	return res, p, err
}

// FindByID ...
func (uc WebCustomerUC) FindByID(c context.Context, parameter models.WebCustomerParameter) (res models.WebCustomer, err error) {

	repo := repository.NewWebCustomerRepository(uc.DB)
	res, err = repo.FindByID(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	if res.CustomerProfilePicture == nil ||
		res.CustomerName == nil ||
		res.CustomerBranchName == nil ||
		res.CustomerBranchCode == nil ||
		res.CustomerPhone == nil ||
		res.CustomerBranchPicPhoneNo == nil ||
		res.CustomerReligion == nil ||
		res.CustomerBirthDate == nil ||
		res.CustomerNik == nil ||
		res.CustomerPhotoKtp == nil ||
		res.CustomerProfilePicture == nil {
		res.CustomerProfileStatus = &models.CustomerProfileStatusIncomplete
	} else {
		res.CustomerProfileStatus = &models.CustomerProfileStatusComplete
	}

	return res, err
}

func (uc WebCustomerUC) Edit(c context.Context, id string, data *requests.WebCustomerRequest, imgProfile *multipart.FileHeader) (res models.WebCustomer, err error) {

	// currentObjectUc, err := uc.FindByID(c, models.MpBankParameter{ID: id})
	currentObjectUc, err := uc.FindByID(c, models.WebCustomerParameter{ID: id})
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
	repo := repository.NewWebCustomerRepository(uc.DB)
	// now := time.Now().UTC()
	// strnow := now.Format(time.RFC3339)

	birthDate, _ := time.Parse("2006-01-02", data.CustomerBirthDate)
	data.CustomerBirthDate = birthDate.Format("2006-01-02")

	res = models.WebCustomer{
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
		CustomerUserID:         &data.CustomerUserID,
		CustomerReligion:       &data.CustomerReligion,
		CustomerLevelID:        &data.CustomerLevelID,
		CustomerGender:         &data.CustomerGender,
		CustomerBirthDate:      &data.CustomerBirthDate,
	}

	res.ID, err = repo.Edit(c, &res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return res, err
}

func (uc WebCustomerUC) Add(c context.Context, data *requests.WebCustomerRequest, imgProfile *multipart.FileHeader) (res models.WebCustomer, err error) {

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

	birthDate, _ := time.Parse("2006-01-02", data.CustomerBirthDate)
	data.CustomerBirthDate = birthDate.Format("2006-01-02")

	repo := repository.NewWebCustomerRepository(uc.DB)
	// now := time.Now().UTC()
	// strnow := now.Format(time.RFC3339)
	res = models.WebCustomer{
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
		CustomerUserID:         &data.CustomerUserID,
		CustomerReligion:       &data.CustomerReligion,
		CustomerLevelID:        &data.CustomerLevelID,
		CustomerGender:         &data.CustomerGender,
		CustomerBirthDate:      &data.CustomerBirthDate,
	}

	res.ID, err = repo.Add(c, &res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return res, err
}
