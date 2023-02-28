package handlers

import (
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/usecase"
)

// CustomerLevelHandler ...
type CustomerLevelHandler struct {
	Handler
}

// SelectAll ...
func (h *CustomerLevelHandler) FindAll(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.CustomerLevelParameter{
		Search: ctx.Query("search"),
		By:     ctx.Query("by"),
		Sort:   ctx.Query("sort"),
	}
	uc := usecase.CustomerLevelUC{ContractUC: h.ContractUC}
	res, err := uc.FindAll(c, parameter)
	if err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusInternalServerError)
	}

	return h.SendResponse(ctx, res, nil, err, http.StatusOK)
}
