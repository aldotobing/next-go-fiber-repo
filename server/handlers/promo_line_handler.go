package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/helper"
	"nextbasis-service-v-0.1/pkg/str"
	"nextbasis-service-v-0.1/server/requests"
	"nextbasis-service-v-0.1/usecase"
)

// PromoLineHandler ...
type PromoLineHandler struct {
	Handler
}

// SelectAll ...
func (h *PromoLineHandler) SelectAll(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.PromoLineParameter{
		ID:      ctx.Query("id"),
		PromoID: ctx.Query("promo_id"),
		Search:  ctx.Query("search"),
		By:      ctx.Query("by"),
		Sort:    ctx.Query("sort"),
	}
	uc := usecase.PromoLineUC{ContractUC: h.ContractUC}
	res, err := uc.SelectAll(c, parameter)

	type StructObject struct {
		ListObject []models.PromoLine `json:"list_promo_line"`
	}

	ObjcetData := new(StructObject)

	if res != nil {
		ObjcetData.ListObject = res
	}

	return h.SendResponse(ctx, ObjcetData, nil, err, 0)
}

// FindAll ...
func (h *PromoLineHandler) FindAll(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.PromoLineParameter{
		ID:      ctx.Query("id"),
		PromoID: ctx.Query("promo_id"),
		Search:  ctx.Query("search"),
		Page:    str.StringToInt(ctx.Query("page")),
		Limit:   str.StringToInt(ctx.Query("limit")),
		By:      ctx.Query("by"),
		Sort:    ctx.Query("sort"),
	}
	uc := usecase.PromoLineUC{ContractUC: h.ContractUC}
	res, meta, err := uc.FindAll(c, parameter)

	type StructObject struct {
		ListObject []models.PromoLine `json:"list_promo_line"`
	}

	ObjectData := new(StructObject)

	if res != nil {
		ObjectData.ListObject = res
	}

	return h.SendResponse(ctx, ObjectData, meta, err, 0)
}

// Add ...
func (h *PromoLineHandler) Add(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	// input := new(requests.PromoLineRequest)
	// err := json.Unmarshal([]byte(ctx.FormValue("form_data", "")), input)
	// if err := ctx.BodyParser(input); err != nil {
	// 	return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	// }
	// if err := h.Validator.Struct(input); err != nil {
	// 	errMessage := h.ExtractErrorValidationMessages(err.(validator.ValidationErrors))
	// 	return h.SendResponse(ctx, nil, nil, errMessage, http.StatusBadRequest)
	// }

	input := new(requests.PromoLineRequest)
	if err := ctx.BodyParser(input); err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}
	if err := h.Validator.Struct(input); err != nil {
		errMessage := h.ExtractErrorValidationMessages(err.(validator.ValidationErrors))
		return h.SendResponse(ctx, nil, nil, errMessage, http.StatusBadRequest)
	}

	uc := usecase.PromoLineUC{ContractUC: h.ContractUC}
	res, err := uc.Add(c, input)

	return h.SendResponse(ctx, res, nil, err, 0)
}

// Edit ...
func (h *PromoLineHandler) Edit(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	id := ctx.Params("id")
	if id == "" {
		return h.SendResponse(ctx, nil, nil, helper.InvalidParameter, http.StatusBadRequest)
	}

	input := new(requests.PromoLineRequest)
	if err := ctx.BodyParser(input); err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}
	if err := h.Validator.Struct(input); err != nil {
		errMessage := h.ExtractErrorValidationMessages(err.(validator.ValidationErrors))
		return h.SendResponse(ctx, nil, nil, errMessage, http.StatusBadRequest)
	}

	uc := usecase.PromoLineUC{ContractUC: h.ContractUC}
	res, err := uc.Edit(c, id, input)

	return h.SendResponse(ctx, res, nil, err, 0)
}

// Delete ...
func (h *PromoLineHandler) Delete(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	id := ctx.Params("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return h.SendResponse(ctx, nil, nil, helper.InvalidParameter, http.StatusBadRequest)
	}

	uc := usecase.PromoLineUC{ContractUC: h.ContractUC}
	res, err := uc.Delete(c, idInt)

	return h.SendResponse(ctx, res, nil, err, 0)
}
