package routers

import (
	"time"

	"nextbasis-service-v-0.1/pkg/str"
	"nextbasis-service-v-0.1/server/handlers"
	"nextbasis-service-v-0.1/server/middlewares"

	"github.com/gofiber/fiber/v2"
)

// FireBaseUIDRoutes ...
type FireBaseUIDRoutes struct {
	RouterGroup fiber.Router
	Handler     handlers.Handler
}

// RegisterRoute register firebaseuid routes
func (route FireBaseUIDRoutes) RegisterRoute() {
	handler := handlers.FireBaseUIDHandler{Handler: route.Handler}
	// jwtMiddleware := middlewares.JwtMiddleware{ContractUC: handler.ContractUC}

	r := route.RouterGroup.Group("/api/apps/firebaseuid")
	// r.Use(jwtMiddleware.VerifyUser)
	r.Use(middlewares.SavingContextValue(time.Duration(str.StringToInt(route.Handler.ContractUC.EnvConfig["APP_TIMEOUT"])) * time.Second))
	r.Get("/", handler.GetFirebaseUIDData)
	r.Get("/sync", handler.SyncData)
}
