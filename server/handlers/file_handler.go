package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/helper"
	"nextbasis-service-v-0.1/pkg/str"
	"nextbasis-service-v-0.1/server/requests"
	"nextbasis-service-v-0.1/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// FileHandler ...
type FileHandler struct {
	Handler
}

// Upload ...
func (h *FileHandler) Upload(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(context.Background(), time.Duration(str.StringToInt(h.ContractUC.EnvConfig["APP_TIMEOUT"]))*time.Second)
	c = context.WithValue(c, "requestid", ctx.Locals("requestid"))
	c = context.WithValue(c, "user_id", ctx.Locals("user_id"))
	defer cancel()

	// Read file type
	fileType := ctx.FormValue("type")
	if !str.Contains(models.FileWhitelist, fileType) {
		return h.SendResponse(ctx, nil, nil, errors.New(helper.InvalidFileType), 0)
	}

	// Upload file to local temporary
	fileHeader, err := ctx.FormFile("file")
	if err != nil {
		return h.SendResponse(ctx, nil, nil, err, 0)
	}
	fileUc := usecase.FileUC{ContractUC: h.ContractUC}
	res, err := fileUc.Upload(c, fileType, fileHeader)

	return h.SendResponse(ctx, res, nil, err, 0)
}

func (h *FileHandler) UploadTes(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(context.Background(), time.Duration(str.StringToInt(h.ContractUC.EnvConfig["APP_TIMEOUT"]))*time.Second)
	c = context.WithValue(c, "requestid", ctx.Locals("requestid"))
	c = context.WithValue(c, "user_id", ctx.Query("user_id"))
	defer cancel()

	input := new(requests.CityRequest)
	err := json.Unmarshal([]byte(ctx.FormValue("arrdata")), input)
	if err := ctx.BodyParser(input); err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}
	if err := h.Validator.Struct(input); err != nil {
		errMessage := h.ExtractErrorValidationMessages(err.(validator.ValidationErrors))
		return h.SendResponse(ctx, nil, nil, errMessage, http.StatusBadRequest)
	}

	// Read file type
	fileType := ctx.FormValue("type")
	if !str.Contains(models.FileWhitelist, fileType) {
		return h.SendResponse(ctx, nil, nil, errors.New(helper.InvalidFileType), 0)
	}

	// Upload file to local temporary
	fileHeader, err := ctx.FormFile("img_ktp")
	if err != nil {
		return h.SendResponse(ctx, nil, nil, err, 0)
	}
	fileUc := usecase.FileUC{ContractUC: h.ContractUC}
	res, err := fileUc.UploadTes(c, fileType, fileHeader)

	return h.SendResponse(ctx, res, nil, err, 0)
}
