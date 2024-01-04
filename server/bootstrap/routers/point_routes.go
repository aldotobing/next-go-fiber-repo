package routers

import (
	"time"

	"nextbasis-service-v-0.1/pkg/str"
	"nextbasis-service-v-0.1/server/handlers"
	"nextbasis-service-v-0.1/server/middlewares"

	"github.com/gofiber/fiber/v2"
)

// PointRoutes ...
type PointRoutes struct {
	RouterGroup fiber.Router
	Handler     handlers.Handler
}

// RegisterRoute register Customer routes
func (route PointRoutes) RegisterRoute() {
	handler := handlers.PointHandler{Handler: route.Handler}
	// jwtMiddleware := middlewares.JwtMiddleware{ContractUC: handler.ContractUC}

	r := route.RouterGroup.Group("/api/point")
	// r.Use(jwtMiddleware.VerifyUser)
	r.Use(middlewares.SavingContextValue(time.Duration(str.StringToInt(route.Handler.ContractUC.EnvConfig["APP_TIMEOUT"])) * time.Second))
	r.Get("/", handler.FindAll)
	r.Get("/select", handler.SelectAll)
	r.Get("/id/:customer_id", handler.FindByID)
	r.Get("/balance/customer_id/:customer_id", handler.GetBalance)
	r.Get("/balance/all/customer_id/:customer_id", handler.GetBalanceAll)
	r.Post("/", handler.Add)
	r.Put("/:id", handler.Update)
	r.Delete("/:id", handler.Delete)
}
