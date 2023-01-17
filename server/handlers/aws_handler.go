package handlers

import (
	"context"
	"time"

	"nextbasis-service-v-0.1/pkg/str"
	"nextbasis-service-v-0.1/usecase"

	"github.com/gofiber/fiber/v2"
)

// AwsHandler ...
type AwsHandler struct {
	Handler
}

// Upload ...
func (h *AwsHandler) Upload(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(context.Background(), time.Duration(str.StringToInt(h.ContractUC.EnvConfig["APP_TIMEOUT"]))*time.Second)
	c = context.WithValue(c, "requestid", ctx.Locals("requestid"))
	c = context.WithValue(c, "user_id", ctx.Locals("user_id"))
	defer cancel()

	AwsHeader, err := ctx.FormFile("upload_image")
	if err != nil {
		return h.SendResponse(ctx, nil, nil, err, 0)
	}
	AwsUc := usecase.AwsUC{ContractUC: h.ContractUC}
	res, err := AwsUc.Upload("", AwsHeader)

	return h.SendResponse(ctx, res, nil, err, 0)
}

func (h *AwsHandler) Delete(ctx *fiber.Ctx) error {

	FileToDelete := ctx.Query("str_delete_file")
	FilePath := ctx.Query("file_path")

	AwsUc := usecase.AwsUC{ContractUC: h.ContractUC}
	res, err := AwsUc.Delete(FilePath, FileToDelete)

	return h.SendResponse(ctx, res, nil, err, 0)
}
