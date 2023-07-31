package handlers

import (
	"context"
	"crypto/sha512"
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/helper"
	"nextbasis-service-v-0.1/pkg/str"
	"nextbasis-service-v-0.1/server/requests"
	"nextbasis-service-v-0.1/usecase"
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

// TransactionVAHandler ...
type TransactionVAHandler struct {
	Handler
}

// SelectAll ...
func (h *TransactionVAHandler) SelectAll(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.TransactionVAParameter{
		ID:     ctx.Query("customer_id"),
		UserId: ctx.Query("admin_user_id"),
		Search: ctx.Query("search"),
		By:     ctx.Query("by"),
		Sort:   ctx.Query("sort"),
	}
	uc := usecase.TransactionVAUC{ContractUC: h.ContractUC}
	res, err := uc.SelectAll(c, parameter)

	type StructObject struct {
		ListObject []models.TransactionVA `json:"list_partner"`
	}

	ObjectData := new(StructObject)

	if res != nil {
		ObjectData.ListObject = res
	}

	return h.SendResponse(ctx, ObjectData, nil, err, 0)
}

// FindAll ...
func (h *TransactionVAHandler) FindAll(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.TransactionVAParameter{
		ID:     ctx.Query("customer_id"),
		UserId: ctx.Query("admin_user_id"),
		Search: ctx.Query("search"),
		Page:   str.StringToInt(ctx.Query("page")),
		Limit:  str.StringToInt(ctx.Query("limit")),
		By:     ctx.Query("by"),
		Sort:   ctx.Query("sort"),
	}
	uc := usecase.TransactionVAUC{ContractUC: h.ContractUC}
	res, meta, err := uc.FindAll(c, parameter)

	type StructObject struct {
		ListObject []models.TransactionVA `json:"list_partner"`
	}

	ObjectData := new(StructObject)

	if res != nil {
		ObjectData.ListObject = res
	}

	return h.SendResponse(ctx, ObjectData, meta, err, 0)
}

// FindByID ...
func (h *TransactionVAHandler) FindByID(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.TransactionVAParameter{
		ID: ctx.Params("partner_id"),
	}
	if parameter.ID == "" {
		return h.SendResponse(ctx, nil, nil, helper.InvalidParameter, http.StatusBadRequest)
	}

	uc := usecase.TransactionVAUC{ContractUC: h.ContractUC}
	res, err := uc.FindByID(c, parameter)

	type StructObject struct {
		ListObject models.TransactionVA `json:"partner"`
	}

	ObjectData := new(StructObject)

	ObjectData.ListObject = res

	return h.SendResponse(ctx, ObjectData, nil, err, 0)
}

// Edit ...
func (h *TransactionVAHandler) Edit(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	id := ctx.Params("partner_id")
	if id == "" {
		return h.SendResponse(ctx, nil, nil, helper.InvalidParameter, http.StatusBadRequest)
	}

	input := new(requests.TransactionVARequest)
	if err := ctx.BodyParser(input); err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}
	if err := h.Validator.Struct(input); err != nil {
		errMessage := h.ExtractErrorValidationMessages(err.(validator.ValidationErrors))
		return h.SendResponse(ctx, nil, nil, errMessage, http.StatusBadRequest)
	}
	uc := usecase.TransactionVAUC{ContractUC: h.ContractUC}
	res, err := uc.Edit(c, id, input)

	return h.SendResponse(ctx, res, nil, err, 0)
}

// Edit ...
func (h *TransactionVAHandler) Add(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	input := new(requests.TransactionVARequest)
	if err := ctx.BodyParser(input); err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}
	if err := h.Validator.Struct(input); err != nil {
		errMessage := h.ExtractErrorValidationMessages(err.(validator.ValidationErrors))
		return h.SendResponse(ctx, nil, nil, errMessage, http.StatusBadRequest)
	}
	uc := usecase.TransactionVAUC{ContractUC: h.ContractUC}
	res, err := uc.Add(c, input)

	return h.SendResponse(ctx, res, nil, err, 0)
}

