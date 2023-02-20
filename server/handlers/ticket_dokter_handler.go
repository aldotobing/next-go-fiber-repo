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

// TicketDokterHandler ...
type TicketDokterHandler struct {
	Handler
}

// SelectAll ...
func (h *TicketDokterHandler) SelectAll(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.TicketDokterParameter{
		ID:         ctx.Query("ticket_id"),
		CustomerID: ctx.Query("customer_id"),
		DoctorID:   ctx.Query("doctor_id"),
		Status:     ctx.Query("status"),
		Search:     ctx.Query("search"),
		By:         ctx.Query("by"),
		Sort:       ctx.Query("sort"),
	}
	uc := usecase.TicketDokterUC{ContractUC: h.ContractUC}
	res, err := uc.SelectAll(c, parameter)

	// if parameter.CustomerID == "" {
	// 	cus_id_err := " : param customer_id is mandatory"
	// 	return h.SendResponse(ctx, nil, nil, helper.InvalidParameter+cus_id_err, http.StatusBadRequest)
	// }

	type StructObject struct {
		ListObject []models.TicketDokter `json:"list_ticket_dokter"`
	}

	ObjcetData := new(StructObject)

	if res != nil {
		ObjcetData.ListObject = res
	}

	return h.SendResponse(ctx, ObjcetData, nil, err, 0)
}

// FindAll ...
func (h *TicketDokterHandler) FindAll(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.TicketDokterParameter{
		ID:         ctx.Query("ticket_id"),
		CustomerID: ctx.Query("customer_id"),
		DoctorID:   ctx.Query("doctor_id"),
		Status:     ctx.Query("status"),
		Search:     ctx.Query("search"),
		Page:       str.StringToInt(ctx.Query("page")),
		Limit:      str.StringToInt(ctx.Query("limit")),
		By:         ctx.Query("by"),
		Sort:       ctx.Query("sort"),
	}
	uc := usecase.TicketDokterUC{ContractUC: h.ContractUC}
	res, meta, err := uc.FindAll(c, parameter)

	// if parameter.CustomerID == "" {
	// 	cus_id_err := " : param customer_id is mandatory"
	// 	return h.SendResponse(ctx, nil, nil, helper.InvalidParameter+cus_id_err, http.StatusBadRequest)
	// }

	type StructObject struct {
		ListObject []models.TicketDokter `json:"list_ticket_dokter"`
	}

	ObjectData := new(StructObject)

	if res != nil {
		ObjectData.ListObject = res
	}

	return h.SendResponse(ctx, ObjectData, meta, err, 0)
}

// Add ...
func (h *TicketDokterHandler) Add(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	input := new(requests.TicketDokterRequest)
	if err := ctx.BodyParser(input); err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}
	if err := h.Validator.Struct(input); err != nil {
		errMessage := h.ExtractErrorValidationMessages(err.(validator.ValidationErrors))
		return h.SendResponse(ctx, nil, nil, errMessage, http.StatusBadRequest)
	}

	uc := usecase.TicketDokterUC{ContractUC: h.ContractUC}
	res, err := uc.Add(c, input)

	return h.SendResponse(ctx, res, nil, err, 0)
}

// Edit ...
func (h *TicketDokterHandler) Edit(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	id := ctx.Params("ticket_id")
	if id == "" {
		return h.SendResponse(ctx, nil, nil, helper.InvalidParameter, http.StatusBadRequest)
	}

	input := new(requests.TicketDokterRequest)
	if err := ctx.BodyParser(input); err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}
	if err := h.Validator.Struct(input); err != nil {
		errMessage := h.ExtractErrorValidationMessages(err.(validator.ValidationErrors))
		return h.SendResponse(ctx, nil, nil, errMessage, http.StatusBadRequest)
	}

	uc := usecase.TicketDokterUC{ContractUC: h.ContractUC}
	res, err := uc.Edit(c, id, input)

	return h.SendResponse(ctx, res, nil, err, 0)
}
