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
)

// SalesOrderCustomerSyncUC ...
type SalesOrderCustomerSyncUC struct {
	*ContractUC
}

// BuildBody ...
func (uc SalesOrderCustomerSyncUC) BuildBody(res *models.SalesOrderCustomerSync) {
}

// FindByID ...
func (uc SalesOrderCustomerSyncUC) FindByID(c context.Context, parameter models.SalesOrderCustomerSyncParameter) (res models.SalesOrderCustomerSync, err error) {
	repo := repository.NewSalesOrderCustomerSyncRepository(uc.DB)
	res, err = repo.FindByID(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}
	uc.BuildBody(&res)

	return res, err
}

// FindByDocumentNo ...
func (uc SalesOrderCustomerSyncUC) FindByDocumentNo(c context.Context, parameter models.SalesOrderCustomerSyncParameter) (res models.SalesOrderCustomerSync, err error) {
	repo := repository.NewSalesOrderCustomerSyncRepository(uc.DB)
	res, err = repo.FindByID(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}
	uc.BuildBody(&res)

	return res, err
}

// SelectAll ...
func (uc SalesOrderCustomerSyncUC) DataSync(c context.Context, parameter models.SalesOrderCustomerSyncParameter) (res []models.SalesOrderCustomerSync, err error) {
	repo := repository.NewSalesOrderCustomerSyncRepository(uc.DB)

	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc).Add(time.Minute * time.Duration(-15))
	strnow := now.Format(time.RFC3339)
	parameter.DateParam = strnow
	parameter.Status = "submitted"
	jsonReq, err := json.Marshal(parameter)
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8084/NEXTbasis-service-agon/rest/salesOrder/data/online_store", bytes.NewBuffer(jsonReq))
	if err != nil {
		fmt.Println("client err")
		fmt.Print(err.Error())
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer C2A5CE6A2292E7745CE5A3F7E68A9")

	resp, err := client.Do(req)
	if err != nil {

		fmt.Print(err.Error())
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}
	// var responseObject http.Response
	json.Unmarshal(bodyBytes, &res)
	// fmt.Printf("API Response as struct %+v\n", &responseObject)

	var resBuilder []models.SalesOrderCustomerSync
	for _, invoiceObject := range res {
		fmt.Printf("%s\n", *invoiceObject.ID)

		_, errinsert := repo.InsertDataWithLine(c, &invoiceObject)

		if errinsert != nil {
			fmt.Print(errinsert)
		}

		resBuilder = append(resBuilder, invoiceObject)

	}

	return resBuilder, err
}

func (uc SalesOrderCustomerSyncUC) RevisedSync(c context.Context, parameter models.SalesOrderCustomerSyncParameter) (res []models.SalesOrderCustomerSync, err error) {
	repo := repository.NewSalesOrderCustomerSyncRepository(uc.DB)

	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc).Add(time.Minute * time.Duration(-15))
	strnow := now.Format(time.RFC3339)
	parameter.DateParam = strnow
	parameter.Status = "revised"
	parameter.HeaderOnly = "1"
	jsonReq, err := json.Marshal(parameter)
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8084/NEXTbasis-service-agon/rest/salesOrder/data/online_store", bytes.NewBuffer(jsonReq))
	if err != nil {
		fmt.Println("client err")
		fmt.Print(err.Error())
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer C2A5CE6A2292E7745CE5A3F7E68A9")

	resp, err := client.Do(req)
	if err != nil {

		fmt.Print(err.Error())
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}
	// var responseObject http.Response
	json.Unmarshal(bodyBytes, &res)
	// fmt.Printf("API Response as struct %+v\n", &responseObject)

	var resBuilder []models.SalesOrderCustomerSync
	for _, invoiceObject := range res {
		fmt.Printf("%s\n", *invoiceObject.ID)

		_, errinsert := repo.RevisedSync(c, &invoiceObject)

		if errinsert != nil {
			fmt.Print(errinsert)
		}

		resBuilder = append(resBuilder, invoiceObject)

	}

	return resBuilder, err
}