func (h *TransactionVAHandler) GetTransactionByVaCode(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	input := new(requests.InquiryVaRequest)
	if err := ctx.BodyParser(input); err != nil {
		return h.SendBasicResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}
	fmt.Println("Masuk Sini")
	if err := h.Validator.Struct(input); err != nil {
		errMessage := h.ExtractErrorValidationMessages(err.(validator.ValidationErrors))
		return h.SendBasicResponse(ctx, nil, nil, errMessage, http.StatusBadRequest)
	}

	parameter := models.TransactionVAParameter{
		VACode:        input.InquiryBody.Billkey1,
		CurrentVaUser: 1,
	}

	fmt.Println(parameter.VACode)

	type InquiryResponseData struct {
		InquiryResult viewmodel.VaBillInfoVM `json:"InquiryResponse"`
	}

	ObjectData := new(InquiryResponseData)

	uc := usecase.TransactionVAUC{ContractUC: h.ContractUC}
	res, err := uc.FindByCode(c, parameter)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			ObjectData.InquiryResult.Status.ErrorCode = "B5"
			ObjectData.InquiryResult.Status.IsError = "true"
			ObjectData.InquiryResult.Status.StatusDescription = "Bill Not Found"
			err = nil
		}
	}

	if res.VACode != nil {

		ObjectDataDetail := new(viewmodel.VaBillDetailVM)

		ObjectData.InquiryResult.BillInfo1 = *res.VACode
		ObjectData.InquiryResult.BillInfo2 = *res.Customername

		ObjectDataDetail.BillCode = "01"
		ObjectDataDetail.BillName = *res.InvoiceCode
		ObjectDataDetail.BillShortName = "Pembayaran"
		ObjectDataDetail.BillAmount = *res.Amount
		ObjectData.InquiryResult.VabillDetails.BillDetail = append(ObjectData.InquiryResult.VabillDetails.BillDetail, *ObjectDataDetail)
		ObjectData.InquiryResult.Currency = "360"
		ObjectData.InquiryResult.Status.IsError = "false"
		ObjectData.InquiryResult.Status.ErrorCode = "00"
		ObjectData.InquiryResult.Status.StatusDescription = "Transaction Success"
		if res.PaidStatus != nil && *res.PaidStatus == "paid" {
			ObjectData.InquiryResult.Status.IsError = "false"
			ObjectData.InquiryResult.Status.ErrorCode = "B8"
			ObjectData.InquiryResult.Status.StatusDescription = "Bill Already Paid"
		}
	}

	return h.SendBasicResponse(ctx, ObjectData, nil, err, 0)
}

func (h *TransactionVAHandler) GetSah(ctx *fiber.Ctx) error {
	SHA512 := sha512.New()
	SHA512.Write([]byte("BMRI_SIDO"))
	basic := sha512.Sum512([]byte("BMRI_SIDO"))
	// newPasswd := base64.StdEncoding.EncodeToString(basic[:])
	token := fmt.Sprintf("sha512: %x", basic)
	return h.SendBasicResponse(ctx, token, nil, nil, 0)
}

