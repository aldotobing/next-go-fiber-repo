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

	parameter := models.DashboardWebParameter{
		StartDate: ctx.Query("start_date"),
		EndDate:   ctx.Query("end_date"),
	}

	uc := usecase.DashboardWebUC{ContractUC: h.ContractUC}
	res, err := uc.GetData(c, parameter)

	// for i, object := range res {
	// 	detail, errdetail := uc.GetRegionDetailData(c, models.DashboardWebRegionParameter{
	// 		GroupID:   *object.RegionGroupID,
	// 		StartDate: ctx.Query("start_date"),
	// 		EndDate:   ctx.Query("end_date"),
	// 	})
	// 	if errdetail == nil {
	// 		res[i].DetailData = detail
	// 	}

	// }

	return h.SendResponse(ctx, res, nil, err, 0)
}

func (h *DashboardWebHandler) GetRegionDetailData(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	uc := usecase.DashboardWebUC{ContractUC: h.ContractUC}
	res, err := uc.GetRegionDetailData(c, models.DashboardWebRegionParameter{
		GroupID:   ctx.Query("group_id"),
		StartDate: ctx.Query("start_date"),
		EndDate:   ctx.Query("end_date"),
	})

	return h.SendResponse(ctx, res, nil, err, 0)
}

func (h *DashboardWebHandler) GetBranchCustomerData(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.DashboardWebBranchParameter{
		Search:    ctx.Query("search"),
		Page:      str.StringToInt(ctx.Query("page")),
		Limit:     str.StringToInt(ctx.Query("limit")),
		By:        ctx.Query("by"),
		Sort:      ctx.Query("sort"),
		BarnchID:  ctx.Query("branch_id"),
		StartDate: ctx.Query("start_date"),
		EndDate:   ctx.Query("end_date"),
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

func (h *DashboardWebHandler) GetAllBranchCustomerData(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.DashboardWebBranchParameter{
		Search:    ctx.Query("search"),
		Page:      str.StringToInt(ctx.Query("page")),
		Limit:     str.StringToInt(ctx.Query("limit")),
		By:        ctx.Query("by"),
		Sort:      ctx.Query("sort"),
		BarnchID:  ctx.Query("branch_id"),
		StartDate: ctx.Query("start_date"),
		EndDate:   ctx.Query("end_date"),
	}

	uc := usecase.DashboardWebUC{ContractUC: h.ContractUC}
	res, meta, err := uc.GetAllBranchDetailCustomerData(c, parameter)

	type StructObject struct {
		ListObjcet []models.DashboardWebBranchDetail `json:"list_dashboard_branch_customer"`
	}

	ObjcetData := new(StructObject)

	if res != nil {
		ObjcetData.ListObjcet = res
	}

	return h.SendResponse(ctx, res, meta, err, 0)
}

func (h *DashboardWebHandler) GetAllCustomerDataByUserID(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.DashboardWebBranchParameter{
		Search:    ctx.Query("search"),
		Page:      str.StringToInt(ctx.Query("page")),
		Limit:     str.StringToInt(ctx.Query("limit")),
		By:        ctx.Query("by"),
		Sort:      ctx.Query("sort"),
		StartDate: ctx.Query("start_date"),
		EndDate:   ctx.Query("end_date"),
		UserID:    ctx.Query("user_id"),
	}

	uc := usecase.DashboardWebUC{ContractUC: h.ContractUC}
	res, meta, err := uc.GetAllDetailCustomerDataWithUserID(c, parameter)
	if err != nil {
		h.SendResponse(ctx, res, meta, err, 0)
	}

	if res == nil {
		res = make([]models.DashboardWebBranchDetail, 0)
	}

	return h.SendResponse(ctx, res, meta, err, 0)
}

func (h *DashboardWebHandler) GetOmzetValue(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.DashboardWebBranchParameter{
		StartDate:       ctx.Query("start_date"),
		EndDate:         ctx.Query("end_date"),
		ItemID:          ctx.Query("item_id"),
		ItemCategoryID:  ctx.Query("item_category_id"),
		ItemIDs:         ctx.Query("item_ids"),
		ItemCategoryIDs: ctx.Query("item_category_ids"),
	}

	uc := usecase.DashboardWebUC{ContractUC: h.ContractUC}
	res, err := uc.GetOmzetValue(c, parameter)
	if err != nil {
		h.SendResponse(ctx, res, nil, err, 0)
	}

	return h.SendResponse(ctx, res, nil, err, 0)
}

func (h *DashboardWebHandler) GetOmzetValueByRegionGroupID(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	regionGroupID := ctx.Query("group_id")

	parameter := models.DashboardWebBranchParameter{
		StartDate:       ctx.Query("start_date"),
		EndDate:         ctx.Query("end_date"),
		ItemID:          ctx.Query("item_id"),
		ItemCategoryID:  ctx.Query("item_category_id"),
		ItemIDs:         ctx.Query("item_ids"),
		ItemCategoryIDs: ctx.Query("item_category_ids"),
	}

	uc := usecase.DashboardWebUC{ContractUC: h.ContractUC}
	res, err := uc.GetOmzetValueByRegionGroupID(c, parameter, regionGroupID)
	if err != nil {
		h.SendResponse(ctx, res, nil, err, 0)
	}

	return h.SendResponse(ctx, res, nil, err, 0)
}

func (h *DashboardWebHandler) GetOmzetValueByRegionID(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	regionID := ctx.Query("region_id")

	parameter := models.DashboardWebBranchParameter{
		StartDate:       ctx.Query("start_date"),
		EndDate:         ctx.Query("end_date"),
		ItemID:          ctx.Query("item_id"),
		ItemCategoryID:  ctx.Query("item_category_id"),
		ItemIDs:         ctx.Query("item_ids"),
		ItemCategoryIDs: ctx.Query("item_category_ids"),
	}

	uc := usecase.DashboardWebUC{ContractUC: h.ContractUC}
	res, err := uc.GetOmzetValueByRegionID(c, parameter, regionID)
	if err != nil {
		h.SendResponse(ctx, res, nil, err, 0)
	}

	return h.SendResponse(ctx, res, nil, err, 0)
}
