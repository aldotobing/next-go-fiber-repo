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
	r.Get("/detail", handler.GetRegionDetailData)
	r.Get("/branch", handler.GetBranchCustomerData)
	r.Get("/branch/select", handler.GetAllBranchCustomerData)
}
