package helper

import (
	"time"
)

func GetExpiredWithInterval(in time.Time, interval int) (res string) {
	if interval == 0 {
		in = in.AddDate(0, 1, 0)
	} else {
		in = in.AddDate(0, 0, interval)
	}
	res = in.Format("2006-01-02")
	return
}
