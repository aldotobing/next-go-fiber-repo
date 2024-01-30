package helper

import (
	"time"
)

func GetExpiredPoint(in time.Time) (res string) {
	in = in.AddDate(1, 6, 0)

	var date time.Time
	if in.Month() < time.July {
		date = time.Date(in.Year(), time.July, 1, 0, 0, 0, 0, in.Location())
	} else {
		date = time.Date(in.Year()+1, time.January, 1, 0, 0, 0, 0, in.Location())
	}

	date = date.AddDate(0, 0, -1)
	res = date.Format("2006-01-02")
	return
}

func GetExpiredOneMonth(in time.Time) (res string) {
	in = in.AddDate(0, 1, 0)
	res = in.Format("2006-01-02")
	return
}
