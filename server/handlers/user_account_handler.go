package handlers

import (
	"context"
	"net/http"

	"nextbasis-service-v-0.1/server/requests"
	"nextbasis-service-v-0.1/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type UserAccountHandler struct {
	Handler
}

func (h *UserAccountHandler) Login(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	input := new(requests.UserAccountLoginRequest)
	if err := ctx.BodyParser(input); err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}
	if err := h.Validator.Struct(input); err != nil {
		errMessage := h.ExtractErrorValidationMessages(err.(validator.ValidationErrors))
		return h.SendResponse(ctx, nil, nil, errMessage, http.StatusBadRequest)
	}

	uc := usecase.UserAccountUC{ContractUC: h.ContractUC}
	res, err := uc.Login(c, input)

	return h.SendResponse(ctx, res, nil, err, 0)
}

func (h *UserAccountHandler) ResendOtp(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)
	id := ctx.Query("user_id")
	input := new(requests.UserOtpRequest)
	if err := ctx.BodyParser(input); err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}
	if err := h.Validator.Struct(input); err != nil {
		errMessage := h.ExtractErrorValidationMessages(err.(validator.ValidationErrors))
		return h.SendResponse(ctx, nil, nil, errMessage, http.StatusBadRequest)
	}

	uc := usecase.UserAccountUC{ContractUC: h.ContractUC}
	res, err := uc.ResendOtp(c, id, input)

	return h.SendResponse(ctx, res, nil, err, 0)
}

func (h *UserAccountHandler) SubmitOtp(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)
	id := ctx.Query("user_id")
	println(id)
	input := new(requests.UserOtpSubmit)
	if err := ctx.BodyParser(input); err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}
	if err := h.Validator.Struct(input); err != nil {
		errMessage := h.ExtractErrorValidationMessages(err.(validator.ValidationErrors))
		return h.SendResponse(ctx, nil, nil, errMessage, http.StatusBadRequest)
	}

	uc := usecase.UserAccountUC{ContractUC: h.ContractUC}
	res, err := uc.SubmitOtpUser(c, id, input)

	return h.SendResponse(ctx, res, nil, err, 0)
}
