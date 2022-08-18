package helper

import (
	"net/http"

	"github.com/robfig/cron/v3"
)

func SetCronTest() {
	c := cron.New()
	c.AddFunc("CRON_TZ=Asia/Jakarta 30 01 * * *", func() {
		url := "https://api.v3.gostellar.id/v1/api/scheduller/expired_package"
		client := &http.Client{}
		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Set("Authorization", "Basic Og==")
		res, _ := client.Do(req)

		if res != nil {
			// log.Fatal(err)
		}

	})
	c.Start()
}
