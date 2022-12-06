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
)

// NewsHandler ...
type NewsHandler struct {
	Handler
}

// SelectAll ...
func (h *NewsHandler) SelectAll(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.NewsParameter{
		Title:     ctx.Query("title"),
		StartDate: ctx.Query("start_date"),
		EndDate:   ctx.Query("end_date"),
		UserId:    ctx.Query("user_id"),
		Search:    ctx.Query("search"),
		By:        ctx.Query("by"),
		Sort:      ctx.Query("sort"),
	}
	uc := usecase.NewsUC{ContractUC: h.ContractUC}
	res, err := uc.SelectAll(c, parameter)

	type InvoceHeaderObjcet struct {
		ListObjcet []models.News `json:"list_news"`
	}

	ObjcetData := new(InvoceHeaderObjcet)

	if res != nil {
		ObjcetData.ListObjcet = res
	}

	return h.SendResponse(ctx, ObjcetData, nil, err, 0)
}

// FindAll ...
func (h *NewsHandler) FindAll(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.NewsParameter{
		Title:  ctx.Query("title"),
		Search: ctx.Query("search"),
		Page:   str.StringToInt(ctx.Query("page")),
		Limit:  str.StringToInt(ctx.Query("limit")),
		By:     ctx.Query("by"),
		Sort:   ctx.Query("sort"),
	}
	uc := usecase.NewsUC{ContractUC: h.ContractUC}
	res, meta, err := uc.FindAll(c, parameter)

	type InvoceHeaderObjcet struct {
		ListObjcet []models.News `json:"list_news"`
	}

	ObjcetData := new(InvoceHeaderObjcet)

	if res != nil {
		ObjcetData.ListObjcet = res
	}

	return h.SendResponse(ctx, ObjcetData, meta, err, 0)
}

// Add ...
func (h *NewsHandler) Add(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	input := new(requests.NewsRequest)
	if err := ctx.BodyParser(input); err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}
	if err := h.Validator.Struct(input); err != nil {
		errMessage := h.ExtractErrorValidationMessages(err.(validator.ValidationErrors))
		return h.SendResponse(ctx, nil, nil, errMessage, http.StatusBadRequest)
	}

	uc := usecase.NewsUC{ContractUC: h.ContractUC}
	res, err := uc.Add(c, input)

	return h.SendResponse(ctx, res, nil, err, 0)
}
