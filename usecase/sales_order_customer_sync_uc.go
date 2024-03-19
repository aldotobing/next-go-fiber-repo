package usecase

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"nextbasis-service-v-0.1/db/repository"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/helper"
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
	req, err := http.NewRequest("GET", "http://nextbasis.id:8080/mysmagonsrv/rest/salesOrder/data/online_store", bytes.NewBuffer(jsonReq))
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

		if errinsert == nil {
			if invoiceObject.Status != nil {
				userrepo := repository.NewCustomerRepository(uc.DB)
				salesorderHeaderrepo := repository.NewSalesOrderHeaderRepository(uc.DB)
				salesorderHeader, errheader := salesorderHeaderrepo.FindByCode(c, models.SalesOrderHeaderParameter{DocumentNo: *invoiceObject.DocumentNo})
				if errheader == nil {
					useraccount, erruser := userrepo.FindByID(c, models.CustomerParameter{ID: *salesorderHeader.CustomerID})

					if erruser == nil && useraccount.CustomerFCMToken != nil && *useraccount.CustomerFCMToken != "" {

						if errheader == nil {
							orderlinerepo := repository.NewSalesOrderLineRepository(uc.DB)
							orderline, errline := orderlinerepo.SelectAll(c, models.SalesOrderLineParameter{
								HeaderID: *salesorderHeader.ID,
								By:       "def.created_date",
							})

							if errline == nil {
								messageTemplate := ""
								messageTitle := ""
								messageType := "2"
								if *invoiceObject.Status == "submitted" {
									messageTemplate = helper.BuildProcessSalesOrderTransactionTemplate(salesorderHeader, orderline, useraccount, 1)
									messageTitle = "Transaksi " + *invoiceObject.DocumentNo + " diproses."
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
										UserID: invoiceObject.CustomerID,
										RowID:  invoiceObject.ID,
									})
									if errnotifinsert == nil {

									}

								}

								if useraccount.CustomerPhone != nil && *useraccount.CustomerPhone != "" {
									// if messageTemplate != "" {
									// senDwaMessage := uc.ContractUC.WhatsApp.SendTransactionWA(*useraccount.CustomerPhone, messageTemplate)
									// if senDwaMessage != nil {
									// 	fmt.Println("sukses")
									// }

									// }

									if useraccount.CustomerSalesmanID != nil {
										salesmanmessageTemplate := ""
										salesmannRepo := repository.NewSalesmanRepository(uc.DB)
										customerSales, errcustsales := salesmannRepo.FindByID(c, models.SalesmanParameter{ID: *useraccount.CustomerSalesmanID})

										salesmanmessageTemplate = helper.BuildProcessSalesOrderTransactionTemplate(salesorderHeader, orderline, useraccount, 2)

										if errcustsales == nil {
											if customerSales.PhoneNo != nil {
												if salesmanmessageTemplate != "" {

													senDwaMessage := uc.ContractUC.WhatsApp.SendTransactionWA(*customerSales.PhoneNo, salesmanmessageTemplate)
													if senDwaMessage != nil {
														fmt.Println("sukses")
													}
												}

											}
										}
									}

								}

								// if useraccount.CustomerBranchPicPhoneNo != nil && useraccount.CustomerBranchPicName != nil {
								// 	picMessageTemplate := helper.BuildProcessSalesOrderTransactionTemplate(salesorderHeader, orderline, useraccount, 3)
								// 	_ = uc.ContractUC.WhatsApp.SendTransactionWA(*useraccount.CustomerBranchPicPhoneNo, picMessageTemplate)
								// }
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

// SelectAll ...
func (uc SalesOrderCustomerSyncUC) PullDataSync(c context.Context, parameter models.SalesOrderCustomerSyncParameter) (res []models.SalesOrderCustomerSync, err error) {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc).Add(time.Minute * time.Duration(-15))
	strnow := now.Format(time.RFC3339)
	parameter.DateParam = strnow
	parameter.Status = "submitted"
	jsonReq, err := json.Marshal(parameter)
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://nextbasis.id:8080/mysmagonsrv/rest/salesOrder/data/online_store", bytes.NewBuffer(jsonReq))
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
		cacheKey := "submitted_so_data_:" + *invoiceObject.DocumentNo
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

		resBuilder = append(resBuilder, invoiceObject)

	}

	return resBuilder, err
}

func (uc SalesOrderCustomerSyncUC) PushDataSync(c context.Context, parameter models.SalesOrderCustomerSyncParameter) (res []models.SalesOrderCustomerSync, err error) {
	repo := repository.NewSalesOrderCustomerSyncRepository(uc.DB)

	cacheKey := "*submitted_so_data*"

	// Try to get data from Redis cache first
	strsoList, err := uc.RedisClient.GetAllKeyFromRedis(cacheKey)

	if err == nil {
		var minLen = 85
		var keyLen = len(strsoList)

		if keyLen < minLen {
			minLen = keyLen
		}
		if minLen > 0 {

			for i := 0; i < minLen; i++ {

				key := strsoList[i]

				soObject := new(models.SalesOrderCustomerSync)
				err = uc.RedisClient.GetFromRedis(key, &soObject)
				if err != nil {
					fmt.Println(err)
				}
				if err == nil {
					fmt.Println("from redis : ", key)
					_, modifyOnly, errinsert := repo.MergeData(c, soObject)

					if errinsert != nil {
						errstr := errinsert.Error()

						if strings.Contains(errstr, "cust_bill_to_id") || strings.Contains(errstr, "uom_id") || strings.Contains(errstr, "item_id") ||
							strings.Contains(errstr, "more than one row returned by a subquery used as an expression") ||
							strings.Contains(errstr, "item_stock_location_id_fkey") {
							cacheKeyerr := "err_so_data:" + *soObject.DocumentNo
							errjsonData, err := json.Marshal(soObject)
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
							return nil, fmt.Errorf("failed to insert data for order %+v: %w", soObject, errinsert)
						}
					}

					if errinsert == nil {

						if soObject.Status != nil && modifyOnly == 0 {
							userrepo := repository.NewCustomerRepository(uc.DB)
							salesorderHeaderrepo := repository.NewSalesOrderHeaderRepository(uc.DB)
							salesorderHeader, errheader := salesorderHeaderrepo.FindByCode(c, models.SalesOrderHeaderParameter{DocumentNo: *soObject.DocumentNo})
							if errheader == nil {
								useraccount, erruser := userrepo.FindByID(c, models.CustomerParameter{ID: *salesorderHeader.CustomerID})

								if erruser == nil && useraccount.CustomerFCMToken != nil && *useraccount.CustomerFCMToken != "" {

									if errheader == nil {
										orderlinerepo := repository.NewSalesOrderLineRepository(uc.DB)
										orderline, errline := orderlinerepo.SelectAll(c, models.SalesOrderLineParameter{
											HeaderID: *salesorderHeader.ID,
											By:       "def.created_date",
										})

										if errline == nil {
											messageTemplate := ""
											messageTitle := ""
											messageType := "2"
											if *soObject.Status == "submitted" {
												messageTemplate = helper.BuildProcessSalesOrderTransactionTemplate(salesorderHeader, orderline, useraccount, 1)
												messageTitle = "Transaksi " + *soObject.DocumentNo + " diproses."
											}

											if useraccount.CustomerFCMToken != nil && *useraccount.CustomerFCMToken != "" {
												userfcmObject := new(models.FcmSo)
												userfcmObject.FcmToken = useraccount.CustomerFCMToken
												userfcmObject.Template = &messageTemplate
												userfcmObject.Title = &messageTitle
												userFcmCacheKey := "fcm_so:" + *soObject.DocumentNo
												jsonData, err := json.Marshal(userfcmObject)
												if err != nil {
													logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "json_marshal", uc.ReqID)
													return res, err
												}
												err = uc.RedisClient.Client.Set(userFcmCacheKey, jsonData, time.Hour*168).Err()
												if err != nil {
													logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "redis_set", uc.ReqID)
													return res, err
												}
												// FcmUc := FCMUC{ContractUC: uc.ContractUC}
												// _, errfcm := FcmUc.SendFCMMessage(c, messageTitle, messageTemplate, *useraccount.CustomerFCMToken)
												// if errfcm == nil {

												// }

												userNotificationRepo := repository.NewUserNotificationRepository(uc.DB)
												_, errnotifinsert := userNotificationRepo.Add(c, &models.UserNotification{
													Title:  &messageTitle,
													Text:   &messageTemplate,
													Type:   &messageType,
													UserID: soObject.CustomerID,
													RowID:  soObject.ID,
												})
												if errnotifinsert == nil {

												}

											}

											if useraccount.CustomerPhone != nil && *useraccount.CustomerPhone != "" {
												// if messageTemplate != "" {
												// senDwaMessage := uc.ContractUC.WhatsApp.SendTransactionWA(*useraccount.CustomerPhone, messageTemplate)
												// if senDwaMessage != nil {
												// 	fmt.Println("sukses")
												// }

												// }

												if useraccount.CustomerSalesmanID != nil {
													salesmanmessageTemplate := ""
													salesmannRepo := repository.NewSalesmanRepository(uc.DB)
													customerSales, errcustsales := salesmannRepo.FindByID(c, models.SalesmanParameter{ID: *useraccount.CustomerSalesmanID})

													salesmanmessageTemplate = helper.BuildProcessSalesOrderTransactionTemplate(salesorderHeader, orderline, useraccount, 2)

													if errcustsales == nil {
														if customerSales.PhoneNo != nil {
															if salesmanmessageTemplate != "" {
																userWaObject := new(models.WaSo)
																userWaObject.Phone = customerSales.PhoneNo
																userWaObject.Template = &messageTemplate

																salesWaCacheKey := "wa_so:" + *soObject.DocumentNo + *customerSales.PhoneNo
																jsonData, err := json.Marshal(userWaObject)
																if err != nil {
																	logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "json_marshal", uc.ReqID)
																	return res, err
																}
																err = uc.RedisClient.Client.Set(salesWaCacheKey, jsonData, time.Hour*168).Err()
																if err != nil {
																	logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "redis_set", uc.ReqID)
																	return res, err
																}
																// senDwaMessage := uc.ContractUC.WhatsApp.SendTransactionWA(*customerSales.PhoneNo, salesmanmessageTemplate)
																// if senDwaMessage != nil {
																// 	fmt.Println("sukses")
																// }
															}

														}
													}
												}

											}

										}
									}

								}
							}
						}

					}
					res = append(res, *soObject)
					_ = uc.RedisClient.Delete(key)
				}

			}
		}
	}

	return res, err
}

