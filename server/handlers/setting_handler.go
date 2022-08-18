package handlers

import (
	"context"
	"net/http"
	"time"

	"nextbasis-service-v-0.1/helper"
	"nextbasis-service-v-0.1/pkg/str"
	"nextbasis-service-v-0.1/server/requests"
	"nextbasis-service-v-0.1/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// SettingHandler ...
type SettingHandler struct {
	Handler
}

// SelectAll ...
func (h *SettingHandler) SelectAll(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(context.Background(), time.Duration(str.StringToInt(h.ContractUC.EnvConfig["APP_TIMEOUT"]))*time.Second)
	c = context.WithValue(c, "requestid", ctx.Locals("requestid"))
	c = context.WithValue(c, "user_id", ctx.Locals("user_id"))
	defer cancel()

	settingUC := usecase.SettingUC{ContractUC: h.ContractUC}
	res, err := settingUC.SelectAll(c)

	return h.SendResponse(ctx, res, nil, err, 0)
}

// FindByID ...
func (h *SettingHandler) FindByID(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(context.Background(), time.Duration(str.StringToInt(h.ContractUC.EnvConfig["APP_TIMEOUT"]))*time.Second)
	c = context.WithValue(c, "requestid", ctx.Locals("requestid"))
	c = context.WithValue(c, "user_id", ctx.Locals("user_id"))
	defer cancel()

	id := ctx.Params("id")
	if id == "" {
		return h.SendResponse(ctx, nil, nil, helper.InvalidParameter, http.StatusBadRequest)
	}

	settingUC := usecase.SettingUC{ContractUC: h.ContractUC}
	res, err := settingUC.FindByID(c, id)

	return h.SendResponse(ctx, res, nil, err, 0)
}

// Add ...
func (h *SettingHandler) Add(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(context.Background(), time.Duration(str.StringToInt(h.ContractUC.EnvConfig["APP_TIMEOUT"]))*time.Second)
	c = context.WithValue(c, "requestid", ctx.Locals("requestid"))
	c = context.WithValue(c, "user_id", ctx.Locals("user_id"))
	defer cancel()

	input := new(requests.SettingRequest)
	if err := ctx.BodyParser(input); err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}
	if err := h.Validator.Struct(input); err != nil {
		errMessage := h.ExtractErrorValidationMessages(err.(validator.ValidationErrors))
		return h.SendResponse(ctx, nil, nil, errMessage, http.StatusBadRequest)
	}

	settingUC := usecase.SettingUC{ContractUC: h.ContractUC}
	res, err := settingUC.Add(c, input)

	return h.SendResponse(ctx, res, nil, err, 0)
}

// Edit ...
func (h *SettingHandler) Edit(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(context.Background(), time.Duration(str.StringToInt(h.ContractUC.EnvConfig["APP_TIMEOUT"]))*time.Second)
	c = context.WithValue(c, "requestid", ctx.Locals("requestid"))
	c = context.WithValue(c, "user_id", ctx.Locals("user_id"))
	defer cancel()

	id := ctx.Params("id")
	if id == "" {
		return h.SendResponse(ctx, nil, nil, helper.InvalidParameter, http.StatusBadRequest)
	}

	input := new(requests.SettingRequest)
	if err := ctx.BodyParser(input); err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}
	if err := h.Validator.Struct(input); err != nil {
		errMessage := h.ExtractErrorValidationMessages(err.(validator.ValidationErrors))
		return h.SendResponse(ctx, nil, nil, errMessage, http.StatusBadRequest)
	}

	settingUC := usecase.SettingUC{ContractUC: h.ContractUC}
	res, err := settingUC.Edit(c, id, input)

	return h.SendResponse(ctx, res, nil, err, 0)
}

// Delete ...
func (h *SettingHandler) Delete(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(context.Background(), time.Duration(str.StringToInt(h.ContractUC.EnvConfig["APP_TIMEOUT"]))*time.Second)
	c = context.WithValue(c, "requestid", ctx.Locals("requestid"))
	c = context.WithValue(c, "user_id", ctx.Locals("user_id"))
	defer cancel()

	id := ctx.Params("id")
	if id == "" {
		return h.SendResponse(ctx, nil, nil, helper.InvalidParameter, http.StatusBadRequest)
	}

	settingUC := usecase.SettingUC{ContractUC: h.ContractUC}
	res, err := settingUC.Delete(c, id)

	return h.SendResponse(ctx, res, nil, err, 0)
}
