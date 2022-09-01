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

// SalesInvoiceHandler ...
type SalesInvoiceHandler struct {
	Handler
}

// SelectAll ...
func (h *SalesInvoiceHandler) SelectAll(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.SalesInvoiceParameter{
		NoInvoice: ctx.Query("document_no"),
		Search:    ctx.Query("search"),
		By:        ctx.Query("by"),
		Sort:      ctx.Query("sort"),
	}
	uc := usecase.SalesInvoiceUC{ContractUC: h.ContractUC}
	res, err := uc.SelectAll(c, parameter)

	type InvoceHeaderObjcet struct {
		ListObjcet []models.SalesInvoice `json:"list_invoice"`
	}

	ObjcetData := new(InvoceHeaderObjcet)

	if res != nil {
		ObjcetData.ListObjcet = res
	}

	return h.SendResponse(ctx, ObjcetData, nil, err, 0)
}

// FindAll ...
func (h *SalesInvoiceHandler) FindAll(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.SalesInvoiceParameter{
		NoInvoice: ctx.Query("document_no"),
		Search:    ctx.Query("search"),
		Page:      str.StringToInt(ctx.Query("page")),
		Limit:     str.StringToInt(ctx.Query("limit")),
		By:        ctx.Query("by"),
		Sort:      ctx.Query("sort"),
	}
	uc := usecase.SalesInvoiceUC{ContractUC: h.ContractUC}
	res, meta, err := uc.FindAll(c, parameter)

	type InvoceHeaderObjcet struct {
		ListObjcet []models.SalesInvoice `json:"list_invoice"`
	}

	ObjcetData := new(InvoceHeaderObjcet)

	if res != nil {
		ObjcetData.ListObjcet = res
	}

	return h.SendResponse(ctx, ObjcetData, meta, err, 0)
}

// FindByID ...
func (h *SalesInvoiceHandler) FindByID(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.SalesInvoiceParameter{
		ID: ctx.Params("id"),
	}
	if parameter.ID == "" {
		return h.SendResponse(ctx, nil, nil, helper.InvalidParameter, http.StatusBadRequest)
	}

	uc := usecase.SalesInvoiceUC{ContractUC: h.ContractUC}
	res, err := uc.FindByID(c, parameter)

	return h.SendResponse(ctx, res, nil, err, 0)
}

// FindByDocumentNo ...
func (h *SalesInvoiceHandler) FindByDocumentNo(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.SalesInvoiceParameter{
		ID: ctx.Params("document_no"),
	}
	if parameter.ID == "" {
		return h.SendResponse(ctx, nil, nil, helper.InvalidParameter, http.StatusBadRequest)
	}

	uc := usecase.SalesInvoiceUC{ContractUC: h.ContractUC}
	res, err := uc.FindByDocumentNo(c, parameter)

	return h.SendResponse(ctx, res, nil, err, 0)
}

// FindByCustomerId ...
func (h *SalesInvoiceHandler) FindByCustomerId(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.SalesInvoiceParameter{
		ID: ctx.Params("cust_bill_to_id"),
	}
	if parameter.ID == "" {
		return h.SendResponse(ctx, nil, nil, helper.InvalidParameter, http.StatusBadRequest)
	}

	uc := usecase.SalesInvoiceUC{ContractUC: h.ContractUC}
	res, err := uc.FindByCustomerId(c, parameter)

	return h.SendResponse(ctx, res, nil, err, 0)
}
