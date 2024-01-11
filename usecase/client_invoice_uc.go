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

	body := `[{
        "tax_amount": "322694.13",
        "customer_point": "325",
        "gross_amount": "2933583.00",
        "tax_calc_method": "E",
        "invoice_date": "2024-01-10 16:12:02.581246",
        "salesman_id": "78",
        "rounding_amount": "0.00",
        "branch_id": "1",
        "document_no": "350290764750",
        "sales_order_id": "4429106",
        "price_list_version_id": "9",
        "transaction_time": "16:12:02",
        "payment_terms_id": "23",
        "customer_code": "3502AP-K045",
        "taxable_amount": "2933583.00",
        "transaction_date": "2024-01-10",
        "global_disc_amount": "0.00",
        "company_id": "2",
        "salesman_code": "SL3502S17",
        "price_list_id": "9",
        "srh_doc_no": "CO3502240108014R01",
        "branch_code": "3502",
        "disc_amount": "0.00",
        "list_line": [
            {
                "disc_percent2": "0.00",
                "tax_amount": "11414.70",
                "item_code": "REMCR0D05I",
                "disc_percent1": "0.00",
                "header_id": "5611172",
                "item_id": "525",
                "line_no": "80",
                "use_disc_percent": "0",
                "gross_amount": "103770.00",
                "unit_price": "1383.60",
                "disc_percent5": "0.00",
                "disc_percent4": "0.00",
                "disc_percent3": "0.00",
                "uom_id": "2",
                "disc_amount": "0.00",
                "category_id": "26",
                "rounding_amount": "0.00",
                "qty": "75.00",
                "stock_qty": "75.00",
                "id_line": "15343382",
                "net_amount": "115184.70",
                "taxable_amount": "103770.00",
                "uom_code": "BKS"
            },
            {
                "disc_percent2": "0.00",
                "tax_amount": "6271.10",
                "item_code": "00TLM0CRD5",
                "disc_percent1": "0.00",
                "header_id": "5611172",
                "item_id": "474",
                "line_no": "70",
                "use_disc_percent": "0",
                "gross_amount": "57010.00",
                "unit_price": "2280.40",
                "disc_percent5": "0.00",
                "disc_percent4": "0.00",
                "disc_percent3": "0.00",
                "uom_id": "2",
                "disc_amount": "0.00",
                "category_id": "21",
                "rounding_amount": "0.00",
                "qty": "25.00",
                "stock_qty": "25.00",
                "id_line": "15343381",
                "net_amount": "63281.10",
                "taxable_amount": "57010.00",
                "uom_code": "BKS"
            },
            {
                "disc_percent2": "0.00",
                "tax_amount": "12542.20",
                "item_code": "00TL00CRD5",
                "disc_percent1": "0.00",
                "header_id": "5611172",
                "item_id": "452",
                "line_no": "60",
                "use_disc_percent": "0",
                "gross_amount": "114020.00",
                "unit_price": "2280.40",
                "disc_percent5": "0.00",
                "disc_percent4": "0.00",
                "disc_percent3": "0.00",
                "uom_id": "2",
                "disc_amount": "0.00",
                "category_id": "21",
                "rounding_amount": "0.00",
                "qty": "50.00",
                "stock_qty": "50.00",
                "id_line": "15343380",
                "net_amount": "126562.20",
                "taxable_amount": "114020.00",
                "uom_code": "BKS"
            },
            {
                "disc_percent2": "0.00",
                "tax_amount": "10190.40",
                "item_code": "RTLCOLD06A",
                "disc_percent1": "0.00",
                "header_id": "5611172",
                "item_id": "527",
                "line_no": "50",
                "use_disc_percent": "0",
                "gross_amount": "92640.00",
                "unit_price": "7720.00",
                "disc_percent5": "0.00",
                "disc_percent4": "0.00",
                "disc_percent3": "0.00",
                "uom_id": "3",
                "disc_amount": "0.00",
                "category_id": "21",
                "rounding_amount": "0.00",
                "qty": "12.00",
                "stock_qty": "12.00",
                "id_line": "15343379",
                "net_amount": "102830.40",
                "taxable_amount": "92640.00",
                "uom_code": "BTL"
            },
            {
                "disc_percent2": "0.00",
                "tax_amount": "28868.40",
                "item_code": "00TANACRD",
                "disc_percent1": "0.00",
                "header_id": "5611172",
                "item_id": "245",
                "line_no": "40",
                "use_disc_percent": "0",
                "gross_amount": "262440.00",
                "unit_price": "2187.00",
                "disc_percent5": "0.00",
                "disc_percent4": "0.00",
                "disc_percent3": "0.00",
                "uom_id": "2",
                "disc_amount": "0.00",
                "category_id": "28",
                "rounding_amount": "0.00",
                "qty": "120.00",
                "stock_qty": "120.00",
                "id_line": "15343378",
                "net_amount": "291308.40",
                "taxable_amount": "262440.00",
                "uom_code": "BKS"
            },
            {
                "disc_percent2": "0.00",
                "tax_amount": "243064.80",
                "item_code": "00TA00CRD12",
                "disc_percent1": "0.00",
                "header_id": "5611172",
                "item_id": "82",
                "line_no": "30",
                "use_disc_percent": "0",
                "gross_amount": "2209680.00",
                "unit_price": "3069.00",
                "disc_percent5": "0.00",
                "disc_percent4": "0.00",
                "disc_percent3": "0.00",
                "uom_id": "2",
                "disc_amount": "0.00",
                "category_id": "2",
                "rounding_amount": "0.00",
                "qty": "720.00",
                "stock_qty": "720.00",
                "id_line": "15343377",
                "net_amount": "2452744.80",
                "taxable_amount": "2209680.00",
                "uom_code": "BKS"
            },
            {
                "disc_percent2": "0.00",
                "tax_amount": "7584.28",
                "item_code": "00TA00PMS20",
                "disc_percent1": "0.00",
                "header_id": "5611172",
                "item_id": "242",
                "line_no": "20",
                "use_disc_percent": "0",
                "gross_amount": "68948.00",
                "unit_price": "34474.00",
                "disc_percent5": "0.00",
                "disc_percent4": "0.00",
                "disc_percent3": "0.00",
                "uom_id": "4",
                "disc_amount": "0.00",
                "category_id": "8",
                "rounding_amount": "0.00",
                "qty": "2.00",
                "stock_qty": "2.00",
                "id_line": "15343376",
                "net_amount": "76532.28",
                "taxable_amount": "68948.00",
                "uom_code": "STP"
            },
            {
                "disc_percent2": "0.00",
                "tax_amount": "2758.25",
                "item_code": "00JW0000DB",
                "disc_percent1": "0.00",
                "header_id": "5611172",
                "item_id": "280",
                "line_no": "10",
                "use_disc_percent": "0",
                "gross_amount": "25075.00",
                "unit_price": "1003.00",
                "disc_percent5": "0.00",
                "disc_percent4": "0.00",
                "disc_percent3": "0.00",
                "uom_id": "2",
                "disc_amount": "0.00",
                "category_id": "36",
                "rounding_amount": "0.00",
                "qty": "25.00",
                "stock_qty": "25.00",
                "id_line": "15343375",
                "net_amount": "27833.25",
                "taxable_amount": "25075.00",
                "uom_code": "BKS"
            }
        ],
        "id_invoice_header": "5611172",
        "document_type_id": "1",
        "paid_amount": "0.00",
        "net_amount": "3256277.13",
        "outstanding_amount": "0.00",
        "customer_id": "14500041",
        "status": "submitted"
    }]`

	var res []models.CilentInvoice
	if err := json.Unmarshal([]byte(body), &res); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	var resBuilder []models.CilentInvoice
	for _, invoiceObject := range res {
		// _, _, err := repo.InsertDataWithLine(c, &invoiceObject)
		// if err != nil {
		// 	// return nil, fmt.Errorf("failed to insert data for invoice %+v: %w", invoiceObject, err)
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
						pointMonthly = 45000
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

		if strings.Contains("CO", *invoiceObject.SalesRequestCode) && *invoiceObject.OutstandingAmount == "0.00" {
			customer, _ := WebCustomerUC{ContractUC: uc.ContractUC}.FindByCodes(c, models.WebCustomerParameter{Code: *invoiceObject.CustomerCode})
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
						if pointMonthly > maxMonthly {
							continue
						}
						minOrder, _ := strconv.ParseFloat(rules.MinOrder, 64)
						netOmount, _ := strconv.ParseFloat(*invoiceObject.NetAmount, 64)
						if netOmount < minOrder {
							continue
						}

						pointConversion, _ := strconv.ParseFloat(rules.PointConversion, 64)
						getPoint := netOmount / pointConversion

						pointUC.Add(c, requests.PointRequest{
							CustomerCodes: []requests.PointCustomerCode{
								{CustomerCode: customer[0].Code},
							},
							InvoiceDocumentNo: *invoiceObject.DocumentNo,
							Point:             strconv.FormatFloat(getPoint, 'f', 2, 64),
							PointType:         "2",
						})
					}
				}
			}
		}
		resBuilder = append(resBuilder, invoiceObject)
	}

	return resBuilder, nil
}
