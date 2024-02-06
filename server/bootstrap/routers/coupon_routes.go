package routers

import (
	"time"

	"nextbasis-service-v-0.1/pkg/str"
	"nextbasis-service-v-0.1/server/handlers"
	"nextbasis-service-v-0.1/server/middlewares"

	"github.com/gofiber/fiber/v2"
)

// CouponRoutes ...
type CouponRoutes struct {
	RouterGroup fiber.Router
	Handler     handlers.Handler
}

// RegisterRoute register Customer routes
func (route CouponRoutes) RegisterRoute() {
	handler := handlers.CouponHandler{Handler: route.Handler}
	// jwtMiddleware := middlewares.JwtMiddleware{ContractUC: handler.ContractUC}

	r := route.RouterGroup.Group("/api/coupon")
	// r.Use(jwtMiddleware.VerifyUser)
	r.Use(middlewares.SavingContextValue(time.Duration(str.StringToInt(route.Handler.ContractUC.EnvConfig["APP_TIMEOUT"])) * time.Second))
	r.Get("/", handler.FindAll)
	r.Get("/select", handler.SelectAll)
	r.Get("/id/:id", handler.FindByID)
	r.Post("/", handler.Add)
	r.Post("/photo", handler.Photo)
	r.Put("/:id", handler.Update)
	r.Delete("/:id", handler.Delete)
}
