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

// ItemPromoHandler ...
type ItemPromoHandler struct {
	Handler
}

// SelectAll ...
func (h *ItemPromoHandler) SelectAll(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.ItemPromoParameter{
		ItemID:    ctx.Query("item_id"),
		StartDate: ctx.Query("start_date"),
		EndDate:   ctx.Query("end_date"),
		Search:    ctx.Query("search"),
		By:        ctx.Query("by"),
		Sort:      ctx.Query("sort"),
	}
	uc := usecase.ItemPromoUC{ContractUC: h.ContractUC}
	res, err := uc.SelectAll(c, parameter)

	type StructObject struct {
		ListObjcet []models.ItemPromo `json:"list_item"`
	}

	ObjcetData := new(StructObject)

	if res != nil {
		ObjcetData.ListObjcet = res
	}

	return h.SendResponse(ctx, ObjcetData, nil, err, 0)
}

// FindAll ...
func (h *ItemPromoHandler) FindAll(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.ItemPromoParameter{
		ItemID:    ctx.Query("item_id"),
		StartDate: ctx.Query("start_date"),
		Search:    ctx.Query("search"),
		Page:      str.StringToInt(ctx.Query("page")),
		Limit:     str.StringToInt(ctx.Query("limit")),
		By:        ctx.Query("by"),
		Sort:      ctx.Query("sort"),
	}
	uc := usecase.ItemPromoUC{ContractUC: h.ContractUC}
	res, meta, err := uc.FindAll(c, parameter)

	type StructObject struct {
		ListObjcet []models.ItemPromo `json:"list_item"`
	}

	ObjcetData := new(StructObject)

	if res != nil {
		ObjcetData.ListObjcet = res
	}

	return h.SendResponse(ctx, ObjcetData, meta, err, 0)
}

// FindByID ...
func (h *ItemPromoHandler) FindByID(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.ItemPromoParameter{
		ItemID: ctx.Params("item_id"),
	}
	if parameter.ItemID == "" {
		return h.SendResponse(ctx, nil, nil, helper.InvalidParameter, http.StatusBadRequest)
	}

	uc := usecase.ItemPromoUC{ContractUC: h.ContractUC}
	res, err := uc.FindByID(c, parameter)

	return h.SendResponse(ctx, res, nil, err, 0)
}
