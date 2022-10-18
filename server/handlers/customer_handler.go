package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/helper"
	"nextbasis-service-v-0.1/pkg/str"
	"nextbasis-service-v-0.1/server/requests"
	"nextbasis-service-v-0.1/usecase"
)

// CustomerHandler ...
type CustomerHandler struct {
	Handler
}

// SelectAll ...
func (h *CustomerHandler) SelectAll(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.CustomerParameter{
		ID:             ctx.Query("customer_id"),
		CustomerTypeId: ctx.Query("customer_type_id"),
		UserId:         ctx.Query("admin_user_id"),
		Search:         ctx.Query("search"),
		By:             ctx.Query("by"),
		Sort:           ctx.Query("sort"),
	}
	uc := usecase.CustomerUC{ContractUC: h.ContractUC}
	res, err := uc.SelectAll(c, parameter)

	type StructObject struct {
		ListObjcet []models.Customer `json:"list_customer"`
	}

	ObjcetData := new(StructObject)

	if res != nil {
		ObjcetData.ListObjcet = res
	}

	return h.SendResponse(ctx, ObjcetData, nil, err, 0)
}

// FindAll ...
func (h *CustomerHandler) FindAll(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.CustomerParameter{
		ID:             ctx.Query("customer_id"),
		CustomerTypeId: ctx.Query("customer_type_id"),
		UserId:         ctx.Query("admin_user_id"),
		Search:         ctx.Query("search"),
		Page:           str.StringToInt(ctx.Query("page")),
		Limit:          str.StringToInt(ctx.Query("limit")),
		By:             ctx.Query("by"),
		Sort:           ctx.Query("sort"),
	}
	uc := usecase.CustomerUC{ContractUC: h.ContractUC}
	res, meta, err := uc.FindAll(c, parameter)

	type StructObject struct {
		ListObject []models.Customer `json:"list_customer"`
	}

	ObjectData := new(StructObject)

	if res != nil {
		ObjectData.ListObject = res
	}

	return h.SendResponse(ctx, ObjectData, meta, err, 0)
}

// FindByID ...
func (h *CustomerHandler) FindByID(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.CustomerParameter{
		ID: ctx.Params("customer_id"),
	}
	if parameter.ID == "" {
		return h.SendResponse(ctx, nil, nil, helper.InvalidParameter, http.StatusBadRequest)
	}

	uc := usecase.CustomerUC{ContractUC: h.ContractUC}
	res, err := uc.FindByID(c, parameter)

	type StructObject struct {
		ListObject models.Customer `json:"customer"`
	}

	ObjectData := new(StructObject)

	// if res != nil {
	ObjectData.ListObject = res
	// }

	return h.SendResponse(ctx, ObjectData, nil, err, 0)
}

// Edit ...
func (h *CustomerHandler) Edit(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	id := ctx.Params("customer_id")
	if id == "" {
		return h.SendResponse(ctx, nil, nil, helper.InvalidParameter, http.StatusBadRequest)
	}

	input := new(requests.CustomerRequest)
	if err := ctx.BodyParser(input); err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}
	if err := h.Validator.Struct(input); err != nil {
		errMessage := h.ExtractErrorValidationMessages(err.(validator.ValidationErrors))
		return h.SendResponse(ctx, nil, nil, errMessage, http.StatusBadRequest)
	}

	fmt.Println("aedit")

	uc := usecase.CustomerUC{ContractUC: h.ContractUC}
	res, err := uc.Edit(c, id, input)

	return h.SendResponse(ctx, res, nil, err, 0)
}

func (h *CustomerHandler) EditAddress(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	id := ctx.Params("customer_id")
	if id == "" {
		return h.SendResponse(ctx, nil, nil, helper.InvalidParameter, http.StatusBadRequest)
	}

	input := new(requests.CustomerRequest)
	if err := ctx.BodyParser(input); err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}
	if err := h.Validator.Struct(input); err != nil {
		errMessage := h.ExtractErrorValidationMessages(err.(validator.ValidationErrors))
		return h.SendResponse(ctx, nil, nil, errMessage, http.StatusBadRequest)
	}

	uc := usecase.CustomerUC{ContractUC: h.ContractUC}
	res, err := uc.EditAddress(c, id, input)

	return h.SendResponse(ctx, res, nil, err, 0)
}
