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

// ItemPriceHandler ...
type ItemPriceHandler struct {
	Handler
}

// SelectAll ...
func (h *ItemPriceHandler) SelectAll(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.ItemPriceParameter{
		ID:                 ctx.Query("id"),
		UomID:              ctx.Query("uom_id"),
		PriceListVersionID: ctx.Query("price_list_version_id"),
		Search:             ctx.Query("search"),
		By:                 ctx.Query("by"),
		Sort:               ctx.Query("sort"),
	}
	uc := usecase.ItemPriceUC{ContractUC: h.ContractUC}
	res, err := uc.SelectAll(c, parameter)

	type StructObject struct {
		ListObject []models.ItemPrice `json:"list_item_price"`
	}

	ObjectData := new(StructObject)

	if res != nil {
		ObjectData.ListObject = res
	}

	return h.SendResponse(ctx, ObjectData, nil, err, 0)
}

// FindAll ...
func (h *ItemPriceHandler) FindAll(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.ItemPriceParameter{
		ID:                 ctx.Query("id"),
		UomID:              ctx.Query("uom_id"),
		PriceListVersionID: ctx.Query("price_list_version_id"),
		Search:             ctx.Query("search"),
		Page:               str.StringToInt(ctx.Query("page")),
		Limit:              str.StringToInt(ctx.Query("limit")),
		By:                 ctx.Query("by"),
		Sort:               ctx.Query("sort"),
	}
	uc := usecase.ItemPriceUC{ContractUC: h.ContractUC}
	res, meta, err := uc.FindAll(c, parameter)

	type StructObject struct {
		ListObject []models.ItemPrice `json:"list_item_price"`
	}

	ObjectData := new(StructObject)

	if res != nil {
		ObjectData.ListObject = res
	}

	return h.SendResponse(ctx, ObjectData, meta, err, 0)
}

// FindByID ...
func (h *ItemPriceHandler) FindByID(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.ItemPriceParameter{
		ID: ctx.Params("id"),
	}
	if parameter.ID == "" {
		price_list_err := " : ID is mandatory"
		return h.SendResponse(ctx, nil, nil, helper.InvalidParameter+price_list_err, http.StatusBadRequest)
	}

	uc := usecase.ItemPriceUC{ContractUC: h.ContractUC}
	res, err := uc.FindByID(c, parameter)

	return h.SendResponse(ctx, res, nil, err, 0)
}
