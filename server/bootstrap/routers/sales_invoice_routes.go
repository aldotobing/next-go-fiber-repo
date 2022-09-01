package routers

import (
	"time"

	"nextbasis-service-v-0.1/pkg/str"
	"nextbasis-service-v-0.1/server/handlers"
	"nextbasis-service-v-0.1/server/middlewares"

	"github.com/gofiber/fiber/v2"
)

// SalesInvoiceRoutes ...
type SalesInvoiceRoutes struct {
	RouterGroup fiber.Router
	Handler     handlers.Handler
}

// RegisterRoute register SalesInvoice routes
func (route SalesInvoiceRoutes) RegisterRoute() {
	handler := handlers.SalesInvoiceHandler{Handler: route.Handler}
	// jwtMiddleware := middlewares.JwtMiddleware{ContractUC: handler.ContractUC}

	r := route.RouterGroup.Group("/api/apps/sales_invoice")
	// r.Use(jwtMiddleware.VerifyUser)
	r.Use(middlewares.SavingContextValue(time.Duration(str.StringToInt(route.Handler.ContractUC.EnvConfig["APP_TIMEOUT"])) * time.Second))
	r.Get("/", handler.FindAll)
	r.Get("/select", handler.SelectAll)
	r.Get("/id/:id", handler.FindByID)
	r.Get("/document_no/:document_no", handler.FindByDocumentNo)
	r.Get("/customer_id/:cust_bill_to_id", handler.FindByCustomerId)
}
