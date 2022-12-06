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

	// // Read Aws type
	// AwsType := ctx.FormValue("type")
	// if !str.Contains(models.FileWhitelist, AwsType) {
	// 	return h.SendResponse(ctx, nil, nil, errors.New(helper.InvalidFileType), 0)
	// }

	// Upload Aws to local temporary
	AwsHeader, err := ctx.FormFile("upload_image")
	if err != nil {
		return h.SendResponse(ctx, nil, nil, err, 0)
	}
	AwsUc := usecase.AwsUC{ContractUC: h.ContractUC}
	res, err := AwsUc.Upload("customer", AwsHeader)

	return h.SendResponse(ctx, res, nil, err, 0)
}
