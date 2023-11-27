package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/helper"
	"nextbasis-service-v-0.1/pkg/str"
	"nextbasis-service-v-0.1/server/requests"
	"nextbasis-service-v-0.1/usecase"
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

// WebCustomerHandler ...
type WebCustomerHandler struct {
	Handler
}

// SelectAll ...
func (h *WebCustomerHandler) SelectAll(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.WebCustomerParameter{
		ID:             ctx.Query("customer_id"),
		CustomerTypeId: ctx.Query("customer_type_id"),
		SalesmanTypeID: ctx.Query("salesman_type_id"),
		UserId:         ctx.Query("admin_user_id"),
		BranchId:       ctx.Query("branch_id"),
		Search:         ctx.Query("search"),
		By:             ctx.Query("by"),
		Sort:           ctx.Query("sort"),
	}
	uc := usecase.WebCustomerUC{ContractUC: h.ContractUC}
	res, err := uc.SelectAll(c, parameter)

	type StructObject struct {
		ListObject []viewmodel.CustomerVM `json:"list_customer"`
	}

	objectData := new(StructObject)

	if res != nil {
		objectData.ListObject = res
	}

	return h.SendResponse(ctx, objectData, nil, err, 0)
}

// FindAll ...
func (h *WebCustomerHandler) FindAll(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)
	parameter := models.WebCustomerParameter{
		ID:             ctx.Query("customer_id"),
		CustomerTypeId: ctx.Query("customer_type_id"),
		UserId:         ctx.Query("admin_user_id"),
		BranchId:       ctx.Query("branch_id"),
		Search:         ctx.Query("search"),
		Page:           str.StringToInt(ctx.Query("page")),
		Limit:          str.StringToInt(ctx.Query("limit")),
		By:             ctx.Query("by"),
		Sort:           ctx.Query("sort"),
		PhoneNumber:    ctx.Query("phone_number"),
		ShowInApp:      ctx.Query("show_in_app"),
		Active:         ctx.Query("active"),
		IsDataComplete: ctx.Query("is_data_complete"),
	}
	uc := usecase.WebCustomerUC{ContractUC: h.ContractUC}
	res, meta, err := uc.FindAll(c, parameter)

	type StructObject struct {
		ListObject []viewmodel.CustomerVM `json:"list_customer"`
	}

	objectData := new(StructObject)

	if res != nil {
		objectData.ListObject = res
	}

	return h.SendResponse(ctx, objectData, meta, err, 0)
}

// FindByID ...
func (h *WebCustomerHandler) FindByID(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.WebCustomerParameter{
		ID: ctx.Params("customer_id"),
	}
	if parameter.ID == "" {
		return h.SendResponse(ctx, nil, nil, helper.InvalidParameter, http.StatusBadRequest)
	}

	uc := usecase.WebCustomerUC{ContractUC: h.ContractUC}
	res, err := uc.FindByID(c, parameter)

	type StructObject struct {
		ListObject          viewmodel.CustomerVM   `json:"customer"`
		CustomerTarget      interface{}            `json:"customer_target"`
		CustomerAchievement map[string]interface{} `json:"customer_achievement"`
		SalesmanVisit       interface{}            `json:"salesman_visit"`
	}

	objectData := new(StructObject)

	objectData.ListObject = res

	target := h.FetchVisitDay(parameter)
	if target != "" {
		objectData.ListObject.VisitDay = target
	}

	objectData.CustomerTarget = helper.FetchClientDataTarget(models.CustomerTargetSemesterParameter{
		ID:   res.ID,
		Code: res.Code,
	})

	achievement := make(map[string]interface{})
	quarterAchievement, _ := usecase.CustomerAchievementQuarterUC{ContractUC: h.ContractUC}.SelectAll(c, models.CustomerAchievementQuarterParameter{
		ID: res.ID,
		By: "cus.created_date",
	})
	if len(quarterAchievement) == 1 {
		achievement["quater_achievement"] = quarterAchievement[0].Achievement
	}
	semesterAchievement, _ := usecase.CustomerAchievementSemesterUC{ContractUC: h.ContractUC}.SelectAll(c, models.CustomerAchievementSemesterParameter{
		ID: res.ID,
		By: "cus.created_date",
	})
	if len(semesterAchievement) == 1 {
		achievement["semester_achievement"] = semesterAchievement[0].Achievement
	}
	yearAchievement, _ := usecase.CustomerAchievementYearUC{ContractUC: h.ContractUC}.SelectAll(c, models.CustomerAchievementYearParameter{
		ID: res.ID,
		By: "cus.created_date",
	})
	if len(yearAchievement) == 1 {
		achievement["year_achievement"] = yearAchievement[0].Achievement
	}
	annualAchievement, _ := usecase.CustomerAchievementUC{ContractUC: h.ContractUC}.SelectAll(c, models.CustomerAchievementParameter{
		ID: res.ID,
		By: "cus.created_date",
	})
	if len(annualAchievement) == 1 {
		achievement["month_achievement"] = annualAchievement[0].Achievement
	}
	objectData.CustomerAchievement = achievement

	objectData.SalesmanVisit = helper.FetchVisitDay(models.CustomerParameter{
		ID:   res.ID,
		Code: res.Code,
	})

	return h.SendResponse(ctx, objectData, nil, err, 0)
}

