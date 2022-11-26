package routers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"nextbasis-service-v-0.1/pkg/str"
	"nextbasis-service-v-0.1/server/handlers"
	"nextbasis-service-v-0.1/server/middlewares"
)

// PromoItemLine Routes ...
type PromoItemLineRoutes struct {
	RouterGroup fiber.Router
	Handler     handlers.Handler
}

// RegisterRoute promo content
func (route PromoItemLineRoutes) RegisterRoute() {
	handler := handlers.PromoItemLineHandler{Handler: route.Handler}
	//jwtMiddleware := middlewares.JwtMiddleware{ContractUC: handler.ContractUC}

	r := route.RouterGroup.Group("/api/apps/promo_item_line")
	//r.Use(jwtMiddleware.VerifyUser)
	r.Use(middlewares.SavingContextValue(time.Duration(str.StringToInt(route.Handler.ContractUC.EnvConfig["APP_TIMEOUT"])) * time.Second))
	r.Get("/", handler.FindAll)
	r.Get("/select", handler.SelectAll)
	r.Post("/add", handler.Add)
}
