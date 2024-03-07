package usecase

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"

	"nextbasis-service-v-0.1/db/repository"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/functioncaller"
	"nextbasis-service-v-0.1/pkg/logruslogger"
	"nextbasis-service-v-0.1/server/requests"
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

	var res []models.CilentInvoice
	if err := json.Unmarshal([]byte(bodyBytes), &res); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	var resBuilder []models.CilentInvoice
	for _, invoiceObject := range res {
		_, _, err := repo.InsertDataWithLine(c, &invoiceObject)
		if err != nil {
			return nil, fmt.Errorf("failed to insert data for invoice %+v: %w", invoiceObject, err)
		}

		if invoiceObject.SalesRequestCode != nil && strings.Contains(*invoiceObject.SalesRequestCode, "CO") &&
			invoiceObject.OutstandingAmount != nil && *invoiceObject.OutstandingAmount == "0.00" &&
			invoiceObject.CustomerCode != nil && invoiceObject.NetAmount != nil &&
			invoiceObject.DocumentNo != nil {
			customer, _ := WebCustomerUC{ContractUC: uc.ContractUC}.FindByCodes(c, models.WebCustomerParameter{Code: `'` + *invoiceObject.CustomerCode + `'`})
			if len(customer) == 1 {
				if customer[0].IndexPoint == 1 {
					pointRules, _ := PointRuleUC{ContractUC: uc.ContractUC}.SelectAll(c, models.PointRuleParameter{
						Now:  time.Now().Format("2006-01-02"),
						By:   "def.id",
						Sort: "asc",
					})
					pointUC := PointUC{ContractUC: uc.ContractUC}
					invoiceDate, _ := time.Parse("2006-01-02 15:04:05.999999999", *invoiceObject.InvoiceDate)
					pointThisMonth, _ := pointUC.GetPointThisMonth(c, customer[0].ID, invoiceDate.Month().String(), strconv.Itoa(invoiceDate.Year()))
					for _, rules := range pointRules {
						pointMonthly, _ := strconv.ParseFloat(pointThisMonth.Balance, 64)

						var maxMonthly float64
						if customer[0].MonthlyMaxPoint != "" && customer[0].MonthlyMaxPoint != "0" {
							maxMonthly, _ = strconv.ParseFloat(customer[0].MonthlyMaxPoint, 64)
						} else {
							maxMonthly, _ = strconv.ParseFloat(rules.MonthlyMaxPoint, 64)
						}

						minOrder, _ := strconv.ParseFloat(rules.MinOrder, 64)
						netOmount, _ := strconv.ParseFloat(*invoiceObject.NetAmount, 64)

						pointConversion, _ := strconv.ParseFloat(rules.PointConversion, 64)
						getPoint := math.Floor(netOmount/minOrder) * pointConversion

						if pointMonthly+getPoint > maxMonthly {
							getPoint = maxMonthly - pointMonthly
						}

						if getPoint > 0 {
							pointUC.Add(c, requests.PointRequest{
								CustomerCodes: []requests.PointCustomerCode{
									{CustomerCode: customer[0].Code},
								},
								InvoiceDocumentNo: *invoiceObject.DocumentNo,
								Point:             strconv.FormatFloat(getPoint, 'f', 0, 64),
								PointType:         "2",
							})
						}
					}
				}
			}
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
		_, _, err := repo.InsertDataWithLine(c, &invoiceObject)
		if err != nil {
			// return nil, fmt.Errorf("failed to insert data for invoice %+v: %w", invoiceObject, err)
		}

		resBuilder = append(resBuilder, invoiceObject)
	}

	return resBuilder, nil
}

func (uc CilentInvoiceUC) PutRedisDataSync(c context.Context, parameter models.CilentInvoiceParameter) ([]models.CilentInvoice, error) {

	// fmt.Println("put redis")
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		return nil, fmt.Errorf("failed to load location: %w", err)
	}

	now := time.Now().In(loc).Add(time.Minute * time.Duration(-30))
	strnow := now.Format(time.RFC3339)
	parameter.DateParam = strnow

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

	var res []models.CilentInvoice
	if err := json.Unmarshal([]byte(bodyBytes), &res); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	var resBuilder []models.CilentInvoice
	for _, invoiceObject := range res {
		cacheKey := "invoice_header_:" + *invoiceObject.DocumentNo
		jsonData, err := json.Marshal(invoiceObject)
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "json_marshal", uc.ReqID)
			return res, err
		}
		err = uc.RedisClient.Client.Set(cacheKey, jsonData, time.Hour*168).Err()
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "redis_set", uc.ReqID)
			return res, err
		}
		// &invoiceObject
		resBuilder = append(resBuilder, invoiceObject)
	}

	return resBuilder, nil
}

