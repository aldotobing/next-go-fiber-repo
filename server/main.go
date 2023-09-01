package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"nextbasis-service-v-0.1/helper"
	redisPkg "nextbasis-service-v-0.1/pkg/redis"

	conf "nextbasis-service-v-0.1/config"
	"nextbasis-service-v-0.1/pkg/str"
	"nextbasis-service-v-0.1/server/bootstrap"
	_ "nextbasis-service-v-0.1/server/docs"
	"nextbasis-service-v-0.1/server/middlewares"
	"nextbasis-service-v-0.1/usecase"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	idTranslations "github.com/go-playground/validator/v10/translations/id"
	redis "github.com/go-redis/redis/v7"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/swagger"
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

	// Set up Redis client
	baseClient := redis.NewClient(&redis.Options{
		Addr:     configs.EnvConfig["REDIS_URL"],
		Password: configs.EnvConfig["REDIS_PASSWORD"], // no password set
		DB:       0,                                   // use default DB
	})
	// fmt.Println("Redis pass: " + configs.EnvConfig["REDIS_PASSWORD"])

	redisStorage := &redisPkg.RedisClient{
		Client: baseClient,
	}

	// Wrap base client in your custom RedisClient type
	configs.RedisClient = redisPkg.RedisClient{
		Client: baseClient,
	}

	pong, err := baseClient.Ping().Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	} else {
		log.Printf("Connected to Redis: %v", pong)
	}

	// init validation driver
	validatorInit(&configs)

	// init fiber app
	app := fiber.New(fiber.Config{
		BodyLimit:         str.StringToInt(configs.EnvConfig["FILE_MAX_UPLOAD_SIZE"]),
		ErrorHandler:      middlewares.InternalServer,
		ReduceMemoryUsage: true,
	})
	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "/swagger/doc.json",
		DeepLinking: false,
	}))

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
		Fcm:         configs.FCM,
	}

	boot := bootstrap.Bootstrap{
		App:        app,
		ContractUC: ContractUC,
		Validator:  validatorDriver,
		Translator: translator,
	}
	boot.App.Use(limiter.New(limiter.Config{
		Max: 20,
		// Max:	100,
		Expiration: 1 * time.Second,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.SendStatus(fiber.StatusTooManyRequests)
		},
		//14-06-2023
		Storage:                redisStorage,
		SkipFailedRequests:     false,
		SkipSuccessfulRequests: false,
		LimiterMiddleware:      limiter.SlidingWindow{},
		//----------------
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
	boot.App.Use(compress.New())

	//DISABLE SCHEDULER
	//helper.SetCronJobs()

	go helper.StartDBEventListener(configs)

	boot.RegisterRouters()
	log.Fatal(boot.App.Listen(configs.EnvConfig["APP_HOST"]))
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current working directory: %v", err)
	}
	log.Printf("Current working directory: %s", cwd)
	// log.Fatal(boot.App.ListenTLS(configs.EnvConfig["APP_HOST"], "cert/api.crt", "cert/api.key"))

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
