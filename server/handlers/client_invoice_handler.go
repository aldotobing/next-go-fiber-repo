package handlers

import (
	"context"
	"net/http"

	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/helper"
	"nextbasis-service-v-0.1/pkg/str"
	"nextbasis-service-v-0.1/usecase"

	"github.com/gofiber/fiber/v2"
)

// CilentInvoiceHandler ...
type CilentInvoiceHandler struct {
	Handler
}

// SelectAll ...
func (h *CilentInvoiceHandler) SelectAll(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.CilentInvoiceParameter{
		CustomerID: ctx.Params("customer_id"),
		Search:     ctx.Query("search"),
		By:         ctx.Query("by"),
		Sort:       ctx.Query("sort"),
	}
	uc := usecase.CilentInvoiceUC{ContractUC: h.ContractUC}
	res, err := uc.SelectAll(c, parameter)

	return h.SendResponse(ctx, res, nil, err, 0)
}

// FindAll ...
func (h *CilentInvoiceHandler) FindAll(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.CilentInvoiceParameter{
		CustomerID: ctx.Params("customer_id"),
		Search:     ctx.Query("search"),
		Page:       str.StringToInt(ctx.Query("page")),
		Limit:      str.StringToInt(ctx.Query("limit")),
		By:         ctx.Query("by"),
		Sort:       ctx.Query("sort"),
	}
	uc := usecase.CilentInvoiceUC{ContractUC: h.ContractUC}
	res, meta, err := uc.FindAll(c, parameter)

	type StructObject struct {
		ListObjcet []models.CilentInvoice `json:"list_customer_order"`
	}

	ObjcetData := new(StructObject)

	if res != nil {
		ObjcetData.ListObjcet = res
	}

	return h.SendResponse(ctx, ObjcetData, meta, err, 0)
}

// FindByID ...
func (h *CilentInvoiceHandler) FindByID(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.CilentInvoiceParameter{
		ID: ctx.Params("id"),
	}
	if parameter.ID == "" {
		return h.SendResponse(ctx, nil, nil, helper.InvalidParameter, http.StatusBadRequest)
	}

	uc := usecase.CilentInvoiceUC{ContractUC: h.ContractUC}
	res, err := uc.FindByID(c, parameter)

	return h.SendResponse(ctx, res, nil, err, 0)
}

func (h *CilentInvoiceHandler) SelectAll3RD(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.CilentInvoiceParameter{
		CustomerID: ctx.Params("customer_id"),
		Search:     ctx.Query("search"),
		By:         ctx.Query("by"),
		Sort:       ctx.Query("sort"),
	}
	uc := usecase.CilentInvoiceUC{ContractUC: h.ContractUC}
	res, err := uc.SelectAll3RD(c, parameter)

	return h.SendResponse(ctx, res, nil, err, 0)
}
