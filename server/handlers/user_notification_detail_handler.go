package handlers

import (
	"context"
	"fmt"

	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/usecase"

	"github.com/gofiber/fiber/v2"
)

// UserNotificationDetailHandler ...
type UserNotificationDetailHandler struct {
	Handler
}

// SelectAll ...
func (h *UserNotificationDetailHandler) GetDetail(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.UserNotificationDetailParameter{
		Type:   ctx.Query("type"),
		RowID:  ctx.Query("row_id"),
		Search: "",
		By:     "def.id",
		Sort:   "desc",
	}
	if parameter.Type == "1" {
		fmt.Println("find transaction")
		uc := usecase.CustomerOrderHeaderUC{ContractUC: h.ContractUC}
		res, err := uc.AppsFindByID(c, models.CustomerOrderHeaderParameter{ID: parameter.RowID})
		if *res.ID != "" {
			uc := usecase.CustomerOrderLineUC{ContractUC: h.ContractUC}
			resline, err2 := uc.SelectAll(c, models.CustomerOrderLineParameter{
				HeaderID: *res.ID,
				Search:   "",
				By:       "def.created_date",
				Sort:     "asc"})
			if err2 != nil {
				fmt.Println("error ", err2)
			}
			if resline != nil {
				res.ListLine = resline
			}
		}

		return h.SendResponse(ctx, res, nil, err, 0)
	}

	return h.SendResponse(ctx, nil, nil, nil, 0)
}
