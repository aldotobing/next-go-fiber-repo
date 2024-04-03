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

// PointPromoHandler ...
type PointPromoHandler struct {
	Handler
}

// FindAll ...
func (h *PointPromoHandler) FindAll(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)
	parameter := models.PointPromoParameter{
		CustomerID: ctx.Query("customer_id"),
		PointType:  ctx.Query("point_type"),
		StartDate:  ctx.Query("start_date"),
		EndDate:    ctx.Query("end_date"),
		Page:       str.StringToInt(ctx.Query("page")),
		Limit:      str.StringToInt(ctx.Query("limit")),
		By:         ctx.Query("by"),
		Sort:       ctx.Query("sort"),
	}
	uc := usecase.PointPromoUC{ContractUC: h.ContractUC}
	res, meta, err := uc.FindAll(c, parameter)
	if err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}

	return h.SendResponse(ctx, res, meta, err, 0)
}

// SelectAll ...
func (h *PointPromoHandler) SelectAll(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)
	parameter := models.PointPromoParameter{
		CustomerID: ctx.Query("customer_id"),
		PointType:  ctx.Query("point_type"),
		By:         ctx.Query("by"),
		Sort:       ctx.Query("sort"),
	}
	uc := usecase.PointPromoUC{ContractUC: h.ContractUC}
	res, err := uc.SelectAll(c, parameter)
	if err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}

	return h.SendResponse(ctx, res, nil, err, 0)
}

// FindByID ...
func (h *PointPromoHandler) FindByID(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)
	parameter := models.PointPromoParameter{
		ID: ctx.Params("id"),
	}
	uc := usecase.PointPromoUC{ContractUC: h.ContractUC}
	res, err := uc.FindByID(c, parameter)
	if err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}

	return h.SendResponse(ctx, res, nil, err, 0)
}

// EligiblePoint ...
func (h *PointPromoHandler) EligiblePoint(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)
	chartIDList := ctx.Query("chart_id_list")

	uc := usecase.PointPromoUC{ContractUC: h.ContractUC}
	res, err := uc.EligiblePoint(c, chartIDList)
	if err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}

	return h.SendResponse(ctx, res, nil, err, 0)
}

// Add ...
func (h *PointPromoHandler) Add(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	input := new(requests.PointPromoRequest)
	if err := ctx.BodyParser(input); err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}
	if err := h.Validator.Struct(input); err != nil {
		errMessage := h.ExtractErrorValidationMessages(err.(validator.ValidationErrors))
		return h.SendResponse(ctx, nil, nil, errMessage, http.StatusBadRequest)
	}

	uc := usecase.PointPromoUC{ContractUC: h.ContractUC}
	res, err := uc.Add(c, *input)
	if err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}

	return h.SendResponse(ctx, res, nil, err, 0)
}

// Photo ...
func (h *PointPromoHandler) Photo(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	image, _ := ctx.FormFile("image")

	uc := usecase.PointPromoUC{ContractUC: h.ContractUC}
	res, err := uc.AddPhoto(c, image)
	if err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}

	return h.SendResponse(ctx, res, nil, err, 0)
}

// Update ...
func (h *PointPromoHandler) Update(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	id := ctx.Params("id")
	if id == "" {
		return h.SendResponse(ctx, nil, nil, helper.InvalidParameter, http.StatusBadRequest)
	}

	input := new(requests.PointPromoRequest)
	if err := ctx.BodyParser(input); err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}
	if err := h.Validator.Struct(input); err != nil {
		errMessage := h.ExtractErrorValidationMessages(err.(validator.ValidationErrors))
		return h.SendResponse(ctx, nil, nil, errMessage, http.StatusBadRequest)
	}

	uc := usecase.PointPromoUC{ContractUC: h.ContractUC}
	res, err := uc.Update(c, id, *input)
	if err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}

	return h.SendResponse(ctx, res, nil, err, 0)
}

// Delete ...
func (h *PointPromoHandler) Delete(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	id := ctx.Params("id")
	if id == "" {
		return h.SendResponse(ctx, nil, nil, helper.InvalidParameter, http.StatusBadRequest)
	}

	uc := usecase.PointPromoUC{ContractUC: h.ContractUC}
	err := uc.Delete(c, id)
	if err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}

	return h.SendResponse(ctx, nil, nil, err, 0)
}
