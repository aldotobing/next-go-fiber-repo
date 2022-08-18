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

// OtpHandler ...
type OtpHandler struct {
	Handler
}

// ResendOtpRequest ...
func (h *OtpHandler) ResendOtpRequest(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	id := fmt.Sprintf("%v", ctx.Locals("claims"))

	input := new(requests.UserOtpRequest)
	if err := ctx.BodyParser(input); err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}
	if err := h.Validator.Struct(input); err != nil {
		errMessage := h.ExtractErrorValidationMessages(err.(validator.ValidationErrors))
		return h.SendResponse(ctx, nil, nil, errMessage, http.StatusBadRequest)
	}

	uc := usecase.OtpUC{ContractUC: h.ContractUC}
	res, err := uc.OtpRequest(c, id, input)

	return h.SendResponse(ctx, res, nil, err, 0)
}