func (uc SalesOrderCustomerSyncUC) SendSubmittedSOFCMNotification(c context.Context, parameter models.SalesOrderCustomerSyncParameter) (res []models.FcmSo, err error) {

	cacheKey := "*fcm_so*"

	// Try to get data from Redis cache first
	strsoList, err := uc.RedisClient.GetAllKeyFromRedis(cacheKey)

	if err == nil {
		var minLen = 50
		var keyLen = len(strsoList)

		if keyLen < minLen {
			minLen = keyLen
		}
		if minLen > 0 {
			for i := 0; i < minLen; i++ {

				key := strsoList[i]
				soObject := new(models.FcmSo)
				err = uc.RedisClient.GetFromRedis(key, &soObject)
				if err != nil {
					fmt.Println(err)
				}
				if err == nil {
					fmt.Println("from redis : ", key)
					// tkn := "dsyFCZqrRVq5PXzLvP-rba:APA91bG3aapCuJGy1Tn4FLxS3TQdKzSw_IJwgo_MDdIB4g00Y68xoOZ8moGMm4tNVKlWJFbMHfKfrHVlXHUgA5fvdYem2wpib77_rQx3DF57QpIoUZc59xCOfQWiDfrE3fpDBx15KOJJ"
					// soObject.FcmToken = &tkn
					// jsonData, err := json.Marshal(soObject)
					// if err != nil {
					// 	logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "json_marshal", uc.ReqID)
					// 	return res, err
					// }
					// err = uc.RedisClient.Client.Set(key, jsonData, time.Hour*168).Err()
					// if err != nil {
					// 	logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "redis_set", uc.ReqID)
					// 	return res, err
					// }
					FcmUc := FCMUC{ContractUC: uc.ContractUC}
					_, errfcm := FcmUc.SendFCMMessage(c, *soObject.Title, *soObject.Template, *soObject.FcmToken)
					if errfcm == nil {
						res = append(res, *soObject)
						_ = uc.RedisClient.Delete(key)
					}

				}
			}
		}
	}

	return res, err
}

