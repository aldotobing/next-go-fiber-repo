package routers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"nextbasis-service-v-0.1/pkg/str"
	"nextbasis-service-v-0.1/server/handlers"
	"nextbasis-service-v-0.1/server/middlewares"
)

// TicketDokterChat Routes ...
type TicketDokterChatRoutes struct {
	RouterGroup fiber.Router
	Handler     handlers.Handler
}

// RegisterRoute promo content
func (route TicketDokterChatRoutes) RegisterRoute() {
	handler := handlers.TicketDokterChatHandler{Handler: route.Handler}
	//jwtMiddleware := middlewares.JwtMiddleware{ContractUC: handler.ContractUC}

	r := route.RouterGroup.Group("/api/ticket_dokter_chat")
	//r.Use(jwtMiddleware.VerifyUser)
	r.Use(middlewares.SavingContextValue(time.Duration(str.StringToInt(route.Handler.ContractUC.EnvConfig["APP_TIMEOUT"])) * time.Second))
	r.Get("/ticket_dokter_id/:ticket_dokter_id", handler.FindByTicketDocterID)
	r.Post("/add", handler.Add)
	r.Put("/id/:id", handler.Delete)
}
