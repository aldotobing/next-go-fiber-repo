package main

import (
	"log"
	"net/http"
	"os"
	"time"

	conf "nextbasis-service-v-0.1/config"
	"nextbasis-service-v-0.1/pkg/str"
	"nextbasis-service-v-0.1/server/bootstrap"
	"nextbasis-service-v-0.1/server/middlewares"
	"nextbasis-service-v-0.1/usecase"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	idTranslations "github.com/go-playground/validator/v10/translations/id"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/rs/xid"
)

var (
	validatorDriver *validator.Validate
	uni             *ut.UniversalTranslator
	translator      ut.Translator
	logFormat       = `{"host":"${host}","pid":"${pid}","time":"${time}","req_id":"${locals:requestid}","status":"${status}","method":"${method}","latency":"${latency}","path":"${path}",` +
		`"user_agent":"${ua}","in":"${bytesReceived}", "req_body":"${body}", "out":"${bytesSent}","res_body":"${resBody}"}`
)

func main() {
	os.Setenv("TZ", "Asia/Jakarta")
	// load all config
	configs, err := conf.LoadConfigs()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer configs.DB.Close()
	defer configs.DBMS.Close()

	// init validation driver
	validatorInit(&configs)

	// init fiber app
	app := fiber.New(fiber.Config{
		BodyLimit:    str.StringToInt(configs.EnvConfig["FILE_MAX_UPLOAD_SIZE"]),
		ErrorHandler: middlewares.InternalServer,
	})

	ContractUC := usecase.ContractUC{
		ReqID:       xid.New().String(),
		EnvConfig:   configs.EnvConfig,
		DB:          configs.DB,
		DBMS:        configs.DBMS,
		RedisClient: configs.RedisClient,
		JweCred:     configs.JweCred,
		Validate:    validatorDriver,
		Translator:  translator,
		JwtCred:     configs.JwtCred,
		Minio:       configs.Minio,
		Aes:         configs.Aes,
		AesFront:    configs.AesFront,
		Firestore:   configs.Firestore,
		Mandrill:    configs.Mandrill,
		Recaptcha:   configs.Recaptcha,
		Mail:        configs.Mail,
		Mailing:     configs.Mailing,
		WhatsApp:    configs.WooWAClient,
		AWSS3:       configs.Aws,
	}

	boot := bootstrap.Bootstrap{
		App:        app,
		ContractUC: ContractUC,
		Validator:  validatorDriver,
		Translator: translator,
	}
	boot.App.Use(limiter.New(limiter.Config{
		Max:        100,
		Expiration: 1 * time.Second,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.SendStatus(fiber.StatusTooManyRequests)
		},
	}))
	boot.App.Use(recover.New())
	boot.App.Use(requestid.New())
	boot.App.Use(cors.New(cors.Config{
		AllowOrigins:     configs.EnvConfig["APP_CORS_DOMAIN"],
		AllowMethods:     http.MethodHead + "," + http.MethodGet + "," + http.MethodPost + "," + http.MethodPut + "," + http.MethodPatch + "," + http.MethodDelete,
		AllowHeaders:     "*",
		AllowCredentials: false,
	}))
	boot.App.Use(logger.New(logger.Config{
		Format:     logFormat + "\n",
		TimeFormat: time.RFC3339,
		TimeZone:   "Asia/Jakarta",
	}))

	boot.RegisterRouters()
	log.Fatal(boot.App.Listen(configs.EnvConfig["APP_HOST"]))

}

func validatorInit(configs *conf.Configs) {
	en := en.New()
	id := id.New()
	uni = ut.New(en, id)

	transEN, _ := uni.GetTranslator("en")
	transID, _ := uni.GetTranslator("id")

	validatorDriver = validator.New()

	enTranslations.RegisterDefaultTranslations(validatorDriver, transEN)
	idTranslations.RegisterDefaultTranslations(validatorDriver, transID)

	switch configs.EnvConfig["APP_LOCALE"] {
	case "en":
		translator = transEN
	case "id":
		translator = transID
	}
}
