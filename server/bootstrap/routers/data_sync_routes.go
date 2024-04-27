package routers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"nextbasis-service-v-0.1/pkg/str"
	"nextbasis-service-v-0.1/server/handlers"
	"nextbasis-service-v-0.1/server/middlewares"
)

// DataSyncRoutes ...
type DataSyncRoutes struct {
	RouterGroup fiber.Router
	Handler     handlers.Handler
}

// RegisterRoute register city routes
func (route DataSyncRoutes) RegisterRoute() {
	handler := handlers.DataSyncHandler{Handler: route.Handler}
	//jwtMiddleware := middlewares.JwtMiddleware{ContractUC: handler.ContractUC}

	r := route.RouterGroup.Group("/api/sync/master")
	//r.Use(jwtMiddleware.VerifyUser)
	r.Use(middlewares.SavingContextValue(time.Duration(str.StringToInt(route.Handler.ContractUC.EnvConfig["APP_TIMEOUT"])) * time.Second))
	r.Get("/item", handler.ItemDataSync)
	r.Get("/price_list", handler.PriceListDataSync)
	r.Get("/price_list_version", handler.PriceListVersionDataSync)
	r.Get("/item_price", handler.ItemPriceDataSync)
	r.Get("/customer", handler.CustomerDataSync)
	r.Get("/salesman", handler.SalesmanDataSync)
	r.Get("/item-uom-line", handler.ItemUomLineDataSync)

	transhandler := handlers.TransactionDataSyncHandler{Handler: route.Handler}
	tr := route.RouterGroup.Group("/api/sync/transaction")
	tr.Use(middlewares.SavingContextValue(time.Duration(str.StringToInt(route.Handler.ContractUC.EnvConfig["APP_TIMEOUT"])) * time.Second))
	tr.Get("/voidedrequest", transhandler.CustomerOrderVoidDataSync)
	tr.Get("/invoice/redis/sfa/pull", transhandler.InvoiceSFAPull)
	tr.Get("/invoice/redis/mysm/pull", transhandler.InvoiceMYSMPull)
	tr.Get("/invoice/redis/sfa/sync", transhandler.InvoiceSyncSFA)
	tr.Get("/invoice/redis/reserve", transhandler.InvoiceReserveSyncGetRedis)
	tr.Get("/invoice/redis/replace", transhandler.InvoiceSyncSFA)
	tr.Get("/invoice/redis/pointonly", transhandler.InvoiceSyncGetRedisPointOnly)
	tr.Get("/return_invoicedata", transhandler.ReturnInvoiceSync)
	tr.Get("/invoicedata/undone", transhandler.UndoneDataSync)
	tr.Get("/sodata/pull", transhandler.SalesOrderCustomerPullData)
	tr.Get("/sodata/push", transhandler.SalesOrderCustomerPushData)
	tr.Get("/sodata/sendfcmmessage", transhandler.SalesOrderCustomerSendFcmMessage)
	tr.Get("/sodata/sendsamesmanwa", transhandler.SalesOrderCustomerSendSalesmanWa)
	tr.Get("/revisedsodata", transhandler.SalesOrderCustomerRevisedSync)

	tr.Get("/invoice/redis/sfa/undone/pull", transhandler.UndoneInvoiceSFAPull)
	tr.Get("/invoice/redis/sfa/undone/sync", transhandler.UndoneInvoiceSyncSFA)

}