func (uc SalesOrderCustomerSyncUC) SendSubmittedSOSalesmanWa(c context.Context, parameter models.SalesOrderCustomerSyncParameter) (res []models.WaSo, err error) {

	cacheKey := "*wa_so*"

	// Try to get data from Redis cache first
	strsoList, err := uc.RedisClient.GetAllKeyFromRedis(cacheKey)

	if err == nil {
		var minLen = 25
		var keyLen = len(strsoList)

		if keyLen < minLen {
			minLen = keyLen
		}
		if minLen > 0 {
			for i := 0; i < minLen; i++ {

				key := strsoList[i]
				soObject := new(models.WaSo)
				err = uc.RedisClient.GetFromRedis(key, &soObject)
				if err != nil {
					fmt.Println(err)
				}
				if err == nil {
					fmt.Println("from redis : ", key)
					senDwaMessage := uc.ContractUC.WhatsApp.SendTransactionWA(*soObject.Phone, *soObject.Template)
					if senDwaMessage == nil {
						res = append(res, *soObject)
						_ = uc.RedisClient.Delete(key)
					}

				}
			}
		}
	}

	return res, err
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
	req, err := http.NewRequest("GET", "http://nextbasis.id:8080/mysmagonsrv/rest/salesOrder/data/online_store", bytes.NewBuffer(jsonReq))
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
