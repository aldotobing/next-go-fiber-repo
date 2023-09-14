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
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

// CilentInvoiceUC ...
type CilentInvoiceUC struct {
	*ContractUC
}

// BuildBody ...
func (uc CilentInvoiceUC) BuildBody(res *models.CilentInvoice) {
}

// SelectAll ...
func (uc CilentInvoiceUC) SelectAll(c context.Context, parameter models.CilentInvoiceParameter) (res []models.CilentInvoice, err error) {
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(0, 0, parameter.By, parameter.Sort, models.CilentInvoiceOrderBy, models.CilentInvoiceOrderByrByString)

	repo := repository.NewCilentInvoiceRepository(uc.DB)
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
func (uc CilentInvoiceUC) FindAll(c context.Context, parameter models.CilentInvoiceParameter) (res []models.CilentInvoice, p viewmodel.PaginationVM, err error) {
	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.CilentInvoiceOrderBy, models.CilentInvoiceOrderByrByString)

	var count int
	repo := repository.NewCilentInvoiceRepository(uc.DB)
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
func (uc CilentInvoiceUC) FindByID(c context.Context, parameter models.CilentInvoiceParameter) (res models.CilentInvoice, err error) {
	repo := repository.NewCilentInvoiceRepository(uc.DB)
	res, err = repo.FindByID(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}
	uc.BuildBody(&res)

	return res, err
}

// FindByDocumentNo ...
func (uc CilentInvoiceUC) FindByDocumentNo(c context.Context, parameter models.CilentInvoiceParameter) (res models.CilentInvoice, err error) {
	repo := repository.NewCilentInvoiceRepository(uc.DB)
	res, err = repo.FindByID(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}
	uc.BuildBody(&res)

	return res, err
}

// SelectAll ...
// func (uc CilentInvoiceUC) DataSync(c context.Context, parameter models.CilentInvoiceParameter) (res []models.CilentInvoice, err error) {
// 	repo := repository.NewCilentInvoiceRepository(uc.DB)

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

// 	var resBuilder []models.CilentInvoice
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

func (uc CilentInvoiceUC) DataSync(c context.Context, parameter models.CilentInvoiceParameter) ([]models.CilentInvoice, error) {
	repo := repository.NewCilentInvoiceRepository(uc.DB)

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
	req, err := http.NewRequest("GET", "http://nextbasis.id:8080/mysmagonsrv/rest/salesInvoice/data/2", bytes.NewBuffer(jsonReq))
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

	var res []models.CilentInvoice
	if err := json.Unmarshal(bodyBytes, &res); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	var resBuilder []models.CilentInvoice
	for _, invoiceObject := range res {
		_, err := repo.InsertDataWithLine(c, &invoiceObject)
		if err != nil {
			return nil, fmt.Errorf("failed to insert data for invoice %+v: %w", invoiceObject, err)
		}
		resBuilder = append(resBuilder, invoiceObject)
	}

	return resBuilder, nil
}

func (uc CilentInvoiceUC) UndoneDataSync(c context.Context, parameter models.CilentInvoiceParameter) ([]models.CilentInvoice, error) {
	repo := repository.NewCilentInvoiceRepository(uc.DB)

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

	var res []models.CilentInvoice
	if err := json.Unmarshal(bodyBytes, &res); err != nil {
		return nil, fmt.Errorf("failed to unmarshal undone response: %w", err)
	}

	var resBuilder []models.CilentInvoice
	for _, invoiceObject := range res {
		_, err := repo.InsertDataWithLine(c, &invoiceObject)
		if err != nil {
			// return nil, fmt.Errorf("failed to insert data for invoice %+v: %w", invoiceObject, err)
		}
		resBuilder = append(resBuilder, invoiceObject)
	}

	return resBuilder, nil
}
