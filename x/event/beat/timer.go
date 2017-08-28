package beat

import (
	"qrcode-bulk/qrcode-bulk-generator/x/event"
)

var daily = event.NewHub(8)

func OnNewDay() *event.Line {
	return daily.NewLine()
}
