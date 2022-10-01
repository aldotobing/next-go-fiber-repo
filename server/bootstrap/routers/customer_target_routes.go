package routers

import (
	"time"

	"nextbasis-service-v-0.1/pkg/str"
	"nextbasis-service-v-0.1/server/handlers"
	"nextbasis-service-v-0.1/server/middlewares"

	"github.com/gofiber/fiber/v2"
)

// CustomerTargetRoutes ...
type CustomerTargetRoutes struct {
	RouterGroup fiber.Router
	Handler     handlers.Handler
}

// RegisterRoute register CustomerTarget routes
func (route CustomerTargetRoutes) RegisterRoute() {
	handler := handlers.CustomerTargetHandler{Handler: route.Handler}
	handlerQuarter := handlers.CustomerTargetQuarterHandler{Handler: route.Handler}
	handlerYear := handlers.CustomerTargetYearHandler{Handler: route.Handler}
	// jwtMiddleware := middlewares.JwtMiddleware{ContractUC: handler.ContractUC}

	r := route.RouterGroup.Group("/api/apps/customer/target")
	// r.Use(jwtMiddleware.VerifyUser)
	r.Use(middlewares.SavingContextValue(time.Duration(str.StringToInt(route.Handler.ContractUC.EnvConfig["APP_TIMEOUT"])) * time.Second))
	r.Get("/", handler.FindAll)
	r.Get("/select", handler.SelectAll)
	// r.Get("/id/:customer_id", handler.SelectAll)

	rQuarter := route.RouterGroup.Group("/api/apps/customer/target/quarter")
	// r.Use(jwtMiddleware.VerifyUser)
	rQuarter.Use(middlewares.SavingContextValue(time.Duration(str.StringToInt(route.Handler.ContractUC.EnvConfig["APP_TIMEOUT"])) * time.Second))
	rQuarter.Get("/", handlerQuarter.FindAll)
	rQuarter.Get("/select", handlerQuarter.SelectAll)
	// r.Get("/id/:customer_id", handler.SelectAll)

	rYear := route.RouterGroup.Group("/api/apps/customer/target/year")
	// r.Use(jwtMiddleware.VerifyUser)
	rYear.Use(middlewares.SavingContextValue(time.Duration(str.StringToInt(route.Handler.ContractUC.EnvConfig["APP_TIMEOUT"])) * time.Second))
	rYear.Get("/", handlerYear.FindAll)
	rYear.Get("/select", handlerYear.SelectAll)
	// r.Get("/id/:customer_id", handler.SelectAll)
}
