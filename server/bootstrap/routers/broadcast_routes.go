package routers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"nextbasis-service-v-0.1/pkg/str"
	"nextbasis-service-v-0.1/server/handlers"
	"nextbasis-service-v-0.1/server/middlewares"
)

// BroadcastRoutes ...
type BroadcastRoutes struct {
	RouterGroup fiber.Router
	Handler     handlers.Handler
}

// RegisterRoute register Broadcast routes
func (route BroadcastRoutes) RegisterRoute() {
	handler := handlers.BroadcastHandler{Handler: route.Handler}
	//jwtMiddleware := middlewares.JwtMiddleware{ContractUC: handler.ContractUC}

	r := route.RouterGroup.Group("/api/broadcast")

	//r.Use(jwtMiddleware.VerifyUser)
	r.Use(middlewares.SavingContextValue(time.Duration(str.StringToInt(route.Handler.ContractUC.EnvConfig["APP_TIMEOUT"])) * time.Minute))
	r.Get("/", handler.FindAll)
	r.Get("/select", handler.SelectAll)
	r.Get("/id/:id", handler.FindByID)
	r.Post("/", handler.Broadcast)
	r.Post("/execute/:id", handler.BroadcastWithID)
	r.Post("/run/scheduler", handler.BroadcastWithScheduler)
	r.Post("/add", handler.Add)
	r.Put("/edit/:id", handler.Update)
	r.Delete("/delete/:id", handler.Delete)
}
