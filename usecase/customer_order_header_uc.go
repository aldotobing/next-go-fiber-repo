package usecase

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"nextbasis-service-v-0.1/db/repository"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/helper"
	"nextbasis-service-v-0.1/pkg/functioncaller"
	"nextbasis-service-v-0.1/pkg/logruslogger"
	"nextbasis-service-v-0.1/pkg/number"
	pkgtime "nextbasis-service-v-0.1/pkg/time"
	"nextbasis-service-v-0.1/server/requests"
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

// CustomerOrderHeaderUC ...
type CustomerOrderHeaderUC struct {
	*ContractUC
}

// BuildBody ...
func (uc CustomerOrderHeaderUC) BuildBody(res *models.CustomerOrderHeader) {
}

// SelectAll ...
func (uc CustomerOrderHeaderUC) SelectAll(c context.Context, parameter models.CustomerOrderHeaderParameter) (res []models.CustomerOrderHeader, err error) {
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(0, 0, parameter.By, parameter.Sort, models.CustomerOrderHeaderOrderBy, models.CustomerOrderHeaderOrderByrByString)

	repo := repository.NewCustomerOrderHeaderRepository(uc.DB)
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
func (uc CustomerOrderHeaderUC) FindAll(c context.Context, parameter models.CustomerOrderHeaderParameter) (res []models.CustomerOrderHeader, p viewmodel.PaginationVM, err error) {
	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.CustomerOrderHeaderOrderBy, models.CustomerOrderHeaderOrderByrByString)

	var count int
	repo := repository.NewCustomerOrderHeaderRepository(uc.DB)
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
func (uc CustomerOrderHeaderUC) FindByID(c context.Context, parameter models.CustomerOrderHeaderParameter) (res models.CustomerOrderHeader, err error) {
	repo := repository.NewCustomerOrderHeaderRepository(uc.DB)
	res, err = repo.FindByID(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}
	uc.BuildBody(&res)

	return res, err
}

// Add ...
func (uc CustomerOrderHeaderUC) CheckOut(c context.Context, data *requests.CustomerOrderHeaderRequest) (res models.CustomerOrderHeader, err error) {

	repo := repository.NewCustomerOrderHeaderRepository(uc.DB)

	chekablerepo := repository.NewShoppingCartRepository(uc.DB)

	checkAble, errcheck := chekablerepo.GetTotal(c, models.ShoppingCartParameter{
		CustomerID: data.CustomerID,
		ListLine:   data.LineList,
	})

	// now := time.Now().UTC()
	// strnow := now.Format(time.RFC3339)
	round_amount := "0"
	gross_amount := "0"
	taxable_amount := "0"
	tax_amount := "0"
	net_amount := "0"
	disc_amount := "0"
	res = models.CustomerOrderHeader{
		TransactionDate:      &data.TransactionDate,
		TransactionTime:      &data.TransactionTime,
		CustomerID:           &data.CustomerID,
		PaymentTermsID:       &data.PaymentTermsID,
		ExpectedDeliveryDate: &data.ExpectedDeliveryDate,
		GrossAmount:          &gross_amount,
		DiscAmount:           &disc_amount,
		TaxableAmount:        &taxable_amount,
		TaxAmount:            &tax_amount,
		RoundingAmount:       &round_amount,
		NetAmount:            &net_amount,
		TaxCalcMethod:        &data.TaxCalcMethod,
		SalesmanID:           &data.SalesmanID,
		BranchID:             &data.BranchID,
		PriceLIstID:          &data.PriceLIstID,
		LineList:             &data.LineList,
	}

	if checkAble.IsAble == nil || *checkAble.IsAble == "0" {
		bayar, _ := strconv.ParseFloat(*checkAble.MinOmzet, 0)
		minOrder := strings.ReplaceAll(number.FormatCurrency(bayar, "IDR", ".", "", 0), "Rp", "")
		return res, errors.New(helper.InvalidMinimumAmountOrder + minOrder + ` rupiah.`)
	}

	if errcheck != nil {
		return res, errors.New("Try Again Latter")
	}

	res.ID, err = repo.CheckOut(c, &res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	userrepo := repository.NewCustomerRepository(uc.DB)

	useraccount, erruser := userrepo.FindByID(c, models.CustomerParameter{ID: *res.CustomerID})

	if erruser == nil {

		FcmUc := FCMUC{ContractUC: uc.ContractUC}
		orderrepo := repository.NewCustomerOrderHeaderRepository(uc.DB)
		orderlinerepo := repository.NewCustomerOrderLineRepository(uc.DB)
		order, errorder := orderrepo.FindByID(c, models.CustomerOrderHeaderParameter{ID: *res.ID})
		if errorder == nil {
			// fmt.Println("tanggal ", *order.TransactionDate)
			dateString := pkgtime.GetDate(*order.TransactionDate+"T00:00:00Z", "02 - 01 - 2006", "Asia/Jakarta")

			messageType := "1"
			bayar, _ := strconv.ParseFloat(*order.NetAmount, 0)
			harga := strings.ReplaceAll(number.FormatCurrency(bayar, "IDR", ".", "", 0), "Rp", "")
			msgtitle := "Checkout " + *order.DocumentNo
			msgbody := `*Kepada Yang Terhormat* \n\n *` + *useraccount.Code + ` - ` + *useraccount.CustomerName + `*`
			msgbody += `\n\n*NO ORDERAN ` + *order.DocumentNo + ` anda pada tanggal ` + dateString + ` oleh Toko : (` + *useraccount.CustomerName + `) telah berhasil dan akan diproses*`
			msgbody += `\n\n*Berikut merupakan rincian pesanan anda:*`
			orderline, errline := orderlinerepo.SelectAll(c, models.CustomerOrderLineParameter{
				HeaderID: *order.ID,
				By:       "def.created_date",
			})
			if errline == nil {
				msgbody += `\n`
				for i := range orderline {
					msgbody += `\n ` + *orderline[i].QTY + ` ` + *orderline[i].UomName + ` ` + *orderline[i].ItemName + `\n`

				}
				ordercount := len(orderline)
				msgbody += `\n`
				msgbody += `Total ` + strconv.Itoa(ordercount) + ` item, senilai ` + harga + ` (belum termasuk potongan/diskon bila ada program potongan/diskon) `
				msgbody += `\n`
				msgbody += `\nTerima kasih atas pemesanan anda`
				msgbody += `\n`
				msgbody += `\nSalam Sehat`
				msgbody += `\n`
				msgbody += `\nNB : Bila ini bukan transaksi dari Toko Bapak/Ibu, silahkan menghubungi Distributor Produk Sido Muncul.`
			}

			if useraccount.CustomerFCMToken != nil && *useraccount.CustomerFCMToken != "" {

				_, errfcm := FcmUc.SendFCMMessage(c, msgtitle, msgbody, *useraccount.CustomerFCMToken)
				if errfcm == nil {

				}

				userNotificationRepo := repository.NewUserNotificationRepository(uc.DB)
				_, errnotifinsert := userNotificationRepo.Add(c, &models.UserNotification{
					Title:  &msgtitle,
					Text:   &msgbody,
					Type:   &messageType,
					UserID: order.CustomerID,
					RowID:  order.ID,
				})
				if errnotifinsert == nil {

				}

			}
			if useraccount.CustomerPhone != nil && *useraccount.CustomerPhone != "" {

				senDwaMessage := uc.ContractUC.WhatsApp.SendTransactionWA(*useraccount.CustomerPhone, msgbody)
				if senDwaMessage != nil {
					fmt.Println("sukses")
				}
				if useraccount.CustomerSalesmanID != nil {
					salesmannRepo := repository.NewSalesmanRepository(uc.DB)
					customerSales, errcustsales := salesmannRepo.FindByID(c, models.SalesmanParameter{ID: *useraccount.CustomerSalesmanID})
					if errcustsales == nil {
						if customerSales.PhoneNo != nil {
							senDwaMessage := uc.ContractUC.WhatsApp.SendTransactionWA(*customerSales.PhoneNo, msgbody)
							if senDwaMessage != nil {
								fmt.Println("sukses")
							}
						}
					}
				}
			}
		}

	}

	return res, err
}

func (uc CustomerOrderHeaderUC) VoidedDataSync(c context.Context, parameter models.CustomerOrderHeaderParameter) (res []models.CustomerOrderHeader, err error) {
	repo := repository.NewCustomerOrderHeaderRepository(uc.DB)

	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc).Add(time.Minute * time.Duration(-15))

	strnow := now.Format(time.RFC3339)
	fmt.Println(strnow)
	parameter.DateParam = strnow
	jsonReq, err := json.Marshal(parameter)
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://nextbasis.id:8080/mysmagonsrv/rest/salesRequest/voideddata/2", bytes.NewBuffer(jsonReq))
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

	var resBuilder []models.CustomerOrderHeader
	for _, invoiceObject := range res {

		orderrepo := repository.NewCustomerOrderHeaderRepository(uc.DB)

		currentOrder, errcurrent := orderrepo.FindByCode(c, models.CustomerOrderHeaderParameter{DocumentNo: *invoiceObject.DocumentNo})

		_, errinsert := repo.SyncVoid(c, &invoiceObject)

		if errinsert != nil {
			fmt.Print(errinsert)
		}
		if errcurrent == nil {

			if currentOrder.Status != nil && *currentOrder.Status != *invoiceObject.Status {
				userrepo := repository.NewCustomerRepository(uc.DB)

				useraccount, erruser := userrepo.FindByID(c, models.CustomerParameter{ID: *currentOrder.CustomerID})
				if erruser == nil {
					orderlinerepo := repository.NewCustomerOrderLineRepository(uc.DB)
					orderline, errline := orderlinerepo.SelectAll(c, models.CustomerOrderLineParameter{
						HeaderID: *currentOrder.ID,
						By:       "def.created_date",
					})

					if errline == nil {
						messageTemplate := ""
						messageTitle := ""
						messageType := "1"
						if *invoiceObject.Status == "voided" {
							messageTemplate = helper.BuildVoidTransactionTemplate(currentOrder, orderline, useraccount)
							messageTitle = "Transaksi " + *currentOrder.DocumentNo + " dibatalkan."
						} else if *invoiceObject.Status == "submitted" {
							messageTemplate = helper.BuildProcessTransactionTemplate(currentOrder, orderline, useraccount)
							messageTitle = "Transaksi " + *currentOrder.DocumentNo + " diproses."
						}

						if useraccount.CustomerFCMToken != nil && *useraccount.CustomerFCMToken != "" {
							FcmUc := FCMUC{ContractUC: uc.ContractUC}
							_, errfcm := FcmUc.SendFCMMessage(c, messageTitle, messageTemplate, *useraccount.CustomerFCMToken)
							if errfcm == nil {

							}

							userNotificationRepo := repository.NewUserNotificationRepository(uc.DB)
							_, errnotifinsert := userNotificationRepo.Add(c, &models.UserNotification{
								Title:  &messageTitle,
								Text:   &messageTemplate,
								Type:   &messageType,
								UserID: currentOrder.CustomerID,
								RowID:  currentOrder.ID,
							})
							if errnotifinsert == nil {

							}

						}

						if useraccount.CustomerPhone != nil && *useraccount.CustomerPhone != "" {
							if messageTemplate != "" {
								senDwaMessage := uc.ContractUC.WhatsApp.SendTransactionWA(*useraccount.CustomerPhone, messageTemplate)
								if senDwaMessage != nil {
									fmt.Println("sukses")
								}

							}

						}
					}

				}

			}
		}

		resBuilder = append(resBuilder, invoiceObject)

	}

	return resBuilder, err
}

// apps

// SelectAll ...
func (uc CustomerOrderHeaderUC) AppsSelectAll(c context.Context, parameter models.CustomerOrderHeaderParameter) (res []models.CustomerOrderHeader, err error) {
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(0, 0, parameter.By, parameter.Sort, models.CustomerOrderHeaderOrderBy, models.CustomerOrderHeaderOrderByrByString)

	repo := repository.NewCustomerOrderHeaderRepository(uc.DB)
	res, err = repo.AppsSelectAll(c, parameter)

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
func (uc CustomerOrderHeaderUC) AppsFindAll(c context.Context, parameter models.CustomerOrderHeaderParameter) (res []models.CustomerOrderHeader, p viewmodel.PaginationVM, err error) {
	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.CustomerOrderHeaderOrderBy, models.CustomerOrderHeaderOrderByrByString)

	var count int
	repo := repository.NewCustomerOrderHeaderRepository(uc.DB)
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
func (uc CustomerOrderHeaderUC) AppsFindByID(c context.Context, parameter models.CustomerOrderHeaderParameter) (res models.CustomerOrderHeader, err error) {
	repo := repository.NewCustomerOrderHeaderRepository(uc.DB)
	res, err = repo.FindByID(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}
	uc.BuildBody(&res)

	return res, err
}
