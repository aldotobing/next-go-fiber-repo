package routers

import (
	"time"

	"nextbasis-service-v-0.1/pkg/str"
	"nextbasis-service-v-0.1/server/handlers"
	"nextbasis-service-v-0.1/server/middlewares"

	"github.com/gofiber/fiber/v2"
)

// CustomerAchievementRoutes ...
type CustomerAchievementRoutes struct {
	RouterGroup fiber.Router
	Handler     handlers.Handler
}

// RegisterRoute register CustomerAchievement routes
func (route CustomerAchievementRoutes) RegisterRoute() {
	handler := handlers.CustomerAchievementHandler{Handler: route.Handler}
	handlerQuarter := handlers.CustomerAchievementQuarterHandler{Handler: route.Handler}
	handlerSemester := handlers.CustomerAchievementSemesterHandler{Handler: route.Handler}
	handlerYear := handlers.CustomerAchievementYearHandler{Handler: route.Handler}
	// jwtMiddleware := middlewares.JwtMiddleware{ContractUC: handler.ContractUC}

	r := route.RouterGroup.Group("/api/apps/customer/achievement")
	// r.Use(jwtMiddleware.VerifyUser)
	r.Use(middlewares.SavingContextValue(time.Duration(str.StringToInt(route.Handler.ContractUC.EnvConfig["APP_TIMEOUT"])) * time.Second))
	r.Get("/", handler.FindAll)
	r.Get("/select", handler.SelectAll)

	rQuarter := route.RouterGroup.Group("/api/apps/customer/achievement/quarter")
	// r.Use(jwtMiddleware.VerifyUser)
	rQuarter.Use(middlewares.SavingContextValue(time.Duration(str.StringToInt(route.Handler.ContractUC.EnvConfig["APP_TIMEOUT"])) * time.Second))
	rQuarter.Get("/", handlerQuarter.FindAll)
	rQuarter.Get("/select", handlerQuarter.SelectAll)

	rSemester := route.RouterGroup.Group("/api/apps/customer/achievement/semester")
	// r.Use(jwtMiddleware.VerifyUser)
	rSemester.Use(middlewares.SavingContextValue(time.Duration(str.StringToInt(route.Handler.ContractUC.EnvConfig["APP_TIMEOUT"])) * time.Second))
	rSemester.Get("/", handlerSemester.FindAll)
	rSemester.Get("/select", handlerSemester.SelectAll)

	rYear := route.RouterGroup.Group("/api/apps/customer/achievement/year")
	// r.Use(jwtMiddleware.VerifyUser)
	rYear.Use(middlewares.SavingContextValue(time.Duration(str.StringToInt(route.Handler.ContractUC.EnvConfig["APP_TIMEOUT"])) * time.Second))
	rYear.Get("/", handlerYear.FindAll)
	rYear.Get("/select", handlerYear.SelectAll)
}
