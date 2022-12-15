package handlers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/usecase"
)

// DashboardWebHandler ...
type DashboardWebHandler struct {
	Handler
}

// FindByID ...
func (h *DashboardWebHandler) GetData(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.DashboardWebParameter{}

	uc := usecase.DashboardWebUC{ContractUC: h.ContractUC}
	res, err := uc.GetData(c, parameter)

	return h.SendResponse(ctx, res, nil, err, 0)
}
