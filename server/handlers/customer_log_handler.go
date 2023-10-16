package handlers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/usecase"
)

// CustomerLogHandler ...
type CustomerLogHandler struct {
	Handler
}

// SelectAll ...
func (h *CustomerLogHandler) SelectAll(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	// user_id: localStorage.getItem('user_id'),
	//   start_date: parameter.start_date,
	//   end_date: parameter.end_date,
	//   region_group_id: parameter.region_group_id,
	//   region_id: parameter.region_id,
	//   branch_id: parameter.branch_id,
	//   customer_level_id: parameter.customer_level_id,
	parameter := models.CustomerLogParameter{
		UserId:          ctx.Query("user_id"),
		StartDate:       ctx.Query("start_date"),
		EndDate:         ctx.Query("end_date"),
		RegionGroupID:   ctx.Query("region_group_id"),
		RegionID:        ctx.Query("region_id"),
		BranchID:        ctx.Query("branch_id"),
		CustomerLevelID: ctx.Query("customer_level_id"),
		Search:          ctx.Query("search"),
		By:              ctx.Query("by"),
		Sort:            ctx.Query("sort"),
	}
	uc := usecase.CustomerLogUC{ContractUC: h.ContractUC}
	res, err := uc.SelectAll(c, parameter)

	return h.SendResponse(ctx, res, nil, err, 0)
}
