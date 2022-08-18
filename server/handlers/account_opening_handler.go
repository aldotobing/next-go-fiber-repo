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

// AccountOpeningHandler ...
type AccountOpeningHandler struct {
	Handler
}

// SelectAll ...
func (h *AccountOpeningHandler) SelectAll(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.AccountOpeningParameter{
		UserID:           ctx.Query("user_id"),
		MaritalStatusID:  ctx.Query("marital_status_id"),
		GenderID:         ctx.Query("gender_id"),
		BirthPlaceCityID: ctx.Query("birth_place_city_id"),
		Search:           ctx.Query("search"),
		Status:           ctx.Query("status"),
		By:               ctx.Query("by"),
		Sort:             ctx.Query("sort"),
	}
	uc := usecase.AccountOpeningUC{ContractUC: h.ContractUC}
	res, err := uc.SelectAll(c, parameter)

	return h.SendResponse(ctx, res, nil, err, 0)
}

// FindAll ...
func (h *AccountOpeningHandler) FindAll(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.AccountOpeningParameter{
		UserID:           ctx.Query("user_id"),
		MaritalStatusID:  ctx.Query("marital_status_id"),
		GenderID:         ctx.Query("gender_id"),
		BirthPlaceCityID: ctx.Query("birth_place_city_id"),
		Search:           ctx.Query("search"),
		Status:           ctx.Query("status"),
		Page:             str.StringToInt(ctx.Query("page")),
		Limit:            str.StringToInt(ctx.Query("limit")),
		By:               ctx.Query("by"),
		Sort:             ctx.Query("sort"),
	}
	uc := usecase.AccountOpeningUC{ContractUC: h.ContractUC}
	res, err := uc.SelectAll(c, parameter)

	return h.SendResponse(ctx, res, nil, err, 0)
}

// FindByID ...
func (h *AccountOpeningHandler) FindByID(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.AccountOpeningParameter{
		ID: ctx.Params("id"),
	}
	if parameter.ID == "" {
		return h.SendResponse(ctx, nil, nil, helper.InvalidParameter, http.StatusBadRequest)
	}

	uc := usecase.AccountOpeningUC{ContractUC: h.ContractUC}
	res, err := uc.FindByID(c, parameter)

	return h.SendResponse(ctx, res, nil, err, 0)
}

// Add ...
func (h *AccountOpeningHandler) Add(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	input := new(requests.AccountOpeningRequest)
	if err := ctx.BodyParser(input); err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}
	if err := h.Validator.Struct(input); err != nil {
		errMessage := h.ExtractErrorValidationMessages(err.(validator.ValidationErrors))
		return h.SendResponse(ctx, nil, nil, errMessage, http.StatusBadRequest)
	}

	uc := usecase.AccountOpeningUC{ContractUC: h.ContractUC}
	res, err := uc.Add(c, input)

	return h.SendResponse(ctx, res, nil, err, 0)
}

// Edit ...
func (h *AccountOpeningHandler) Edit(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	id := ctx.Params("id")
	if id == "" {
		return h.SendResponse(ctx, nil, nil, helper.InvalidParameter, http.StatusBadRequest)
	}

	input := new(requests.AccountOpeningRequest)
	if err := ctx.BodyParser(input); err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}
	if err := h.Validator.Struct(input); err != nil {
		errMessage := h.ExtractErrorValidationMessages(err.(validator.ValidationErrors))
		return h.SendResponse(ctx, nil, nil, errMessage, http.StatusBadRequest)
	}

	uc := usecase.AccountOpeningUC{ContractUC: h.ContractUC}
	res, err := uc.Edit(c, id, input)

	return h.SendResponse(ctx, res, nil, err, 0)
}

// Delete ...
func (h *AccountOpeningHandler) Delete(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	id := ctx.Params("id")
	if id == "" {
		return h.SendResponse(ctx, nil, nil, helper.InvalidParameter, http.StatusBadRequest)
	}

	uc := usecase.AccountOpeningUC{ContractUC: h.ContractUC}
	res, err := uc.Delete(c, id)

	return h.SendResponse(ctx, res, nil, err, 0)
}
