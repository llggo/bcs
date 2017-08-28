package business

import (
	"fmt"
)

type BusinessConfig struct {
	General  GeneralConfig  `json:"general"`
	Service  ServiceConfig  `json:"service"`
	Priority PriorityConfig `json:"priority"`
}

func (c BusinessConfig) String() string {
	return fmt.Sprintf("business=[%s][%s][%s]", c.General, c.Service, c.Priority)
}

func (c *BusinessConfig) Check() {
	c.General.Check()
	c.Service.Check()
	c.Priority.Check()
}
