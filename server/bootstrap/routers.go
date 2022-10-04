package bootstrap

import (
	"net/http"

	"nextbasis-service-v-0.1/server/bootstrap/routers"
	"nextbasis-service-v-0.1/server/handlers"

	"github.com/gofiber/fiber/v2"
)

// RegisterRouters ...
func (boot Bootstrap) RegisterRouters() {
	handler := handlers.Handler{
		FiberApp:   boot.App,
		ContractUC: &boot.ContractUC,
		Validator:  boot.Validator,
		Translator: boot.Translator,
	}

	// Testing
	boot.App.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(http.StatusOK).JSON("work")
	})

	apiV1 := boot.App.Group("/v1")

	// auth routes
	authRoutes := routers.AuthRoutes{RouterGroup: apiV1, Handler: handler}
	authRoutes.RegisterRoute()

	// otp routes
	otpRoutes := routers.OtpRoutes{RouterGroup: apiV1, Handler: handler}
	otpRoutes.RegisterRoute()

	// user routes
	userRoutes := routers.UserRoutes{RouterGroup: apiV1, Handler: handler}
	userRoutes.RegisterRoute()

	// account opening routes
	// AccountOpeningRoutes := routers.AccountOpeningRoutes{RouterGroup: apiV1, Handler: handler}
	// AccountOpeningRoutes.RegisterRoute()

	// role routes
	RoleRoutes := routers.RoleRoutes{RouterGroup: apiV1, Handler: handler}
	RoleRoutes.RegisterRoute()

	// file routes
	fileRoutes := routers.FileRoutes{RouterGroup: apiV1, Handler: handler}
	fileRoutes.RegisterRoute()

	// settings routes
	settingRoutes := routers.SettingRoutes{RouterGroup: apiV1, Handler: handler}
	settingRoutes.RegisterRoute()

	// marital status routes
	MaritalStatusRoutes := routers.MaritalStatusRoutes{RouterGroup: apiV1, Handler: handler}
	MaritalStatusRoutes.RegisterRoute()

	// religion routes
	ReligionRoutes := routers.ReligionRoutes{RouterGroup: apiV1, Handler: handler}
	ReligionRoutes.RegisterRoute()

	// education level routes
	EducationLevelRoutes := routers.EducationLevelRoutes{RouterGroup: apiV1, Handler: handler}
	EducationLevelRoutes.RegisterRoute()

	// income routes
	IncomeRoutes := routers.IncomeRoutes{RouterGroup: apiV1, Handler: handler}
	IncomeRoutes.RegisterRoute()

	// investment purpose routes
	InvestmentPurposeRoutes := routers.InvestmentPurposeRoutes{RouterGroup: apiV1, Handler: handler}
	InvestmentPurposeRoutes.RegisterRoute()

	// line of business routes
	LineOfBusinessRoutes := routers.LineOfBusinessRoutes{RouterGroup: apiV1, Handler: handler}
	LineOfBusinessRoutes.RegisterRoute()

	// occupation routes
	OccupationRoutes := routers.OccupationRoutes{RouterGroup: apiV1, Handler: handler}
	OccupationRoutes.RegisterRoute()

	// residence ownership routes
	ResidenceOwnershipRoutes := routers.ResidenceOwnershipRoutes{RouterGroup: apiV1, Handler: handler}
	ResidenceOwnershipRoutes.RegisterRoute()

	// gender routes
	GenderRoutes := routers.GenderRoutes{RouterGroup: apiV1, Handler: handler}
	GenderRoutes.RegisterRoute()

	// province routes
	ProvinceRoutes := routers.ProvinceRoutes{RouterGroup: apiV1, Handler: handler}
	ProvinceRoutes.RegisterRoute()

	// city routes
	CityRoutes := routers.CityRoutes{RouterGroup: apiV1, Handler: handler}
	CityRoutes.RegisterRoute()

	// district routes
	DistrictRoutes := routers.DistrictRoutes{RouterGroup: apiV1, Handler: handler}
	DistrictRoutes.RegisterRoute()

	// sub district routes
	SubDistrictRoutes := routers.SubDistrictRoutes{RouterGroup: apiV1, Handler: handler}
	SubDistrictRoutes.RegisterRoute()

	//useracount routes
	UserAccountRoutes := routers.UserAccountRoutes{RouterGroup: apiV1, Handler: handler}
	UserAccountRoutes.RegisterRoute()

	//muuser
	MuUserRoutes := routers.MuUserRoutes{RouterGroup: apiV1, Handler: handler}
	MuUserRoutes.RegisterRoute()

	//item routes
	ItemRoutes := routers.ItemRoutes{RouterGroup: apiV1, Handler: handler}
	ItemRoutes.RegisterRoute()

	//ItemCategory Routes
	ItemCategoryRoutes := routers.ItemCategoryRoutes{RouterGroup: apiV1, Handler: handler}
	ItemCategoryRoutes.RegisterRoute()

	//SalesInvoice Routes
	SalesInvoiceRoutes := routers.SalesInvoiceRoutes{RouterGroup: apiV1, Handler: handler}
	SalesInvoiceRoutes.RegisterRoute()

	//Customer routes
	CustomerRoutes := routers.CustomerRoutes{RouterGroup: apiV1, Handler: handler}
	CustomerRoutes.RegisterRoute()

	ShoppingCartRoutes := routers.ShoppingCartRoutes{RouterGroup: apiV1, Handler: handler}
	ShoppingCartRoutes.RegisterRoute()

	CustomerOrderRoutes := routers.CustomerOrderHeaderRoutes{RouterGroup: apiV1, Handler: handler}
	CustomerOrderRoutes.RegisterRoute()

	ItemPromoRoutes := routers.ItemPromoRoutes{RouterGroup: apiV1, Handler: handler}
	ItemPromoRoutes.RegisterRoute()

	//ProductFocusCategory Routes
	ProductFocusCategoryRoutes := routers.ProductFocusCategoryRoutes{RouterGroup: apiV1, Handler: handler}
	ProductFocusCategoryRoutes.RegisterRoute()

	//ItemProductFocus Routes
	ItemProductFocusRoutes := routers.ItemProductFocusRoutes{RouterGroup: apiV1, Handler: handler}
	ItemProductFocusRoutes.RegisterRoute()

	//Invoice Routes
	InvoiceRoutes := routers.CilentInvoiceRoutes{RouterGroup: apiV1, Handler: handler}
	InvoiceRoutes.RegisterRoute()

	//Item MostSold Routes
	ItemMostSoldRoutes := routers.ItemMostSoldRoutes{RouterGroup: apiV1, Handler: handler}
	ItemMostSoldRoutes.RegisterRoute()

	//Item Details Routes
	ItemDetailsRoutes := routers.ItemDetailsRoutes{RouterGroup: apiV1, Handler: handler}
	ItemDetailsRoutes.RegisterRoute()

	//Item Details Routes
	ItemSearchRoutes := routers.ItemSearchRoutes{RouterGroup: apiV1, Handler: handler}
	ItemSearchRoutes.RegisterRoute()

	//Data Sync Routes
	DataSyncRoutes := routers.DataSyncRoutes{RouterGroup: apiV1, Handler: handler}
	DataSyncRoutes.RegisterRoute()

	CustomerTargetRoutes := routers.CustomerTargetRoutes{RouterGroup: apiV1, Handler: handler}
	CustomerTargetRoutes.RegisterRoute()

	CustomerAchievementRoutes := routers.CustomerAchievementRoutes{RouterGroup: apiV1, Handler: handler}
	CustomerAchievementRoutes.RegisterRoute()

	PromoContentRoutes := routers.PromoContentRoutes{RouterGroup: apiV1, Handler: handler}
	PromoContentRoutes.RegisterRoute()
}
