package usecase

import (
	"bytes"
	"context"
	"encoding/json"
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
	// repo := repository.NewCilentInvoiceRepository(uc.DB)

	// loc, err := time.LoadLocation("Asia/Jakarta")
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to load location: %w", err)
	// }

	// now := time.Now().In(loc).Add(time.Minute * time.Duration(-30))
	// strnow := now.Format(time.RFC3339)
	// parameter.DateParam = strnow

	// fmt.Println("Get Date" + parameter.DateParam)

	// jsonReq, err := json.Marshal(parameter)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to marshal json: %w", err)
	// }

	// client := &http.Client{}
	// req, err := http.NewRequest("GET", "http://nextbasis.id:8080/mysmagonsrv/rest/salesInvoice/data/2", bytes.NewBuffer(jsonReq))
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to create new request: %w", err)
	// }

	// req.Header.Add("Accept", "application/json")
	// req.Header.Add("Content-Type", "application/json")
	// req.Header.Add("Authorization", "Bearer C2A5CE6A2292E7745CE5A3F7E68A9")

	// resp, err := client.Do(req)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to execute request: %w", err)
	// }

	// if resp == nil || resp.Body == nil {
	// 	return nil, errors.New("response or response body is nil")
	// }
	// defer resp.Body.Close()

	// bodyBytes, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to read response body: %w", err)
	// }

	// fmt.Println("Response Body:", string(bodyBytes))

	body := `[
        {
            "id_invoice_header": "3766153",
            "document_no": "3602AA-24010109",
            "document_type_id": "1",
            "transaction_date": "2024-01-02",
            "transaction_time": "17:19:14",
            "customer_id": "43227783",
            "customer_code": "3602DD-S0077",
            "customer_name": null,
            "tax_calc_method": "E",
            "salesman_id": "1207",
            "salesman_code": "SL3602A047",
            "srh_doc_no": "CO3602240102002",
            "customer_point": "154",
            "salesman_name": null,
            "payment_terms_id": "1",
            "payment_terms_name": null,
            "sales_order_id": "4400689",
            "company_id": "2",
            "branch_id": "32",
            "branch_name": null,
            "price_list_id": "2",
            "price_list_name": null,
            "price_list_version_id": "2",
            "price_list_version_name": null,
            "status": "submitted",
            "gross_amount": "1388086.80",
            "taxable_amount": "1388086.80",
            "tax_amount": "152689.55",
            "rounding_amount": "0.00",
            "outstanding_amount": "0.00",
            "net_amount": "1540776.35",
            "due_date": null,
            "disc_amount": "0.00",
            "paid_amount": "1540776.60",
            "no_ppn": null,
            "global_disc_amount": "0.00",
            "list_line": [
                {
                    "id_line": "15225677",
                    "header_id": "5576722",
                    "line_no": "10",
                    "category_id": "1",
                    "item_id": "268",
                    "item_code": "KBENAG00BX",
                    "qty": "720.00",
                    "uom_id": "7",
                    "uom_code": "SACHET",
                    "stock_qty": "720.00",
                    "unit_price": "831.00",
                    "gross_amount": "598320.00",
                    "use_disc_percent": "0",
                    "disc_percent1": "0.00",
                    "disc_percent2": "0.00",
                    "disc_percent3": "0.00",
                    "disc_percent4": "0.00",
                    "disc_percent5": "0.00",
                    "disc_amount": "0.00",
                    "taxable_amount": "598320.00",
                    "tax_amount": "65815.20",
                    "rounding_amount": "0.00",
                    "net_amount": "664135.20",
                    "sales_order_line_id": null,
                    "debt": null,
                    "paid": null
                },
                {
                    "id_line": "15225678",
                    "header_id": "5576722",
                    "line_no": "20",
                    "category_id": "3",
                    "item_id": "394",
                    "item_code": "JAHESS00R",
                    "qty": "120.00",
                    "uom_id": "7",
                    "uom_code": "SACHET",
                    "stock_qty": "120.00",
                    "unit_price": "1203.59",
                    "gross_amount": "144430.80",
                    "use_disc_percent": "0",
                    "disc_percent1": "0.00",
                    "disc_percent2": "0.00",
                    "disc_percent3": "0.00",
                    "disc_percent4": "0.00",
                    "disc_percent5": "0.00",
                    "disc_amount": "0.00",
                    "taxable_amount": "144430.80",
                    "tax_amount": "15887.39",
                    "rounding_amount": "0.00",
                    "net_amount": "160318.19",
                    "sales_order_line_id": null,
                    "debt": null,
                    "paid": null
                },
                {
                    "id_line": "15225679",
                    "header_id": "5576722",
                    "line_no": "30",
                    "category_id": "28",
                    "item_id": "245",
                    "item_code": "00TANACRD",
                    "qty": "60.00",
                    "uom_id": "7",
                    "uom_code": "SACHET",
                    "stock_qty": "60.00",
                    "unit_price": "2169.00",
                    "gross_amount": "130140.00",
                    "use_disc_percent": "0",
                    "disc_percent1": "0.00",
                    "disc_percent2": "0.00",
                    "disc_percent3": "0.00",
                    "disc_percent4": "0.00",
                    "disc_percent5": "0.00",
                    "disc_amount": "0.00",
                    "taxable_amount": "130140.00",
                    "tax_amount": "14315.40",
                    "rounding_amount": "0.00",
                    "net_amount": "144455.40",
                    "sales_order_line_id": null,
                    "debt": null,
                    "paid": null
                },
                {
                    "id_line": "15225680",
                    "header_id": "5576722",
                    "line_no": "40",
                    "category_id": "2",
                    "item_id": "82",
                    "item_code": "00TA00CRD12",
                    "qty": "120.00",
                    "uom_id": "7",
                    "uom_code": "SACHET",
                    "stock_qty": "120.00",
                    "unit_price": "3042.00",
                    "gross_amount": "365040.00",
                    "use_disc_percent": "0",
                    "disc_percent1": "0.00",
                    "disc_percent2": "0.00",
                    "disc_percent3": "0.00",
                    "disc_percent4": "0.00",
                    "disc_percent5": "0.00",
                    "disc_amount": "0.00",
                    "taxable_amount": "365040.00",
                    "tax_amount": "40154.40",
                    "rounding_amount": "0.00",
                    "net_amount": "405194.40",
                    "sales_order_line_id": null,
                    "debt": null,
                    "paid": null
                },
                {
                    "id_line": "15225681",
                    "header_id": "5576722",
                    "line_no": "50",
                    "category_id": "22",
                    "item_id": "391",
                    "item_code": "KOPIJH00R",
                    "qty": "120.00",
                    "uom_id": "7",
                    "uom_code": "SACHET",
                    "stock_qty": "120.00",
                    "unit_price": "1251.30",
                    "gross_amount": "150156.00",
                    "use_disc_percent": "0",
                    "disc_percent1": "0.00",
                    "disc_percent2": "0.00",
                    "disc_percent3": "0.00",
                    "disc_percent4": "0.00",
                    "disc_percent5": "0.00",
                    "disc_amount": "0.00",
                    "taxable_amount": "150156.00",
                    "tax_amount": "16517.16",
                    "rounding_amount": "0.00",
                    "net_amount": "166673.16",
                    "sales_order_line_id": null,
                    "debt": null,
                    "paid": null
                }
            ],
            "invoice_date": "2024-01-03 17:19:14.920762",
            "paid_date": "2024-01-03"
        },
        {
            "id_invoice_header": "3766154",
            "document_no": "3602AA-24010091",
            "document_type_id": "1",
            "transaction_date": "2024-01-02",
            "transaction_time": "17:19:14",
            "customer_id": "43227782",
            "customer_code": "3602DD-U0088",
            "customer_name": null,
            "tax_calc_method": "E",
            "salesman_id": "1207",
            "salesman_code": "SL3602A047",
            "srh_doc_no": "CO3602240102001",
            "customer_point": "66",
            "salesman_name": null,
            "payment_terms_id": "1",
            "payment_terms_name": null,
            "sales_order_id": "4400591",
            "company_id": "2",
            "branch_id": "32",
            "branch_name": null,
            "price_list_id": "2",
            "price_list_name": null,
            "price_list_version_id": "2",
            "price_list_version_name": null,
            "status": "submitted",
            "gross_amount": "602854.80",
            "taxable_amount": "602854.80",
            "tax_amount": "66314.03",
            "rounding_amount": "0.00",
            "outstanding_amount": "0.00",
            "net_amount": "669168.83",
            "due_date": null,
            "disc_amount": "0.00",
            "paid_amount": "669169.08",
            "no_ppn": null,
            "global_disc_amount": "0.00",
            "list_line": [
                {
                    "id_line": "15225672",
                    "header_id": "5576721",
                    "line_no": "10",
                    "category_id": "3",
                    "item_id": "394",
                    "item_code": "JAHESS00R",
                    "qty": "120.00",
                    "uom_id": "7",
                    "uom_code": "SACHET",
                    "stock_qty": "120.00",
                    "unit_price": "1203.59",
                    "gross_amount": "144430.80",
                    "use_disc_percent": "0",
                    "disc_percent1": "0.00",
                    "disc_percent2": "0.00",
                    "disc_percent3": "0.00",
                    "disc_percent4": "0.00",
                    "disc_percent5": "0.00",
                    "disc_amount": "0.00",
                    "taxable_amount": "144430.80",
                    "tax_amount": "15887.39",
                    "rounding_amount": "0.00",
                    "net_amount": "160318.19",
                    "sales_order_line_id": null,
                    "debt": null,
                    "paid": null
                },
                {
                    "id_line": "15225673",
                    "header_id": "5576721",
                    "line_no": "20",
                    "category_id": "1",
                    "item_id": "268",
                    "item_code": "KBENAG00BX",
                    "qty": "120.00",
                    "uom_id": "7",
                    "uom_code": "SACHET",
                    "stock_qty": "120.00",
                    "unit_price": "831.00",
                    "gross_amount": "99720.00",
                    "use_disc_percent": "0",
                    "disc_percent1": "0.00",
                    "disc_percent2": "0.00",
                    "disc_percent3": "0.00",
                    "disc_percent4": "0.00",
                    "disc_percent5": "0.00",
                    "disc_amount": "0.00",
                    "taxable_amount": "99720.00",
                    "tax_amount": "10969.20",
                    "rounding_amount": "0.00",
                    "net_amount": "110689.20",
                    "sales_order_line_id": null,
                    "debt": null,
                    "paid": null
                },
                {
                    "id_line": "15225674",
                    "header_id": "5576721",
                    "line_no": "30",
                    "category_id": "28",
                    "item_id": "245",
                    "item_code": "00TANACRD",
                    "qty": "12.00",
                    "uom_id": "7",
                    "uom_code": "SACHET",
                    "stock_qty": "12.00",
                    "unit_price": "2169.00",
                    "gross_amount": "26028.00",
                    "use_disc_percent": "0",
                    "disc_percent1": "0.00",
                    "disc_percent2": "0.00",
                    "disc_percent3": "0.00",
                    "disc_percent4": "0.00",
                    "disc_percent5": "0.00",
                    "disc_amount": "0.00",
                    "taxable_amount": "26028.00",
                    "tax_amount": "2863.08",
                    "rounding_amount": "0.00",
                    "net_amount": "28891.08",
                    "sales_order_line_id": null,
                    "debt": null,
                    "paid": null
                },
                {
                    "id_line": "15225675",
                    "header_id": "5576721",
                    "line_no": "40",
                    "category_id": "2",
                    "item_id": "82",
                    "item_code": "00TA00CRD12",
                    "qty": "60.00",
                    "uom_id": "7",
                    "uom_code": "SACHET",
                    "stock_qty": "60.00",
                    "unit_price": "3042.00",
                    "gross_amount": "182520.00",
                    "use_disc_percent": "0",
                    "disc_percent1": "0.00",
                    "disc_percent2": "0.00",
                    "disc_percent3": "0.00",
                    "disc_percent4": "0.00",
                    "disc_percent5": "0.00",
                    "disc_amount": "0.00",
                    "taxable_amount": "182520.00",
                    "tax_amount": "20077.20",
                    "rounding_amount": "0.00",
                    "net_amount": "202597.20",
                    "sales_order_line_id": null,
                    "debt": null,
                    "paid": null
                },
                {
                    "id_line": "15225676",
                    "header_id": "5576721",
                    "line_no": "50",
                    "category_id": "22",
                    "item_id": "391",
                    "item_code": "KOPIJH00R",
                    "qty": "120.00",
                    "uom_id": "7",
                    "uom_code": "SACHET",
                    "stock_qty": "120.00",
                    "unit_price": "1251.30",
                    "gross_amount": "150156.00",
                    "use_disc_percent": "0",
                    "disc_percent1": "0.00",
                    "disc_percent2": "0.00",
                    "disc_percent3": "0.00",
                    "disc_percent4": "0.00",
                    "disc_percent5": "0.00",
                    "disc_amount": "0.00",
                    "taxable_amount": "150156.00",
                    "tax_amount": "16517.16",
                    "rounding_amount": "0.00",
                    "net_amount": "166673.16",
                    "sales_order_line_id": null,
                    "debt": null,
                    "paid": null
                }
            ],
            "invoice_date": "2024-01-03 17:19:14.28773",
            "paid_date": "2024-01-03"
        }
    ]`

	var res []models.CilentInvoice
	if err := json.Unmarshal([]byte(body), &res); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	var resBuilder []models.CilentInvoice
	for _, invoiceObject := range res {
		// _, _, err := repo.InsertDataWithLine(c, &invoiceObject)
		// if err != nil {
		// 	return nil, fmt.Errorf("failed to insert data for invoice %+v: %w", invoiceObject, err)
		// }

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
					pointThisMonth, _ := pointUC.GetPointThisMonth(c, customer[0].ID)
					for _, rules := range pointRules {
						pointMonthly, _ := strconv.ParseFloat(pointThisMonth.Balance, 64)
						maxMonthly, _ := strconv.ParseFloat(rules.MonthlyMaxPoint, 64)

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
