package routers

import (
	"time"

	"nextbasis-service-v-0.1/pkg/str"
	"nextbasis-service-v-0.1/server/handlers"
	"nextbasis-service-v-0.1/server/middlewares"

	"github.com/gofiber/fiber/v2"
)

type MuUserRoutes struct {
	RouterGroup fiber.Router
	Handler     handlers.Handler
}

func (route MuUserRoutes) RegisterRoute() {
	handler := handlers.MuUserHandler{Handler: route.Handler}

	// jwtMiddleware := middlewares.JwtMiddleware{ContractUC: handler.ContractUC}

	r := route.RouterGroup.Group("/api/mu_user")
	r.Use(middlewares.SavingContextValue(time.Duration(str.StringToInt(route.Handler.ContractUC.EnvConfig["APP_TIMEOUT"])) * time.Second))
	// r.Use(jwtMiddleware.VerifyUser)
	r.Get("/", handler.FindAll)
	r.Get("/id/:id", handler.FindByID)

}
