package handlers

import (
	"context"

	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/usecase"

	"github.com/gofiber/fiber/v2"
)

// DataSyncHandler ...
type DataSyncHandler struct {
	Handler
}

// ItemDataSync ...
func (h *DataSyncHandler) ItemDataSync(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.ItemSyncParameter{
		ID:     ctx.Params("id"),
		Search: ctx.Query("search"),
		By:     ctx.Query("by"),
		Sort:   ctx.Query("sort"),
	}

	uc := usecase.ItemSyncUC{ContractUC: h.ContractUC}
	res, err := uc.DataSync(c, parameter)

	return h.SendResponse(ctx, res, nil, err, 0)
}

// PriceListDataSync ...
func (h *DataSyncHandler) PriceListDataSync(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.PriceListSyncParameter{
		ID:     ctx.Params("id"),
		Search: ctx.Query("search"),
		By:     ctx.Query("by"),
		Sort:   ctx.Query("sort"),
	}

	uc := usecase.PriceListSyncUC{ContractUC: h.ContractUC}
	res, err := uc.DataSync(c, parameter)

	return h.SendResponse(ctx, res, nil, err, 0)
}

// PriceListVersionDataSync ...
func (h *DataSyncHandler) PriceListVersionDataSync(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.PriceListVersionSyncParameter{
		ID:     ctx.Params("id"),
		Search: ctx.Query("search"),
		By:     ctx.Query("by"),
		Sort:   ctx.Query("sort"),
	}

	uc := usecase.PriceListVersionSyncUC{ContractUC: h.ContractUC}
	res, err := uc.DataSync(c, parameter)

	return h.SendResponse(ctx, res, nil, err, 0)
}

// ItemPriceDataSync ...
func (h *DataSyncHandler) ItemPriceDataSync(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.ItemPriceSyncParameter{
		ID:     ctx.Params("id"),
		Search: ctx.Query("search"),
		By:     ctx.Query("by"),
		Sort:   ctx.Query("sort"),
	}

	uc := usecase.ItemPriceSyncUC{ContractUC: h.ContractUC}
	res, err := uc.DataSync(c, parameter)

	return h.SendResponse(ctx, res, nil, err, 0)
}

// CustomerDataSync ...
func (h *DataSyncHandler) CustomerDataSync(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.CustomerDataSyncParameter{
		ID:     ctx.Params("id"),
		Search: ctx.Query("search"),
		By:     ctx.Query("by"),
		Sort:   ctx.Query("sort"),
	}

	uc := usecase.CustomerDataSyncUC{ContractUC: h.ContractUC}
	res, err := uc.DataSync(c, parameter)

	return h.SendResponse(ctx, res, nil, err, 0)
}

// CustomerDataSync ...
func (h *DataSyncHandler) SalesmanDataSync(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.SalesmanDataSyncParameter{
		ID:     ctx.Params("id"),
		Search: ctx.Query("search"),
		By:     ctx.Query("by"),
		Sort:   ctx.Query("sort"),
	}

	uc := usecase.SalesmanDataSyncUC{ContractUC: h.ContractUC}
	res, err := uc.DataSync(c, parameter)

	return h.SendResponse(ctx, res, nil, err, 0)
}

// ItemDataSync ...
func (h *DataSyncHandler) ItemUomLineDataSync(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.ItemUomLineSyncParameter{
		ID:     ctx.Params("id"),
		Search: ctx.Query("search"),
		By:     ctx.Query("by"),
		Sort:   ctx.Query("sort"),
	}

	uc := usecase.ItemUomLineSyncUC{ContractUC: h.ContractUC}
	res, err := uc.DataSync(c, parameter)

	return h.SendResponse(ctx, res, nil, err, 0)
}
