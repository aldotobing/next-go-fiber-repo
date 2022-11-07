package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/str"
	"nextbasis-service-v-0.1/usecase"
)

// CustomerTargetQuarterHandler ...
type CustomerTargetQuarterHandler struct {
	Handler
}

// SelectAll ...
func (h *CustomerTargetQuarterHandler) SelectAll(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.CustomerTargetQuarterParameter{
		ID:     ctx.Query("customer_id"),
		Search: ctx.Query("search"),
		By:     ctx.Query("by"),
		Sort:   ctx.Query("sort"),
	}
	uc := usecase.CustomerTargetQuarterUC{ContractUC: h.ContractUC}
	res, err := uc.SelectAll(c, parameter)

	type StructObject struct {
		ListObjcet []models.CustomerTargetQuarter `json:"list_customer"`
	}

	ObjcetData := new(StructObject)

	if res != nil {
		target := h.FetchClientDataTarget(parameter)
		if target != "" {
			for i, _ := range res {
				res[i].CustomerTargetQuarter = &target
			}
		}
		ObjcetData.ListObjcet = res
	}

	return h.SendResponse(ctx, ObjcetData, nil, err, 0)
}

// FindAll ...
func (h *CustomerTargetQuarterHandler) FindAll(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.CustomerTargetQuarterParameter{
		ID:     ctx.Query("customer_id"),
		Search: ctx.Query("search"),
		Page:   str.StringToInt(ctx.Query("page")),
		Limit:  str.StringToInt(ctx.Query("limit")),
		By:     ctx.Query("by"),
		Sort:   ctx.Query("sort"),
	}
	uc := usecase.CustomerTargetQuarterUC{ContractUC: h.ContractUC}
	res, meta, err := uc.FindAll(c, parameter)

	type StructObject struct {
		ListObject []models.CustomerTargetQuarter `json:"list_customer"`
	}

	ObjectData := new(StructObject)

	if res != nil {
		ObjectData.ListObject = res
	}

	return h.SendResponse(ctx, ObjectData, meta, err, 0)
}

func (h *CustomerTargetQuarterHandler) FetchClientDataTarget(params models.CustomerTargetQuarterParameter) string {
	jsonReq, err := json.Marshal(params)
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://nextbasis.id:8080/mysmagonsrv/rest/customertarget/data/1", bytes.NewBuffer(jsonReq))
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

	type resutlData struct {
		QuartalTarget string `json:"quartal_target"`
		CurrentTarget string `json:"current_target"`
		AnualTarget   string `json:"anual_target"`
	}

	ObjectData := new(resutlData)

	// var responseObject http.Response
	json.Unmarshal(bodyBytes, &ObjectData)
	fmt.Println("data ", ObjectData)

	return ObjectData.QuartalTarget
}

// FindByID ...
// func (h *CustomerTargetQuarterHandler) FindByID(ctx *fiber.Ctx) error {
// 	c := ctx.Locals("ctx").(context.Context)

// 	parameter := models.CustomerTargetQuarterParameter{
// 		ID: ctx.Params("customer_id"),
// 	}
// 	if parameter.ID == "" {
// 		return h.SendResponse(ctx, nil, nil, helper.InvalidParameter, http.StatusBadRequest)
// 	}

// 	uc := usecase.CustomerTargetQuarterUC{ContractUC: h.ContractUC}
// 	res, err := uc.FindByID(c, parameter)

// 	type StructObject struct {
// 		ListObject models.CustomerTargetQuarter `json:"customer"`
// 	}

// 	ObjectData := new(StructObject)

// 	// if res != nil {
// 	ObjectData.ListObject = res
// 	// }

// 	return h.SendResponse(ctx, ObjectData, nil, err, 0)
// }
