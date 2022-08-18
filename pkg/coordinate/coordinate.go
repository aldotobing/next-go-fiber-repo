package coordinate

import (
	"math"
)

// Distance ...
func Distance(lat1 float64, lng1 float64, lat2 float64, lng2 float64, unit ...string) float64 {
	radlat1 := float64(math.Pi * lat1 / 180)
	radlat2 := float64(math.Pi * lat2 / 180)

	theta := float64(lng1 - lng2)
	radtheta := float64(math.Pi * theta / 180)

	dist := math.Sin(radlat1)*math.Sin(radlat2) + math.Cos(radlat1)*math.Cos(radlat2)*math.Cos(radtheta)
	if dist > 1 {
		dist = 1
	}

	dist = math.Acos(dist)
	dist = dist * 180 / math.Pi
	dist = dist * 60 * 1.1515

	if len(unit) > 0 {
		if unit[0] == "K" {
			dist = dist * 1.609344
		} else if unit[0] == "N" {
			dist = dist * 0.8684
		}
	}

	return dist
}

// Destination ...
func Destination(latitude, longitude float64, bearing, distance float64) (float64, float64) {
	R := 6378.1
	bearing = bearing * math.Pi / 180

	lat1 := latitude * math.Pi / 180
	lon1 := longitude * math.Pi / 180

	lat2 := math.Asin(math.Sin(lat1)*math.Cos(distance/R) +
		math.Cos(lat1)*math.Sin(distance/R)*math.Cos(bearing))
	lon2 := lon1 + math.Atan2(math.Sin(bearing)*math.Sin(distance/R)*math.Cos(lat1),
		math.Cos(distance/R)-math.Sin(lat1)*math.Sin(lat2))

	lat2 = lat2 * 180 / math.Pi
	lon2 = lon2 * 180 / math.Pi

	return lat2, lon2
}
