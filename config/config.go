package config

import (
	"database/sql"
	"fmt"

	"nextbasis-service-v-0.1/pkg/aes"
	"nextbasis-service-v-0.1/pkg/aesfront"
	"nextbasis-service-v-0.1/pkg/aws"
	"nextbasis-service-v-0.1/pkg/fcm"
	"nextbasis-service-v-0.1/pkg/jwe"
	"nextbasis-service-v-0.1/pkg/jwt"
	"nextbasis-service-v-0.1/pkg/mail"
	"nextbasis-service-v-0.1/pkg/mailing"
	"nextbasis-service-v-0.1/pkg/mandrill"
	"nextbasis-service-v-0.1/pkg/whatsapp"

	// miniopkg "nextbasis-service-v-0.1/pkg/minio"
	"log"
	"time"

	miniopkg "nextbasis-service-v-0.1/pkg/minio"
	"nextbasis-service-v-0.1/pkg/mssqldb"
	postgresqlPkg "nextbasis-service-v-0.1/pkg/postgresql"
	"nextbasis-service-v-0.1/pkg/recaptcha"
	redisPkg "nextbasis-service-v-0.1/pkg/redis"
	"nextbasis-service-v-0.1/pkg/str"
	twilioPkg "nextbasis-service-v-0.1/pkg/twilio"

	awsConf "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"

	"cloud.google.com/go/firestore"
	"github.com/go-redis/redis/v7"
	"github.com/joho/godotenv"
	"github.com/minio/minio-go/v7"
)

// Configs ...
type Configs struct {
	EnvConfig    map[string]string
	DB           *sql.DB
	DBMS         *sql.DB
	RedisClient  redisPkg.RedisClient
	JweCred      jwe.Credential
	JwtCred      jwt.Credential
	Aes          aes.Credential
	AesFront     aesfront.Credential
	Minio        *minio.Client
	Aws          aws.AWSS3
	Firestore    *firestore.Client
	Mandrill     mandrill.Credential
	Recaptcha    recaptcha.Credential
	Mail         mail.Connection
	Mailing      mailing.GoMailConfig
	TwilioClient *twilioPkg.Client
	WooWAClient  *whatsapp.Client
	FCM          fcm.Connection
}

var (
	envConfig, _ = godotenv.Read("../.env")
	ImagePath    = envConfig["MINIO_BASE_URL"] + `/` + envConfig["MINIO_DEFAULT_BUCKET"] + `/`
)

