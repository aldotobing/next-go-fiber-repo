package handlers

import (
	"context"
	"fmt"
	"net/http"

	"nextbasis-service-v-0.1/server/requests"
	"nextbasis-service-v-0.1/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// AuthHandler ...
type AuthHandler struct {
	Handler
}

// Register ...
func (h *AuthHandler) Register(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	input := new(requests.RegisterRequest)
	if err := ctx.BodyParser(input); err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}
	if err := h.Validator.Struct(input); err != nil {
		errMessage := h.ExtractErrorValidationMessages(err.(validator.ValidationErrors))
		return h.SendResponse(ctx, nil, nil, errMessage, http.StatusBadRequest)
	}

	uc := usecase.AuthUC{ContractUC: h.ContractUC}
	res, err := uc.Register(c, input)
	if err != nil {
		return h.SendResponse(ctx, nil, nil, err.Error(), http.StatusBadRequest)
	}

	return h.SendResponse(ctx, res, nil, err, 0)
}

// SubmitOtpRegister ...
func (h *AuthHandler) SubmitOtpRegister(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	id := fmt.Sprintf("%v", ctx.Locals("claims"))

	input := new(requests.UserOtpSubmit)
	if err := ctx.BodyParser(input); err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}
	if err := h.Validator.Struct(input); err != nil {
		errMessage := h.ExtractErrorValidationMessages(err.(validator.ValidationErrors))
		return h.SendResponse(ctx, nil, nil, errMessage, http.StatusBadRequest)
	}
	uc := usecase.AuthUC{ContractUC: h.ContractUC}
	res, err := uc.SubmitOtpRegister(c, id, input)
	if err != nil {
		return h.SendResponse(ctx, nil, nil, err.Error(), http.StatusBadRequest)
	}
	return h.SendResponse(ctx, res, nil, err, 0)
}

// VerifyMail ...
func (h *AuthHandler) VerifyMail(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	input := new(requests.VerifyMailRequest)
	if err := ctx.BodyParser(input); err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}
	if err := h.Validator.Struct(input); err != nil {
		errMessage := h.ExtractErrorValidationMessages(err.(validator.ValidationErrors))
		return h.SendResponse(ctx, nil, nil, errMessage, http.StatusBadRequest)
	}

	uc := usecase.AuthUC{ContractUC: h.ContractUC}
	res, err := uc.ReqVerifyMail(c, input)

	return h.SendResponse(ctx, res, nil, err, 0)
}

// VerifyUser ...
func (h *AuthHandler) VerifyUserMail(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	input := new(requests.VerifyUserKeyRequest)
	if err := ctx.BodyParser(input); err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}
	if err := h.Validator.Struct(input); err != nil {
		errMessage := h.ExtractErrorValidationMessages(err.(validator.ValidationErrors))
		return h.SendResponse(ctx, nil, nil, errMessage, http.StatusBadRequest)
	}

	uc := usecase.AuthUC{ContractUC: h.ContractUC}
	res, err := uc.VerifyMail(c, input.Key)

	return h.SendResponse(ctx, res, nil, err, 0)
}
