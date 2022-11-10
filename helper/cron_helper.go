package helper

import (
	"net/http"

	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
)

func SetCronJobs() {
	c := cron.New()
	var envConfig, _ = godotenv.Read("../.env")
	c.AddFunc("CRON_TZ=Asia/Jakarta 0/2 * * * *", func() {
		url := envConfig["APP_BASE_URL"] + "/v1/api/apps/firebaseuid/sync"
		client := &http.Client{}
		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Set("Authorization", "Basic Og==")
		res, _ := client.Do(req)

		if res != nil {
			// fmt.Println("error")
		}

	})
	c.Start()
}
