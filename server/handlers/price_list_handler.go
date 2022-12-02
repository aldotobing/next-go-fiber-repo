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

// PriceListHandler ...
type PriceListHandler struct {
	Handler
}

// SelectAll ...
func (h *PriceListHandler) SelectAll(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.PriceListParameter{
		ID:       ctx.Query("id"),
		BranchID: ctx.Query("branch_id"),
		Search:   ctx.Query("search"),
		By:       ctx.Query("by"),
		Sort:     ctx.Query("sort"),
		ExceptId: ctx.Query("except_id"),
	}
	uc := usecase.PriceListUC{ContractUC: h.ContractUC}
	res, err := uc.SelectAll(c, parameter)

	type StructObject struct {
		ListObject []models.PriceList `json:"list_pricelist"`
	}

	ObjectData := new(StructObject)

	if res != nil {
		ObjectData.ListObject = res
	}

	return h.SendResponse(ctx, ObjectData, nil, err, 0)
}

// FindAll ...
func (h *PriceListHandler) FindAll(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.PriceListParameter{
		ID:       ctx.Query("id"),
		BranchID: ctx.Query("branch_id"),
		Search:   ctx.Query("search"),
		Page:     str.StringToInt(ctx.Query("page")),
		Limit:    str.StringToInt(ctx.Query("limit")),
		By:       ctx.Query("by"),
		Sort:     ctx.Query("sort"),
		ExceptId: ctx.Query("except_id"),
	}
	uc := usecase.PriceListUC{ContractUC: h.ContractUC}
	res, meta, err := uc.FindAll(c, parameter)

	type StructObject struct {
		ListObjcet []models.PriceList `json:"list_pricelist"`
	}

	ObjcetData := new(StructObject)

	if res != nil {
		ObjcetData.ListObjcet = res
	}

	return h.SendResponse(ctx, ObjcetData, meta, err, 0)
}

// FindByID ...
func (h *PriceListHandler) FindByID(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.PriceListParameter{
		ID: ctx.Params("id"),
	}
	if parameter.ID == "" {
		return h.SendResponse(ctx, nil, nil, helper.InvalidParameter, http.StatusBadRequest)
	}

	uc := usecase.PriceListUC{ContractUC: h.ContractUC}
	res, err := uc.FindByID(c, parameter)

	return h.SendResponse(ctx, res, nil, err, 0)
}
