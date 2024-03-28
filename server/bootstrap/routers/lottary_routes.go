package routers

import (
	"time"

	"nextbasis-service-v-0.1/pkg/str"
	"nextbasis-service-v-0.1/server/handlers"
	"nextbasis-service-v-0.1/server/middlewares"

	"github.com/gofiber/fiber/v2"
)

// LottaryRoutes ...
type LottaryRoutes struct {
	RouterGroup fiber.Router
	Handler     handlers.Handler
}

// RegisterRoute register Customer routes
func (route LottaryRoutes) RegisterRoute() {
	handler := handlers.LottaryHandler{Handler: route.Handler}
	// jwtMiddleware := middlewares.JwtMiddleware{ContractUC: handler.ContractUC}

	r := route.RouterGroup.Group("/api/lottary")
	// r.Use(jwtMiddleware.VerifyUser)
	r.Use(middlewares.SavingContextValue(time.Duration(str.StringToInt(route.Handler.ContractUC.EnvConfig["APP_TIMEOUT"])) * time.Second))
	r.Get("/paginate", handler.FindAll)
	r.Get("/id/:id", handler.FindByID)
	r.Get("/select", handler.SelectAll)
	r.Get("/reportdata", handler.ReportData)
	r.Post("/import", handler.Import)
	r.Put("/:id", handler.Update)
	r.Delete("/:id", handler.Delete)
}
