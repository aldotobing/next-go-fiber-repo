package handlers

import (
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/helper"
	"nextbasis-service-v-0.1/pkg/str"
	"nextbasis-service-v-0.1/usecase"
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

// ItemHandler ...
type ItemHandler struct {
	Handler
}

// SelectAll ...
func (h *ItemHandler) SelectAll(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.ItemParameter{
		ID:                 ctx.Query("item_id"),
		ItemCategoryId:     ctx.Query("item_category_id"),
		UomID:              ctx.Query("uom_id"),
		PriceListVersionId: ctx.Query("price_list_version_id"),
		CustomerTypeId:     ctx.Query("customer_type_id"),
		Search:             ctx.Query("search"),
		By:                 ctx.Query("by"),
		Sort:               ctx.Query("sort"),
		ExceptId:           ctx.Query("except_id"),
	}
	uc := usecase.ItemUC{ContractUC: h.ContractUC}
	res, err := uc.SelectAll(c, parameter)

	type StructObject struct {
		ListObject []models.Item `json:"list_item"`
	}

	ObjectData := new(StructObject)

	if res != nil {
		ObjectData.ListObject = res
	}

	return h.SendResponse(ctx, ObjectData, nil, err, 0)
}

// SelectAllV2 ...
func (h *ItemHandler) SelectAllV2(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.ItemParameter{
		ID:                 ctx.Query("item_id"),
		ItemCategoryId:     ctx.Query("item_category_id"),
		UomID:              ctx.Query("uom_id"),
		PriceListVersionId: ctx.Query("price_list_version_id"),
		PriceListId:        ctx.Query("price_list_id"),
		CustomerTypeId:     ctx.Query("customer_type_id"),
		Search:             ctx.Query("search"),
		By:                 ctx.Query("by"),
		Sort:               ctx.Query("sort"),
		ExceptId:           ctx.Query("except_id"),
	}
	uc := usecase.ItemUC{ContractUC: h.ContractUC}
	res, err := uc.SelectAllV2(c, parameter, false, false)

	type StructObject struct {
		ListObject []viewmodel.ItemVM `json:"list_item"`
	}

	ObjectData := new(StructObject)

	if res != nil {
		ObjectData.ListObject = res
	}

	return h.SendResponse(ctx, ObjectData, nil, err, 0)
}

// FindAll ...
func (h *ItemHandler) FindAll(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.ItemParameter{
		ID:                 ctx.Query("item_id"),
		ItemCategoryId:     ctx.Query("item_category_id"),
		PriceListVersionId: ctx.Query("price_list_version_id"),
		UomID:              ctx.Query("uom_id"),
		CustomerTypeId:     ctx.Query("customer_type_id"),
		Search:             ctx.Query("search"),
		Page:               str.StringToInt(ctx.Query("page")),
		Limit:              str.StringToInt(ctx.Query("limit")),
		By:                 ctx.Query("by"),
		Sort:               ctx.Query("sort"),
		ExceptId:           ctx.Query("except_id"),
	}
	uc := usecase.ItemUC{ContractUC: h.ContractUC}
	res, meta, err := uc.FindAll(c, parameter)

	type StructObject struct {
		ListObject []models.Item `json:"list_item"`
	}

	ObjectData := new(StructObject)

	if res != nil {
		ObjectData.ListObject = res
	}

	return h.SendResponse(ctx, ObjectData, meta, err, 0)
}

// FindByID ...
func (h *ItemHandler) FindByID(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.ItemParameter{
		ID: ctx.Params("id"),
	}
	if parameter.ID == "" {
		return h.SendResponse(ctx, nil, nil, helper.InvalidParameter, http.StatusBadRequest)
	}

	uc := usecase.ItemUC{ContractUC: h.ContractUC}
	res, err := uc.FindByID(c, parameter)

	return h.SendResponse(ctx, res, nil, err, 0)
}
