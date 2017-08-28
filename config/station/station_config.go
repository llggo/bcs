package station

import (
	"fmt"
	"qrcode-bulk/qrcode-bulk-generator/config/shared"
	"qrcode-bulk/qrcode-bulk-generator/x/conf"
)

var logger = shared.ConfigLog

type StationConfig struct {
	Sync    SyncConfig
	Storage StorageConfig
	Static  StaticConfig
	Log     conf.LogConfig
	Server  conf.ServerConfig
	Sys     SystemConfig
}

func (s *StationConfig) Check() {
	s.Log.Init()
	s.Sync.Check()
	s.Storage.Check()
	s.Static.Check()
	s.Sys.check()
}

func (s StationConfig) String() string {
	return fmt.Sprintf("station:[%s][%s][%s][%s][%s]", s.Sync, s.Storage, s.Static, s.Log, s.Server)
}
