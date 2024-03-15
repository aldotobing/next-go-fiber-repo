package routers

import (
	"time"

	"nextbasis-service-v-0.1/pkg/str"
	"nextbasis-service-v-0.1/server/handlers"
	"nextbasis-service-v-0.1/server/middlewares"

	"github.com/gofiber/fiber/v2"
)

// DashboardWebRoutes ...
type DashboardWebRoutes struct {
	RouterGroup fiber.Router
	Handler     handlers.Handler
}

// RegisterRoute register DashboardWeb routes
func (route DashboardWebRoutes) RegisterRoute() {
	handler := handlers.DashboardWebHandler{Handler: route.Handler}
	// jwtMiddleware := middlewares.JwtMiddleware{ContractUC: handler.ContractUC}

	r := route.RouterGroup.Group("/api/web/dashboard")
	// r.Use(jwtMiddleware.VerifyUser)
	r.Use(middlewares.SavingContextValue(time.Duration(str.StringToInt(route.Handler.ContractUC.EnvConfig["APP_TIMEOUT"])) * time.Second))
	r.Get("/", handler.GetData)
	r.Get("/group", handler.GetDataByGroupID)
	r.Get("/detail", handler.GetRegionDetailData)
	// r.Get("/detail/total/registered-user", handler.GetUserByRegionDetailData)
	// r.Get("/branch", handler.GetBranchCustomerData)
	// r.Get("/branch/select", handler.GetAllBranchCustomerData)
	// r.Get("/branch/select/report", handler.GetAllReportBranchCustomerData)

	// r.Get("/select/branch", handler.GetAllBranchDataByUserID)
	// r.Get("/select/customer", handler.GetAllCustomerDataByUserID)

	// r.Get("/omzet", handler.GetOmzetValue)
	// r.Get("/omzet/group", handler.GetOmzetValueByRegionGroupID)
	// r.Get("/omzet/region", handler.GetOmzetValueByRegionID)
	// r.Get("/omzet/branch", handler.GetOmzetValueByBranchID)
	// r.Get("/omzet/customer", handler.GetOmzetValueByCustomerID)

	// r.Get("/omzet/graph", handler.GetOmzetValueGraph)

	// r.Get("/tracking/invoice", handler.GetTrackingInvoiceData)
	// r.Get("/virtual_account", handler.GetVirtualAccountData)
}
