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

	// jwtMiddleware := middlewares.JwtMiddleware{ContractUC: handler.ContractUC}

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

	//Customer routes
	customerLogRoutes := routers.CustomerLogRoutes{RouterGroup: apiV1, Handler: handler}
	customerLogRoutes.RegisterRoute()

	//Customer Level routes
	CustomerLevelRoutes := routers.CustomerLevelRoutes{RouterGroup: apiV1, Handler: handler}
	CustomerLevelRoutes.RegisterRoute()

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

	//News Routes
	NewsRoutes := routers.NewsRoutes{RouterGroup: apiV1, Handler: handler}
	NewsRoutes.RegisterRoute()

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

	BranchRoutes := routers.BranchRoutes{RouterGroup: apiV1, Handler: handler}
	BranchRoutes.RegisterRoute()

	FireBaseUIDRoutes := routers.FireBaseUIDRoutes{RouterGroup: apiV1, Handler: handler}
	FireBaseUIDRoutes.RegisterRoute()

	PromoLineRoutes := routers.PromoLineRoutes{RouterGroup: apiV1, Handler: handler}
	PromoLineRoutes.RegisterRoute()

	PromoItemLineRoutes := routers.PromoItemLineRoutes{RouterGroup: apiV1, Handler: handler}
	PromoItemLineRoutes.RegisterRoute()

	UomRoutes := routers.UomRoutes{RouterGroup: apiV1, Handler: handler}
	UomRoutes.RegisterRoute()

	PriceListRoutes := routers.PriceListRoutes{RouterGroup: apiV1, Handler: handler}
	PriceListRoutes.RegisterRoute()

	ItemPriceRoutes := routers.ItemPriceRoutes{RouterGroup: apiV1, Handler: handler}
	ItemPriceRoutes.RegisterRoute()

	UserNotificationRoutes := routers.UserNotificationRoutes{RouterGroup: apiV1, Handler: handler}
	UserNotificationRoutes.RegisterRoute()

	VideoPromoteRoutes := routers.VideoPromoteRoutes{RouterGroup: apiV1, Handler: handler}
	VideoPromoteRoutes.RegisterRoute()

	DashboardWebRoutes := routers.DashboardWebRoutes{RouterGroup: apiV1, Handler: handler}
	DashboardWebRoutes.RegisterRoute()

	SalesmanRoutes := routers.SalesmanRoutes{RouterGroup: apiV1, Handler: handler}
	SalesmanRoutes.RegisterRoute()

	CustomerTypeRoutes := routers.CustomerTypeRoutes{RouterGroup: apiV1, Handler: handler}
	CustomerTypeRoutes.RegisterRoute()

	WebPromoItemLineRoutes := routers.WebPromoItemLineRoutes{RouterGroup: apiV1, Handler: handler}
	WebPromoItemLineRoutes.RegisterRoute()

	WebItemRoutes := routers.WebItemRoutes{RouterGroup: apiV1, Handler: handler}
	WebItemRoutes.RegisterRoute()

	WebItemUomLineRoutes := routers.WebItemUomLineRoutes{RouterGroup: apiV1, Handler: handler}
	WebItemUomLineRoutes.RegisterRoute()

	WebPromoRoutes := routers.WebPromoRoutes{RouterGroup: apiV1, Handler: handler}
	WebPromoRoutes.RegisterRoute()

	WebPromoBonusItemLineRoutes := routers.WebPromoBonusItemLineRoutes{RouterGroup: apiV1, Handler: handler}
	WebPromoBonusItemLineRoutes.RegisterRoute()

	WebRoleGroupRoutes := routers.WebRoleGroupRoutes{RouterGroup: apiV1, Handler: handler}
	WebRoleGroupRoutes.RegisterRoute()

	WebRoleGroupRoleLineRoutes := routers.WebRoleGroupRoleLineRoutes{RouterGroup: apiV1, Handler: handler}
	WebRoleGroupRoleLineRoutes.RegisterRoute()

	WebRoleRoutes := routers.WeebRoleRoutes{RouterGroup: apiV1, Handler: handler}
	WebRoleRoutes.RegisterRoute()

	WebUserRoutes := routers.WebUserRoutes{RouterGroup: apiV1, Handler: handler}
	WebUserRoutes.RegisterRoute()

	UserNotificationDetailRoutes := routers.UserNotificationDetailRoutes{RouterGroup: apiV1, Handler: handler}
	UserNotificationDetailRoutes.RegisterRoute()

	WebCustomerRoutes := routers.WebCustomerRoutes{RouterGroup: apiV1, Handler: handler}
	WebCustomerRoutes.RegisterRoute()

	WebPartnerRoutes := routers.WebPartnerRoutes{RouterGroup: apiV1, Handler: handler}
	WebPartnerRoutes.RegisterRoute()

	WebDoctorRoutes := routers.WebDoctorRoutes{RouterGroup: apiV1, Handler: handler}
	WebDoctorRoutes.RegisterRoute()

	TransactionVaRoutes := routers.TransactionVARoutes{RouterGroup: apiV1, Handler: handler}
	TransactionVaRoutes.RegisterRoute()

	WebSalesmanRoutes := routers.WebSalesmanRoutes{RouterGroup: apiV1, Handler: handler}
	WebSalesmanRoutes.RegisterRoute()

	TicketDokterRoutes := routers.TicketDokterRoutes{RouterGroup: apiV1, Handler: handler}
	TicketDokterRoutes.RegisterRoute()

	ticketDokterChatRoutes := routers.TicketDokterChatRoutes{RouterGroup: apiV1, Handler: handler}
	ticketDokterChatRoutes.RegisterRoute()

	UserCheckinActivityRoute := routers.UserCheckinActivityRoutes{RouterGroup: apiV1, Handler: handler}
	UserCheckinActivityRoute.RegisterRoute()

	WebRegionAreaRoute := routers.WebRegionAreaRoutes{RouterGroup: apiV1, Handler: handler}
	WebRegionAreaRoute.RegisterRoute()

	BroadcastRoute := routers.BroadcastRoutes{RouterGroup: apiV1, Handler: handler}
	BroadcastRoute.RegisterRoute()

	voucherRoute := routers.VoucherRoutes{RouterGroup: apiV1, Handler: handler}
	voucherRoute.RegisterRoute()

	voucherReedeemRoute := routers.VoucherRedeemRoutes{RouterGroup: apiV1, Handler: handler}
	voucherReedeemRoute.RegisterRoute()

	itemOldPriceRoute := routers.ItemOldPriceRoutes{RouterGroup: apiV1, Handler: handler}
	itemOldPriceRoute.RegisterRoute()

	pointRoute := routers.PointRoutes{RouterGroup: apiV1, Handler: handler}
	pointRoute.RegisterRoute()

	pointMaxCustomerRoute := routers.PointMaxCustomerRoutes{RouterGroup: apiV1, Handler: handler}
	pointMaxCustomerRoute.RegisterRoute()

	pointRulesRoute := routers.PointRulesRoutes{RouterGroup: apiV1, Handler: handler}
	pointRulesRoute.RegisterRoute()

	pointPromoRoute := routers.PointPromoRoutes{RouterGroup: apiV1, Handler: handler}
	pointPromoRoute.RegisterRoute()

	couponRoute := routers.CouponRoutes{RouterGroup: apiV1, Handler: handler}
	couponRoute.RegisterRoute()

	couponRedeemRoute := routers.CouponRedeemRoutes{RouterGroup: apiV1, Handler: handler}
	couponRedeemRoute.RegisterRoute()

	lottaryRoute := routers.LottaryRoutes{RouterGroup: apiV1, Handler: handler}
	lottaryRoute.RegisterRoute()
}
