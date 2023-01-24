package handlers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/str"
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

	for i, object := range res {
		detail, errdetail := uc.GetRegionDetailData(c, models.DashboardWebRegionParameter{GroupID: *object.RegionGroupID})
		if errdetail == nil {
			res[i].DetailData = detail
		}

	}

	return h.SendResponse(ctx, res, nil, err, 0)
}

func (h *DashboardWebHandler) GetBranchCustomerData(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.DashboardWebBranchParameter{
		Search:   ctx.Query("search"),
		Page:     str.StringToInt(ctx.Query("page")),
		Limit:    str.StringToInt(ctx.Query("limit")),
		By:       ctx.Query("by"),
		Sort:     ctx.Query("sort"),
		BarnchID: ctx.Query("branch_id"),
	}

	uc := usecase.DashboardWebUC{ContractUC: h.ContractUC}
	res, meta, err := uc.GetBranchDetailCustomerData(c, parameter)

	type StructObject struct {
		ListObjcet []models.DashboardWebBranchDetail `json:"list_dashboard_branch_customer"`
	}

	ObjcetData := new(StructObject)

	if res != nil {
		ObjcetData.ListObjcet = res
	}

	return h.SendResponse(ctx, res, meta, err, 0)
}
