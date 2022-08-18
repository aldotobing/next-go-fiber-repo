package routers

import (
	"time"

	"nextbasis-service-v-0.1/pkg/str"
	"nextbasis-service-v-0.1/server/handlers"
	"nextbasis-service-v-0.1/server/middlewares"

	"github.com/gofiber/fiber/v2"
)

// AccountOpeningRoutes ...
type AccountOpeningRoutes struct {
	RouterGroup fiber.Router
	Handler     handlers.Handler
}

// RegisterRoute register account opening routes
func (route AccountOpeningRoutes) RegisterRoute() {
	handler := handlers.AccountOpeningHandler{Handler: route.Handler}
	jwtMiddleware := middlewares.JwtMiddleware{ContractUC: handler.ContractUC}

	r := route.RouterGroup.Group("/api/account-opening")
	r.Use(jwtMiddleware.VerifyUser)
	r.Use(middlewares.SavingContextValue(time.Duration(str.StringToInt(route.Handler.ContractUC.EnvConfig["APP_TIMEOUT"])) * time.Second))
	r.Get("/", handler.FindAll)
	r.Get("/select", handler.SelectAll)
	r.Get("/id/:id", handler.FindByID)
	r.Post("/", handler.Add)
	r.Put("/id/:id", handler.Edit)
	r.Delete("/id/:id", handler.Delete)
}
