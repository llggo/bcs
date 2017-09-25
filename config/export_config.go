package config

import (
	"bar-code/bcs/config/business"
	"bar-code/bcs/config/database"
	"bar-code/bcs/config/station"
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
