package middlewares

import (
	"crypto/sha512"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"nextbasis-service-v-0.1/helper"
	"nextbasis-service-v-0.1/pkg/functioncaller"
	"nextbasis-service-v-0.1/pkg/interfacepkg"
	"nextbasis-service-v-0.1/pkg/logruslogger"
	"nextbasis-service-v-0.1/server/handlers"
	"nextbasis-service-v-0.1/usecase"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// JwtMiddleware ...
type JwtMiddleware struct {
	*usecase.ContractUC
}

// VerifyBasic ...
func (jwtMiddleware JwtMiddleware) VerifyBasic(ctx *fiber.Ctx) (err error) {
	basic := base64.StdEncoding.EncodeToString([]byte(jwtMiddleware.EnvConfig["BASIC_USERNAME"] + ":" + jwtMiddleware.EnvConfig["BASIC_PASSWORD"]))

	header := ctx.Get("Authorization")
	if !strings.Contains(header, "Basic") {
		logruslogger.Log(logruslogger.WarnLevel, helper.HeaderNotPresent, functioncaller.PrintFuncName(), "middleware-jwt-header")
		return errors.New(helper.HeaderNotPresent)
	}
	token := strings.Replace(header, "Basic ", "", -1)
	if token != basic {
		logruslogger.Log(logruslogger.WarnLevel, basic, functioncaller.PrintFuncName(), "invalid-token")
		return errors.New(helper.UnexpectedClaims)
	}

	return ctx.Next()
}

// verify jwt middleware
func (jwtMiddleware JwtMiddleware) verify(ctx *fiber.Ctx, role string) (res map[string]interface{}, err error) {
	claims := &jwt.StandardClaims{}

	header := ctx.Get("Authorization")
	if !strings.Contains(header, "Bearer") {
		logruslogger.Log(logruslogger.WarnLevel, helper.HeaderNotPresent, functioncaller.PrintFuncName(), "middleware-jwt-header")
		return res, errors.New(helper.HeaderNotPresent)
	}

	//check claims and signing method
	token := strings.Replace(header, "Bearer ", "", -1)
	_, err = jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		if jwt.SigningMethodHS256 != token.Method {
			logruslogger.Log(logruslogger.WarnLevel, helper.UnexpectedSigningMethod, functioncaller.PrintFuncName(), "middleware-jwt-checkSigningMethod")
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		secret := jwtMiddleware.EnvConfig["TOKEN_SECRET"]
		return []byte(secret), nil
	})
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "middleware-jwt-checkClaims")
		return res, errors.New(helper.UnexpectedClaims)
	}

	//check token live time
	if claims.ExpiresAt < time.Now().Unix() {
		logruslogger.Log(logruslogger.WarnLevel, helper.ExpiredToken, functioncaller.PrintFuncName(), "middleware-jwt-checkTokenLiveTime")
		return res, errors.New(helper.ExpiredToken)
	}

	//jwe roll back encrypted id
	res, err = jwtMiddleware.JweCred.Rollback(claims.Id)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, helper.Unauthorized, functioncaller.PrintFuncName(), "pkg-jwe-rollback")
		return res, errors.New(helper.Unauthorized)
	}
	if res == nil {
		logruslogger.Log(logruslogger.WarnLevel, helper.Unauthorized, functioncaller.PrintFuncName(), "pkg-jwe-resultNil")
		return res, errors.New(helper.Unauthorized)
	}

	if role != "" && fmt.Sprintf("%v", res["role"]) != role {
		logruslogger.Log(logruslogger.WarnLevel, helper.InvalidRole, functioncaller.PrintFuncName(), "pkg-jwe-resultNil")
		return res, errors.New(helper.InvalidRole)
	}

	logruslogger.Log(logruslogger.InfoLevel, interfacepkg.Marshal(res), functioncaller.PrintFuncName(), "user", ctx.Locals("requestid"))

	return res, nil
}

// VerifyUser jwt middleware
func (jwtMiddleware JwtMiddleware) VerifyUser(ctx *fiber.Ctx) (err error) {
	handler := handlers.Handler{ContractUC: jwtMiddleware.ContractUC}

	jweRes, err := jwtMiddleware.verify(ctx, "")
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "verify")
		return handler.SendResponse(ctx, nil, nil, err.Error(), http.StatusUnauthorized)
	}

	// set id to uce case contract
	ctx.Locals("user_id", jweRes["user_id"].(string))

	return ctx.Next()
}

// VerifyBasic ...
func (jwtMiddleware JwtMiddleware) VerifySignature(ctx *fiber.Ctx) (err error) {
	// fmt.Println(ctx.Body())
	basicCode := `888`
	codeBase := string(ctx.Body()[:])
	MandiriCode := `BMRI_SIDO`
	finalCode := basicCode + `:` + codeBase + `:` + MandiriCode
	sha_512 := sha512.Sum512([]byte(finalCode))
	basic := fmt.Sprintf("%x", sha_512)
	// fmt.Println("res ", codeBase)
	// fmt.Println("code nya" + basic)

	header := ctx.Get("Authorization")
	if !strings.Contains(header, "signature") {
		logruslogger.Log(logruslogger.WarnLevel, helper.HeaderNotPresent, functioncaller.PrintFuncName(), "middleware-jwt-header")
		return errors.New(helper.HeaderNotPresent)
	}
	token := strings.Replace(header, "signature ", "", -1)
	if token != basic {
		logruslogger.Log(logruslogger.WarnLevel, basic, functioncaller.PrintFuncName(), "invalid-token")
		return errors.New(helper.UnexpectedClaims)
	}

	return ctx.Next()
}
