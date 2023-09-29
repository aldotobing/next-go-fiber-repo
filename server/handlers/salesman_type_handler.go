package handlers

import (
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/usecase"
)

// SalesmanTypeHandler ...
type SalesmanTypeHandler struct {
	Handler
}

// SelectAll ...
func (h *SalesmanTypeHandler) SelectAll(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.SalesmanParameter{
		Search: ctx.Query("search"),
		By:     ctx.Query("by"),
		Sort:   ctx.Query("sort"),
	}
	uc := usecase.SalesmanTypeUC{ContractUC: h.ContractUC}
	res, err := uc.SelectAll(c, parameter)
	if err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}

	return h.SendResponse(ctx, res, nil, err, 0)
}
