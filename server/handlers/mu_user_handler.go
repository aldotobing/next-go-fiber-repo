package handlers

import (
	"context"
	"net/http"

	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/helper"
	"nextbasis-service-v-0.1/pkg/str"
	"nextbasis-service-v-0.1/usecase"

	"github.com/gofiber/fiber/v2"
)

type MuUserHandler struct {
	Handler
}

func (h *MuUserHandler) FindAll(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.MuUserParameter{
		Search:      ctx.Query("search"),
		Page:        str.StringToInt(ctx.Query("page")),
		Limit:       str.StringToInt(ctx.Query("limit")),
		By:          ctx.Query("by"),
		Sort:        ctx.Query("sort"),
		RoleGroupID: ctx.Query("role_group_id"),
	}
	uc := usecase.MuUserUC{ContractUC: h.ContractUC}
	res, meta, err := uc.FindAll(c, parameter)

	return h.SendResponse(ctx, res, meta, err, 0)
}

func (h *MuUserHandler) FindByID(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.MuUserParameter{
		ID: ctx.Params("id"),
	}
	if parameter.ID == "" {
		return h.SendResponse(ctx, nil, nil, helper.InvalidParameter, http.StatusBadRequest)
	}

	uc := usecase.MuUserUC{ContractUC: h.ContractUC}
	res, err := uc.FindByID(c, parameter)

	return h.SendResponse(ctx, res, nil, err, 0)
}

// func (h *MuUserHandler) Add(ctx *fiber.Ctx) error {
// 	c := ctx.Locals("ctx").(context.Context)
// 	input := new(requests.MuUserRequest)
// 	err := json.Unmarshal([]byte(ctx.FormValue("input_data")), input)
// 	if err := ctx.BodyParser(input); err != nil {
// 		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
// 	}
// 	if err := h.Validator.Struct(input); err != nil {
// 		errMessage := h.ExtractErrorValidationMessages(err.(validator.ValidationErrors))
// 		return h.SendResponse(ctx, nil, nil, errMessage, http.StatusBadRequest)
// 	}
// 	uc := usecase.MuUserUC{ContractUC: h.ContractUC}
// 	res, err := uc.Add(c, input)
// 	return h.SendResponse(ctx, res, nil, err, 0)
// }

// func (h *MuUserHandler) Edit(ctx *fiber.Ctx) error {
// 	c := ctx.Locals("ctx").(context.Context)

// 	id := ctx.Params("id")
// 	if id == "" {
// 		return h.SendResponse(ctx, nil, nil, helper.InvalidParameter, http.StatusBadRequest)
// 	}

// 	input := new(requests.MuUserRequest)
// 	err := json.Unmarshal([]byte(ctx.FormValue("form_data")), input)
// 	if err := ctx.BodyParser(input); err != nil {
// 		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
// 	}

// 	if err := h.Validator.Struct(input); err != nil {
// 		errMessage := h.ExtractErrorValidationMessages(err.(validator.ValidationErrors))
// 		return h.SendResponse(ctx, nil, nil, errMessage, http.StatusBadRequest)
// 	}

// 	imgKtp, _ := ctx.FormFile("img_ktp")

// 	imgProfile, _ := ctx.FormFile("img_profile")

// 	uc := usecase.MuUserUC{ContractUC: h.ContractUC}
// 	res, err := uc.Edit(c, id, input, imgKtp, imgProfile)

// 	return h.SendResponse(ctx, res, nil, err, 0)
// }

// func (h *MuUserHandler) Delete(ctx *fiber.Ctx) error {

// 	c := ctx.Locals("ctx").(context.Context)

// 	id := ctx.Params("id")
// 	DeletedBy := ctx.Params("user_id")
// 	if id == "" {
// 		return h.SendResponse(ctx, nil, nil, helper.InvalidParameter, http.StatusBadRequest)
// 	}

// 	uc := usecase.MuUserUC{ContractUC: h.ContractUC}
// 	res, err := uc.Delete(c, id, DeletedBy)

// 	return h.SendResponse(ctx, res, nil, err, 0)
// }

// func (h *MuUserHandler) UpdatePassword(ctx *fiber.Ctx) error {
// 	c := ctx.Locals("ctx").(context.Context)

// 	id := ctx.Params("id")
// 	if id == "" {
// 		return h.SendResponse(ctx, nil, nil, helper.InvalidParameter, http.StatusBadRequest)
// 	}

// 	input := new(requests.MuUserRequest)
// 	if err := ctx.BodyParser(input); err != nil {
// 		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
// 	}
// 	if err := h.Validator.Struct(input); err != nil {
// 		errMessage := h.ExtractErrorValidationMessages(err.(validator.ValidationErrors))
// 		return h.SendResponse(ctx, nil, nil, errMessage, http.StatusBadRequest)
// 	}

// 	uc := usecase.MuUserUC{ContractUC: h.ContractUC}
// 	res, err := uc.UpdatePassword(c, id, input)

// 	return h.SendResponse(ctx, res, nil, err, 0)
// }
