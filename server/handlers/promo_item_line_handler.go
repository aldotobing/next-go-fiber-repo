package handlers

import (
	"context"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/str"
	"nextbasis-service-v-0.1/server/requests"
	"nextbasis-service-v-0.1/usecase"
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

// PromoItemLineHandler ...
type PromoItemLineHandler struct {
	Handler
}

// SelectAll ...
func (h *PromoItemLineHandler) SelectAll(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.PromoItemLineParameter{
		ID:          ctx.Query("id"),
		PromoID:     ctx.Query("promo_id"),
		PromoLineID: ctx.Query("promo_line_id"),
		CustomerID:  ctx.Query("customer_id"),
		Search:      ctx.Query("search"),
		By:          ctx.Query("by"),
		Sort:        ctx.Query("sort"),
	}
	uc := usecase.PromoItemLineUC{ContractUC: h.ContractUC}
	res, err := uc.SelectAll(c, parameter)

	// if parameter.CustomerID == "" {
	// 	cus_id_err := " : param customer_id is mandatory"
	// 	return h.SendResponse(ctx, nil, nil, helper.InvalidParameter+cus_id_err, http.StatusBadRequest)
	// }

	type StructObject struct {
		ListObject []viewmodel.PromoItemLineVM `json:"list_promo_item_line"`
	}

	ObjcetData := new(StructObject)

	if res != nil {
		ObjcetData.ListObject = res
	}

	return h.SendResponse(ctx, ObjcetData, nil, err, 0)
}

// FindAll ...
func (h *PromoItemLineHandler) FindAll(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.PromoItemLineParameter{
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
	uc := usecase.PromoItemLineUC{ContractUC: h.ContractUC}
	res, meta, err := uc.FindAll(c, parameter)

	// if parameter.CustomerID == "" {
	// 	cus_id_err := " : param customer_id is mandatory"
	// 	return h.SendResponse(ctx, nil, nil, helper.InvalidParameter+cus_id_err, http.StatusBadRequest)
	// }

	type StructObject struct {
		ListObject []models.PromoItemLine `json:"list_promo_item_line"`
	}

	ObjectData := new(StructObject)

	if res != nil {
		ObjectData.ListObject = res
	}

	return h.SendResponse(ctx, ObjectData, meta, err, 0)
}

// Add ...
func (h *PromoItemLineHandler) Add(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	// input := new(requests.PromoItemLineRequest)
	// err := json.Unmarshal([]byte(ctx.FormValue("form_data", "")), input)
	// if err := ctx.BodyParser(input); err != nil {
	// 	return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	// }
	// if err := h.Validator.Struct(input); err != nil {
	// 	errMessage := h.ExtractErrorValidationMessages(err.(validator.ValidationErrors))
	// 	return h.SendResponse(ctx, nil, nil, errMessage, http.StatusBadRequest)
	// }

	input := new(requests.PromoItemLineRequest)
	if err := ctx.BodyParser(input); err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}
	if err := h.Validator.Struct(input); err != nil {
		errMessage := h.ExtractErrorValidationMessages(err.(validator.ValidationErrors))
		return h.SendResponse(ctx, nil, nil, errMessage, http.StatusBadRequest)
	}

	uc := usecase.PromoItemLineUC{ContractUC: h.ContractUC}
	res, err := uc.Add(c, input)

	return h.SendResponse(ctx, res, nil, err, 0)
}
