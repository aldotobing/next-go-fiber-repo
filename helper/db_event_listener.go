package helper

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/lib/pq"
	"nextbasis-service-v-0.1/config"
	"nextbasis-service-v-0.1/pkg/str"
)

func StartDBEventListener(configs config.Configs) {
	connStr := fmt.Sprintf("host=%s dbname=%s user=%s password=%s port=%d sslmode=disable",
		configs.EnvConfig["DATABASE_HOST"],
		configs.EnvConfig["DATABASE_DB"],
		configs.EnvConfig["DATABASE_USER"],
		configs.EnvConfig["DATABASE_PASSWORD"],
		str.StringToInt(configs.EnvConfig["DATABASE_PORT"]))

	reportProblem := func(ev pq.ListenerEventType, err error) {
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	listener := pq.NewListener(connStr, 10*time.Second, time.Minute, reportProblem)

	err := listener.Listen("customer_changes")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Listening for notifications on 'customer' table changes...")
	for {
		select {
		case n := <-listener.Notify:
			fmt.Println("Received notification:", n.Extra)
			// Extract the ID from the notification
			splitMsg := strings.Split(n.Extra, " ")
			if len(splitMsg) >= 4 {
				customerID := splitMsg[3]
				cacheKey := "customer:" + customerID
				configs.RedisClient.Client.Del(cacheKey)
				fmt.Printf("Deleted Redis cache for key: %s\n", cacheKey)
			}
		case <-time.After(90 * time.Second):
			fmt.Println("Received no events for 90 seconds, checking connection")
			go func() {
				listener.Ping()
			}()
		}
	}
}
