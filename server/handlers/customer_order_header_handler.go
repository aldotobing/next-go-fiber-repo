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

// CustomerOrderHeaderHandler ...
type CustomerOrderHeaderHandler struct {
	Handler
}

// SelectAll ...
func (h *CustomerOrderHeaderHandler) SelectAll(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.CustomerOrderHeaderParameter{
		CustomerID: ctx.Params("customer_id"),
		Search:     ctx.Query("search"),
		By:         ctx.Query("by"),
		Sort:       ctx.Query("sort"),
	}
	uc := usecase.CustomerOrderHeaderUC{ContractUC: h.ContractUC}
	res, err := uc.SelectAll(c, parameter)

	for i := range res {
		lineuc := usecase.CustomerOrderLineUC{ContractUC: h.ContractUC}
		lineparameter := models.CustomerOrderLineParameter{
			HeaderID: *res[i].ID,
			Search:   ctx.Query("search"),
			By:       ctx.Query("by"),
			Sort:     ctx.Query("sort"),
		}
		listLine, _ := lineuc.SelectAll(c, lineparameter)

		if listLine != nil {
			res[i].ListLine = listLine
		}

	}

	type StructObject struct {
		ListObjcet []models.CustomerOrderHeader `json:"list_customer_order"`
	}

	ObjcetData := new(StructObject)

	if res != nil {
		ObjcetData.ListObjcet = res
	}

	return h.SendResponse(ctx, ObjcetData, nil, err, 0)
}

// FindAll ...
func (h *CustomerOrderHeaderHandler) FindAll(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.CustomerOrderHeaderParameter{
		CustomerID: ctx.Params("customer_id"),
		Search:     ctx.Query("search"),
		Page:       str.StringToInt(ctx.Query("page")),
		Limit:      str.StringToInt(ctx.Query("limit")),
		By:         ctx.Query("by"),
		Sort:       ctx.Query("sort"),
	}
	uc := usecase.CustomerOrderHeaderUC{ContractUC: h.ContractUC}
	res, meta, err := uc.FindAll(c, parameter)

	type StructObject struct {
		ListObjcet []models.CustomerOrderHeader `json:"list_customer_order"`
	}

	ObjcetData := new(StructObject)

	if res != nil {
		ObjcetData.ListObjcet = res
	}

	return h.SendResponse(ctx, ObjcetData, meta, err, 0)
}

// FindByID ...
func (h *CustomerOrderHeaderHandler) FindByID(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.CustomerOrderHeaderParameter{
		ID: ctx.Params("id"),
	}
	if parameter.ID == "" {
		return h.SendResponse(ctx, nil, nil, helper.InvalidParameter, http.StatusBadRequest)
	}

	uc := usecase.CustomerOrderHeaderUC{ContractUC: h.ContractUC}
	res, err := uc.FindByID(c, parameter)

	lineuc := usecase.CustomerOrderLineUC{ContractUC: h.ContractUC}
	lineparameter := models.CustomerOrderLineParameter{
		HeaderID: *res.ID,
		Search:   ctx.Query("search"),
		By:       "def.created_date",
		Sort:     ctx.Query("sort"),
	}
	listLine, _ := lineuc.SelectAll(c, lineparameter)

	if listLine != nil {
		res.ListLine = listLine
	}

	type StructObject struct {
		ListObjcet []models.CustomerOrderHeader `json:"list_customer_order"`
	}

	ObjcetData := new(StructObject)
	ObjcetData.ListObjcet = append(ObjcetData.ListObjcet, res)

	return h.SendResponse(ctx, ObjcetData, nil, err, 0)
}

// rest
func (h *CustomerOrderHeaderHandler) RestSelectAll(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.CustomerOrderHeaderParameter{
		DateParam: ctx.Query("date_param"),
		Search:    ctx.Query("search"),
		By:        ctx.Query("by"),
		Sort:      ctx.Query("sort"),
	}
	uc := usecase.CustomerOrderHeaderUC{ContractUC: h.ContractUC}
	res, err := uc.SelectAll(c, parameter)

	for i := range res {
		lineuc := usecase.CustomerOrderLineUC{ContractUC: h.ContractUC}
		lineparameter := models.CustomerOrderLineParameter{
			HeaderID: *res[i].ID,
			Search:   ctx.Query("search"),
			By:       ctx.Query("by"),
			Sort:     ctx.Query("sort"),
		}
		listLine, _ := lineuc.SelectAll(c, lineparameter)

		if listLine != nil {
			res[i].ListLine = listLine
		}

	}

	type StructObject struct {
		ListObjcet []models.CustomerOrderHeader `json:"list_customer_order"`
	}

	ObjcetData := new(StructObject)

	if res != nil {
		ObjcetData.ListObjcet = res
	}

	return h.SendResponse(ctx, ObjcetData, nil, err, 0)
}

