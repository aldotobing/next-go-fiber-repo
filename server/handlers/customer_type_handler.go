package handlers

import (
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/helper"
	"nextbasis-service-v-0.1/pkg/str"
	"nextbasis-service-v-0.1/usecase"
)

// CustomerTypeHandler ...
type CustomerTypeHandler struct {
	Handler
}

// SelectAll ...
func (h *CustomerTypeHandler) SelectAll(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.CustomerTypeParameter{
		UserID:   ctx.Query("user_id"),
		Search:   ctx.Query("search"),
		By:       ctx.Query("by"),
		Sort:     ctx.Query("sort"),
		ExceptId: ctx.Query("except_id"),
	}
	uc := usecase.CustomerTypeUC{ContractUC: h.ContractUC}
	res, err := uc.SelectAll(c, parameter)

	type StructObject struct {
		ListObjcet []models.CustomerType `json:"list_customertype"`
	}

	ObjcetData := new(StructObject)

	if res != nil {
		ObjcetData.ListObjcet = res
	}

	return h.SendResponse(ctx, ObjcetData, nil, err, 0)
}

// FindAll ...
func (h *CustomerTypeHandler) FindAll(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.CustomerTypeParameter{
		UserID:   ctx.Query("user_id"),
		Search:   ctx.Query("search"),
		Page:     str.StringToInt(ctx.Query("page")),
		Limit:    str.StringToInt(ctx.Query("limit")),
		By:       ctx.Query("by"),
		Sort:     ctx.Query("sort"),
		ExceptId: ctx.Query("except_id"),
	}
	uc := usecase.CustomerTypeUC{ContractUC: h.ContractUC}
	res, meta, err := uc.FindAll(c, parameter)

	type StructObject struct {
		ListObjcet []models.CustomerType `json:"list_customertype"`
	}

	ObjcetData := new(StructObject)

	if res != nil {
		ObjcetData.ListObjcet = res
	}

	return h.SendResponse(ctx, ObjcetData, meta, err, 0)
}

// FindByID ...
func (h *CustomerTypeHandler) FindByID(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.CustomerTypeParameter{
		ID: ctx.Params("id"),
	}
	if parameter.ID == "" {
		return h.SendResponse(ctx, nil, nil, helper.InvalidParameter, http.StatusBadRequest)
	}

	uc := usecase.CustomerTypeUC{ContractUC: h.ContractUC}
	res, err := uc.FindByID(c, parameter)

	return h.SendResponse(ctx, res, nil, err, 0)
}
