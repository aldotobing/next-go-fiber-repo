package main

import (
	"flag"

	amqpPkg "nextbasis-service-v-0.1/pkg/amqp"
)

var (
	uri          *string
	formURL      = flag.String("form_url", "http://localhost", "The URL that requests are sent to")
	logFile      = flag.String("log_file", "system.log", "The file where errors are logged")
	threads      = flag.Int("threads", 1, "The max amount of go routines that you would like the process to use")
	maxprocs     = flag.Int("max_procs", 1, "The max amount of processors that your application should use")
	paymentsKey  = flag.String("payments_key", "secret", "Access key")
	exchange     = flag.String("exchange", amqpPkg.OtpExchange, "The exchange we will be binding to")
	exchangeType = flag.String("exchange_type", "direct", "Type of exchange we are binding to | topic | direct| etc..")
	queue        = flag.String("queue", amqpPkg.Otp, "Name of the queue that you would like to connect to")
	routingKey   = flag.String("routing_key", amqpPkg.OtpDeadLetter, "queue to route messages to")
	workerName   = flag.String("worker_name", "worker.name", "name to identify worker by")
	verbosity    = flag.Bool("verbos", false, "Set true if you would like to log EVERYTHING")

	// Hold consumer so our go routine can listen to
	// it's done error chan and trigger reconnects
	// if it's ever returned
	// conn      *amqpconsumer.Consumer
	// envConfig map[string]string
)

// func init() {
// 	flag.Parse()
// 	runtime.GOMAXPROCS(*maxprocs)
// 	envConfig, err := godotenv.Read("../.env")
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}
// 	uri = flag.String("uri", envConfig["AMQP_URL"], "The rabbitmq endpoint")
// }

// func main() {
// 	// load all config
// 	configs, err := conf.LoadConfigs()
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}
// 	defer configs.DB.Close()

// 	file := false
// 	// Open a system file to start logging to
// 	if file {
// 		f, err := os.OpenFile(*logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
// 		defer f.Close()
// 		if err != nil {
// 			log.Printf("error opening file: %v", err.Error())
// 		}
// 		log.SetOutput(f)
// 	}

// 	conn := amqpconsumer.NewConsumer(*workerName, *uri, *exchange, *exchangeType, *queue)
// 	if err := conn.Connect(); err != nil {
// 		log.Printf("Error: %v", err)
// 	}

// 	deliveries, err := conn.AnnounceQueue(*queue, *routingKey)
// 	if err != nil {
// 		log.Printf("Error when calling AnnounceQueue(): %v", err.Error())
// 	}

// 	cUC := usecase.ContractUC{
// 		DB:           configs.DB,
// 		RedisClient:  configs.RedisClient,
// 		EnvConfig:    configs.EnvConfig,
// 		Aes:          configs.Aes,
// 		Mandrill:     configs.Mandrill,
// 		TwilioClient: configs.TwilioClient,
// 	}

// 	conn.Handle(deliveries, handler, *threads, *queue, *routingKey, cUC)
// }

// func handler(deliveries <-chan amqp.Delivery, uc *usecase.ContractUC) {
// 	ctx := "RegistrationOtp"

// 	for d := range deliveries {
// 		var formData map[string]interface{}

// 		err := json.Unmarshal(d.Body, &formData)
// 		if err != nil {
// 			log.Printf("Error unmarshaling data: %s", err.Error())
// 		}

// 		if formData["phone"] == nil {
// 			logruslogger.Log(logruslogger.WarnLevel, interfacepkg.Marshal(formData), ctx, "receiver_phone_empty", formData["qid"].(string))
// 			d.Reject(false)
// 			continue
// 		}

// 		if formData["sender"] == nil {
// 			logruslogger.Log(logruslogger.WarnLevel, interfacepkg.Marshal(formData), ctx, "sender_empty", formData["qid"].(string))
// 			d.Reject(false)
// 			continue
// 		}

// 		if formData["otp"] == nil {
// 			logruslogger.Log(logruslogger.WarnLevel, interfacepkg.Marshal(formData), ctx, "otp_code_empty", formData["qid"].(string))
// 			d.Reject(false)
// 			continue
// 		}

// 		if formData["type"] == nil {
// 			logruslogger.Log(logruslogger.WarnLevel, interfacepkg.Marshal(formData), ctx, "type_empty", formData["qid"].(string))
// 			d.Reject(false)
// 			continue
// 		}

// 		uc.ReqID = formData["qid"].(string)
// 		// smsUc := usecase.SmsUC{ContractUC: uc}
// 		// err = smsUc.SendOtpSMS(formData["sender"].(string), formData["phone"].(string), formData["otp"].(string), formData["type"].(string))
// 		// if err != nil {
// 		// 	logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "err", formData["qid"].(string))

// 		// 	// Get fail counter from redis
// 		// 	failCounter := amqpconsumer.FailCounter{}
// 		// 	err = uc.GetFromRedis("amqpFail"+formData["qid"].(string), &failCounter)
// 		// 	if err != nil {
// 		// 		failCounter = amqpconsumer.FailCounter{
// 		// 			Counter: 1,
// 		// 		}
// 		// 	}

// 		// 	if failCounter.Counter > amqpconsumer.MaxFailCounter {
// 		// 		logruslogger.Log(logruslogger.WarnLevel, strconv.Itoa(failCounter.Counter), ctx, "rejected", formData["qid"].(string))
// 		// 		d.Reject(false)
// 		// 	} else {
// 		// 		// Save the new counter to redis
// 		// 		failCounter.Counter++
// 		// 		err = uc.StoreToRedisExp("amqpFail"+formData["qid"].(string), failCounter, "10m")

// 		// 		logruslogger.Log(logruslogger.WarnLevel, strconv.Itoa(failCounter.Counter), ctx, "failed", formData["qid"].(string))
// 		// 		d.Nack(false, true)
// 		// 	}
// 		// } else {
// 		// 	logruslogger.Log(logruslogger.InfoLevel, string(d.Body), ctx, "success", formData["qid"].(string))
// 		// 	d.Ack(false)
// 		// }
// 	}

// 	return
// }

// func handleError(err error, msg string) {
// 	if err != nil {
// 		log.Fatalf("%s: %s", msg, err)
// 	}
// }