// FindAll ...
func (h *CustomerOrderHeaderHandler) FindAllForWeb(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.CustomerOrderHeaderParameter{
		UserID: ctx.Query("admin_user_id"),
		Search: ctx.Query("search"),
		Page:   str.StringToInt(ctx.Query("page")),
		Limit:  str.StringToInt(ctx.Query("limit")),
		By:     ctx.Query("by"),
		Sort:   ctx.Query("sort"),
	}
	uc := usecase.CustomerOrderHeaderUC{ContractUC: h.ContractUC}
	res, meta, err := uc.FindAll(c, parameter)

	for i := range res {
		lineuc := usecase.CustomerOrderLineUC{ContractUC: h.ContractUC}
		lineparameter := models.CustomerOrderLineParameter{
			HeaderID: *res[i].ID,
			Search:   ctx.Query("search"),
			By:       ctx.Query("by"),
			Sort:     ctx.Query("sort"),
		}
		listLine, _ := lineuc.SelectAll(c, lineparameter)

		if listLine != nil {
			res[i].ListLine = listLine
		}

	}

	type StructObject struct {
		ListObjcet []models.CustomerOrderHeader `json:"list_customer_order"`
	}

	ObjcetData := new(StructObject)

	if res != nil {
		ObjcetData.ListObjcet = res
	}

	return h.SendResponse(ctx, ObjcetData, meta, err, 0)
}

// SelectAll ...
func (h *CustomerOrderHeaderHandler) AppsSelectAll(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.CustomerOrderHeaderParameter{
		CustomerID: ctx.Params("customer_id"),
		Search:     ctx.Query("search"),
		By:         ctx.Query("by"),
		Sort:       ctx.Query("sort"),
		StartDate:  ctx.Query("start_date"),
		EndDate:    ctx.Query("end_date"),
	}
	uc := usecase.CustomerOrderHeaderUC{ContractUC: h.ContractUC}
	res, err := uc.AppsSelectAll(c, parameter)

	for i := range res {
		lineuc := usecase.CustomerOrderLineUC{ContractUC: h.ContractUC}
		lineparameter := models.CustomerOrderLineParameter{
			HeaderID: *res[i].ID,
			Search:   ctx.Query("search"),
			By:       ctx.Query("by"),
			Sort:     ctx.Query("sort"),
		}

		if *res[i].OrderSource == "2" {
			listLine, _ := lineuc.SFASelectAll(c, lineparameter)

			if listLine != nil {
				res[i].ListLine = listLine
			}
		} else {
			listLine, _ := lineuc.SelectAll(c, lineparameter)

			if listLine != nil {
				res[i].ListLine = listLine
			}
		}

	}

	type StructObject struct {
		ListObjcet []models.CustomerOrderHeader `json:"list_customer_order"`
	}

	ObjcetData := new(StructObject)

	if res != nil {
		ObjcetData.ListObjcet = res
	}

	return h.SendResponse(ctx, ObjcetData, nil, err, 0)
}

// FindAll ...
func (h *CustomerOrderHeaderHandler) AppsFindAll(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.CustomerOrderHeaderParameter{
		CustomerID: ctx.Params("customer_id"),
		Search:     ctx.Query("search"),
		Page:       str.StringToInt(ctx.Query("page")),
		Limit:      str.StringToInt(ctx.Query("limit")),
		By:         ctx.Query("by"),
		Sort:       ctx.Query("sort"),
	}
	uc := usecase.CustomerOrderHeaderUC{ContractUC: h.ContractUC}
	res, meta, err := uc.FindAll(c, parameter)

	type StructObject struct {
		ListObjcet []models.CustomerOrderHeader `json:"list_customer_order"`
	}

	ObjcetData := new(StructObject)

	if res != nil {
		ObjcetData.ListObjcet = res
	}

	return h.SendResponse(ctx, ObjcetData, meta, err, 0)
}

// FindByID ...
func (h *CustomerOrderHeaderHandler) AppsFindByID(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.CustomerOrderHeaderParameter{
		ID: ctx.Params("id"),
	}
	if parameter.ID == "" {
		return h.SendResponse(ctx, nil, nil, helper.InvalidParameter, http.StatusBadRequest)
	}

	uc := usecase.CustomerOrderHeaderUC{ContractUC: h.ContractUC}
	res, err := uc.AppsFindByID(c, parameter)

	return h.SendResponse(ctx, res, nil, err, 0)
}

// FindByID ...
func (h *CustomerOrderHeaderHandler) ReUpdateDate(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)
	uc := usecase.CustomerOrderHeaderUC{ContractUC: h.ContractUC}
	res, err := uc.ReUpdateModifiedDate(c)

	return h.SendResponse(ctx, res, nil, err, 0)
}
