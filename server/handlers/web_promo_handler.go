package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/helper"
	"nextbasis-service-v-0.1/pkg/str"
	"nextbasis-service-v-0.1/server/requests"
	"nextbasis-service-v-0.1/usecase"
)

// WebPromoHandler ...
type WebPromoHandler struct {
	Handler
}

// SelectAll ...
func (h *WebPromoHandler) SelectAll(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.WebPromoParameter{
		ID:        ctx.Query("promo_id"),
		Code:      ctx.Query("promo_code"),
		StartDate: ctx.Query("start_date"),
		EndDate:   ctx.Query("end_date"),
		Search:    ctx.Query("search"),
		By:        ctx.Query("by"),
		Sort:      ctx.Query("sort"),
	}
	uc := usecase.WebPromoUC{ContractUC: h.ContractUC}
	res, err := uc.SelectAll(c, parameter)

	type StructObject struct {
		ListObjcet []models.WebPromo `json:"list_promo"`
	}

	ObjcetData := new(StructObject)

	if res != nil {
		ObjcetData.ListObjcet = res
	}

	return h.SendResponse(ctx, ObjcetData, nil, err, 0)
}

// FindAll ...
func (h *WebPromoHandler) FindAll(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.WebPromoParameter{
		ID:     ctx.Query("customer_id"),
		Code:   ctx.Query("customer_code"),
		Search: ctx.Query("search"),
		Page:   str.StringToInt(ctx.Query("page")),
		Limit:  str.StringToInt(ctx.Query("limit")),
		By:     ctx.Query("by"),
		Sort:   ctx.Query("sort"),
	}
	uc := usecase.WebPromoUC{ContractUC: h.ContractUC}
	res, meta, err := uc.FindAll(c, parameter)

	type StructObject struct {
		ListObject []models.WebPromo `json:"list_customer"`
	}

	ObjectData := new(StructObject)

	if res != nil {
		ObjectData.ListObject = res
	}

	return h.SendResponse(ctx, ObjectData, meta, err, 0)
}

// Add ...
func (h *WebPromoHandler) Add(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	input := new(requests.WebPromoRequest)
	err := json.Unmarshal([]byte(ctx.FormValue("form_data")), input)
	if err := ctx.BodyParser(input); err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}
	if err := h.Validator.Struct(input); err != nil {
		errMessage := h.ExtractErrorValidationMessages(err.(validator.ValidationErrors))
		return h.SendResponse(ctx, nil, nil, errMessage, http.StatusBadRequest)
	}

	imgBanner, _ := ctx.FormFile("img_banner")

	uc := usecase.WebPromoUC{ContractUC: h.ContractUC}
	res, err := uc.Add(c, input, imgBanner)

	return h.SendResponse(ctx, res, nil, err, 0)
}

// Edit ...
func (h *WebPromoHandler) Edit(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	id := ctx.Params("id")
	if id == "" {
		return h.SendResponse(ctx, nil, nil, helper.InvalidParameter, http.StatusBadRequest)
	}

	input := new(requests.WebPromoRequest)
	err := json.Unmarshal([]byte(ctx.FormValue("form_data")), input)
	if err := ctx.BodyParser(input); err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}
	if err := h.Validator.Struct(input); err != nil {
		errMessage := h.ExtractErrorValidationMessages(err.(validator.ValidationErrors))
		return h.SendResponse(ctx, nil, nil, errMessage, http.StatusBadRequest)
	}

	imgBanner, _ := ctx.FormFile("img_banner")

	uc := usecase.WebPromoUC{ContractUC: h.ContractUC}
	res, err := uc.Edit(c, input, imgBanner, id)

	return h.SendResponse(ctx, res, nil, err, 0)
}

// Delete ...
func (h *WebPromoHandler) Delete(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	id := ctx.Params("id")
	if id == "" {
		return h.SendResponse(ctx, nil, nil, helper.InvalidParameter, http.StatusBadRequest)
	}

	uc := usecase.WebPromoUC{ContractUC: h.ContractUC}
	res, err := uc.Delete(c, id)

	return h.SendResponse(ctx, res, nil, err, 0)
}

