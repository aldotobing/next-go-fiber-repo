package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/helper"
	"nextbasis-service-v-0.1/pkg/str"
	"nextbasis-service-v-0.1/server/requests"
	"nextbasis-service-v-0.1/usecase"
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

// WebCustomerHandler ...
type WebCustomerHandler struct {
	Handler
}

// SelectAll ...
func (h *WebCustomerHandler) SelectAll(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.WebCustomerParameter{
		ID:             ctx.Query("customer_id"),
		CustomerTypeId: ctx.Query("customer_type_id"),
		UserId:         ctx.Query("admin_user_id"),
		Search:         ctx.Query("search"),
		By:             ctx.Query("by"),
		Sort:           ctx.Query("sort"),
	}
	uc := usecase.WebCustomerUC{ContractUC: h.ContractUC}
	res, err := uc.SelectAll(c, parameter)

	type StructObject struct {
		ListObject []viewmodel.CustomerVM `json:"list_customer"`
	}

	objectData := new(StructObject)

	if res != nil {
		objectData.ListObject = res
	}

	return h.SendResponse(ctx, objectData, nil, err, 0)
}

// FindAll ...
func (h *WebCustomerHandler) FindAll(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)
	parameter := models.WebCustomerParameter{
		ID:             ctx.Query("customer_id"),
		CustomerTypeId: ctx.Query("customer_type_id"),
		UserId:         ctx.Query("admin_user_id"),
		Search:         ctx.Query("search"),
		Page:           str.StringToInt(ctx.Query("page")),
		Limit:          str.StringToInt(ctx.Query("limit")),
		By:             ctx.Query("by"),
		Sort:           ctx.Query("sort"),
	}
	uc := usecase.WebCustomerUC{ContractUC: h.ContractUC}
	res, meta, err := uc.FindAll(c, parameter)

	type StructObject struct {
		ListObject []viewmodel.CustomerVM `json:"list_customer"`
	}

	objectData := new(StructObject)

	if res != nil {
		objectData.ListObject = res
	}

	return h.SendResponse(ctx, objectData, meta, err, 0)
}

// FindByID ...
func (h *WebCustomerHandler) FindByID(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.WebCustomerParameter{
		ID: ctx.Params("customer_id"),
	}
	if parameter.ID == "" {
		return h.SendResponse(ctx, nil, nil, helper.InvalidParameter, http.StatusBadRequest)
	}

	uc := usecase.WebCustomerUC{ContractUC: h.ContractUC}
	res, err := uc.FindByID(c, parameter)

	type StructObject struct {
		ListObject models.WebCustomer `json:"customer"`
	}

	objectData := new(StructObject)

	objectData.ListObject = res

	target := h.FetchVisitDay(parameter)
	if target != "" {
		objectData.ListObject.VisitDay = &target
	}

	return h.SendResponse(ctx, objectData, nil, err, 0)
}

func (h *WebCustomerHandler) FetchVisitDay(params models.WebCustomerParameter) string {
	jsonReq, err := json.Marshal(params)
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://nextbasis.id:8080/mysmagonsrv/rest/customer/visitday/1", bytes.NewBuffer(jsonReq))
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
		VisitDay string `json:"visit_day"`
	}

	objectData := new(resutlData)

	// var responseObject http.Response
	json.Unmarshal(bodyBytes, &objectData)

	return objectData.VisitDay
}

// Edit ...
func (h *WebCustomerHandler) Edit(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	id := ctx.Params("customer_id")
	if id == "" {
		return h.SendResponse(ctx, nil, nil, helper.InvalidParameter, http.StatusBadRequest)
	}

	input := new(requests.WebCustomerRequest)
	err := json.Unmarshal([]byte(ctx.FormValue("form_data")), &input)
	if err := ctx.BodyParser(input); err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}
	if err := h.Validator.Struct(input); err != nil {
		errMessage := h.ExtractErrorValidationMessages(err.(validator.ValidationErrors))
		return h.SendResponse(ctx, nil, nil, errMessage, http.StatusBadRequest)
	}

	if input.CustomerGender != "" && !str.Contains(models.CustomerGenderList, strings.ToLower(input.CustomerGender)) {
		return h.SendResponse(ctx, nil, nil, errors.New(helper.InvalidGender), http.StatusBadRequest)
	}

	imgProfile, _ := ctx.FormFile("img_profile")
	uc := usecase.WebCustomerUC{ContractUC: h.ContractUC}
	res, err := uc.Edit(c, id, input, imgProfile)

	return h.SendResponse(ctx, res, nil, err, 0)
}

// Edit ...
func (h *WebCustomerHandler) Add(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	input := new(requests.WebCustomerRequest)
	err := json.Unmarshal([]byte(ctx.FormValue("form_data")), input)
	if err := ctx.BodyParser(input); err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}
	if err := h.Validator.Struct(input); err != nil {
		errMessage := h.ExtractErrorValidationMessages(err.(validator.ValidationErrors))
		return h.SendResponse(ctx, nil, nil, errMessage, http.StatusBadRequest)
	}

	if input.CustomerGender != "" && !str.Contains(models.CustomerGenderList, strings.ToLower(input.CustomerGender)) {
		return h.SendResponse(ctx, nil, nil, errors.New(helper.InvalidGender), http.StatusBadRequest)
	}

	imgProfile, _ := ctx.FormFile("img_profile")
	uc := usecase.WebCustomerUC{ContractUC: h.ContractUC}
	res, err := uc.Add(c, input, imgProfile)

	return h.SendResponse(ctx, res, nil, err, 0)
}
