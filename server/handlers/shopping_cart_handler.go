package handlers

import (
	"context"
	"fmt"
	"net/http"

	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/helper"
	"nextbasis-service-v-0.1/pkg/str"
	"nextbasis-service-v-0.1/server/requests"
	"nextbasis-service-v-0.1/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// ShoppingCartHandler ...
type ShoppingCartHandler struct {
	Handler
}

// SelectAll ...
func (h *ShoppingCartHandler) SelectAll(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	cusid := ctx.Params("customer_id")
	if cusid == "" {
		return h.SendResponse(ctx, nil, nil, helper.InvalidParameter, http.StatusBadRequest)
	}

	parameter := models.ShoppingCartParameter{
		CustomerID: ctx.Params("customer_id"),
		Search:     ctx.Query("search"),
		By:         ctx.Query("by"),
		Sort:       ctx.Query("sort"),
	}
	uc := usecase.ShoppingCartUC{ContractUC: h.ContractUC}
	res, err := uc.SelectAll(c, parameter)

	type StructObject struct {
		ListObjcet []models.ShoppingCart `json:"list_sopping_cart"`
	}

	ObjcetData := new(StructObject)

	if res != nil {
		ObjcetData.ListObjcet = res
	}

	return h.SendResponse(ctx, ObjcetData, nil, err, 0)
}

// FindAll ...
func (h *ShoppingCartHandler) FindAll(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.ShoppingCartParameter{
		CustomerID: ctx.Params("customer_id"),
		Search:     ctx.Query("search"),
		Page:       str.StringToInt(ctx.Query("page")),
		Limit:      str.StringToInt(ctx.Query("limit")),
		By:         ctx.Query("by"),
		Sort:       ctx.Query("sort"),
	}
	uc := usecase.ShoppingCartUC{ContractUC: h.ContractUC}
	res, meta, err := uc.FindAll(c, parameter)

	type StructObject struct {
		ListObjcet []models.ShoppingCart `json:"list_sopping_cart"`
	}

	ObjcetData := new(StructObject)

	if res != nil {
		ObjcetData.ListObjcet = res
	}

	return h.SendResponse(ctx, ObjcetData, meta, err, 0)
}

// FindByID ...
func (h *ShoppingCartHandler) FindByID(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.ShoppingCartParameter{
		ID: ctx.Params("id"),
	}
	if parameter.ID == "" {
		return h.SendResponse(ctx, nil, nil, helper.InvalidParameter, http.StatusBadRequest)
	}

	uc := usecase.ShoppingCartUC{ContractUC: h.ContractUC}
	res, err := uc.FindByID(c, parameter)

	return h.SendResponse(ctx, res, nil, err, 0)
}

// Add ...
func (h *ShoppingCartHandler) Add(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	input := new(requests.ShoppingCartRequest)
	if err := ctx.BodyParser(input); err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}
	if err := h.Validator.Struct(input); err != nil {
		errMessage := h.ExtractErrorValidationMessages(err.(validator.ValidationErrors))
		return h.SendResponse(ctx, nil, nil, errMessage, http.StatusBadRequest)
	}

	uc := usecase.ShoppingCartUC{ContractUC: h.ContractUC}
	res, err := uc.Add(c, input)

	return h.SendResponse(ctx, res, nil, err, 0)
}

// Edit ...
func (h *ShoppingCartHandler) MultipleEdit(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	listInput := new([]requests.ShoppingCartRequest)
	if err := ctx.BodyParser(listInput); err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}

	if err := h.Validator.Struct(listInput); err != nil {
		errMessage := h.ExtractErrorValidationMessages(err.(validator.ValidationErrors))
		return h.SendResponse(ctx, nil, nil, errMessage, http.StatusBadRequest)
	}

	uc := usecase.ShoppingCartUC{ContractUC: h.ContractUC}
	res, err := uc.MultipleEdit(c, listInput)

	return h.SendResponse(ctx, res, nil, err, 0)
}

// MultipleEditByCartID ...
func (h *ShoppingCartHandler) MultipleEditByCartID(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	cartID := ctx.Params("cart_id")
	if cartID == "" {
		return h.SendResponse(ctx, nil, nil, helper.InvalidParameter, http.StatusBadRequest)
	}

	listInput := new([]requests.ShoppingCartRequest)
	if err := ctx.BodyParser(listInput); err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}

	for _, input := range *listInput {
		if err := h.Validator.Struct(input); err != nil {
			errMessage := h.ExtractErrorValidationMessages(err.(validator.ValidationErrors))
			return h.SendResponse(ctx, nil, nil, errMessage, http.StatusBadRequest)
		}
	}

	uc := usecase.ShoppingCartUC{ContractUC: h.ContractUC}
	res, err := uc.MultipleEditByCartID(c, listInput)

	return h.SendResponse(ctx, res, nil, err, 0)
}

