package routers

import (
	"time"

	"nextbasis-service-v-0.1/pkg/str"
	"nextbasis-service-v-0.1/server/handlers"
	"nextbasis-service-v-0.1/server/middlewares"

	"github.com/gofiber/fiber/v2"
)

// TransactionVARoutes ...
type TransactionVARoutes struct {
	RouterGroup fiber.Router
	Handler     handlers.Handler
}

// RegisterRoute register Customer routes
func (route TransactionVARoutes) RegisterRoute() {
	handler := handlers.TransactionVAHandler{Handler: route.Handler}
	jwtMiddleware := middlewares.JwtMiddleware{ContractUC: handler.ContractUC}

	r := route.RouterGroup.Group("/api/transaction_va")
	// r.Use(jwtMiddleware.VerifyUser)
	r.Use(middlewares.SavingContextValue(time.Duration(str.StringToInt(route.Handler.ContractUC.EnvConfig["APP_TIMEOUT"])) * time.Second))
	r.Get("/", handler.FindAll)
	r.Get("/select", handler.SelectAll)
	r.Get("/id/:partner_id", handler.FindByID)
	r.Put("/id/:partner_id", handler.Edit)
	r.Post("/request_va", handler.Add)
	r.Get("/authsah521", handler.GetSah)
	// r.Get("/tesgen", handler.GetTransactionByVaCode)

	r2 := route.RouterGroup.Group("/api/mysido/va/inquiry")
	r2.Use(jwtMiddleware.VerifySignature)
	r2.Use(middlewares.SavingContextValue(time.Duration(str.StringToInt(route.Handler.ContractUC.EnvConfig["APP_TIMEOUT"])) * time.Second))
	r2.Post("/", handler.GetTransactionByVaCode)
	r2.Post("/paid", handler.PaidTransactionByVaCode)

}
