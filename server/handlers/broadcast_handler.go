package handlers

import (
	"context"
	"net/http"

	"nextbasis-service-v-0.1/server/requests"
	"nextbasis-service-v-0.1/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// BroadcastHandler ...
type BroadcastHandler struct {
	Handler
}

// FindAll ...
func (h *BroadcastHandler) Broadcast(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	input := new(requests.BroadcastRequest)
	if err := ctx.BodyParser(input); err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}
	if err := h.Validator.Struct(input); err != nil {
		errMessage := h.ExtractErrorValidationMessages(err.(validator.ValidationErrors))
		return h.SendResponse(ctx, nil, nil, errMessage, http.StatusBadRequest)
	}

	uc := usecase.BroadcastUC{ContractUC: h.ContractUC}
	err := uc.Broadcast(c, input)

	return h.SendResponse(ctx, nil, nil, err, 0)
}
