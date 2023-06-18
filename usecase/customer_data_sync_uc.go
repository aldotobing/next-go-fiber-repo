package usecase

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"nextbasis-service-v-0.1/db/repository"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/functioncaller"
	"nextbasis-service-v-0.1/pkg/logruslogger"
	"nextbasis-service-v-0.1/server/requests"
)

// CustomerDataSyncUC ...
type CustomerDataSyncUC struct {
	*ContractUC
}

// BuildBody ...
func (uc CustomerDataSyncUC) BuildBody(res *models.CustomerDataSync) {
}

// FindByID ...
func (uc CustomerDataSyncUC) FindByCode(c context.Context, parameter models.CustomerDataSyncParameter) (res models.CustomerDataSync, err error) {
	repo := repository.NewCustomerDataSyncRepository(uc.DB)
	res, err = repo.FindByCode(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}
	uc.BuildBody(&res)

	return res, err
}

// Add ...
func (uc CustomerDataSyncUC) Add(c context.Context, data *requests.CustomerDataSyncRequest) (res models.CustomerDataSync, err error) {

	repo := repository.NewCustomerDataSyncRepository(uc.DB)
	// now := time.Now().UTC()
	// strnow := now.Format(time.RFC3339)
	res = models.CustomerDataSync{
		// Name: &data.Name,
	}
	res.ID, err = repo.Add(c, &res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return res, err
}

// Edit ,...
func (uc CustomerDataSyncUC) Edit(c context.Context, id string, data *requests.CustomerDataSyncRequest) (res models.CustomerDataSync, err error) {
	repo := repository.NewCustomerDataSyncRepository(uc.DB)
	res = models.CustomerDataSync{
		ID:                &id,
		Name:              data.Name,
		Address:           data.Address,
		PhoneNo:           data.PhoneNo,
		Code:              data.Code,
		TermOfPaymentCode: data.TermOfPaymentCode,
		PriceListCode:     data.PriceListCode,
		SalesmanCode:      data.SalesmanCode,
		BranchID:          data.BranchID,
	}

	res.ID, err = repo.Edit(c, &res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return res, err
}

// SelectAll ...
func (uc CustomerDataSyncUC) DataSync(c context.Context, parameter models.CustomerDataSyncParameter) (res []models.CustomerDataSync, err error) {
	repo := repository.NewCustomerDataSyncRepository(uc.DB)

	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc).Add(
		time.Minute * time.Duration(-15))
	strnow := now.Format(time.RFC3339)
	parameter.DateParam = strnow
	parameter.MysmOnly = "1"
	jsonReq, err := json.Marshal(parameter)
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://nextbasis.id:8080/mysmagonsrv/rest/mastercustomer/get", bytes.NewBuffer(jsonReq))
	if err != nil {
		fmt.Println("client err")
		fmt.Print(err.Error())
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer C2A5CE6A2292E7745CE5A3F7E68A9")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error client")
		fmt.Print(err.Error())
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error ")
		fmt.Print(err.Error())
	}

	// var responseObject http.Response
	json.Unmarshal(bodyBytes, &res)
	fmt.Printf("API Response as struct %+v\n", &res)

	var resBuilder []models.CustomerDataSync
	for _, itemObject := range res {
		fmt.Println("masuk perulangan")
		currentItem, _ := uc.FindByCode(c, models.CustomerDataSyncParameter{Code: *itemObject.Code})

		if currentItem.ID != nil {
			fmt.Println("not null")
			itemObject.ID = currentItem.ID
			_, errupdate := repo.Edit(c, &itemObject)
			if errupdate != nil {
				fmt.Print(errupdate)
			}
		} else {
			_, errinsert := repo.Add(c, &itemObject)
			if errinsert != nil {
				fmt.Print(errinsert)
			}
		}

		resBuilder = append(resBuilder, itemObject)
	}

	return res, err
}
