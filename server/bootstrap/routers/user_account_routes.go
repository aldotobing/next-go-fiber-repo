package routers

import (
	"time"

	"nextbasis-service-v-0.1/pkg/str"
	"nextbasis-service-v-0.1/server/handlers"
	"nextbasis-service-v-0.1/server/middlewares"

	"github.com/gofiber/fiber/v2"
)

type UserAccountRoutes struct {
	RouterGroup fiber.Router
	Handler     handlers.Handler
}

func (route UserAccountRoutes) RegisterRoute() {
	handler := handlers.UserAccountHandler{Handler: route.Handler}

	jwtMiddleware := middlewares.JwtMiddleware{ContractUC: handler.ContractUC}

	r := route.RouterGroup.Group("/api/akun-register-login")
	r.Use(middlewares.SavingContextValue(time.Duration(str.StringToInt(route.Handler.ContractUC.EnvConfig["APP_TIMEOUT"])) * time.Second))

	r.Post("/login", handler.Login)

	r2 := route.RouterGroup.Group("/api/akun-verfikasi")
	r2.Use(jwtMiddleware.VerifyUser)
	r2.Use(middlewares.SavingContextValue(time.Duration(str.StringToInt(route.Handler.ContractUC.EnvConfig["APP_TIMEOUT"])) * time.Second))
	r2.Post("/submitOtp", handler.SubmitOtp)
	r2.Post("/resendOtp", handler.ResendOtp)

}
