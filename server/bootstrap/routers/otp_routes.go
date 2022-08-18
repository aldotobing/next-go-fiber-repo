package routers

import (
	"time"

	"nextbasis-service-v-0.1/pkg/str"
	"nextbasis-service-v-0.1/server/handlers"
	"nextbasis-service-v-0.1/server/middlewares"

	"github.com/gofiber/fiber/v2"
)

// OtpRoutes ...
type OtpRoutes struct {
	RouterGroup fiber.Router
	Handler     handlers.Handler
}

// RegisterRoute register otp routes
func (route OtpRoutes) RegisterRoute() {
	handler := handlers.OtpHandler{Handler: route.Handler}
	limitMiddleware := middlewares.LimitInit{ContractUC: handler.ContractUC, MaxLimit: 3, Duration: "24h"}
	jwtMiddleware := middlewares.JwtMiddleware{ContractUC: handler.ContractUC}

	r := route.RouterGroup.Group("/api/otp")
	r.Use(jwtMiddleware.VerifyUser)
	r.Use(limitMiddleware.LimitOtpRequest)
	r.Use(middlewares.SavingContextValue(time.Duration(str.StringToInt(route.Handler.ContractUC.EnvConfig["APP_TIMEOUT"])) * time.Second))
	r.Post("/resend", handler.ResendOtpRequest)
}
