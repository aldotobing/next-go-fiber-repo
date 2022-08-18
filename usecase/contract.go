package usecase

import (
	"database/sql"
	"encoding/json"
	"errors"
	"math/rand"
	"strings"
	"time"

	"nextbasis-service-v-0.1/pkg/aes"
	"nextbasis-service-v-0.1/pkg/aesfront"
	"nextbasis-service-v-0.1/pkg/aws"
	"nextbasis-service-v-0.1/pkg/fcm"
	"nextbasis-service-v-0.1/pkg/jwe"
	"nextbasis-service-v-0.1/pkg/jwt"
	"nextbasis-service-v-0.1/pkg/logruslogger"
	"nextbasis-service-v-0.1/pkg/mail"
	"nextbasis-service-v-0.1/pkg/mailing"
	"nextbasis-service-v-0.1/pkg/mandrill"
	"nextbasis-service-v-0.1/pkg/messages"
	"nextbasis-service-v-0.1/pkg/pusher"
	"nextbasis-service-v-0.1/pkg/recaptcha"
	"nextbasis-service-v-0.1/pkg/str"
	twilioHelper "nextbasis-service-v-0.1/pkg/twilio"
	"nextbasis-service-v-0.1/pkg/whatsapp"
	"nextbasis-service-v-0.1/usecase/viewmodel"

	"nextbasis-service-v-0.1/pkg/redis"

	"cloud.google.com/go/firestore"
	ut "github.com/go-playground/universal-translator"
	validator "github.com/go-playground/validator/v10"
	jwtFiber "github.com/gofiber/jwt/v2"
	"github.com/minio/minio-go/v7"
	"github.com/streadway/amqp"
)

var (
	defaultLimit    = 10
	maxLimit        = 50
	defaultSort     = "asc"
	sortWhitelist   = []string{"asc", "desc"}
	passwordLength  = 6
	defaultLastPage = 0

	// DefaultLocation ...
	DefaultLocation = "Asia/Jakarta"
	// DefaultTimezone ...
	DefaultTimezone = "+07:00"

	// AmqpConnection ...
	AmqpConnection *amqp.Connection
	// AmqpChannel ...
	AmqpChannel *amqp.Channel

	// OtpLifetime ...
	OtpLifetime = "8m"
	// MaxOtpSubmitRetry ...
	MaxOtpSubmitRetry = 5.0
)

// ContractUC ...
type ContractUC struct {
	ReqID        string
	UserID       string
	EnvConfig    map[string]string
	DB           *sql.DB
	DBMS         *sql.DB
	TX           *sql.Tx
	Aes          aes.Credential
	AesFront     aesfront.Credential
	AmqpConn     *amqp.Connection
	AmqpChannel  *amqp.Channel
	RedisClient  redis.RedisClient
	JweCred      jwe.Credential
	Validate     *validator.Validate
	Translator   ut.Translator
	JwtCred      jwt.Credential
	JwtConfig    jwtFiber.Config
	AWSS3        aws.AWSS3
	Pusher       pusher.Credential
	Mailing      mailing.GoMailConfig
	Fcm          fcm.Connection
	TwilioClient *twilioHelper.Client
	Mandrill     mandrill.Credential
	Minio        *minio.Client
	Firestore    *firestore.Client
	Recaptcha    recaptcha.Credential
	Mail         mail.Connection
	WhatsApp     *whatsapp.Client
}

func (uc ContractUC) setPaginationParameter(page, limit int, orderBy, sort string, orderByWhiteLists, orderByStringWhiteLists []string) (int, int, int, string, string) {
	if page <= 0 {
		page = 1
	}

	if limit <= 0 || limit > maxLimit {
		limit = defaultLimit
	}

	orderBy = uc.checkWhiteList(orderBy, orderByWhiteLists)
	if str.Contains(orderByStringWhiteLists, orderBy) {
		orderBy = `LOWER(` + orderBy + `)`
	}

	if !str.Contains(sortWhitelist, sort) {
		sort = defaultSort
	}
	offset := (page - 1) * limit

	return offset, limit, page, orderBy, sort
}

func (uc ContractUC) checkWhiteList(orderBy string, whiteLists []string) string {
	for _, whiteList := range whiteLists {
		if orderBy == whiteList {
			return orderBy
		}
	}

	return "def.updated_at"
}

