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

// ItemSearchHandler ...
type ItemSearchHandler struct {
	Handler
}

// SelectAll ...
func (h *ItemSearchHandler) SelectAll(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.ItemSearchParameter{
		Name:               ctx.Query("item_name"),
		PriceListVersionId: ctx.Query("price_list_version_id"),
		CustomerTypeId:     ctx.Query("customer_type_id"),
		Search:             ctx.Query("search"),
		By:                 ctx.Query("by"),
		Sort:               ctx.Query("sort"),
	}
	uc := usecase.ItemSearchUC{ContractUC: h.ContractUC}
	res, err := uc.SelectAll(c, parameter)

	if parameter.PriceListVersionId == "" {
		price_list_err := " : price_list_version_id is mandatory"
		return h.SendResponse(ctx, nil, nil, helper.InvalidParameter+price_list_err, http.StatusBadRequest)
	}

	type StructObject struct {
		ListObjcet []models.ItemSearch `json:"list_item"`
	}

	ObjcetData := new(StructObject)

	if res != nil {
		ObjcetData.ListObjcet = res
	}

	return h.SendResponse(ctx, ObjcetData, nil, err, 0)
}

// FindAll ...
func (h *ItemSearchHandler) FindAll(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.ItemSearchParameter{
		Name:               ctx.Query("item_name"),
		PriceListVersionId: ctx.Query("price_list_version_id"),
		CustomerTypeId:     ctx.Query("customer_type_id"),
		Search:             ctx.Query("search"),
		Page:               str.StringToInt(ctx.Query("page")),
		Limit:              str.StringToInt(ctx.Query("limit")),
		By:                 ctx.Query("by"),
		Sort:               ctx.Query("sort"),
	}
	uc := usecase.ItemSearchUC{ContractUC: h.ContractUC}
	res, meta, err := uc.FindAll(c, parameter)

	type StructObject struct {
		ListObjcet []models.ItemSearch `json:"list_item"`
	}

	ObjcetData := new(StructObject)

	if res != nil {
		ObjcetData.ListObjcet = res
	}

	return h.SendResponse(ctx, ObjcetData, meta, err, 0)
}
