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

// ProductFocusCategoryHandler ...
type ProductFocusCategoryHandler struct {
	Handler
}

// SelectAll ...
func (h *ProductFocusCategoryHandler) SelectAll(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.ProductFocusCategoryParameter{
		Search: ctx.Query("search"),
		By:     ctx.Query("by"),
		Sort:   ctx.Query("sort"),
	}
	uc := usecase.ProductFocusCategoryUC{ContractUC: h.ContractUC}
	res, err := uc.SelectAll(c, parameter)

	type ResponseObject struct {
		ListObject []models.ProductFocusCategory `json:"list_category"`
	}

	ObjectData := new(ResponseObject)

	if res != nil {
		ObjectData.ListObject = res
	}

	return h.SendResponse(ctx, ObjectData, nil, err, 0)
}

// FindAll ...
func (h *ProductFocusCategoryHandler) FindAll(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.ProductFocusCategoryParameter{
		Search: ctx.Query("search"),
		Page:   str.StringToInt(ctx.Query("page")),
		Limit:  str.StringToInt(ctx.Query("limit")),
		By:     ctx.Query("by"),
		Sort:   ctx.Query("sort"),
	}
	uc := usecase.ProductFocusCategoryUC{ContractUC: h.ContractUC}
	res, meta, err := uc.FindAll(c, parameter)

	return h.SendResponse(ctx, res, meta, err, 0)
}

// FindByID ...
func (h *ProductFocusCategoryHandler) FindByID(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.ProductFocusCategoryParameter{
		ID: ctx.Params("id"),
	}
	if parameter.ID == "" {
		return h.SendResponse(ctx, nil, nil, helper.InvalidParameter, http.StatusBadRequest)
	}

	uc := usecase.ProductFocusCategoryUC{ContractUC: h.ContractUC}
	res, err := uc.FindByID(c, parameter)

	return h.SendResponse(ctx, res, nil, err, 0)
}

// FindByBranchID ...
func (h *ProductFocusCategoryHandler) FindByBranchID(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.ProductFocusCategoryParameter{
		BRANCHID: ctx.Params("branchid"),
	}
	if parameter.ID == "" {
		return h.SendResponse(ctx, nil, nil, helper.InvalidParameter, http.StatusBadRequest)
	}

	uc := usecase.ProductFocusCategoryUC{ContractUC: h.ContractUC}
	res, err := uc.FindByBranchID(c, parameter)

	return h.SendResponse(ctx, res, nil, err, 0)
}

// FindByCategoryID ...
func (h *ProductFocusCategoryHandler) FindByCategoryID(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.ProductFocusCategoryParameter{
		Search: ctx.Query("search"),
		Page:   str.StringToInt(ctx.Query("page")),
		Limit:  str.StringToInt(ctx.Query("limit")),
		By:     ctx.Query("by"),
		Sort:   ctx.Query("sort"),
	}
	uc := usecase.ProductFocusCategoryUC{ContractUC: h.ContractUC}
	res, meta, err := uc.FindByCategoryID(c, parameter)

	return h.SendResponse(ctx, res, meta, err, 0)
}
