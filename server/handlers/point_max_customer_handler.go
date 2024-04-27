package handlers

import (
	"context"
	"net/http"

	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/helper"
	"nextbasis-service-v-0.1/pkg/str"
	"nextbasis-service-v-0.1/server/requests"
	"nextbasis-service-v-0.1/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// PointMaxCustomerHandler ...
type PointMaxCustomerHandler struct {
	Handler
}

// FindAll ...
func (h *PointMaxCustomerHandler) FindAll(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)
	parameter := models.PointMaxCustomerParameter{
		ShowAll: ctx.Query("show_all"),
		Page:    str.StringToInt(ctx.Query("page")),
		Limit:   str.StringToInt(ctx.Query("limit")),
		Search:  ctx.Query("search"),
		By:      ctx.Query("by"),
		Sort:    ctx.Query("sort"),
	}
	uc := usecase.PointMaxCustomerUC{ContractUC: h.ContractUC}
	res, meta, err := uc.FindAll(c, parameter)
	if err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}

	return h.SendResponse(ctx, res, meta, err, 0)
}

// SelectAll ...
func (h *PointMaxCustomerHandler) SelectAll(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)
	parameter := models.PointMaxCustomerParameter{
		ShowAll: ctx.Query("show_all"),
		Search:  ctx.Query("search"),
		By:      ctx.Query("by"),
		Sort:    ctx.Query("sort"),
	}
	uc := usecase.PointMaxCustomerUC{ContractUC: h.ContractUC}
	res, err := uc.SelectAll(c, parameter)
	if err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}

	return h.SendResponse(ctx, res, nil, err, 0)
}

// FindByID ...
func (h *PointMaxCustomerHandler) FindByID(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)
	parameter := models.PointMaxCustomerParameter{
		ID: ctx.Params("id"),
	}
	uc := usecase.PointMaxCustomerUC{ContractUC: h.ContractUC}
	res, err := uc.FindByID(c, parameter)
	if err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}

	return h.SendResponse(ctx, res, nil, err, 0)
}

// Add ...
func (h *PointMaxCustomerHandler) Add(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	input := new(requests.PointMaxCustomerRequestHeader)
	if err := ctx.BodyParser(input); err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}
	if err := h.Validator.Struct(input); err != nil {
		errMessage := h.ExtractErrorValidationMessages(err.(validator.ValidationErrors))
		return h.SendResponse(ctx, nil, nil, errMessage, http.StatusBadRequest)
	}

	uc := usecase.PointMaxCustomerUC{ContractUC: h.ContractUC}
	res, err := uc.Add(c, *input)
	if err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}

	return h.SendResponse(ctx, res, nil, err, 0)
}

// Update ...
func (h *PointMaxCustomerHandler) Update(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	id := ctx.Params("id")
	if id == "" {
		return h.SendResponse(ctx, nil, nil, helper.InvalidParameter, http.StatusBadRequest)
	}

	input := new(requests.PointMaxCustomerRequest)
	if err := ctx.BodyParser(input); err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}
	if err := h.Validator.Struct(input); err != nil {
		errMessage := h.ExtractErrorValidationMessages(err.(validator.ValidationErrors))
		return h.SendResponse(ctx, nil, nil, errMessage, http.StatusBadRequest)
	}

	uc := usecase.PointMaxCustomerUC{ContractUC: h.ContractUC}
	res, err := uc.Update(c, id, *input)
	if err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}

	return h.SendResponse(ctx, res, nil, err, 0)
}

// Delete ...
func (h *PointMaxCustomerHandler) Delete(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	id := ctx.Params("id")
	if id == "" {
		return h.SendResponse(ctx, nil, nil, helper.InvalidParameter, http.StatusBadRequest)
	}

	uc := usecase.PointMaxCustomerUC{ContractUC: h.ContractUC}
	err := uc.Delete(c, id)
	if err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}

	return h.SendResponse(ctx, nil, nil, err, 0)
}