func (h *ShoppingCartHandler) MultipleDelete(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	id := ctx.Params("customer_id")
	if id == "" {
		return h.SendResponse(ctx, nil, nil, helper.InvalidParameter, http.StatusBadRequest)
	}

	listInput := new([]requests.ShoppingCartRequest)
	if err := ctx.BodyParser(listInput); err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}

	for _, input := range *listInput {
		if err := h.Validator.Struct(input); err != nil {
			errMessage := h.ExtractErrorValidationMessages(err.(validator.ValidationErrors))
			return h.SendResponse(ctx, nil, nil, errMessage, http.StatusBadRequest)
		}
	}

	uc := usecase.ShoppingCartUC{ContractUC: h.ContractUC}
	res, err := uc.MultipleDelete(c, listInput)

	return h.SendResponse(ctx, res, nil, err, 0)
}

// Delete ...
func (h *ShoppingCartHandler) Delete(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	id := ctx.Params("id")
	if id == "" {
		return h.SendResponse(ctx, nil, nil, helper.InvalidParameter, http.StatusBadRequest)
	}

	uc := usecase.ShoppingCartUC{ContractUC: h.ContractUC}
	res, err := uc.Delete(c, id)

	return h.SendResponse(ctx, res, nil, err, 0)
}

func (h *ShoppingCartHandler) CheckOut(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	input := new(requests.CustomerOrderHeaderRequest)
	if err := ctx.BodyParser(input); err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}
	if err := h.Validator.Struct(input); err != nil {
		errMessage := h.ExtractErrorValidationMessages(err.(validator.ValidationErrors))
		return h.SendResponse(ctx, nil, nil, errMessage, http.StatusBadRequest)
	}

	uc := usecase.CustomerOrderHeaderUC{ContractUC: h.ContractUC}
	res, err := uc.CheckOut(c, input)

	return h.SendResponse(ctx, res, nil, err, 0)
}

func (h *ShoppingCartHandler) SelecGroupedtAll(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	cusid := ctx.Params("customer_id")
	if cusid == "" {
		return h.SendResponse(ctx, nil, nil, helper.InvalidParameter, http.StatusBadRequest)
	}

	parameter := models.ShoppingCartParameter{
		CustomerID: ctx.Params("customer_id"),
		Search:     ctx.Query("search"),
		By:         ctx.Query("by"),
		Sort:       ctx.Query("sort"),
	}
	uc := usecase.ShoppingCartUC{ContractUC: h.ContractUC}
	res, err := uc.SelectAllForGroup(c, parameter)

	type StructObject struct {
		ListObjcet []models.GroupedShoppingCart `json:"list_shopping_cart"`
	}

	ObjcetData := new(StructObject)

	if res != nil {
		ObjcetData.ListObjcet = res
	}

	if ObjcetData.ListObjcet != nil && len(ObjcetData.ListObjcet) > 0 {
		for i, object := range ObjcetData.ListObjcet {
			lineparameter := models.ShoppingCartParameter{
				CustomerID:     ctx.Params("customer_id"),
				ItemCategoryID: object.CategoryID,
				Search:         ctx.Query("search"),
				By:             ctx.Query("by"),
				Sort:           ctx.Query("sort"),
			}
			lineres, err := uc.SelectAll(c, lineparameter)
			if err != nil {
				fmt.Println("error bind")
			}
			ObjcetData.ListObjcet[i].ListShoppingChart = lineres

		}
	}

	return h.SendResponse(ctx, ObjcetData, nil, err, 0)
}

// SelectAll ...
func (h *ShoppingCartHandler) SelectAllBonus(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)
	parameter := models.ShoppingCartParameter{
		ListID: ctx.Query("chart_id_list"),
		Search: ctx.Query("search"),
		By:     ctx.Query("by"),
		Sort:   ctx.Query("sort"),
	}
	uc := usecase.ShoppingCartUC{ContractUC: h.ContractUC}
	res, err := uc.SelectAllBonus(c, parameter)

	if err != nil {
		fmt.Println("error", err.Error())
	}

	type StructObject struct {
		ListObjcet []models.ShoppingCartItemBonus `json:"list_item_bonus"`
	}

	ObjcetData := new(StructObject)

	if res != nil {
		ObjcetData.ListObjcet = res
	}

	return h.SendResponse(ctx, ObjcetData, nil, err, 0)
}
