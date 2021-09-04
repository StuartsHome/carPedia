package utils

import (
	"time"
)

// returns the timestamp to the hour
func UnixTimeHour(t int64) int64 {
	temp := time.Unix(t, 0).UTC()
	return temp.Truncate(time.Hour).UTC().Unix()
}
