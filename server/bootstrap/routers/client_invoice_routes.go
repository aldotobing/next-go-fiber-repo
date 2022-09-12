package routers

import (
	"time"

	"nextbasis-service-v-0.1/pkg/str"
	"nextbasis-service-v-0.1/server/handlers"
	"nextbasis-service-v-0.1/server/middlewares"

	"github.com/gofiber/fiber/v2"
)

// CilentInvoiceRoutes ...
type CilentInvoiceRoutes struct {
	RouterGroup fiber.Router
	Handler     handlers.Handler
}

// RegisterRoute register CilentInvoice routes
func (route CilentInvoiceRoutes) RegisterRoute() {
	handler := handlers.CilentInvoiceHandler{Handler: route.Handler}
	// jwtMiddleware := middlewares.JwtMiddleware{ContractUC: handler.ContractUC}

	r := route.RouterGroup.Group("/api/apps/invoice")
	// r.Use(jwtMiddleware.VerifyUser)
	r.Use(middlewares.SavingContextValue(time.Duration(str.StringToInt(route.Handler.ContractUC.EnvConfig["APP_TIMEOUT"])) * time.Second))
	r.Get("/:customer_id", handler.FindAll)
	r.Get("/select/:customer_id", handler.SelectAll)
	r.Get("/id/:id", handler.FindByID)
	r.Get("/", handler.SelectAll3RD)

}
