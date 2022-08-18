package handlers

import (
	"context"

	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/usecase"

	"github.com/gofiber/fiber/v2"
)

// SchedullerHandler ...
type SchedullerHandler struct {
	Handler
}

// SelectAll ...
func (h *SchedullerHandler) ProcessExpiredPackage(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.SchedullerExpiredPackageParameter{}
	uc := usecase.SchedullerUC{ContractUC: h.ContractUC}
	res, err := uc.ProccessExpiredPackage(c, parameter)

	return h.SendResponse(ctx, res, nil, err, 0)
}
