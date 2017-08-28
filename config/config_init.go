package config

import (
	"flag"
	"github.com/BurntSushi/toml"
)

var projectConfig = &ProjectConfig{}
var defaultConfigFile = "qrcode-pba.toml"

func init() {
	var configFile string
	flag.StringVar(&configFile, "conf", defaultConfigFile, "Config File")
	flag.Parse()
	if _, err := toml.DecodeFile(configFile, projectConfig); err != nil {
		logger.Fatalf("Read config %v", err)
	}
	projectConfig.Check()
	logger.Infof(0, "Config %s", projectConfig)
}
