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

// CouponRedeemHandler ...
type CouponRedeemHandler struct {
	Handler
}

// FindAll ...
func (h *CouponRedeemHandler) FindAll(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)
	parameter := models.CouponRedeemParameter{
		Search:     ctx.Query("search"),
		Now:        ctx.Query("now"),
		CustomerID: ctx.Query("customer_id"),
		Page:       str.StringToInt(ctx.Query("page")),
		Limit:      str.StringToInt(ctx.Query("limit")),
		By:         ctx.Query("by"),
		Sort:       ctx.Query("sort"),
	}
	uc := usecase.CouponRedeemUC{ContractUC: h.ContractUC}
	res, meta, err := uc.FindAll(c, parameter)
	if err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}

	return h.SendResponse(ctx, res, meta, err, 0)
}

// SelectAll ...
func (h *CouponRedeemHandler) SelectAll(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)
	parameter := models.CouponRedeemParameter{
		Search:     ctx.Query("search"),
		ShowAll:    ctx.Query("show_all"),
		CustomerID: ctx.Query("customer_id"),
		By:         ctx.Query("by"),
		Sort:       ctx.Query("sort"),
	}
	uc := usecase.CouponRedeemUC{ContractUC: h.ContractUC}
	res, err := uc.SelectAll(c, parameter)
	if err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}

	return h.SendResponse(ctx, res, nil, err, 0)
}

// FindByID ...
func (h *CouponRedeemHandler) FindByID(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)
	parameter := models.CouponRedeemParameter{
		ID: ctx.Params("id"),
	}
	uc := usecase.CouponRedeemUC{ContractUC: h.ContractUC}
	res, err := uc.FindByID(c, parameter)
	if err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}

	return h.SendResponse(ctx, res, nil, err, 0)
}

// SendOTP ...
func (h *CouponRedeemHandler) SendOTP(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	input := new(requests.CouponRedeemRequest)
	if err := ctx.BodyParser(input); err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}
	if err := h.Validator.Struct(input); err != nil {
		errMessage := h.ExtractErrorValidationMessages(err.(validator.ValidationErrors))
		return h.SendResponse(ctx, nil, nil, errMessage, http.StatusBadRequest)
	}

	uc := usecase.CouponRedeemUC{ContractUC: h.ContractUC}
	_, err := uc.SendOTP(c, *input)
	if err != nil {
		return h.SendResponse(ctx, nil, nil, err, 0)
	}

	return h.SendResponse(ctx, nil, nil, err, 0)
}

// VerifyOTP ...
func (h *CouponRedeemHandler) VerifyOTP(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	input := new(requests.CouponRedeemRequest)
	if err := ctx.BodyParser(input); err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}
	if err := h.Validator.Struct(input); err != nil {
		errMessage := h.ExtractErrorValidationMessages(err.(validator.ValidationErrors))
		return h.SendResponse(ctx, nil, nil, errMessage, http.StatusBadRequest)
	}

	uc := usecase.CouponRedeemUC{ContractUC: h.ContractUC}
	res, err := uc.VerifyOTP(c, *input)
	if err != nil {
		return h.SendResponse(ctx, nil, nil, err, 0)
	}

	return h.SendResponse(ctx, res, nil, err, 0)
}

// Add ...
func (h *CouponRedeemHandler) Add(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	input := new(requests.CouponRedeemRequest)
	if err := ctx.BodyParser(input); err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}
	if err := h.Validator.Struct(input); err != nil {
		errMessage := h.ExtractErrorValidationMessages(err.(validator.ValidationErrors))
		return h.SendResponse(ctx, nil, nil, errMessage, http.StatusBadRequest)
	}

	uc := usecase.CouponRedeemUC{ContractUC: h.ContractUC}
	res, err := uc.Add(c, *input)
	if err != nil {
		return h.SendResponse(ctx, nil, nil, err, 0)
	}

	return h.SendResponse(ctx, res, nil, err, 0)
}

// Revert ...
func (h *CouponRedeemHandler) Revert(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	invoiceNo := ctx.Params("invoice_no")
	if invoiceNo == "" {
		return h.SendResponse(ctx, nil, nil, helper.InvalidParameter, http.StatusBadRequest)
	}

	uc := usecase.CouponRedeemUC{ContractUC: h.ContractUC}
	res, err := uc.Revert(c, invoiceNo)
	if err != nil {
		return h.SendResponse(ctx, nil, nil, err, 0)
	}

	return h.SendResponse(ctx, res, nil, err, 0)
}

// Photo ...
func (h *CouponRedeemHandler) Photo(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	image, _ := ctx.FormFile("image")

	uc := usecase.CouponRedeemUC{ContractUC: h.ContractUC}
	res, err := uc.AddPhoto(c, image)
	if err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}

	return h.SendResponse(ctx, res, nil, err, 0)
}

// ReportSelect ...
func (h *CouponRedeemHandler) ReportSelect(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)
	parameter := models.CouponRedeemParameter{
		Search:          ctx.Query("search"),
		ShowAll:         ctx.Query("show_all"),
		CustomerID:      ctx.Query("customer_id"),
		CustomerLevelID: ctx.Query("customer_level_id"),
		CouponStatus:    ctx.Query("coupon_status"),
		StartDate:       ctx.Query("start_date"),
		EndDate:         ctx.Query("end_date"),
		BranchID:        ctx.Query("branch_id"),
		RegionID:        ctx.Query("region_id"),
		RegionGroupID:   ctx.Query("region_group_id"),
		By:              ctx.Query("by"),
		Sort:            ctx.Query("sort"),
	}
	uc := usecase.CouponRedeemUC{ContractUC: h.ContractUC}
	res, err := uc.SelectReport(c, parameter)
	if err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}

	return h.SendResponse(ctx, res, nil, err, 0)
}
