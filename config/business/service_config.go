package business

import (
	"fmt"
)

type ServiceConfig struct {
	MaxWaitingMinute      int64 `json:"max_waiting_minute"`
	MaxServingMinute      int64 `json:"max_serving_minute"`
	AutoFinishMinute      int64 `json:"auto_finish_minute"`
	WaitLongAlertPercent  int64 `json:"wait_long_alert_percent"`
	ServeLongAlertPercent int64 `json:"serve_long_alert_percent"`
}

func (c ServiceConfig) String() string {
	return fmt.Sprintf("service:wait=%d;serve=%d;fin=%d;wait=%d;serve=%d", c.MaxWaitingMinute, c.MaxServingMinute,
		c.AutoFinishMinute, c.WaitLongAlertPercent, c.ServeLongAlertPercent)
}

func (c *ServiceConfig) Check() {
	if c.MaxWaitingMinute < 1 {
		c.MaxWaitingMinute = 15
	}
	if c.MaxServingMinute < 1 {
		c.MaxServingMinute = 15
	}
	if c.AutoFinishMinute < 1 {
		c.AutoFinishMinute = 120
	}
	if c.AutoFinishMinute < c.MaxServingMinute+30 {
		c.AutoFinishMinute = c.MaxServingMinute + 30
	}
	if c.WaitLongAlertPercent < 0 {
		c.WaitLongAlertPercent = 0
	}
	if c.WaitLongAlertPercent > 100 {
		c.WaitLongAlertPercent = 100
	}
	if c.ServeLongAlertPercent < 0 {
		c.ServeLongAlertPercent = 0
	}
	if c.ServeLongAlertPercent > 100 {
		c.ServeLongAlertPercent = 100
	}
}
