package beat

import (
	"bar-code/bcs/x/event"
)

var daily = event.NewHub(8)

func OnNewDay() *event.Line {
	return daily.NewLine()
}
