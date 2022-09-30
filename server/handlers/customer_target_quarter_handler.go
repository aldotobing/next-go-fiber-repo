package handlers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/str"
	"nextbasis-service-v-0.1/usecase"
)

// CustomerTargetQuarterHandler ...
type CustomerTargetQuarterHandler struct {
	Handler
}

// SelectAll ...
func (h *CustomerTargetQuarterHandler) SelectAll(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.CustomerTargetQuarterParameter{
		ID:     ctx.Query("customer_id"),
		Search: ctx.Query("search"),
		By:     ctx.Query("by"),
		Sort:   ctx.Query("sort"),
	}
	uc := usecase.CustomerTargetQuarterUC{ContractUC: h.ContractUC}
	res, err := uc.SelectAll(c, parameter)

	type StructObject struct {
		ListObjcet []models.CustomerTargetQuarter `json:"list_customer"`
	}

	ObjcetData := new(StructObject)

	if res != nil {
		ObjcetData.ListObjcet = res
	}

	return h.SendResponse(ctx, ObjcetData, nil, err, 0)
}

// FindAll ...
func (h *CustomerTargetQuarterHandler) FindAll(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.CustomerTargetQuarterParameter{
		ID:     ctx.Query("customer_id"),
		Search: ctx.Query("search"),
		Page:   str.StringToInt(ctx.Query("page")),
		Limit:  str.StringToInt(ctx.Query("limit")),
		By:     ctx.Query("by"),
		Sort:   ctx.Query("sort"),
	}
	uc := usecase.CustomerTargetQuarterUC{ContractUC: h.ContractUC}
	res, meta, err := uc.FindAll(c, parameter)

	type StructObject struct {
		ListObject []models.CustomerTargetQuarter `json:"list_customer"`
	}

	ObjectData := new(StructObject)

	if res != nil {
		ObjectData.ListObject = res
	}

	return h.SendResponse(ctx, ObjectData, meta, err, 0)
}

// FindByID ...
// func (h *CustomerTargetQuarterHandler) FindByID(ctx *fiber.Ctx) error {
// 	c := ctx.Locals("ctx").(context.Context)

// 	parameter := models.CustomerTargetQuarterParameter{
// 		ID: ctx.Params("customer_id"),
// 	}
// 	if parameter.ID == "" {
// 		return h.SendResponse(ctx, nil, nil, helper.InvalidParameter, http.StatusBadRequest)
// 	}

// 	uc := usecase.CustomerTargetQuarterUC{ContractUC: h.ContractUC}
// 	res, err := uc.FindByID(c, parameter)

// 	type StructObject struct {
// 		ListObject models.CustomerTargetQuarter `json:"customer"`
// 	}

// 	ObjectData := new(StructObject)

// 	// if res != nil {
// 	ObjectData.ListObject = res
// 	// }

// 	return h.SendResponse(ctx, ObjectData, nil, err, 0)
// }
