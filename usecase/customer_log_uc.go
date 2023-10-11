package usecase

import (
	"context"
	"encoding/json"
	"time"

	"nextbasis-service-v-0.1/db/repository"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/functioncaller"
	"nextbasis-service-v-0.1/pkg/logruslogger"
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

// CustomerLogUC ...
type CustomerLogUC struct {
	*ContractUC
}

// BuildBody ...
func (uc CustomerLogUC) BuildBody(data *models.CustomerLog, res *viewmodel.CustomerLogVM) {
	res.ID = data.ID
	res.CustomerID = data.CustomerID
	res.CustomerCode = data.CustomerCode
	res.CustomerName = data.CustomerName
	res.UserID = data.UserID
	res.UserName = data.UserName
	res.CreatedAt = data.CreatedAt

	var oldData, newData viewmodel.CustomerVM
	json.Unmarshal([]byte(data.OldData), &oldData)
	json.Unmarshal([]byte(data.NewData), &newData)
	var typeChanges, oldDataChanges, newDataChanges string
	if *oldData.CustomerPhone != *newData.CustomerPhone {
		typeChanges += "Phone Number"
		oldDataChanges += *oldData.CustomerPhone
		newDataChanges += *newData.CustomerPhone
	}

	if *oldData.CustomerName != *newData.CustomerName {
		if typeChanges != "" {
			typeChanges += ", "
			oldDataChanges += ", "
			newDataChanges += ", "
		}
		typeChanges += "Name"
		oldDataChanges += *oldData.CustomerName
		newDataChanges += *newData.CustomerName
	}

	if *oldData.CustomerNik != *newData.CustomerNik {
		if typeChanges != "" {
			typeChanges += ", "
			oldDataChanges += ", "
			newDataChanges += ", "
		}
		typeChanges += "NIK"
		oldDataChanges += *oldData.CustomerNik
		newDataChanges += *newData.CustomerNik
	}

	if *oldData.CustomerReligion != *newData.CustomerReligion {
		if typeChanges != "" {
			typeChanges += ", "
			oldDataChanges += ", "
			newDataChanges += ", "
		}
		typeChanges += "Religion"
		oldDataChanges += *oldData.CustomerReligion
		newDataChanges += *newData.CustomerReligion
	}

	oldBirthdate, _ := time.Parse("2006-01-02", *oldData.CustomerBirthDate)
	newBirthdate, _ := time.Parse("2006-01-02", *newData.CustomerBirthDate)

	if oldBirthdate != newBirthdate {
		if typeChanges != "" {
			typeChanges += ", "
			oldDataChanges += ", "
			newDataChanges += ", "
		}
		typeChanges += "Date of Birth"
		oldDataChanges += oldBirthdate.Format("2006-01-02")
		newDataChanges += newBirthdate.Format("2006-01-02")
	}

	res.TypeChanges = typeChanges
	res.OldData = oldDataChanges
	res.NewData = newDataChanges
}

// SelectAll ...
func (uc CustomerLogUC) SelectAll(c context.Context, parameter models.CustomerLogParameter) (res []viewmodel.CustomerLogVM, err error) {
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(0, 0, parameter.By, parameter.Sort, models.WebCustomerOrderBy, models.WebCustomerOrderByrByString)

	repo := repository.NewCustomerLogRepository(uc.DB)
	data, err := repo.SelectAll(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	for i := range data {
		var temp viewmodel.CustomerLogVM

		uc.BuildBody(&data[i], &temp)
		if temp.TypeChanges != "" {
			res = append(res, temp)
		}
	}

	return res, err
}

func (uc CustomerLogUC) Add(c context.Context, oldData, newData interface{}, customerID string, userID int) (err error) {
	repo := repository.NewCustomerLogRepository(uc.DB)
	err = repo.Add(c, oldData, newData, customerID, userID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return err
	}

	return err
}
