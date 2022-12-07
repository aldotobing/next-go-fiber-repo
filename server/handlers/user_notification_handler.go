package handlers

import (
	"context"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/helper"
	"nextbasis-service-v-0.1/pkg/str"
	"nextbasis-service-v-0.1/server/requests"
	"nextbasis-service-v-0.1/usecase"
)

type UserNotificationHandler struct {
	Handler
}

// SelectAll ...
func (h *UserNotificationHandler) SelectAll(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.UserNotificationParameter{
		Search: "",
		By:     "def.created_date",
		Sort:   "desc",
		UserID: ctx.Query("user_id"),
	}
	uc := usecase.UserNotificationUC{ContractUC: h.ContractUC}
	res, err := uc.SelectAll(c, parameter)

	return h.SendResponse(ctx, res, nil, err, 0)
}

// FindAll ...
func (h *UserNotificationHandler) FindAll(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.UserNotificationParameter{
		Search: ctx.Query("search"),
		Page:   str.StringToInt(ctx.Query("page")),
		Limit:  str.StringToInt(ctx.Query("limit")),
		By:     ctx.Query("by"),
		Sort:   "asc",
	}
	uc := usecase.UserNotificationUC{ContractUC: h.ContractUC}
	res, meta, err := uc.FindAll(c, parameter)

	return h.SendResponse(ctx, res, meta, err, 0)
}

func (h *UserNotificationHandler) FindByID(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.UserNotificationParameter{
		ID: ctx.Params("id"),
	}
	if parameter.ID == "" {
		return h.SendResponse(ctx, nil, nil, helper.InvalidParameter, http.StatusBadRequest)
	}

	uc := usecase.UserNotificationUC{ContractUC: h.ContractUC}
	res, err := uc.FindByID(c, parameter)

	return h.SendResponse(ctx, res, nil, err, 0)
}

// Add ...
func (h *UserNotificationHandler) Add(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	input := new(requests.UserNotificationRequest)
	if err := ctx.BodyParser(input); err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}
	if err := h.Validator.Struct(input); err != nil {
		errMessage := h.ExtractErrorValidationMessages(err.(validator.ValidationErrors))
		return h.SendResponse(ctx, nil, nil, errMessage, http.StatusBadRequest)
	}

	uc := usecase.UserNotificationUC{ContractUC: h.ContractUC}
	res, err := uc.Add(c, input)

	return h.SendResponse(ctx, res, nil, err, 0)
}

// Add ...
func (h *UserNotificationHandler) UpdateStatus(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)
	id := ctx.Params("id")
	if id == "" {
		return h.SendResponse(ctx, nil, nil, helper.InvalidParameter, http.StatusBadRequest)
	}
	input := new(requests.UserNotificationRequest)
	if err := ctx.BodyParser(input); err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}
	if err := h.Validator.Struct(input); err != nil {
		errMessage := h.ExtractErrorValidationMessages(err.(validator.ValidationErrors))
		return h.SendResponse(ctx, nil, nil, errMessage, http.StatusBadRequest)
	}

	uc := usecase.UserNotificationUC{ContractUC: h.ContractUC}
	res, err := uc.UpdateStatus(c, id, input)

	return h.SendResponse(ctx, res, nil, err, 0)
}

func (h *UserNotificationHandler) UpdateAllStatus(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)
	user_id := ctx.Params("user_id")
	if user_id == "" {
		return h.SendResponse(ctx, nil, nil, helper.InvalidParameter, http.StatusBadRequest)
	}
	input := new(requests.UserNotificationRequest)
	if err := ctx.BodyParser(input); err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}
	if err := h.Validator.Struct(input); err != nil {
		errMessage := h.ExtractErrorValidationMessages(err.(validator.ValidationErrors))
		return h.SendResponse(ctx, nil, nil, errMessage, http.StatusBadRequest)
	}

	uc := usecase.UserNotificationUC{ContractUC: h.ContractUC}
	res, err := uc.UpdateAllStatus(c, user_id, input)

	return h.SendResponse(ctx, res, nil, err, 0)
}

func (h *UserNotificationHandler) DeleteStatus(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)
	id := ctx.Params("id")
	if id == "" {
		return h.SendResponse(ctx, nil, nil, helper.InvalidParameter, http.StatusBadRequest)
	}
	input := new(requests.UserNotificationRequest)
	if err := ctx.BodyParser(input); err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}
	if err := h.Validator.Struct(input); err != nil {
		errMessage := h.ExtractErrorValidationMessages(err.(validator.ValidationErrors))
		return h.SendResponse(ctx, nil, nil, errMessage, http.StatusBadRequest)
	}

	uc := usecase.UserNotificationUC{ContractUC: h.ContractUC}
	res, err := uc.DeleteStatus(c, id, input)

	return h.SendResponse(ctx, res, nil, err, 0)
}

func (h *UserNotificationHandler) DeleteAllStatus(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)
	user_id := ctx.Params("user_id")
	if user_id == "" {
		return h.SendResponse(ctx, nil, nil, helper.InvalidParameter, http.StatusBadRequest)
	}
	input := new(requests.UserNotificationRequest)
	if err := ctx.BodyParser(input); err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}
	if err := h.Validator.Struct(input); err != nil {
		errMessage := h.ExtractErrorValidationMessages(err.(validator.ValidationErrors))
		return h.SendResponse(ctx, nil, nil, errMessage, http.StatusBadRequest)
	}

	uc := usecase.UserNotificationUC{ContractUC: h.ContractUC}
	res, err := uc.DeleteAllStatus(c, user_id, input)

	return h.SendResponse(ctx, res, nil, err, 0)
}
