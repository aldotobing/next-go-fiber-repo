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

// ItemDetailsHandler ...
type ItemDetailsHandler struct {
	Handler
}

// SelectAll ...
func (h *ItemDetailsHandler) SelectAll(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.ItemDetailsParameter{
		ItemDetailsCategoryId: ctx.Query("item_category_id"),
		UomID:                 ctx.Query("uom_id"),
		PriceListVersionId:    ctx.Query("price_list_version_id"),
		Search:                ctx.Query("search"),
		By:                    ctx.Query("by"),
		Sort:                  ctx.Query("sort"),
		ExceptId:              ctx.Query("except_id"),
	}
	uc := usecase.ItemDetailsUC{ContractUC: h.ContractUC}
	res, err := uc.SelectAll(c, parameter)

	type StructObject struct {
		ListObjcet []models.ItemDetails `json:"list_item"`
	}

	ObjcetData := new(StructObject)

	if res != nil {
		ObjcetData.ListObjcet = res
	}

	return h.SendResponse(ctx, ObjcetData, nil, err, 0)
}

// FindAll ...
func (h *ItemDetailsHandler) FindAll(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.ItemDetailsParameter{
		ItemDetailsCategoryId: ctx.Query("item_category_id"),
		PriceListVersionId:    ctx.Query("price_list_version_id"),
		UomID:                 ctx.Query("uom_id"),
		Search:                ctx.Query("search"),
		Page:                  str.StringToInt(ctx.Query("page")),
		Limit:                 str.StringToInt(ctx.Query("limit")),
		By:                    ctx.Query("by"),
		Sort:                  ctx.Query("sort"),
		ExceptId:              ctx.Query("except_id"),
	}
	uc := usecase.ItemDetailsUC{ContractUC: h.ContractUC}
	res, meta, err := uc.FindAll(c, parameter)

	type StructObject struct {
		ListObjcet []models.ItemDetails `json:"list_item"`
	}

	ObjcetData := new(StructObject)

	if res != nil {
		ObjcetData.ListObjcet = res
	}

	return h.SendResponse(ctx, ObjcetData, meta, err, 0)
}

// FindByID ...
func (h *ItemDetailsHandler) FindByID(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.ItemDetailsParameter{
		ID:                 ctx.Params("id"),
		PriceListVersionId: ctx.Query("price_list_version_id"),
		UomID:              ctx.Query("uom_id"),
	}
	if parameter.ID == "" {
		return h.SendResponse(ctx, nil, nil, helper.InvalidParameter, http.StatusBadRequest)
	}

	if parameter.PriceListVersionId == "" {
		price_list_err := " : price_list_version_id is mandatory"
		return h.SendResponse(ctx, nil, nil, helper.InvalidParameter+price_list_err, http.StatusBadRequest)
	}

	uc := usecase.ItemDetailsUC{ContractUC: h.ContractUC}
	res, err := uc.FindByID(c, parameter)

	return h.SendResponse(ctx, res, nil, err, 0)
}

// FindByIDV2 ...
func (h *ItemDetailsHandler) FindByIDV2(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.ItemDetailsParameter{
		ID:                 ctx.Params("id"),
		PriceListVersionId: ctx.Query("price_list_version_id"),
		UomID:              ctx.Query("uom_id"),
	}
	if parameter.ID == "" {
		return h.SendResponse(ctx, nil, nil, helper.InvalidParameter, http.StatusBadRequest)
	}

	if parameter.PriceListVersionId == "" {
		price_list_err := " : price_list_version_id is mandatory"
		return h.SendResponse(ctx, nil, nil, helper.InvalidParameter+price_list_err, http.StatusBadRequest)
	}

	uc := usecase.ItemDetailsUC{ContractUC: h.ContractUC}
	res, err := uc.FindByIDV2(c, parameter)

	return h.SendResponse(ctx, res, nil, err, 0)
}