func (uc ContractUC) setPaginationResponse(page, limit, total int) (paginationResponse viewmodel.PaginationVM) {
	var lastPage int

	if total > 0 {
		lastPage = total / limit

		if total%limit != 0 {
			lastPage = lastPage + 1
		}
	} else {
		lastPage = defaultLastPage
	}

	paginationResponse = viewmodel.PaginationVM{
		CurrentPage: page,
		LastPage:    lastPage,
		Total:       total,
		PerPage:     limit,
	}

	return paginationResponse
}

// GetRandomString ...
func (uc ContractUC) GetRandomString(length int) string {
	if length == 0 {
		length = passwordLength
	}

	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZÅÄÖ" +
		"abcdefghijklmnopqrstuvwxyzåäö" +
		"0123456789")

	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	password := b.String()

	return password
}

// LimitRetryByKey ...
func (uc ContractUC) LimitRetryByKey(key string, limit float64) (err error) {
	var count float64
	res := map[string]interface{}{}

	err = uc.RedisClient.GetFromRedis(key, &res)
	if err != nil {
		err = nil
		res = map[string]interface{}{
			"counter": count,
		}
	}
	count = res["counter"].(float64) + 1
	if count > limit {
		uc.RedisClient.RemoveFromRedis(key)

		return errors.New(messages.MaxRetryKey)
	}

	res["counter"] = count
	uc.RedisClient.StoreToRedistWithExpired(key, res, "24h")

	return err
}

// PushToQueue ...
// func (uc ContractUC) PushToQueue(queueBody map[string]interface{}, queueType, deadLetterType string) (err error) {
// 	mqueue := queue.NewQueue(AmqpConnection, AmqpChannel)

// 	_, _, err = mqueue.PushQueueReconnect(uc.EnvConfig["AMQP_URL"], queueBody, queueType, deadLetterType)
// 	if err != nil {
// 		return err
// 	}

// 	return err
// }

func (uc ContractUC) ErrorHandler(ctx, scope string, err error) error {
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, scope, uc.ReqID)
		return err
	}

	return nil
}

// StoreToRedis save data to redis with key key
func (uc ContractUC) StoreToRedis(key string, val interface{}) error {
	ctx := "ContractUC.StoreToRedis"

	b, err := json.Marshal(val)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "json_marshal", uc.ReqID)
		return err
	}

	err = uc.RedisClient.Client.Set(key, string(b), 0).Err()
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "redis_set", uc.ReqID)
		return err
	}

	return err
}

// StoreToRedisExp save data to redis with key and exp time
func (uc ContractUC) StoreToRedisExp(key string, val interface{}, duration string) error {
	ctx := "ContractUC.StoreToRedisExp"

	dur, err := time.ParseDuration(duration)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "parse_duration", uc.ReqID)
		return err
	}

	b, err := json.Marshal(val)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "json_marshal", uc.ReqID)
		return err
	}

	err = uc.RedisClient.Client.Set(key, string(b), dur).Err()
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "redis_set", uc.ReqID)
		return err
	}

	return err
}

// GetFromRedis get value from redis by key
func (uc ContractUC) GetFromRedis(key string, cb interface{}) error {
	ctx := "ContractUC.GetFromRedis"

	res, err := uc.RedisClient.Client.Get(key).Result()
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "redis_get", uc.ReqID)
		return err
	}

	if res == "" {
		logruslogger.Log(logruslogger.WarnLevel, "", ctx, "redis_empty", uc.ReqID)
		return errors.New("[Redis] Value of " + key + " is empty.")
	}

	err = json.Unmarshal([]byte(res), &cb)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "json_unmarshal", uc.ReqID)
		return err
	}

	return err
}

// RemoveFromRedis remove a key from redis
func (uc ContractUC) RemoveFromRedis(key string) error {
	ctx := "ContractUC.RemoveFromRedis"

	err := uc.RedisClient.Client.Del(key).Err()
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "redis_delete", uc.ReqID)
		return err
	}

	return err
}

// LimitByKey global function to limit action using redis
func (uc ContractUC) LimitByKey(key string, limit float64, duration, errMsg string) (err error) {
	var count float64
	resRedis := map[string]interface{}{}
	err = uc.GetFromRedis(key, &resRedis)
	if err != nil {
		err = nil
		resRedis = map[string]interface{}{
			"count": count,
		}
	}

	count = resRedis["count"].(float64) + 1
	if count > limit {
		return errors.New(errMsg)
	}

	resRedis["count"] = count
	uc.StoreToRedisExp(key, resRedis, duration)

	return err
}
