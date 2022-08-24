package routers

import (
	"time"

	"nextbasis-service-v-0.1/pkg/str"
	"nextbasis-service-v-0.1/server/handlers"
	"nextbasis-service-v-0.1/server/middlewares"

	"github.com/gofiber/fiber/v2"
)

// ShoppingCartRoutes ...
type ShoppingCartRoutes struct {
	RouterGroup fiber.Router
	Handler     handlers.Handler
}

// RegisterRoute register ShoppingCart routes
func (route ShoppingCartRoutes) RegisterRoute() {
	handler := handlers.ShoppingCartHandler{Handler: route.Handler}
	// jwtMiddleware := middlewares.JwtMiddleware{ContractUC: handler.ContractUC}

	r := route.RouterGroup.Group("/api/apps/shopping_cart")
	// r.Use(jwtMiddleware.VerifyUser)
	r.Use(middlewares.SavingContextValue(time.Duration(str.StringToInt(route.Handler.ContractUC.EnvConfig["APP_TIMEOUT"])) * time.Second))
	r.Get("/", handler.FindAll)
	r.Get("/select/:customer_id", handler.SelectAll)
	r.Get("/id/:id", handler.FindByID)
	r.Post("/", handler.Add)
	r.Post("/checkout", handler.CheckOut)
	r.Put("/:customer_id", handler.MultipleEdit)
	r.Delete("/:customer_id", handler.MultipleDelete)
	r.Delete("/id/:id", handler.Delete)
	// r.Post("/databreakdown/", handler.AddAll)
}
