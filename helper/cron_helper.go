package helper

import (
	"fmt"
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

	c.AddFunc("CRON_TZ=Asia/Jakarta 0/14 * * * *", func() {
		url := envConfig["APP_BASE_URL"] + "/v1/api/sync/transaction/invoicedata"
		client := &http.Client{}
		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Set("Authorization", "Basic Og==")
		res, _ := client.Do(req)

		if res != nil {
			// fmt.Println("error")
		}

	})

	c.AddFunc("CRON_TZ=Asia/Jakarta 0/2 * * * *", func() {
		url := envConfig["APP_BASE_URL"] + "/v1/api/sync/transaction/voidedrequest"
		client := &http.Client{}
		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Set("Authorization", "Basic Og==")
		res, _ := client.Do(req)

		if res != nil {
			// fmt.Println("error")
		}

	})

	c.AddFunc("CRON_TZ=Asia/Jakarta 0/14 * * * *", func() {
		url := envConfig["APP_BASE_URL"] + "/v1/api/sync/transaction/sodata"
		client := &http.Client{}
		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Set("Authorization", "Basic Og==")
		res, _ := client.Do(req)

		if res != nil {
			// fmt.Println("error")
		}

	})

	c.AddFunc("CRON_TZ=Asia/Jakarta 0/4 * * * *", func() {
		url := envConfig["APP_BASE_URL"] + "/v1/api/sync/transaction/revisedsodata"
		client := &http.Client{}
		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Set("Authorization", "Basic Og==")
		res, _ := client.Do(req)

		if res != nil {
			// fmt.Println("error")
		}

	})

	c.AddFunc("CRON_TZ=Asia/Jakarta 0/5 * * * *", func() {
		url := envConfig["APP_BASE_URL"] + "/v1/api/sync/master/customer"
		client := &http.Client{}
		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Set("Authorization", "Basic Og==")
		res, _ := client.Do(req)

		if res != nil {
			// fmt.Println("error")
		}

	})

	c.AddFunc("CRON_TZ=Asia/Jakarta 0/5 * * * *", func() {
		url := envConfig["APP_BASE_URL"] + "/v1/api/sync/master/salesman"
		client := &http.Client{}
		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Set("Authorization", "Basic Og==")
		res, _ := client.Do(req)

		if res != nil {
			// fmt.Println("error")
		}

	})

	c.AddFunc("CRON_TZ=Asia/Jakarta 0/10 * * * *", func() {
		url := envConfig["APP_BASE_URL"] + "/v1/api/sync/master/item"
		client := &http.Client{}
		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Set("Authorization", "Basic Og==")
		res, _ := client.Do(req)

		if res != nil {
			// fmt.Println("error")
		}

	})

	c.AddFunc("CRON_TZ=Asia/Jakarta 0/10 * * * *", func() {
		url := envConfig["APP_BASE_URL"] + "/v1/api/sync/master/price_list"
		client := &http.Client{}
		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Set("Authorization", "Basic Og==")
		res, _ := client.Do(req)

		if res != nil {
			// fmt.Println("error")
		}

	})

	c.AddFunc("CRON_TZ=Asia/Jakarta 0/10 * * * *", func() {
		url := envConfig["APP_BASE_URL"] + "/v1/api/sync/master/price_list_version"
		client := &http.Client{}
		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Set("Authorization", "Basic Og==")
		res, _ := client.Do(req)

		if res != nil {
			// fmt.Println("error")
		}

	})

	c.AddFunc("CRON_TZ=Asia/Jakarta 0/10 * * * *", func() {
		url := envConfig["APP_BASE_URL"] + "/v1/api/sync/master/item_price"
		client := &http.Client{}
		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Set("Authorization", "Basic Og==")
		res, _ := client.Do(req)

		if res != nil {
			// fmt.Println("error")
		}

	})

	c.AddFunc("CRON_TZ=Asia/Jakarta 15 7 * * *", func() {
		fmt.Println("execute procedure reupdate co modifieddate")
		url := envConfig["APP_BASE_URL"] + "/v1/api/apps/customerorder/reupdate"
		client := &http.Client{}
		req, _ := http.NewRequest("PUT", url, nil)
		req.Header.Set("Authorization", "Basic Og==")
		res, _ := client.Do(req)

		if res != nil {
			// fmt.Println("error")
		}

	})

	c.Start()
}
