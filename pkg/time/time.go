package time

import (
	"strings"
	"time"
)

// In ...
func In(t time.Time, name string) (time.Time, error) {
	loc, err := time.LoadLocation(name)
	if err == nil {
		t = t.In(loc)
	}
	return t, err
}

// InFormat ...
func InFormat(t time.Time, name, format string) (string, error) {
	loc, err := time.LoadLocation(name)
	if err != nil {
		return "", err
	}

	return t.In(loc).Format(format), err
}

// InFormatLang ...
func InFormatLang(t time.Time, name, format, lang string) string {
	loc, err := time.LoadLocation(name)
	if err != nil {
		return ""
	}

	res := t.In(loc).Format(format)
	if lang == "id" {
		res = strings.Replace(res, "January", "Januari", 1)
		res = strings.Replace(res, "February", "Februari", 1)
		res = strings.Replace(res, "March", "Maret", 1)
		res = strings.Replace(res, "April", "April", 1)
		res = strings.Replace(res, "May", "Mei", 1)
		res = strings.Replace(res, "Juny", "Juni", 1)
		res = strings.Replace(res, "July", "Juli", 1)
		res = strings.Replace(res, "August", "Agustus", 1)
		res = strings.Replace(res, "September", "September", 1)
		res = strings.Replace(res, "October", "Oktober", 1)
		res = strings.Replace(res, "November", "November", 1)
		res = strings.Replace(res, "December", "Desember", 1)
	}

	return res
}

// InFormatNoErr ...
func InFormatNoErr(t time.Time, name, format string) string {
	loc, err := time.LoadLocation(name)
	if err != nil {
		return ""
	}

	return t.In(loc).Format(format)
}

// Convert ...
func Convert(t, fromFormat, toFormat string) string {
	timeConvert, err := time.Parse(fromFormat, t)
	if err != nil {
		return ""
	}

	return timeConvert.Format(toFormat)
}

// ParseNoErr ...
func ParseNoErr(t, format string) time.Time {
	res, err := time.Parse(format, t)
	if err != nil {
		return res
	}

	return res
}

// Diff ...
func Diff(a, b time.Time) (year, month, day, hour, min, sec int) {
	if a.Location() != b.Location() {
		b = b.In(a.Location())
	}
	if a.After(b) {
		a, b = b, a
	}
	y1, M1, d1 := a.Date()
	y2, M2, d2 := b.Date()

	h1, m1, s1 := a.Clock()
	h2, m2, s2 := b.Clock()

	year = int(y2 - y1)
	month = int(M2 - M1)
	day = int(d2 - d1)
	hour = int(h2 - h1)
	min = int(m2 - m1)
	sec = int(s2 - s1)

	// Normalize negative values
	if sec < 0 {
		sec += 60
		min--
	}
	if min < 0 {
		min += 60
		hour--
	}
	if hour < 0 {
		hour += 24
		day--
	}
	if day < 0 {
		// days in month:
		t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)
		day += 32 - t.Day()
		month--
	}
	if month < 0 {
		month += 12
		year--
	}

	return
}

// DiffCustom ...
func DiffCustom(start string, b time.Time) (year, month, day, hour, min, sec int) {
	if start == "" {
		return
	}
	a, err := time.Parse(time.RFC3339, start)
	if err != nil {
		return
	}

	if a.Location() != b.Location() {
		b = b.In(a.Location())
	}
	if a.After(b) {
		a, b = b, a
	}
	y1, M1, d1 := a.Date()
	y2, M2, d2 := b.Date()

	h1, m1, s1 := a.Clock()
	h2, m2, s2 := b.Clock()

	year = int(y2 - y1)
	month = int(M2 - M1)
	day = int(d2 - d1)
	hour = int(h2 - h1)
	min = int(m2 - m1)
	sec = int(s2 - s1)

	// Normalize negative values
	if sec < 0 {
		sec += 60
		min--
	}
	if min < 0 {
		min += 60
		hour--
	}
	if hour < 0 {
		hour += 24
		day--
	}
	if day < 0 {
		// days in month:
		t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)
		day += 32 - t.Day()
		month--
	}
	if month < 0 {
		month += 12
		year--
	}

	return
}

// AddTimezone ...
func AddTimezone(date, timezone string) string {
	if date == "" {
		return ""
	}

	return date + "T00:00:00" + timezone
}

// AddDays ...
func AddDays(date string, days int) string {
	newDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return ""
	}

	res := newDate.AddDate(0, 0, days)

	return res.Format("2006-01-02")
}

// CheckDate ...
func CheckDate(data, format string) string {
	_, err := time.Parse(format, data)
	if err != nil {
		return ""
	}

	return data
}

// ToString ...
func ToString(data time.Time, format string) string {
	if data.IsZero() {
		return ""
	}

	return data.Format(format)
}

// StartDate ...
func StartDate(input time.Time, location, timezone string) time.Time {
	todayString, _ := InFormat(input, location, "2006-01-02")
	date, _ := time.Parse(time.RFC3339, todayString+"T00:00:00"+timezone)

	return date
}

//weekstart
func WeekStart() time.Time {

	t := time.Now().UTC()

	// Roll back to Monday:
	if wd := t.Weekday(); wd == time.Sunday {
		t = t.AddDate(0, 0, -6)
	} else {
		t = t.AddDate(0, 0, -int(wd)+1)
	}

	return t
}

//weekstart
func WeekEnd() time.Time {

	t := time.Now().UTC()

	// Roll back to Monday:
	if wd := t.Weekday(); wd == time.Sunday {
		t = t.AddDate(0, 0, -6)
	} else {
		t = t.AddDate(0, 0, -int(wd)+1)
	}

	t = t.AddDate(0, 0, 6)

	return t
}

func GetDate(currentDate, format, location string) string {
	loc, err := time.LoadLocation(location)
	if err != nil {
		return ""
	}
	curdate, error := time.ParseInLocation("2006-01-02T00:00:00Z", currentDate, loc)
	if error != nil {
		return ""
	}
	dateStr := "" + curdate.Format(format)
	res := dateStr
	res = strings.ReplaceAll(res, "January", "Januari")
	res = strings.ReplaceAll(res, "February", "Februari")
	res = strings.ReplaceAll(res, "March", "Maret")
	res = strings.ReplaceAll(res, "April", "April")
	res = strings.ReplaceAll(res, "May", "Mei")
	res = strings.ReplaceAll(res, "Juny", "Juni")
	res = strings.ReplaceAll(res, "July", "Juli")
	res = strings.ReplaceAll(res, "August", "Agustus")
	res = strings.ReplaceAll(res, "September", "September")
	res = strings.ReplaceAll(res, "October", "Oktober")
	res = strings.ReplaceAll(res, "November", "November")
	res = strings.ReplaceAll(res, "December", "Desember")

	res = strings.ReplaceAll(res, "Sunday", "Minggu")
	res = strings.ReplaceAll(res, "Monday", "Senin")
	res = strings.ReplaceAll(res, "Tuesday", "Selasa")
	res = strings.ReplaceAll(res, "Wednesday", "Rabu")
	res = strings.ReplaceAll(res, "Thursday", "Kamis")
	res = strings.ReplaceAll(res, "Friday", "Jum'at")
	res = strings.ReplaceAll(res, "Saturday", "Sabtu")

	return res

}
