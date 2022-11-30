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

// UomHandler ...
type UomHandler struct {
	Handler
}

// SelectAll ...
func (h *UomHandler) SelectAll(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.UomParameter{
		ID:     ctx.Query("id"),
		Search: ctx.Query("search"),
		By:     ctx.Query("by"),
		Sort:   ctx.Query("sort"),
	}
	uc := usecase.UomUC{ContractUC: h.ContractUC}
	res, err := uc.SelectAll(c, parameter)

	type StructObject struct {
		ListObject []models.Uom `json:"list_uom"`
	}

	ObjectData := new(StructObject)

	if res != nil {
		ObjectData.ListObject = res
	}

	return h.SendResponse(ctx, ObjectData, nil, err, 0)
}

// FindAll ...
func (h *UomHandler) FindAll(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.UomParameter{
		ID:     ctx.Query("id"),
		Search: ctx.Query("search"),
		Page:   str.StringToInt(ctx.Query("page")),
		Limit:  str.StringToInt(ctx.Query("limit")),
		By:     ctx.Query("by"),
		Sort:   ctx.Query("sort"),
	}
	uc := usecase.UomUC{ContractUC: h.ContractUC}
	res, meta, err := uc.FindAll(c, parameter)

	type StructObject struct {
		ListObject []models.Uom `json:"list_branch"`
	}

	ObjectData := new(StructObject)

	if res != nil {
		ObjectData.ListObject = res
	}

	return h.SendResponse(ctx, ObjectData, meta, err, 0)
}

// FindByID ...
func (h *UomHandler) FindByID(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.UomParameter{
		ID: ctx.Params("id"),
	}
	if parameter.ID == "" {
		return h.SendResponse(ctx, nil, nil, helper.InvalidParameter, http.StatusBadRequest)
	}

	uc := usecase.UomUC{ContractUC: h.ContractUC}
	res, err := uc.FindByID(c, parameter)

	return h.SendResponse(ctx, res, nil, err, 0)
}
