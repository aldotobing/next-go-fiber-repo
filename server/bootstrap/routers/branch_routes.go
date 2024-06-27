package routers

import (
	"time"

	"nextbasis-service-v-0.1/pkg/str"
	"nextbasis-service-v-0.1/server/handlers"
	"nextbasis-service-v-0.1/server/middlewares"

	"github.com/gofiber/fiber/v2"
)

// BranchRoutes ...
type BranchRoutes struct {
	RouterGroup fiber.Router
	Handler     handlers.Handler
}

// RegisterRoute register Branch routes
func (route BranchRoutes) RegisterRoute() {
	handler := handlers.BranchHandler{Handler: route.Handler}
	jwtMiddleware := middlewares.JwtMiddleware{ContractUC: handler.ContractUC}

	r := route.RouterGroup.Group("/api/apps/branch")
	r.Use(jwtMiddleware.VerifyUser)
	r.Use(middlewares.SavingContextValue(time.Duration(str.StringToInt(route.Handler.ContractUC.EnvConfig["APP_TIMEOUT"])) * time.Second))
	r.Get("/", handler.FindAll)
	r.Get("/select", handler.SelectAll)
	r.Get("/id/:id", handler.FindByID)
	r.Get("/except", handler.SelectAll)

	r2 := route.RouterGroup.Group("/api/web/branch")
	r.Use(jwtMiddleware.VerifyUser)
	r2.Use(middlewares.SavingContextValue(time.Duration(str.StringToInt(route.Handler.ContractUC.EnvConfig["APP_TIMEOUT"])) * time.Second))
	r2.Get("/", handler.FindAll)
	r2.Get("/select", handler.SelectAll)
	r2.Get("/id/:id", handler.FindByID)
	r2.Get("/except", handler.SelectAll)
	r2.Put("/id/:id", handler.Update)
	r2.Post("/generatealluser/:id", handler.GenerateAllUser)

}
