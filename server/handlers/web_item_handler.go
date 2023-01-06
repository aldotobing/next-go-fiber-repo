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

// WebItemHandler ...
type WebItemHandler struct {
	Handler
}

// SelectAll ...
func (h *WebItemHandler) SelectAll(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.WebItemParameter{
		ID:             ctx.Query("item_id"),
		ItemCategoryId: ctx.Query("item_category_id"),
		Search:         ctx.Query("search"),
		By:             ctx.Query("by"),
		Sort:           ctx.Query("sort"),
		ExceptId:       ctx.Query("except_id"),
	}
	uc := usecase.WebItemUC{ContractUC: h.ContractUC}
	res, err := uc.SelectAll(c, parameter)

	type StructObject struct {
		ListObject []models.WebItem `json:"list_item"`
	}

	ObjectData := new(StructObject)

	if res != nil {
		ObjectData.ListObject = res
	}

	return h.SendResponse(ctx, ObjectData, nil, err, 0)
}

// FindAll ...
func (h *WebItemHandler) FindAll(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.WebItemParameter{
		ID:             ctx.Query("item_id"),
		ItemCategoryId: ctx.Query("item_category_id"),
		Search:         ctx.Query("search"),
		Page:           str.StringToInt(ctx.Query("page")),
		Limit:          str.StringToInt(ctx.Query("limit")),
		By:             ctx.Query("by"),
		Sort:           ctx.Query("sort"),
		ExceptId:       ctx.Query("except_id"),
	}
	uc := usecase.WebItemUC{ContractUC: h.ContractUC}
	res, meta, err := uc.FindAll(c, parameter)

	type StructObject struct {
		ListObject []models.WebItem `json:"list_item"`
	}

	ObjectData := new(StructObject)

	if res != nil {
		ObjectData.ListObject = res
	}

	return h.SendResponse(ctx, ObjectData, meta, err, 0)
}

// FindByID ...
func (h *WebItemHandler) FindByID(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.WebItemParameter{
		ID: ctx.Params("id"),
	}
	if parameter.ID == "" {
		return h.SendResponse(ctx, nil, nil, helper.InvalidParameter, http.StatusBadRequest)
	}

	uc := usecase.WebItemUC{ContractUC: h.ContractUC}
	res, err := uc.FindByID(c, parameter)

	return h.SendResponse(ctx, res, nil, err, 0)
}
