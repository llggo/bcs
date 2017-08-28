package timer

import (
	"time"
)

const (
	secondsPerMinute = 60
	secondsPerHour   = 60 * 60
	secondsPerDay    = 24 * secondsPerHour
	secondsPerWeek   = 7 * secondsPerDay
	daysPer400Years  = 365*400 + 97
	daysPer100Years  = 365*100 + 24
	daysPer4Years    = 365*4 + 1
)

var offset int64

func init() {
	var _, d = time.Now().Zone()
	offset = int64(d)
}

func LocalStartOfToday() int64 {
	var now = time.Now().Unix()
	var ellapsed = now % secondsPerDay
	return now - ((ellapsed + offset) % secondsPerDay)
}
