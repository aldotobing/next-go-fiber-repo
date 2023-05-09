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

// CustomerTargetSemesterHandler ...
type CustomerTargetSemesterHandler struct {
	Handler
}

// SelectAll ...
func (h *CustomerTargetSemesterHandler) SelectAll(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.CustomerTargetSemesterParameter{
		ID:     ctx.Query("customer_id"),
		Code:   ctx.Query("customer_code"),
		Search: ctx.Query("search"),
		By:     ctx.Query("by"),
		Sort:   ctx.Query("sort"),
	}
	uc := usecase.CustomerTargetSemesterUC{ContractUC: h.ContractUC}
	res, err := uc.SelectAll(c, parameter)

	type StructObject struct {
		ListObjcet []models.CustomerTargetSemester `json:"list_customer"`
	}

	ObjcetData := new(StructObject)

	if res != nil {
		target := h.FetchClientDataTarget(parameter)
		if target != "" {
			for i := range res {
				res[i].CustomerTargetSemester = &target
			}
		}
		ObjcetData.ListObjcet = res
	}

	return h.SendResponse(ctx, ObjcetData, nil, err, 0)
}

// FindAll ...
func (h *CustomerTargetSemesterHandler) FindAll(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.CustomerTargetSemesterParameter{
		ID:     ctx.Query("customer_id"),
		Search: ctx.Query("search"),
		Page:   str.StringToInt(ctx.Query("page")),
		Limit:  str.StringToInt(ctx.Query("limit")),
		By:     ctx.Query("by"),
		Sort:   ctx.Query("sort"),
	}
	uc := usecase.CustomerTargetSemesterUC{ContractUC: h.ContractUC}
	res, meta, err := uc.FindAll(c, parameter)

	type StructObject struct {
		ListObject []models.CustomerTargetSemester `json:"list_customer"`
	}

	ObjectData := new(StructObject)

	if res != nil {
		ObjectData.ListObject = res
	}

	return h.SendResponse(ctx, ObjectData, meta, err, 0)
}

func (h *CustomerTargetSemesterHandler) FetchClientDataTarget(params models.CustomerTargetSemesterParameter) string {
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
		SemesterTarget string `json:"semester_target"`
		CurrentTarget  string `json:"current_target"`
		AnualTarget    string `json:"anual_target"`
	}

	ObjectData := new(resutlData)

	// var responseObject http.Response
	json.Unmarshal(bodyBytes, &ObjectData)

	return ObjectData.SemesterTarget
}

// FindByID ...
// func (h *CustomerTargetSemesterHandler) FindByID(ctx *fiber.Ctx) error {
// 	c := ctx.Locals("ctx").(context.Context)

// 	parameter := models.CustomerTargetSemesterParameter{
// 		ID: ctx.Params("customer_id"),
// 	}
// 	if parameter.ID == "" {
// 		return h.SendResponse(ctx, nil, nil, helper.InvalidParameter, http.StatusBadRequest)
// 	}

// 	uc := usecase.CustomerTargetSemesterUC{ContractUC: h.ContractUC}
// 	res, err := uc.FindByID(c, parameter)

// 	type StructObject struct {
// 		ListObject models.CustomerTargetSemester `json:"customer"`
// 	}

// 	ObjectData := new(StructObject)

// 	// if res != nil {
// 	ObjectData.ListObject = res
// 	// }

// 	return h.SendResponse(ctx, ObjectData, nil, err, 0)
// }
