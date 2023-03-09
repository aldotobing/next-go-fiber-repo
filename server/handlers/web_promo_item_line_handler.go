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

// WebPromoItemLineHandler ...
type WebPromoItemLineHandler struct {
	Handler
}

// SelectAll ...
func (h *WebPromoItemLineHandler) SelectAll(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.WebPromoItemLineParameter{
		ID:          ctx.Query("id"),
		PromoID:     ctx.Query("promo_id"),
		PromoLineID: ctx.Query("promo_line_id"),
		CustomerID:  ctx.Query("customer_id"),
		Search:      ctx.Query("search"),
		By:          ctx.Query("by"),
		Sort:        ctx.Query("sort"),
	}
	uc := usecase.WebPromoItemLineUC{ContractUC: h.ContractUC}
	res, err := uc.SelectAll(c, parameter)

	// if parameter.CustomerID == "" {
	// 	cus_id_err := " : param customer_id is mandatory"
	// 	return h.SendResponse(ctx, nil, nil, helper.InvalidParameter+cus_id_err, http.StatusBadRequest)
	// }

	type StructObject struct {
		ListObject []models.WebPromoItemLine `json:"list_promo_item_line"`
	}

	ObjcetData := new(StructObject)

	if res != nil {
		ObjcetData.ListObject = res
	}

	return h.SendResponse(ctx, ObjcetData, nil, err, 0)
}

// FindAll ...
func (h *WebPromoItemLineHandler) FindAll(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.WebPromoItemLineParameter{
		ID:          ctx.Query("id"),
		PromoID:     ctx.Query("promo_id"),
		PromoLineID: ctx.Query("promo_line_id"),
		CustomerID:  ctx.Query("customer_id"),
		Search:      ctx.Query("search"),
		Page:        str.StringToInt(ctx.Query("page")),
		Limit:       str.StringToInt(ctx.Query("limit")),
		By:          ctx.Query("by"),
		Sort:        ctx.Query("sort"),
	}
	uc := usecase.WebPromoItemLineUC{ContractUC: h.ContractUC}
	res, meta, err := uc.FindAll(c, parameter)

	// if parameter.CustomerID == "" {
	// 	cus_id_err := " : param customer_id is mandatory"
	// 	return h.SendResponse(ctx, nil, nil, helper.InvalidParameter+cus_id_err, http.StatusBadRequest)
	// }

	type StructObject struct {
		ListObject []models.WebPromoItemLine `json:"list_promo_item_line"`
	}

	ObjectData := new(StructObject)

	if res != nil {
		ObjectData.ListObject = res
	}

	return h.SendResponse(ctx, ObjectData, meta, err, 0)
}

// Add ...
func (h *WebPromoItemLineHandler) Add(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	input := new(requests.WebPromoItemLineRequest)
	if err := ctx.BodyParser(input); err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}
	if err := h.Validator.Struct(input); err != nil {
		errMessage := h.ExtractErrorValidationMessages(err.(validator.ValidationErrors))
		return h.SendResponse(ctx, nil, nil, errMessage, http.StatusBadRequest)
	}

	uc := usecase.WebPromoItemLineUC{ContractUC: h.ContractUC}
	res, err := uc.Add(c, input)
	if err != nil {
		return h.SendResponse(ctx, nil, nil, err.Error(), http.StatusBadRequest)
	}

	return h.SendResponse(ctx, res, nil, err, 0)
}

// Delete ...
func (h *WebPromoItemLineHandler) Delete(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	id := ctx.Params("id")
	if id == "" {
		return h.SendResponse(ctx, nil, nil, helper.InvalidParameter, http.StatusBadRequest)
	}

	uc := usecase.WebPromoItemLineUC{ContractUC: h.ContractUC}
	res, err := uc.Delete(c, id)
	if err != nil {
		return h.SendResponse(ctx, nil, nil, err.Error(), http.StatusBadRequest)
	}

	return h.SendResponse(ctx, res, nil, err, 0)
}
