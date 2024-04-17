package handlers

import (
	"context"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/helper"
	"nextbasis-service-v-0.1/pkg/str"
	"nextbasis-service-v-0.1/server/requests"
	"nextbasis-service-v-0.1/usecase"
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

// ItemProductFocusHandler ...
type ItemProductFocusHandler struct {
	Handler
}

// SelectAll ...
func (h *ItemProductFocusHandler) SelectAll(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.ItemProductFocusParameter{
		ItemCategoryId:     ctx.Query("item_category_id"),
		PriceListVersionId: ctx.Query("price_list_version_id"),
		CustomerTypeId:     ctx.Query("customer_type_id"),
		Search:             ctx.Query("search"),
		By:                 ctx.Query("by"),
		Sort:               ctx.Query("sort"),
	}

	uc := usecase.ItemProductFocusUC{ContractUC: h.ContractUC}
	res, err := uc.SelectAll(c, parameter)

	type StructObject struct {
		ListObjcet []models.ItemProductFocus `json:"list_item"`
	}

	ObjcetData := new(StructObject)

	if res != nil {
		ObjcetData.ListObjcet = res
	}

	return h.SendResponse(ctx, ObjcetData, nil, err, 0)
}

// FindAll ...
func (h *ItemProductFocusHandler) FindAll(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.ItemProductFocusParameter{
		ItemCategoryId: ctx.Query("item_category_id"),
		CustomerTypeId: ctx.Query("customer_type_id"),
		Search:         ctx.Query("search"),
		Page:           str.StringToInt(ctx.Query("page")),
		Limit:          str.StringToInt(ctx.Query("limit")),
		By:             ctx.Query("by"),
		Sort:           ctx.Query("sort"),
	}
	uc := usecase.ItemProductFocusUC{ContractUC: h.ContractUC}
	res, meta, err := uc.FindAll(c, parameter)

	type StructObject struct {
		ListObjcet []models.ItemProductFocus `json:"list_item"`
	}

	ObjcetData := new(StructObject)

	if res != nil {
		ObjcetData.ListObjcet = res
	}

	return h.SendResponse(ctx, ObjcetData, meta, err, 0)
}

// FindByID ...
func (h *ItemProductFocusHandler) FindByID(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.ItemProductFocusParameter{
		ID: ctx.Params("id"),
	}
	if parameter.ID == "" {
		return h.SendResponse(ctx, nil, nil, helper.InvalidParameter, http.StatusBadRequest)
	}

	uc := usecase.ItemProductFocusUC{ContractUC: h.ContractUC}
	res, err := uc.FindByID(c, parameter)

	return h.SendResponse(ctx, res, nil, err, 0)
}

// SelectAllV2 ...
func (h *ItemProductFocusHandler) SelectAllV2(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.ItemProductFocusParameter{
		ItemCategoryId:     ctx.Query("item_category_id"),
		PriceListVersionId: ctx.Query("price_list_version_id"),
		CustomerTypeId:     ctx.Query("customer_type_id"),
		Search:             ctx.Query("search"),
		CustomerID:         ctx.Query("customer_id"),
		By:                 ctx.Query("by"),
		Sort:               ctx.Query("sort"),
	}

	if parameter.CustomerID == "" {
		return h.SendResponse(ctx, nil, nil, helper.InvalidParameter, http.StatusBadRequest)
	}

	uc := usecase.ItemProductFocusUC{ContractUC: h.ContractUC}
	res, err := uc.SelectAllV2(c, parameter)

	type StructObject struct {
		ListObjcet []viewmodel.ItemVM `json:"list_item"`
	}

	ObjcetData := new(StructObject)

	if res != nil {
		ObjcetData.ListObjcet = res
	}

	return h.SendResponse(ctx, ObjcetData, nil, err, 0)
}

// Add ...
func (h *ItemProductFocusHandler) Add(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	input := new(requests.ProductFocusRequest)
	if err := ctx.BodyParser(input); err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}
	if err := h.Validator.Struct(input); err != nil {
		errMessage := h.ExtractErrorValidationMessages(err.(validator.ValidationErrors))
		return h.SendResponse(ctx, nil, nil, errMessage, http.StatusBadRequest)
	}

	uc := usecase.ItemProductFocusUC{ContractUC: h.ContractUC}
	_, err := uc.Add(c, *input)
	if err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}

	return h.SendResponse(ctx, nil, nil, err, 200)
}
