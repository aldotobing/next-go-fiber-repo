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

// ItemPriceSyncUC ...
type ItemPriceSyncUC struct {
	*ContractUC
}

// BuildBody ...
func (uc ItemPriceSyncUC) BuildBody(res *models.ItemPriceSync) {
}

// FindByID ...
func (uc ItemPriceSyncUC) FindByID(c context.Context, parameter models.ItemPriceSyncParameter) (res models.ItemPriceSync, err error) {
	repo := repository.NewItemPriceSyncRepository(uc.DB)
	res, err = repo.FindByID(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}
	uc.BuildBody(&res)

	return res, err
}

// FindByID ...
func (uc ItemPriceSyncUC) FindItemPrice(c context.Context, parameter models.ItemPriceSyncParameter) (res models.ItemPriceSync, err error) {
	repo := repository.NewItemPriceSyncRepository(uc.DB)
	res, err = repo.FindItemPrice(c, parameter)
	if err != nil {
		// logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		// return res, err
	}
	uc.BuildBody(&res)

	return res, err
}

// Add ...
func (uc ItemPriceSyncUC) Add(c context.Context, data *requests.ItemPriceSyncRequest) (res models.ItemPriceSync, err error) {

	repo := repository.NewItemPriceSyncRepository(uc.DB)
	// now := time.Now().UTC()
	// strnow := now.Format(time.RFC3339)
	res = models.ItemPriceSync{
		ItemCode: &data.ItemCode,
	}
	res.ID, err = repo.Add(c, &res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return res, err
}

// Edit ,...
func (uc ItemPriceSyncUC) Edit(c context.Context, id string, data *requests.ItemPriceSyncRequest) (res models.ItemPriceSync, err error) {
	repo := repository.NewItemPriceSyncRepository(uc.DB)
	// now := time.Now().UTC()
	// strnow := now.Format(time.RFC3339)
	res = models.ItemPriceSync{
		ID: &id,
		// ProvinceID: &data.ProvinceID,
		ItemCode: &data.ItemCode,
		// Longitude:  &data.Lat,
		// Latitude:   &data.Long,
		// CreatedAt:  &strnow,
		// UpdatedAt:  &strnow,
		// CreatedBy:  &data.CreatedBy,
		// UpdatedBy:  &data.UpdatedBy,
	}

	res.ID, err = repo.Edit(c, &res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	return res, err
}

// SelectAll ...
func (uc ItemPriceSyncUC) DataSync(c context.Context, parameter models.ItemPriceSyncParameter) (res []models.ItemPriceSync, err error) {
	repo := repository.NewItemPriceSyncRepository(uc.DB)
	fmt.Println("get plv")
	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc).Add(
		time.Minute * time.Duration(-15))
	strnow := now.Format(time.RFC3339)
	parameter.DateParam = strnow
	jsonReq, err := json.Marshal(parameter)
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://nextbasis.id:8080/mysmagonsrv/rest/masteritemprice/get", bytes.NewBuffer(jsonReq))
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

	var resBuilder []models.ItemPriceSync
	for _, itemObject := range res {

		currentItem, _ := uc.FindItemPrice(c, models.ItemPriceSyncParameter{
			ItemCode:             *itemObject.ItemCode,
			PriceListVersionCode: *itemObject.PriceListVersionCode,
			UomCode:              *itemObject.UomCode,
			PriceListCode:        *itemObject.PriceListCode,
		})

		if currentItem.ID != nil {
			fmt.Println("present")
			itemObject.ID = currentItem.ID
			_, errupdate := repo.Edit(c, &itemObject)
			if errupdate != nil {
				fmt.Print(errupdate)
			}
		} else {
			fmt.Println("not present")
			_, errinsert := repo.Add(c, &itemObject)
			if errinsert != nil {
				fmt.Print(errinsert)
			}
		}

		// _, errinsert := repo.InsertDataWithLine(c, &invoiceObject)
		// if errinsert != nil {
		// 	fmt.Print("")
		// }

		resBuilder = append(resBuilder, itemObject)
	}

	return res, err
}
