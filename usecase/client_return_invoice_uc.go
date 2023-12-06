package usecase

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"nextbasis-service-v-0.1/db/repository"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/functioncaller"
	"nextbasis-service-v-0.1/pkg/logruslogger"
)

// CilentReturnInvoiceUC ...
type CilentReturnInvoiceUC struct {
	*ContractUC
}

// BuildBody ...
func (uc CilentReturnInvoiceUC) BuildBody(res *models.CilentReturnInvoice) {
}

// SelectAll ...
func (uc CilentReturnInvoiceUC) SelectAll(c context.Context, parameter models.CilentReturnInvoiceParameter) (res []models.CilentReturnInvoice, err error) {
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(0, 0, parameter.By, parameter.Sort, models.CilentReturnInvoiceOrderBy, models.CilentReturnInvoiceOrderByrByString)

	repo := repository.NewCilentReturnInvoiceRepository(uc.DB)
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

// FindByID ...
func (uc CilentReturnInvoiceUC) FindByID(c context.Context, parameter models.CilentReturnInvoiceParameter) (res models.CilentReturnInvoice, err error) {
	repo := repository.NewCilentReturnInvoiceRepository(uc.DB)
	res, err = repo.FindByID(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}
	uc.BuildBody(&res)

	return res, err
}

// FindByDocumentNo ...
func (uc CilentReturnInvoiceUC) FindByDocumentNo(c context.Context, parameter models.CilentReturnInvoiceParameter) (res models.CilentReturnInvoice, err error) {
	repo := repository.NewCilentReturnInvoiceRepository(uc.DB)
	res, err = repo.FindByID(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}
	uc.BuildBody(&res)

	return res, err
}

// SelectAll ...
// func (uc CilentReturnInvoiceUC) DataSync(c context.Context, parameter models.CilentReturnInvoiceParameter) (res []models.CilentReturnInvoice, err error) {
// 	repo := repository.NewCilentReturnInvoiceRepository(uc.DB)

// 	loc, _ := time.LoadLocation("Asia/Jakarta")
// 	now := time.Now().In(loc).Add(time.Minute * time.Duration(-30))
// 	strnow := now.Format(time.RFC3339)
// 	parameter.DateParam = strnow
// 	jsonReq, err := json.Marshal(parameter)
// 	client := &http.Client{}
// 	req, err := http.NewRequest("GET", "http://nextbasis.id:8080/mysmagonsrv/rest/salesInvoice/data/2", bytes.NewBuffer(jsonReq))
// 	if err != nil {
// 		fmt.Println("client err")
// 		fmt.Print(err.Error())
// 	}

// 	req.Header.Add("Accept", "application/json")
// 	req.Header.Add("Content-Type", "application/json")
// 	req.Header.Add("Authorization", "Bearer C2A5CE6A2292E7745CE5A3F7E68A9")

// 	resp, err := client.Do(req)
// 	if err != nil {

// 		fmt.Print(err.Error())
// 	}
// 	defer resp.Body.Close()
// 	bodyBytes, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Print(err.Error())
// 	}
// 	// var responseObject http.Response
// 	json.Unmarshal(bodyBytes, &res)
// 	// fmt.Printf("API Response as struct %+v\n", &responseObject)

// 	var resBuilder []models.CilentReturnInvoice
// 	for _, invoiceObject := range res {
// 		fmt.Printf("%s\n", *invoiceObject.ID)

// 		_, errinsert := repo.InsertDataWithLine(c, &invoiceObject)

// 		if errinsert != nil {
// 			fmt.Print(errinsert)
// 		}

// 		resBuilder = append(resBuilder, invoiceObject)

// 	}

// 	return resBuilder, err
// }

func (uc CilentReturnInvoiceUC) DataSync(c context.Context, parameter models.CilentReturnInvoiceParameter) ([]models.CilentReturnInvoice, error) {
	repo := repository.NewCilentReturnInvoiceRepository(uc.DB)

	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		return nil, fmt.Errorf("failed to load location: %w", err)
	}

	now := time.Now().In(loc).Add(time.Minute * time.Duration(-30))
	strnow := now.Format(time.RFC3339)
	parameter.DateParam = strnow

	fmt.Println("Get Date" + parameter.DateParam)

	jsonReq, err := json.Marshal(parameter)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal json: %w", err)
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://nextbasis.id:8080/mysmagonsrv/rest/salesReturnInvoice/data", bytes.NewBuffer(jsonReq))
	if err != nil {
		return nil, fmt.Errorf("failed to create new request: %w", err)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer C2A5CE6A2292E7745CE5A3F7E68A9")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}

	if resp == nil || resp.Body == nil {
		return nil, errors.New("response or response body is nil")
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	fmt.Println("Response Body:", string(bodyBytes))

	var res []models.CilentReturnInvoice
	if err := json.Unmarshal(bodyBytes, &res); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	var resBuilder []models.CilentReturnInvoice
	for _, invoiceObject := range res {
		_, err := repo.InsertDataWithLine(c, &invoiceObject)
		if err != nil {
			// return nil, fmt.Errorf("failed to insert data for invoice %+v: %w", invoiceObject, err)
		}
		resBuilder = append(resBuilder, invoiceObject)
	}

	return resBuilder, nil
}

func (uc CilentReturnInvoiceUC) UndoneDataSync(c context.Context, parameter models.CilentReturnInvoiceParameter) ([]models.CilentReturnInvoice, error) {
	repo := repository.NewCilentReturnInvoiceRepository(uc.DB)

	// parameter.StartDate = `2023-08-01`
	// parameter.EndDate = `2023-08-31`

	jsonReq, err := json.Marshal(parameter)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal json: %w", err)
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://nextbasis.id:8080/mysmagonsrv/rest/salesInvoice/transaction", bytes.NewBuffer(jsonReq))
	if err != nil {
		return nil, fmt.Errorf("failed to create new undone request: %w", err)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer C2A5CE6A2292E7745CE5A3F7E68A9")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute undone request: %w", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read undone response body: %w", err)
	}

	var res []models.CilentReturnInvoice
	if err := json.Unmarshal(bodyBytes, &res); err != nil {
		return nil, fmt.Errorf("failed to unmarshal undone response: %w", err)
	}

	var resBuilder []models.CilentReturnInvoice
	for _, invoiceObject := range res {
		_, err := repo.InsertDataWithLine(c, &invoiceObject)
		if err != nil {
			// return nil, fmt.Errorf("failed to insert data for invoice %+v: %w", invoiceObject, err)
		}
		resBuilder = append(resBuilder, invoiceObject)
	}

	return resBuilder, nil
}
