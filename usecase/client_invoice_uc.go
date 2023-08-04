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
func (uc CilentInvoiceUC) DataSync(c context.Context, parameter models.CilentInvoiceParameter) (res []models.CilentInvoice, err error) {
	repo := repository.NewCilentInvoiceRepository(uc.DB)

	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc).Add(time.Minute * time.Duration(-30))
	strnow := now.Format(time.RFC3339)
	parameter.DateParam = strnow
	jsonReq, err := json.Marshal(parameter)
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://nextbasis.id:8080/mysmagonsrv/rest/salesInvoice/data/2", bytes.NewBuffer(jsonReq))
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

	var resBuilder []models.CilentInvoice
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