// LoadConfigs ...
func LoadConfigs() (res Configs, err error) {
	res.EnvConfig, err = godotenv.Read("../.env")
	if err != nil {
		log.Fatal("Error loading ..env file")
	}

	//postgresql conn
	dbConn := postgresqlPkg.Connection{
		Host:                    res.EnvConfig["DATABASE_HOST"],
		DbName:                  res.EnvConfig["DATABASE_DB"],
		User:                    res.EnvConfig["DATABASE_USER"],
		Password:                res.EnvConfig["DATABASE_PASSWORD"],
		Port:                    str.StringToInt(res.EnvConfig["DATABASE_PORT"]),
		SslMode:                 res.EnvConfig["DATABASE_SSL_MODE"],
		DBMaxConnection:         str.StringToInt(res.EnvConfig["DATABASE_MAX_CONNECTION"]),
		DBMAxIdleConnection:     str.StringToInt(res.EnvConfig["DATABASE_MAX_IDLE_CONNECTION"]),
		DBMaxLifeTimeConnection: str.StringToInt(res.EnvConfig["DATABASE_MAX_LIFETIME_CONNECTION"]),
	}
	res.DB, err = dbConn.Connect()
	if err != nil {
		log.Fatal(err.Error())
	}
	res.DB.SetMaxOpenConns(dbConn.DBMaxConnection)
	res.DB.SetMaxIdleConns(dbConn.DBMAxIdleConnection)
	res.DB.SetConnMaxLifetime(time.Duration(dbConn.DBMaxLifeTimeConnection) * time.Second)

	// SQL server OAO2 DB connection
	dbmsInfo := mssqldb.Connection{
		Server:   res.EnvConfig["MSSQL_HOST"],
		Port:     str.StringToInt(res.EnvConfig["MSSQL_PORT"]),
		User:     res.EnvConfig["MSSQL_USERNAME"],
		Password: res.EnvConfig["MSSQL_PASSWORD"],
		DB:       res.EnvConfig["MSSQL_DB"],
		Debug:    str.StringToBool(res.EnvConfig["MSSQL_DEBUG"]),
	}
	res.DBMS, err = dbmsInfo.Connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	// redis conn
	redisOption := &redis.Options{
		Addr:     res.EnvConfig["REDIS_HOST"],
		Password: res.EnvConfig["REDIS_PASSWORD"],
		DB:       0,
	}
	res.RedisClient = redisPkg.RedisClient{Client: redis.NewClient(redisOption)}

	// jwe
	res.JweCred = jwe.Credential{
		KeyLocation: res.EnvConfig["APP_PRIVATE_KEY_LOCATION"],
		Passphrase:  res.EnvConfig["APP_PRIVATE_KEY_PASSPHRASE"],
	}

	// jwt
	res.JwtCred = jwt.Credential{
		Secret:           res.EnvConfig["TOKEN_SECRET"],
		ExpSecret:        str.StringToInt(res.EnvConfig["TOKEN_EXP_SECRET"]),
		RefreshSecret:    res.EnvConfig["TOKEN_REFRESH_SECRET"],
		RefreshExpSecret: str.StringToInt(res.EnvConfig["TOKEN_EXP_REFRESH_SECRET"]),
	}

	// aes
	res.Aes = aes.Credential{
		Key: res.EnvConfig["AES_KEY"],
	}

	// aes front
	res.AesFront = aesfront.Credential{
		Key: res.EnvConfig["AES_FRONT_KEY"],
		Iv:  res.EnvConfig["AES_FRONT_IV"],
	}

	// Minio connection
	minioInfo := miniopkg.Connection{
		EndSuperhub: res.EnvConfig["MINIO_ENDPOINT"],
		AccessKey:   res.EnvConfig["MINIO_ACCESS_KEY_ID"],
		SecretKey:   res.EnvConfig["MINIO_SECRET_ACCESS_KEY"],
		UseSSL:      str.StringToBool(res.EnvConfig["MINIO_USE_SSL"]),
	}
	res.Minio, err = minioInfo.InitClient()
	if err != nil {
		return res, err
	}

	// Aws connection
	// awsInfo := awspkg.AWSS3{
	// 	Region:    res.EnvConfig["AWS_REGION"],
	// 	AccessKey: res.EnvConfig["AWS_ACCESS_KEY_ID"],
	// 	SecretKey: res.EnvConfig["AWS_SECRET_ACCESS_KEY"],
	// 	Bucket:    res.EnvConfig["BUCKET_NAME"],
	// }
	// res.Aws, err = awsInfo.
	// if err != nil {
	// 	return res, err
	// }

	res.Mandrill = mandrill.Credential{
		Key:      res.EnvConfig["MANDRILL_KEY"],
		FromMail: res.EnvConfig["MANDRILL_FROM_MAIL"],
		FromName: res.EnvConfig["MANDRILL_FROM_NAME"],
	}

	res.Recaptcha = recaptcha.Credential{
		Secret: res.EnvConfig["RECAPTCHA_SECRET"],
	}

	res.Mail = mail.Connection{
		Host:     res.EnvConfig["SMTP_HOST"],
		Port:     str.StringToInt(res.EnvConfig["SMTP_PORT"]),
		Username: res.EnvConfig["SMTP_USERNAME"],
		Password: res.EnvConfig["SMTP_PASSWORD"],
	}

	res.Mailing = mailing.GoMailConfig{
		SMTPHost: res.EnvConfig["SMTP_HOST"],
		SMTPPort: str.StringToInt(res.EnvConfig["SMTP_PORT"]),
		Sender:   res.EnvConfig["SMTP_FROM"],
		Username: res.EnvConfig["SMTP_USERNAME"],
		Password: res.EnvConfig["SMTP_PASSWORD"],
	}

	region := res.EnvConfig["AWS_REGION"]
	// AccessKey := res.EnvConfig["AWS_ACCESS_KEY_ID"]
	// SecretKey := res.EnvConfig["AWS_SECRET_ACCESS_KEY"]
	// Bucket := res.EnvConfig["AWS_BUCKET_NAME"]
	ACredentials := credentials.NewStaticCredentials(res.EnvConfig["AWS_ACCESS_KEY_ID"], res.EnvConfig["AWS_SECRET_ACCESS_KEY"], "")
	checkForbs := true
	awsCon := awsConf.Config{
		Region:                        &region,
		CredentialsChainVerboseErrors: &checkForbs,
		Credentials:                   ACredentials,
	}

	res.Aws = aws.AWSS3{
		AWSConfig: awsCon,
		Region:    res.EnvConfig["AWS_REGION"],
		AccessKey: res.EnvConfig["AWS_ACCESS_KEY_ID"],
		SecretKey: res.EnvConfig["AWS_SECRET_ACCESS_KEY"],
		Bucket:    res.EnvConfig["AWS_BUCKET_NAME"],
	}

	// setup twilio
	res.TwilioClient = twilioPkg.NewTwilioClient(res.EnvConfig["TWILIO_SID"], res.EnvConfig["TWILIO_TOKEN"], res.EnvConfig["TWILIO_SEND_FROM"])
	res.WooWAClient = whatsapp.NewWooWAClient(res.EnvConfig["WA_URL"], res.EnvConfig["WA_KEY"])

	fmt.Printf("+%v", res.Aws)
	res.FCM.APIKey = res.EnvConfig["FCM_API_KEY"]
	fmt.Printf("+%v", res.TwilioClient)
	fmt.Printf("+%v", res.WooWAClient)
	return res, err
}
