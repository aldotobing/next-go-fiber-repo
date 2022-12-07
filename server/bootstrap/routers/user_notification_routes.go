package routers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"nextbasis-service-v-0.1/pkg/str"
	"nextbasis-service-v-0.1/server/handlers"
	"nextbasis-service-v-0.1/server/middlewares"
)

type UserNotificationRoutes struct {
	RouterGroup fiber.Router
	Handler     handlers.Handler
}

func (route UserNotificationRoutes) RegisterRoute() {
	handler := handlers.UserNotificationHandler{Handler: route.Handler}
	// jwtMiddleware := middlewares.JwtMiddleware{ContractUC: handler.ContractUC}

	r := route.RouterGroup.Group("/api/notification")
	// r.Use(jwtMiddleware.VerifyUser)
	r.Use(middlewares.SavingContextValue(time.Duration(str.StringToInt(route.Handler.ContractUC.EnvConfig["APP_TIMEOUT"])) * time.Second))
	r.Get("/", handler.FindAll)
	r.Get("/select", handler.SelectAll)
	r.Get("/id/:id", handler.FindByID)
	r.Post("/", handler.Add)

	r2 := route.RouterGroup.Group("/api/apps/notification")
	r2.Use(middlewares.SavingContextValue(time.Duration(str.StringToInt(route.Handler.ContractUC.EnvConfig["APP_TIMEOUT"])) * time.Second))
	r2.Get("/", handler.FindAll)
	r2.Get("/select", handler.SelectAll)
	r2.Get("/id/:id", handler.FindByID)
	r2.Post("/", handler.Add)
	r2.Put("/id/:id", handler.UpdateStatus)
	r2.Put("/all/user_id/:user_id", handler.UpdateAllStatus)
	r2.Delete("/id/:id", handler.DeleteStatus)
	r2.Delete("/all/user_id/:user_id", handler.DeleteAllStatus)
}
