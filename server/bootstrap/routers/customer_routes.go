package routers

import (
	"time"

	"nextbasis-service-v-0.1/pkg/str"
	"nextbasis-service-v-0.1/server/handlers"
	"nextbasis-service-v-0.1/server/middlewares"

	"github.com/gofiber/fiber/v2"
)

// CustomerRoutes ...
type CustomerRoutes struct {
	RouterGroup fiber.Router
	Handler     handlers.Handler
}

// RegisterRoute register Customer routes
func (route CustomerRoutes) RegisterRoute() {
	handler := handlers.CustomerHandler{Handler: route.Handler}
	// jwtMiddleware := middlewares.JwtMiddleware{ContractUC: handler.ContractUC}

	r := route.RouterGroup.Group("/api/apps/customer")
	// r.Use(jwtMiddleware.VerifyUser)
	r.Use(middlewares.SavingContextValue(time.Duration(str.StringToInt(route.Handler.ContractUC.EnvConfig["APP_TIMEOUT"])) * time.Second))
	r.Get("/", handler.FindAll)
	r.Get("/select", handler.SelectAll)
	r.Get("/id/:customer_id", handler.FindByID)
	r.Put("/id/:customer_id", handler.Edit)
	r.Put("/address/id/:customer_id", handler.EditAddress)

	r2 := route.RouterGroup.Group("/api/web/customer")
	// r.Use(jwtMiddleware.VerifyUser)
	r2.Use(middlewares.SavingContextValue(time.Duration(str.StringToInt(route.Handler.ContractUC.EnvConfig["APP_TIMEOUT"])) * time.Second))
	r2.Get("/", handler.FindAll)
	r2.Get("/select", handler.SelectAll)
	r2.Get("/id/:customer_id", handler.FindByID)
	r2.Put("/id/:customer_id", handler.Edit)
	r2.Put("/address/id/:customer_id", handler.EditAddress)

	//AWS UPLOAD IMG ROUTE
	handlerAws := handlers.AwsHandler{Handler: route.Handler}
	// jwtMiddleware := middlewares.JwtMiddleware{ContractUC: handler.ContractUC}

	raws := route.RouterGroup.Group("/api/apps/upload")
	// r.Use(jwtMiddleware.VerifyUser)
	raws.Use(middlewares.SavingContextValue(time.Duration(str.StringToInt(route.Handler.ContractUC.EnvConfig["APP_TIMEOUT"])) * time.Second))
	raws.Post("/", handlerAws.Upload)

}
