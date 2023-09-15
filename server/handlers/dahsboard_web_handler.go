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

// GetDataByGroupID ...
func (h *DashboardWebHandler) GetDataByGroupID(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.DashboardWebParameter{
		StartDate: ctx.Query("start_date"),
		EndDate:   ctx.Query("end_date"),
		GroupID:   ctx.Query("group_id"),
	}

	uc := usecase.DashboardWebUC{ContractUC: h.ContractUC}
	res, err := uc.GetDataByGroupID(c, parameter)

	return h.SendResponse(ctx, res, nil, err, 0)
}

func (h *DashboardWebHandler) GetRegionDetailData(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	uc := usecase.DashboardWebUC{ContractUC: h.ContractUC}
	res, err := uc.GetRegionDetailData(c, models.DashboardWebRegionParameter{
		GroupID:   ctx.Query("group_id"),
		RegionID:  ctx.Query("region_id"),
		StartDate: ctx.Query("start_date"),
		EndDate:   ctx.Query("end_date"),
	})

	return h.SendResponse(ctx, res, nil, err, 0)
}

func (h *DashboardWebHandler) GetUserByRegionDetailData(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	uc := usecase.DashboardWebUC{ContractUC: h.ContractUC}
	res, err := uc.GetUserByRegionDetailData(c, models.DashboardWebRegionParameter{
		BranchID:  ctx.Query("branch_id"),
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
		BranchID:  ctx.Query("branch_id"),
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
		BranchID:  ctx.Query("branch_id"),
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

func (h *DashboardWebHandler) GetAllReportBranchCustomerData(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.DashboardWebBranchParameter{
		BranchID:        ctx.Query("branch_id"),
		CustomerLevelID: ctx.Query("customer_level_id"),
		UserID:          ctx.Query("user_id"),
		StartDate:       ctx.Query("start_date"),
		EndDate:         ctx.Query("end_date"),
	}

	uc := usecase.DashboardWebUC{ContractUC: h.ContractUC}
	res, err := uc.GetAllReportBranchDetailCustomerData(c, parameter)

	type StructObject struct {
		ListObjcet []models.DashboardWebBranchDetail `json:"list_dashboard_branch_customer"`
	}

	ObjcetData := new(StructObject)

	if res != nil {
		ObjcetData.ListObjcet = res
	}

	return h.SendResponse(ctx, res, nil, err, 0)
}

func (h *DashboardWebHandler) GetAllBranchDataByUserID(ctx *fiber.Ctx) error {
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
	res, err := uc.GetAllBranchDataWithUserID(c, parameter)
	if err != nil {
		h.SendResponse(ctx, res, nil, err, 0)
	}

	return h.SendResponse(ctx, res, nil, err, 0)
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
		BranchID:  ctx.Query("branch_id"),
	}

	uc := usecase.DashboardWebUC{ContractUC: h.ContractUC}
	res, err := uc.GetAllDetailCustomerDataWithUserID(c, parameter)
	if err != nil {
		h.SendResponse(ctx, res, nil, err, 0)
	}

	return h.SendResponse(ctx, res, nil, err, 0)
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

func (h *DashboardWebHandler) GetOmzetValueByBranchID(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	branchID := ctx.Query("branch_id")

	parameter := models.DashboardWebBranchParameter{
		StartDate:       ctx.Query("start_date"),
		EndDate:         ctx.Query("end_date"),
		ItemID:          ctx.Query("item_id"),
		ItemCategoryID:  ctx.Query("item_category_id"),
		ItemIDs:         ctx.Query("item_ids"),
		ItemCategoryIDs: ctx.Query("item_category_ids"),
	}

	uc := usecase.DashboardWebUC{ContractUC: h.ContractUC}
	res, err := uc.GetOmzetValueByBranchID(c, parameter, branchID)
	if err != nil {
		h.SendResponse(ctx, res, nil, err, 0)
	}

	return h.SendResponse(ctx, res, nil, err, 0)
}

func (h *DashboardWebHandler) GetOmzetValueByCustomerID(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	customerID := ctx.Query("customer_id")

	parameter := models.DashboardWebBranchParameter{
		StartDate:       ctx.Query("start_date"),
		EndDate:         ctx.Query("end_date"),
		ItemID:          ctx.Query("item_id"),
		ItemCategoryID:  ctx.Query("item_category_id"),
		ItemIDs:         ctx.Query("item_ids"),
		ItemCategoryIDs: ctx.Query("item_category_ids"),
	}

	uc := usecase.DashboardWebUC{ContractUC: h.ContractUC}
	res, err := uc.GetOmzetValueByCustomerID(c, parameter, customerID)
	if err != nil {
		h.SendResponse(ctx, res, nil, err, 0)
	}

	return h.SendResponse(ctx, res, nil, err, 0)
}

func (h *DashboardWebHandler) GetTrackingInvoiceData(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)

	parameter := models.DashboardWebBranchParameter{
		By:              ctx.Query("by"),
		Sort:            ctx.Query("sort"),
		StartDate:       ctx.Query("start_date"),
		EndDate:         ctx.Query("end_date"),
		RegionGroupID:   ctx.Query("region_group_id"),
		RegionID:        ctx.Query("region_id"),
		BranchID:        ctx.Query("branch_id"),
		BranchArea:      ctx.Query("branch_area"),
		CustomerLevelID: ctx.Query("customer_level_id"),
		CustomerTypeID:  ctx.Query("customer_type_id"),
		UserID:          ctx.Query("user_id"),
	}

	uc := usecase.DashboardWebUC{ContractUC: h.ContractUC}
	res, err := uc.GetTrackingInvoiceData(c, parameter)
	if err != nil {
		h.SendResponse(ctx, res, nil, err, 0)
	}

	return h.SendResponse(ctx, res, nil, err, 0)
}
