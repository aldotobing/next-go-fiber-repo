package routers

import (
	"time"

	"nextbasis-service-v-0.1/pkg/str"
	"nextbasis-service-v-0.1/server/handlers"
	"nextbasis-service-v-0.1/server/middlewares"

	"github.com/gofiber/fiber/v2"
)

// AuthRoutes ...
type AuthRoutes struct {
	RouterGroup fiber.Router
	Handler     handlers.Handler
}

// RegisterRoute register auth routes
func (route AuthRoutes) RegisterRoute() {
	handler := handlers.AuthHandler{Handler: route.Handler}
	jwtMiddleware := middlewares.JwtMiddleware{ContractUC: handler.ContractUC}

	r := route.RouterGroup.Group("/api/auth")
	r.Use(jwtMiddleware.VerifyBasic)
	r.Use(middlewares.SavingContextValue(time.Duration(str.StringToInt(route.Handler.ContractUC.EnvConfig["APP_TIMEOUT"])) * time.Second))
	r.Post("/register", handler.Register)
	r.Post("/otp", handler.SubmitOtpRegister)

	r2 := route.RouterGroup.Group("/api/verify")
	r2.Use(jwtMiddleware.VerifyUser)
	r2.Use(middlewares.SavingContextValue(time.Duration(str.StringToInt(route.Handler.ContractUC.EnvConfig["APP_TIMEOUT"])) * time.Second))
	r2.Post("/", handler.VerifyMail)
	r2.Post("/mail", handler.VerifyUserMail)
}
