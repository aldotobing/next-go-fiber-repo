package handlers

import (
	"context"

	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/usecase"

	"github.com/gofiber/fiber/v2"
)

// DataSyncHandler ...
type TransactionDataSyncHandler struct {
	Handler
}

func (h *TransactionDataSyncHandler) CustomerOrderVoidDataSync(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.CustomerOrderHeaderParameter{
		CustomerID: ctx.Params("customer_id"),
		Search:     ctx.Query("search"),
		By:         ctx.Query("by"),
		Sort:       ctx.Query("sort"),
	}
	uc := usecase.CustomerOrderHeaderUC{ContractUC: h.ContractUC}
	res, err := uc.VoidedDataSync(c, parameter)

	return h.SendResponse(ctx, res, nil, err, 0)
}

func (h *TransactionDataSyncHandler) InvoiceSync(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.CilentInvoiceParameter{
		CustomerID: ctx.Params("customer_id"),
		Search:     ctx.Query("search"),
		By:         ctx.Query("by"),
		Sort:       ctx.Query("sort"),
	}
	uc := usecase.CilentInvoiceUC{ContractUC: h.ContractUC}
	res, err := uc.DataSync(c, parameter)

	return h.SendResponse(ctx, res, nil, err, 0)
}
