package routers

import (
	"time"

	"nextbasis-service-v-0.1/pkg/str"
	"nextbasis-service-v-0.1/server/handlers"
	"nextbasis-service-v-0.1/server/middlewares"

	"github.com/gofiber/fiber/v2"
)

// CustomerOrderHeaderRoutes ...
type CustomerOrderHeaderRoutes struct {
	RouterGroup fiber.Router
	Handler     handlers.Handler
}

// RegisterRoute register CustomerOrderHeader routes
func (route CustomerOrderHeaderRoutes) RegisterRoute() {
	handler := handlers.CustomerOrderHeaderHandler{Handler: route.Handler}
	// jwtMiddleware := middlewares.JwtMiddleware{ContractUC: handler.ContractUC}

	r := route.RouterGroup.Group("/api/apps/customerorder")
	// r.Use(jwtMiddleware.VerifyUser)
	r.Use(middlewares.SavingContextValue(time.Duration(str.StringToInt(route.Handler.ContractUC.EnvConfig["APP_TIMEOUT"])) * time.Second))
	r.Get("/:customer_id", handler.AppsFindAll)
	r.Get("/select/:customer_id", handler.AppsSelectAll)
	r.Get("/id/:id", handler.AppsFindByID)
	r.Put("/reupdate", handler.ReUpdateDate)

	r2 := route.RouterGroup.Group("/api/rest/customerorder")
	// r.Use(jwtMiddleware.VerifyUser)
	r2.Use(middlewares.SavingContextValue(time.Duration(str.StringToInt(route.Handler.ContractUC.EnvConfig["APP_TIMEOUT"])) * time.Second))
	r2.Get("/select", handler.RestSelectAll)

	r3 := route.RouterGroup.Group("/api/web/customerorder")
	// r.Use(jwtMiddleware.VerifyUser)
	r3.Use(middlewares.SavingContextValue(time.Duration(str.StringToInt(route.Handler.ContractUC.EnvConfig["APP_TIMEOUT"])) * time.Second))
	r3.Get("/", handler.FindAllForWeb)
	r3.Get("/select", handler.SelectAll)
	r3.Get("/id/:id", handler.FindByID)
	r3.Get("/document_no/:document_no", handler.FindByDocumentNo)
}
