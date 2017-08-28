package config

import (
	"qrcode-bulk/qrcode-bulk-generator/config/business"
	"qrcode-bulk/qrcode-bulk-generator/config/database"
	"qrcode-bulk/qrcode-bulk-generator/config/station"
)

func Station() *station.StationConfig {
	return &projectConfig.Station
}

func Database() *database.DatabaseConfig {
	return &projectConfig.Database
}

func Business() *business.BusinessConfig {
	return &projectConfig.Business
}

func Project() *ProjectConfig {
	return projectConfig
}

func SystemConfig() *station.SystemConfig {
	return &projectConfig.Station.Sys
}