func (uc CilentInvoiceUC) GetRedisDataSync(c context.Context) (res []models.CilentInvoice, err error) {
	cacheKey := "*invoice_header*"
	repo := repository.NewCilentInvoiceRepository(uc.DB)

	// Try to get data from Redis cache first
	strinvList, err := uc.RedisClient.GetAllKeyFromRedis(cacheKey)

	if err == nil {
		fmt.Println("list key ", strinvList)
		for i := 0; i < 125; i++ {

			key := strinvList[i]
			fmt.Println("key", key)
			invoiceObject := new(models.CilentInvoice)
			err = uc.RedisClient.GetFromRedis(key, &invoiceObject)
			if err != nil {
				fmt.Println(err)
			}
			if err == nil {
				fmt.Println("from redis : ", key)
				_, _, err := repo.InsertDataWithLine(c, invoiceObject)
				if err != nil {
					errstr := err.Error()
					if strings.Contains(errstr, "cust_bill_to_id") || strings.Contains(errstr, "uom_id") || strings.Contains(errstr, "item_id") ||
						strings.Contains(errstr, "more than one row returned by a subquery used as an expression") {
						cacheKeyerr := "err_cus_item_uom_inv:" + *invoiceObject.DocumentNo
						errjsonData, err := json.Marshal(invoiceObject)
						if err != nil {
							logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "json_marshal", uc.ReqID)
							return res, err
						}
						err = uc.RedisClient.Client.Set(cacheKeyerr, errjsonData, time.Hour*168).Err()
						if err != nil {
							logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "redis_set", uc.ReqID)
							return res, err
						}
					} else {
						return nil, fmt.Errorf("failed to insert data for invoice %+v: %w", invoiceObject, err)
					}
				}
				if err == nil {
					if invoiceObject.SalesRequestCode != nil && strings.Contains(*invoiceObject.SalesRequestCode, "CO") &&
						invoiceObject.OutstandingAmount != nil && *invoiceObject.OutstandingAmount == "0.00" &&
						invoiceObject.CustomerCode != nil && invoiceObject.NetAmount != nil &&
						invoiceObject.DocumentNo != nil {
						customer, _ := WebCustomerUC{ContractUC: uc.ContractUC}.FindByCodes(c, models.WebCustomerParameter{Code: `'` + *invoiceObject.CustomerCode + `'`})
						if len(customer) == 1 {
							if customer[0].IndexPoint == 1 {
								pointRules, _ := PointRuleUC{ContractUC: uc.ContractUC}.SelectAll(c, models.PointRuleParameter{
									Now:  time.Now().Format("2006-01-02"),
									By:   "def.id",
									Sort: "asc",
								})
								pointUC := PointUC{ContractUC: uc.ContractUC}
								invoiceDate, _ := time.Parse("2006-01-02 15:04:05.999999999", *invoiceObject.InvoiceDate)
								pointThisMonth, _ := pointUC.GetPointThisMonth(c, customer[0].ID, invoiceDate.Month().String(), strconv.Itoa(invoiceDate.Year()))
								for _, rules := range pointRules {
									pointMonthly, _ := strconv.ParseFloat(pointThisMonth.Balance, 64)

									var maxMonthly float64
									if customer[0].MonthlyMaxPoint != "" && customer[0].MonthlyMaxPoint != "0" {
										maxMonthly, _ = strconv.ParseFloat(customer[0].MonthlyMaxPoint, 64)
									} else {
										maxMonthly, _ = strconv.ParseFloat(rules.MonthlyMaxPoint, 64)
									}

									minOrder, _ := strconv.ParseFloat(rules.MinOrder, 64)
									netOmount, _ := strconv.ParseFloat(*invoiceObject.NetAmount, 64)

									pointConversion, _ := strconv.ParseFloat(rules.PointConversion, 64)
									getPoint := math.Floor(netOmount/minOrder) * pointConversion

									if pointMonthly+getPoint > maxMonthly {
										getPoint = maxMonthly - pointMonthly
									}

									if getPoint > 0 {
										pointUC.Add(c, requests.PointRequest{
											CustomerCodes: []requests.PointCustomerCode{
												{CustomerCode: customer[0].Code},
											},
											InvoiceDocumentNo: *invoiceObject.DocumentNo,
											Point:             strconv.FormatFloat(getPoint, 'f', 0, 64),
											PointType:         "2",
										})
									}
								}
							}
						}
					}
				}

				res = append(res, *invoiceObject)
				fmt.Println(key)
				_ = uc.RedisClient.Delete(key)
			}
		}

	}

	return res, nil
}

