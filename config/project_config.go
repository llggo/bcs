package config

import (
	"fmt"
	"qrcode-bulk/qrcode-bulk-generator/config/business"
	"qrcode-bulk/qrcode-bulk-generator/config/database"
	"qrcode-bulk/qrcode-bulk-generator/config/shared"
	"qrcode-bulk/qrcode-bulk-generator/config/station"
)

var logger = shared.ConfigLog

type ProjectConfig struct {
	Business business.BusinessConfig `json:"business"`
	Database database.DatabaseConfig `json:"database"`
	Station  station.StationConfig   `json:"station"`
}

func (p ProjectConfig) String() string {
	return fmt.Sprintf("config:[%s][%s][%s]", p.Database, p.Station, p.Business)
}

func (p *ProjectConfig) Check() {
	p.Station.Check()
	p.Database.Check()
	p.Business.Check()
}