// FindByID ...
func (h *WebPromoHandler) FindByID(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.WebPromoParameter{
		ID: ctx.Params("id"),
	}
	if parameter.ID == "" {
		return h.SendResponse(ctx, nil, nil, helper.InvalidParameter, http.StatusBadRequest)
	}

	uc := usecase.WebPromoUC{ContractUC: h.ContractUC}
	res, err := uc.FindByID(c, parameter)

	ucEligible := usecase.WebCustomerTypeEligiblePromoUC{ContractUC: h.ContractUC}
	resEligible, errEligible := ucEligible.SelectAll(c, models.WebCustomerTypeEligiblePromoParameter{
		PromoID: *res.ID,
		By:      "pr._name",
	})

	if resEligible == nil {
		resEligible = make([]models.WebCustomerTypeEligiblePromo, 0)
	}
	var customerTypeIDList string
	for i := range resEligible {
		if resEligible[i].CustomerTypeId != nil {
			if customerTypeIDList == "" {
				customerTypeIDList += *resEligible[i].CustomerTypeId
			} else {
				customerTypeIDList += "," + *resEligible[i].CustomerTypeId
			}
		}
	}
	if errEligible == nil {
		res.CustomerTypeList = &resEligible
		res.CustomerTypeIdList = &customerTypeIDList
	}

	ucRegionEligible := usecase.WebRegionAreaEligiblePromoUC{ContractUC: h.ContractUC}
	resRegionEligible, errRegionEligible := ucRegionEligible.SelectAll(c, models.WebRegionAreaEligiblePromoParameter{
		PromoID: *res.ID,
		By:      "pr._name",
	})

	if resRegionEligible == nil {
		resRegionEligible = make([]models.WebRegionAreaEligiblePromo, 0)
	}

	var regionAreaIDList string
	for i := range resRegionEligible {
		if resRegionEligible[i].RegionID != nil {
			if regionAreaIDList == "" {
				regionAreaIDList += *resRegionEligible[i].RegionID
			} else {
				regionAreaIDList += "," + *resRegionEligible[i].RegionID
			}
		}
	}
	if errRegionEligible == nil {
		res.RegionAreaList = &resRegionEligible
		res.RegionAreaIdList = &regionAreaIDList
	}

	customerLevelUCEligible := usecase.WebCustomerLevelEligiblePromoUC{ContractUC: h.ContractUC}
	customerLevelresEligible, errEligible := customerLevelUCEligible.SelectAll(c, models.WebCustomerLevelEligiblePromoParameter{
		PromoID: *res.ID,
		By:      "pr._name",
	})

	if customerLevelresEligible == nil {
		customerLevelresEligible = make([]models.WebCustomerLevelEligiblePromo, 0)
	}
	var customerLevelIDList string
	for i := range customerLevelresEligible {
		if customerLevelresEligible[i].CustomerLevelId != nil {
			if customerLevelIDList == "" {
				customerLevelIDList += *customerLevelresEligible[i].CustomerLevelId
			} else {
				customerLevelIDList += "," + *customerLevelresEligible[i].CustomerLevelId
			}
		}
	}
	if errEligible == nil {
		res.CustomerLevelList = &customerLevelresEligible
		res.CustomerLevelIdList = &customerLevelIDList
	}

	branchUCEligible := usecase.WebBranchEligiblePromoUC{ContractUC: h.ContractUC}
	branchresEligible, errEligible := branchUCEligible.SelectAll(c, models.WebBranchEligiblePromoParameter{
		PromoID: *res.ID,
		By:      "pr._name",
	})
	if branchresEligible == nil {
		branchresEligible = make([]models.WebBranchEligiblePromo, 0)
	}
	var branchIDList string
	for i := range branchresEligible {
		if customerLevelresEligible[i].CustomerLevelId != nil {
			if branchIDList == "" {
				branchIDList += *customerLevelresEligible[i].CustomerLevelId
			} else {
				branchIDList += "," + *customerLevelresEligible[i].CustomerLevelId
			}
		}
	}
	if errEligible == nil {
		res.BranchList = &branchresEligible
		res.BranchIdList = &branchIDList
	}

	return h.SendResponse(ctx, res, nil, err, 0)
}
