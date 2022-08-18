package usecase

import (
	"context"
	"errors"

	"nextbasis-service-v-0.1/helper"
	"nextbasis-service-v-0.1/pkg/logruslogger"
	"nextbasis-service-v-0.1/usecase/viewmodel"

	"github.com/rs/xid"
)

// JwtUC ...
type JwtUC struct {
	*ContractUC
}

// GenerateToken ...
func (uc JwtUC) GenerateToken(c context.Context, payload map[string]interface{}, res *viewmodel.JwtVM) (err error) {
	ctx := "JwtUC.GenerateToken"

	deviceID := xid.New().String()
	payload["device_id"] = deviceID
	err = uc.StoreToRedisExp("userDeviceID"+payload["user_id"].(string), deviceID, uc.EnvConfig["TOKEN_EXP_SECRET"]+"m")
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "device_id", c.Value("requestid"))
		return errors.New(helper.InternalServer)
	}

	jwePayload, err := uc.ContractUC.JweCred.Generate(payload)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "jwe", c.Value("requestid"))
		return errors.New(helper.JWT)
	}
	res.Token, res.ExpiredDate, err = uc.ContractUC.JwtCred.GetToken(jwePayload)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "jwt", c.Value("requestid"))
		return errors.New(helper.JWT)
	}
	res.RefreshToken, res.RefreshExpiredDate, err = uc.ContractUC.JwtCred.GetRefreshToken(jwePayload)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "refresh_jwt", c.Value("requestid"))
		return errors.New(helper.JWT)
	}

	return err
}
