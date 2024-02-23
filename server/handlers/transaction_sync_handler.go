package handlers

import (
	"context"

	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/str"
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

func (h *TransactionDataSyncHandler) InvoiceSyncPutToRedis(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.CilentInvoiceParameter{
		CustomerID: ctx.Params("customer_id"),
		Search:     ctx.Query("search"),
		By:         ctx.Query("by"),
		Sort:       ctx.Query("sort"),
	}
	uc := usecase.CilentInvoiceUC{ContractUC: h.ContractUC}
	res, err := uc.PutRedisDataSync(c, parameter)

	return h.SendResponse(ctx, res, nil, err, 0)
}

func (h *TransactionDataSyncHandler) InvoiceSyncGetRedis(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	uc := usecase.CilentInvoiceUC{ContractUC: h.ContractUC}
	res, err := uc.GetRedisDataSync(c)

	return h.SendResponse(ctx, res, nil, err, 0)
}

func (h *TransactionDataSyncHandler) InvoiceReserveSyncGetRedis(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	uc := usecase.CilentInvoiceUC{ContractUC: h.ContractUC}
	res, err := uc.GetRedisDataReserveSync(c)

	return h.SendResponse(ctx, res, nil, err, 0)
}

func (h *TransactionDataSyncHandler) ReturnInvoiceSync(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.CilentReturnInvoiceParameter{
		CustomerID: ctx.Params("customer_id"),
		Search:     ctx.Query("search"),
		By:         ctx.Query("by"),
		Sort:       ctx.Query("sort"),
	}
	uc := usecase.CilentReturnInvoiceUC{ContractUC: h.ContractUC}
	res, err := uc.DataSync(c, parameter)

	return h.SendResponse(ctx, res, nil, err, 0)
}

func (h *TransactionDataSyncHandler) UndoneDataSync(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.CilentInvoiceParameter{
		CustomerID: ctx.Params("customer_id"),
		Search:     ctx.Query("search"),
		By:         ctx.Query("by"),
		Sort:       ctx.Query("sort"),
		StartDate:  ctx.Query("start_date"),
		EndDate:    ctx.Query("end_date"),
		Page:       str.StringToInt(ctx.Query("page")),
		Limit:      str.StringToInt(ctx.Query("limit")),
	}
	uc := usecase.CilentInvoiceUC{ContractUC: h.ContractUC}
	res, err := uc.UndoneDataSync(c, parameter)

	return h.SendResponse(ctx, res, nil, err, 0)
}

func (h *TransactionDataSyncHandler) SalesOrderCustomerSync(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.SalesOrderCustomerSyncParameter{
		CustomerID: ctx.Params("customer_id"),
		Search:     ctx.Query("search"),
		By:         ctx.Query("by"),
		Sort:       ctx.Query("sort"),
	}
	uc := usecase.SalesOrderCustomerSyncUC{ContractUC: h.ContractUC}
	res, err := uc.DataSync(c, parameter)

	return h.SendResponse(ctx, res, nil, err, 0)
}

func (h *TransactionDataSyncHandler) SalesOrderCustomerRevisedSync(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.SalesOrderCustomerSyncParameter{
		CustomerID: ctx.Params("customer_id"),
		Search:     ctx.Query("search"),
		By:         ctx.Query("by"),
		Sort:       ctx.Query("sort"),
	}
	uc := usecase.SalesOrderCustomerSyncUC{ContractUC: h.ContractUC}
	res, err := uc.RevisedSync(c, parameter)

	return h.SendResponse(ctx, res, nil, err, 0)
}