func (h *WebCustomerHandler) FetchVisitDay(params models.WebCustomerParameter) string {
	jsonReq, err := json.Marshal(params)
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://nextbasis.id:8080/mysmagonsrvxxx/rest/customer/visitday/1", bytes.NewBuffer(jsonReq))
	if err != nil {
		fmt.Println("client err")
		fmt.Print(err.Error())
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer C2A5CE6A2292E7745CE5A3F7E68A9")

	resp, err := client.Do(req)
	if err != nil {

		fmt.Print(err.Error())
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}

	type resutlData struct {
		VisitDay string `json:"visit_day"`
	}

	objectData := new(resutlData)

	// var responseObject http.Response
	json.Unmarshal(bodyBytes, &objectData)

	return objectData.VisitDay
}

// Edit ...
func (h *WebCustomerHandler) Edit(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	id := ctx.Params("customer_id")
	if id == "" {
		return h.SendResponse(ctx, nil, nil, helper.InvalidParameter, http.StatusBadRequest)
	}

	input := new(requests.WebCustomerRequest)
	err := json.Unmarshal([]byte(ctx.FormValue("form_data")), &input)
	if err := ctx.BodyParser(input); err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}
	if err := h.Validator.Struct(input); err != nil {
		errMessage := h.ExtractErrorValidationMessages(err.(validator.ValidationErrors))
		return h.SendResponse(ctx, nil, nil, errMessage, http.StatusBadRequest)
	}

	if input.CustomerGender != "" && !str.Contains(models.CustomerGenderList, strings.ToLower(input.CustomerGender)) {
		return h.SendResponse(ctx, nil, nil, errors.New(helper.InvalidGender), http.StatusBadRequest)
	}

	imgProfile, _ := ctx.FormFile("img_profile")
	imgKtp, _ := ctx.FormFile("img_ktp")
	uc := usecase.WebCustomerUC{ContractUC: h.ContractUC}
	res, err := uc.Edit(c, id, input, imgProfile, imgKtp)

	return h.SendResponse(ctx, res, nil, err, 0)
}

// EditBulk ...
func (h *WebCustomerHandler) EditBulk(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	input := new(requests.WebCustomerBulkRequest)
	if err := ctx.BodyParser(input); err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}
	if err := h.Validator.Struct(input); err != nil {
		errMessage := h.ExtractErrorValidationMessages(err.(validator.ValidationErrors))
		return h.SendResponse(ctx, nil, nil, errMessage, http.StatusBadRequest)
	}

	uc := usecase.WebCustomerUC{ContractUC: h.ContractUC}
	err := uc.EditBulk(c, *input)
	if err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}

	return h.SendResponse(ctx, nil, nil, err, 0)
}

// Edit ...
func (h *WebCustomerHandler) Add(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	input := new(requests.WebCustomerRequest)
	err := json.Unmarshal([]byte(ctx.FormValue("form_data")), input)
	if err := ctx.BodyParser(input); err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}
	if err := h.Validator.Struct(input); err != nil {
		errMessage := h.ExtractErrorValidationMessages(err.(validator.ValidationErrors))
		return h.SendResponse(ctx, nil, nil, errMessage, http.StatusBadRequest)
	}

	if input.CustomerGender != "" && !str.Contains(models.CustomerGenderList, strings.ToLower(input.CustomerGender)) {
		return h.SendResponse(ctx, nil, nil, errors.New(helper.InvalidGender), http.StatusBadRequest)
	}

	imgProfile, _ := ctx.FormFile("img_profile")
	uc := usecase.WebCustomerUC{ContractUC: h.ContractUC}
	_, err = uc.Add(c, input, imgProfile)

	return h.SendResponse(ctx, nil, nil, err, 0)
}

// ReportSelect ...
func (h *WebCustomerHandler) ReportSelect(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.WebCustomerReportParameter{
		RegionGroupID:         ctx.Query("region_group_id"),
		RegionID:              ctx.Query("region_id"),
		BranchArea:            ctx.Query("branch_area"),
		CustomerTypeID:        ctx.Query("customer_type_id"),
		BranchIDs:             ctx.Query("branch_ids"),
		CustomerLevelID:       ctx.Query("customer_level_id"),
		CustomerProfileStatus: ctx.Query("customer_profile_status"),
		AdminUserID:           ctx.Query("admin_user_id"),
	}
	uc := usecase.WebCustomerUC{ContractUC: h.ContractUC}
	res, err := uc.ReportSelect(c, parameter)
	if err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}

	if res == nil {
		err = errors.New("There is no customer with this filter")
		return h.SendResponse(ctx, res, nil, err, 0)
	}

	return h.SendResponse(ctx, res, nil, err, 0)
}
