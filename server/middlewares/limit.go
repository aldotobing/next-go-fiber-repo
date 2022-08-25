package middlewares

import (
	"net/http"
	"time"

	"nextbasis-service-v-0.1/pkg/functioncaller"
	"nextbasis-service-v-0.1/pkg/logruslogger"
	"nextbasis-service-v-0.1/server/handlers"
	"nextbasis-service-v-0.1/server/requests"
	"nextbasis-service-v-0.1/usecase"

	"github.com/gofiber/fiber/v2"
)

// LimitInit ...
type LimitInit struct {
	*usecase.ContractUC
	MaxLimit float64
	Duration string
}

// LimitOtpRequest ...
func (li LimitInit) LimitOtpRequest(ctx *fiber.Ctx) (err error) {
	handler := handlers.Handler{ContractUC: li.ContractUC}

	_, err = time.ParseDuration(li.Duration)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "verify")
		return handler.SendResponse(ctx, nil, nil, err.Error(), http.StatusUnauthorized)
	}

	input := new(requests.UserOtpRequest)
	if err := ctx.BodyParser(input); err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "verify")
		return handler.SendResponse(ctx, nil, nil, err.Error(), http.StatusUnauthorized)
	}

	var res float64
	key := "OtpRequest" + input.Phone
	err = li.GetFromRedis(key, &res)
	if err != nil {
		li.StoreToRedisExp(key, 1, li.Duration)
		return ctx.Next()
	}

	if res >= li.MaxLimit {
		logruslogger.Log(logruslogger.WarnLevel, "limit", functioncaller.PrintFuncName(), "verify")
		return handler.SendResponse(ctx, nil, nil, "limit", http.StatusUnauthorized)
	}

	li.StoreToRedisExp(key, res+1, li.Duration)

	return ctx.Next()
}
