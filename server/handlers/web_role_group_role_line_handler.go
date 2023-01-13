package handlers

import (
	"context"
	"net/http"

	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/helper"
	"nextbasis-service-v-0.1/pkg/str"
	"nextbasis-service-v-0.1/server/requests"
	"nextbasis-service-v-0.1/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// WebRoleGroupRoleLineHandler ...
type WebRoleGroupRoleLineHandler struct {
	Handler
}

// SelectAll ...
func (h *WebRoleGroupRoleLineHandler) SelectAll(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.WebRoleGroupRoleLineParameter{
		RoleGroupID: ctx.Query("role_group_id"),
		Search:      ctx.Query("search"),
		By:          ctx.Query("by"),
		Sort:        ctx.Query("sort"),
	}
	uc := usecase.WebRoleGroupRoleLineUC{ContractUC: h.ContractUC}
	res, err := uc.SelectAll(c, parameter)

	type StructObject struct {
		ListObjcet []models.WebRoleGroupRoleLine `json:"list_role_group_role_line"`
	}

	ObjcetData := new(StructObject)

	if res != nil {
		ObjcetData.ListObjcet = res
	}

	return h.SendResponse(ctx, ObjcetData, nil, err, 0)
}

// FindAll ...
func (h *WebRoleGroupRoleLineHandler) FindAll(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.WebRoleGroupRoleLineParameter{
		Search:      ctx.Query("search"),
		RoleGroupID: ctx.Query("role_group_id"),
		Page:        str.StringToInt(ctx.Query("page")),
		Limit:       str.StringToInt(ctx.Query("limit")),
		By:          ctx.Query("by"),
		Sort:        ctx.Query("sort"),
	}
	uc := usecase.WebRoleGroupRoleLineUC{ContractUC: h.ContractUC}
	res, meta, err := uc.FindAll(c, parameter)

	type StructObject struct {
		ListObjcet []models.WebRoleGroupRoleLine `json:"list_role_group_role_line"`
	}

	ObjcetData := new(StructObject)

	if res != nil {
		ObjcetData.ListObjcet = res
	}

	return h.SendResponse(ctx, ObjcetData, meta, err, 0)
}

// FindByID ...
func (h *WebRoleGroupRoleLineHandler) FindByID(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.WebRoleGroupRoleLineParameter{
		ID: ctx.Params("id"),
	}
	if parameter.ID == "" {
		return h.SendResponse(ctx, nil, nil, helper.InvalidParameter, http.StatusBadRequest)
	}

	uc := usecase.WebRoleGroupRoleLineUC{ContractUC: h.ContractUC}
	res, err := uc.FindByID(c, parameter)

	return h.SendResponse(ctx, res, nil, err, 0)
}

// Add ...
func (h *WebRoleGroupRoleLineHandler) Add(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	input := new(requests.WebRoleGroupRoleLineRequest)
	if err := ctx.BodyParser(input); err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}
	if err := h.Validator.Struct(input); err != nil {
		errMessage := h.ExtractErrorValidationMessages(err.(validator.ValidationErrors))
		return h.SendResponse(ctx, nil, nil, errMessage, http.StatusBadRequest)
	}

	uc := usecase.WebRoleGroupRoleLineUC{ContractUC: h.ContractUC}
	res, err := uc.Add(c, input)

	return h.SendResponse(ctx, res, nil, err, 0)
}

// Delete ...
func (h *WebRoleGroupRoleLineHandler) Delete(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	id := ctx.Params("id")
	if id == "" {
		return h.SendResponse(ctx, nil, nil, helper.InvalidParameter, http.StatusBadRequest)
	}

	uc := usecase.WebRoleGroupRoleLineUC{ContractUC: h.ContractUC}
	res, err := uc.Delete(c, id)

	return h.SendResponse(ctx, res, nil, err, 0)
}
