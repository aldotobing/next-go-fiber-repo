package handlers

import (
	"context"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/helper"
	"nextbasis-service-v-0.1/server/requests"
	"nextbasis-service-v-0.1/usecase"
)

// TicketDokterChatHandler ...
type TicketDokterChatHandler struct {
	Handler
}

// FindByTicketDocterID ...
func (h *TicketDokterChatHandler) FindByTicketDocterID(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	ticketDokterID := ctx.Params("ticket_dokter_id")
	if ticketDokterID == "" {
		return h.SendResponse(ctx, nil, nil, helper.InvalidParameter, http.StatusBadRequest)
	}

	parameter := models.TicketDokterChatParameter{
		By:   ctx.Query("by"),
		Sort: ctx.Query("sort"),
	}
	uc := usecase.TicketDokterChatUC{ContractUC: h.ContractUC}
	res, err := uc.FindByTicketDocterID(c, parameter, ticketDokterID)
	if err != nil {
		return h.SendResponse(ctx, nil, nil, helper.InvalidParameter, http.StatusBadRequest)
	}

	return h.SendResponse(ctx, res, nil, err, 0)
}

// Add ...
func (h *TicketDokterChatHandler) Add(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	input := new(requests.TicketDokterChatRequest)
	if err := ctx.BodyParser(input); err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}
	if err := h.Validator.Struct(input); err != nil {
		errMessage := h.ExtractErrorValidationMessages(err.(validator.ValidationErrors))
		return h.SendResponse(ctx, nil, nil, errMessage, http.StatusBadRequest)
	}

	uc := usecase.TicketDokterChatUC{ContractUC: h.ContractUC}
	res, err := uc.Add(c, input)

	return h.SendResponse(ctx, res, nil, err, 0)
}

// Delete ...
func (h *TicketDokterChatHandler) Delete(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	id := ctx.Params("id")
	if id == "" {
		return h.SendResponse(ctx, nil, nil, helper.InvalidParameter, http.StatusBadRequest)
	}

	uc := usecase.TicketDokterChatUC{ContractUC: h.ContractUC}
	_, err := uc.Delete(c, id)
	if err != nil {
		h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}

	return h.SendResponse(ctx, nil, nil, err, 0)
}
