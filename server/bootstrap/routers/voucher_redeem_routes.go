package routers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"nextbasis-service-v-0.1/pkg/str"
	"nextbasis-service-v-0.1/server/handlers"
	"nextbasis-service-v-0.1/server/middlewares"
)

// VoucherRedeemRoutes ...
type VoucherRedeemRoutes struct {
	RouterGroup fiber.Router
	Handler     handlers.Handler
}

// RegisterRoute register Broadcast routes
func (route VoucherRedeemRoutes) RegisterRoute() {
	handler := handlers.VoucherRedeemHandler{Handler: route.Handler}
	//jwtMiddleware := middlewares.JwtMiddleware{ContractUC: handler.ContractUC}

	r := route.RouterGroup.Group("/api/voucher_redeem")
	//r.Use(jwtMiddleware.VerifyUser)
	r.Use(middlewares.SavingContextValue(time.Duration(str.StringToInt(route.Handler.ContractUC.EnvConfig["APP_TIMEOUT"])) * time.Second))
	r.Get("/", handler.FindAll)
	r.Get("/select", handler.SelectAll)
	r.Get("/id/:id", handler.FindByID)
	r.Get("/document_no/:document_no", handler.FindByDocumentNo)
	r.Post("/add", handler.Add)
	r.Post("/add_bulk", handler.AddBulk)
	r.Put("/edit/:id", handler.Update)
	r.Put("/redeem/:id", handler.Redeem)
	r.Delete("/delete/:id", handler.Delete)
}
