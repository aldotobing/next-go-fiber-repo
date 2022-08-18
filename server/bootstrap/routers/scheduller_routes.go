package routers

import (
	"time"

	"nextbasis-service-v-0.1/pkg/str"
	"nextbasis-service-v-0.1/server/handlers"
	"nextbasis-service-v-0.1/server/middlewares"

	"github.com/gofiber/fiber/v2"
)

// SchedullerRoutes ...
type SchedullerRoutes struct {
	RouterGroup fiber.Router
	Handler     handlers.Handler
}

// RegisterRoute register Scheduller routes
func (route SchedullerRoutes) RegisterRoute() {
	handler := handlers.SchedullerHandler{Handler: route.Handler}
	jwtMiddleware := middlewares.JwtMiddleware{ContractUC: handler.ContractUC}
	r := route.RouterGroup.Group("/api/scheduller")
	r.Use(middlewares.SavingContextValue(time.Duration(str.StringToInt(route.Handler.ContractUC.EnvConfig["APP_TIMEOUT"])) * time.Second))
	r.Use(jwtMiddleware.VerifyBasic)
	r.Get("/expired_package", handler.ProcessExpiredPackage)
}
