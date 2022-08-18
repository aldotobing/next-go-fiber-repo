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
	exchange     = flag.String("exchange", amqpPkg.VerifyMailExchange, "The exchange we will be binding to")
	exchangeType = flag.String("exchange_type", "direct", "Type of exchange we are binding to | topic | direct| etc..")
	queue        = flag.String("queue", amqpPkg.VerifyMail, "Name of the queue that you would like to connect to")
	routingKey   = flag.String("routing_key", amqpPkg.VerifyMailDeadLetter, "queue to route messages to")
	workerName   = flag.String("worker_name", "worker.name", "name to identify worker by")
	verbosity    = flag.Bool("verbos", false, "Set true if you would like to log EVERYTHING")

	// Hold consumer so our go routine can listen to
	// it's done error chan and trigger reconnects
	// if it's ever returned
	// conn *amqpconsumer.Consumer
)

// func main() {
// 	file := false
// 	if file {
// 		f, err := os.OpenFile(*logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
// 		defer f.Close()
// 		if err != nil {
// 			log.Printf("error opening file: %v", err.Error())
// 		}
// 		log.SetOutput(f)
// 	}

// 	// load all config
// 	configs, err := conf.LoadConfigs()
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}
// 	defer configs.DB.Close()

// 	uri = flag.String("uri", configs.EnvConfig["AMQP_URL"], "The rabbitmq endpoint")
// 	conn = amqpconsumer.NewConsumer(*workerName, *uri, *exchange, *exchangeType, *queue)
// 	if err := conn.Connect(); err != nil {
// 		log.Printf("Error: %v", err)
// 	}

// 	deliveries, err := conn.AnnounceQueue(*queue, *routingKey)
// 	if err != nil {
// 		log.Printf("Error when calling AnnounceQueue(): %v", err.Error())
// 	}

// 	uc := usecase.ContractUC{
// 		EnvConfig:   configs.EnvConfig,
// 		DB:          configs.DB,
// 		RedisClient: configs.RedisClient,
// 		JweCred:     configs.JweCred,
// 		JwtCred:     configs.JwtCred,
// 		Minio:       configs.Minio,
// 		Aes:         configs.Aes,
// 		AesFront:    configs.AesFront,
// 		Firestore:   configs.Firestore,
// 		Mandrill:    configs.Mandrill,
// 	}

// 	conn.Handle(deliveries, handler, *threads, *queue, *routingKey, uc)
// }

// func handler(deliveries <-chan amqp.Delivery, uc *usecase.ContractUC) {
// 	ctx := "VerifyMail"

// 	for d := range deliveries {
// 		var formData map[string]interface{}

// 		err := json.Unmarshal(d.Body, &formData)
// 		if err != nil {
// 			log.Printf("Error unmarshaling data: %s", err.Error())
// 		}

// 		logruslogger.Log(logruslogger.InfoLevel, interfacepkg.Marshal(formData), ctx, "begin", uc.ReqID)

// 		uc.ReqID = formData["qid"].(string)
// 		mailUc := usecase.MailUC{ContractUC: uc}
// 		if formData["type"].(string) == "user" {
// 			err = mailUc.UserVerifyMail(context.Background(), formData["id"].(string))
// 		}
// 		if err != nil {
// 			logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "err", uc.ReqID)

// 			// Get fail counter from redis
// 			failCounter := amqpconsumer.FailCounter{}
// 			err = uc.GetFromRedis("amqpFail"+uc.ReqID, &failCounter)
// 			if err != nil {
// 				failCounter = amqpconsumer.FailCounter{
// 					Counter: 1,
// 				}
// 			}

// 			if failCounter.Counter > amqpconsumer.MaxFailCounter {
// 				logruslogger.Log(logruslogger.WarnLevel, strconv.Itoa(failCounter.Counter), ctx, "rejected", uc.ReqID)
// 				d.Reject(false)
// 			} else {
// 				// Save the new counter to redis
// 				failCounter.Counter++
// 				uc.StoreToRedisExp("amqpFail"+uc.ReqID, failCounter, "10m")

// 				logruslogger.Log(logruslogger.WarnLevel, strconv.Itoa(failCounter.Counter), ctx, "failed", uc.ReqID)
// 				d.Nack(false, true)
// 			}
// 		} else {
// 			logruslogger.Log(logruslogger.InfoLevel, string(d.Body), ctx, "success", uc.ReqID)
// 			d.Ack(false)
// 		}
// 	}
// }