func (uc CilentInvoiceUC) GetRedisDataReserveSync(c context.Context) (res []models.CilentInvoice, err error) {
	cacheKey := "*invoice_header*"
	repo := repository.NewCilentInvoiceRepository(uc.DB)

	// Try to get data from Redis cache first
	strinvList, err := uc.RedisClient.GetAllKeyFromRedis(cacheKey)

	if err == nil {
		fmt.Println("list key ", strinvList)
		for i := range strinvList {
			key := strinvList[len(strinvList)-1-i]
			invoiceObject := new(models.CilentInvoice)
			err = uc.RedisClient.GetFromRedis(key, &invoiceObject)
			if err != nil {
				fmt.Println(err)
			}
			if err == nil {
				fmt.Println("from redis : ", key)
				_, _, err := repo.InsertDataWithLine(c, invoiceObject)
				if err != nil {
					errstr := err.Error()
					if strings.Contains(errstr, "cust_bill_to_id") || strings.Contains(errstr, "uom_id") || strings.Contains(errstr, "item_id") ||
						strings.Contains(errstr, "more than one row returned by a subquery used as an expression") {
						cacheKeyerr := "err_cus_item_uom_inv:" + *invoiceObject.DocumentNo
						errjsonData, err := json.Marshal(invoiceObject)
						if err != nil {
							logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "json_marshal", uc.ReqID)
							return res, err
						}
						err = uc.RedisClient.Client.Set(cacheKeyerr, errjsonData, time.Hour*168).Err()
						if err != nil {
							logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "redis_set", uc.ReqID)
							return res, err
						}
					} else {
						return nil, fmt.Errorf("failed to insert data for invoice %+v: %w", invoiceObject, err)
					}
					// jsonDatas, errj := json.Marshal(invoiceObject)
					// if errj != nil {
					// 	logruslogger.Log(logruslogger.WarnLevel, errj.Error(), functioncaller.PrintFuncName(), "json_marshal", uc.ReqID)
					// 	// return res, err
					// }
					// errrd := uc.RedisClient.Client.Set("error_data_inv_:"+*invoiceObject.DocumentNo, jsonDatas, time.Hour*24).Err()
					// if errrd != nil {
					// 	logruslogger.Log(logruslogger.WarnLevel, errrd.Error(), functioncaller.PrintFuncName(), "redis_set", uc.ReqID)
					// 	// return res, err
					// }

				}
				if err == nil {
					if invoiceObject.SalesRequestCode != nil && strings.Contains(*invoiceObject.SalesRequestCode, "CO") &&
						invoiceObject.OutstandingAmount != nil && *invoiceObject.OutstandingAmount == "0.00" &&
						invoiceObject.CustomerCode != nil && invoiceObject.NetAmount != nil &&
						invoiceObject.DocumentNo != nil {
						customer, _ := WebCustomerUC{ContractUC: uc.ContractUC}.FindByCodes(c, models.WebCustomerParameter{Code: `'` + *invoiceObject.CustomerCode + `'`})
						if len(customer) == 1 {
							if customer[0].IndexPoint == 1 {
								invoiceDate, _ := time.Parse("2006-01-02 15:04:05.999999999", *invoiceObject.InvoiceDate)

								pointRules, _ := PointRuleUC{ContractUC: uc.ContractUC}.SelectAll(c, models.PointRuleParameter{
									Now:  invoiceDate.Format("2006-01-02"),
									By:   "def.id",
									Sort: "asc",
								})
								pointUC := PointUC{ContractUC: uc.ContractUC}

								

								pointThisMonth, _ := pointUC.GetPointThisMonth(c, customer[0].ID, invoiceDate.Month().String(), strconv.Itoa(invoiceDate.Year()))
								for _, rules := range pointRules {
									pointMonthly, _ := strconv.ParseFloat(pointThisMonth.Balance, 64)

									var maxMonthly float64
									if customer[0].MonthlyMaxPoint != "" {
										maxMonthly, _ = strconv.ParseFloat(customer[0].MonthlyMaxPoint, 64)
									} else {
										maxMonthly, _ = strconv.ParseFloat(rules.MonthlyMaxPoint, 64)
									}

									minOrder, _ := strconv.ParseFloat(rules.MinOrder, 64)
									netOmount, _ := strconv.ParseFloat(*invoiceObject.NetAmount, 64)

									pointConversion, _ := strconv.ParseFloat(rules.PointConversion, 64)
									getPoint := math.Floor(netOmount/minOrder) * pointConversion

									if pointMonthly+getPoint > maxMonthly {
										getPoint = maxMonthly - pointMonthly
									}

									if getPoint > 0 {
										pointUC.Add(c, requests.PointRequest{
											CustomerCodes: []requests.PointCustomerCode{
												{CustomerCode: customer[0].Code},
											},
											InvoiceDocumentNo: *invoiceObject.DocumentNo,
											Point:             strconv.FormatFloat(getPoint, 'f', 0, 64),
											PointType:         "2",
										})
									}
								}
							}
						}
					}
				}

				res = append(res, *invoiceObject)
				fmt.Println(key)
				err = uc.RedisClient.Delete(key)
			}

		}

	}

	return res, nil
}