func (h *TransactionVAHandler) PaidTransactionByVaCode(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	input := new(requests.PaymentVaRequest)
	if err := ctx.BodyParser(input); err != nil {
		return h.SendBasicResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}

	if err := h.Validator.Struct(input); err != nil {
		errMessage := h.ExtractErrorValidationMessages(err.(validator.ValidationErrors))
		return h.SendBasicResponse(ctx, nil, nil, errMessage, http.StatusBadRequest)
	}

	parameter := models.TransactionVAParameter{
		VACode:        input.PaymentRequestBody.Billkey1,
		CurrentVaUser: 0,
	}

	fmt.Println(parameter.VACode)

	type InquiryResponseData struct {
		InquiryResult viewmodel.PaymentVaBillInfoVM `json:"paymentResponse"`
	}

	ObjectData := new(InquiryResponseData)

	if parameter.VACode == "8954102230400004" {

		fmt.Println("before sleep")
		time.Sleep(30 * time.Second)
		fmt.Println("after sleep")
		return nil
	}

	if parameter.VACode == "8954102230400005" {
		ObjectData.InquiryResult.Status.ErrorCode = "87"
		ObjectData.InquiryResult.Status.IsError = "true"
		ObjectData.InquiryResult.Status.StatusDescription = "Provider Database Problem"
		return h.SendBasicResponse(ctx, ObjectData, nil, nil, 0)
	}

	if parameter.VACode == "8954102230400006" {
		ObjectData.InquiryResult.Status.ErrorCode = "91"
		ObjectData.InquiryResult.Status.IsError = "true"
		ObjectData.InquiryResult.Status.StatusDescription = "Link Down"
		return h.SendBasicResponse(ctx, ObjectData, nil, nil, 0)
	}

	if parameter.VACode == "8954102230400007" {
		ObjectData.InquiryResult.Status.ErrorCode = "89"
		ObjectData.InquiryResult.Status.IsError = "true"
		ObjectData.InquiryResult.Status.StatusDescription = "Time Out"
		return h.SendBasicResponse(ctx, ObjectData, nil, nil, 0)
	}

	if parameter.VACode == "8954102230400008" {
		ObjectData.InquiryResult.Status.ErrorCode = "01"
		ObjectData.InquiryResult.Status.IsError = "true"
		ObjectData.InquiryResult.Status.StatusDescription = "General Error"
		return h.SendBasicResponse(ctx, ObjectData, nil, nil, 0)
	}
	if parameter.VACode == "8954102230400009" {
		ObjectData.InquiryResult.Status.ErrorCode = "C0"
		ObjectData.InquiryResult.Status.IsError = "true"
		ObjectData.InquiryResult.Status.StatusDescription = "Bill Suspended"
		return h.SendBasicResponse(ctx, ObjectData, nil, nil, 0)
	}

	uc := usecase.TransactionVAUC{ContractUC: h.ContractUC}
	res, err := uc.FindByCode(c, parameter)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			ObjectData.InquiryResult.Status.ErrorCode = "B5"
			ObjectData.InquiryResult.Status.IsError = "true"
			ObjectData.InquiryResult.Status.StatusDescription = "Bill Not Found"
			err = nil
		}
	}

	if res.VACode != nil {
		uc := usecase.TransactionVAUC{ContractUC: h.ContractUC}
		inputUpdate := new(requests.TransactionVARequest)
		inputUpdate.VaRef1 = input.PaymentRequestBody.Reference1
		inputUpdate.VaRef2 = input.PaymentRequestBody.Reference2
		inputUpdate.VaPairID = input.PaymentRequestBody.TransactionID
		inputUpdate.Amount = input.PaymentRequestBody.Billkey2
		_, errpaid := uc.PaidTransaction(c, *res.ID, inputUpdate, res)

		if errpaid == nil {

			resafterupdate, errafter := uc.FindByCode(c, parameter)

			if errafter != nil {
				if errafter.Error() == "sql: no rows in result set" {
					ObjectData.InquiryResult.Status.ErrorCode = "B5"
					ObjectData.InquiryResult.Status.IsError = "true"
					ObjectData.InquiryResult.Status.StatusDescription = "Bill Not Found"
					errafter = nil
				}
			}

			ObjectData.InquiryResult.BillInfo1 = *resafterupdate.VACode
			ObjectData.InquiryResult.BillInfo2 = *resafterupdate.Customername

			ObjectData.InquiryResult.Status.IsError = "false"
			ObjectData.InquiryResult.Status.ErrorCode = "00"
			ObjectData.InquiryResult.Status.StatusDescription = "Transaction Success"

		}
	}

	return h.SendBasicResponse(ctx, ObjectData, nil, err, 0)
}
