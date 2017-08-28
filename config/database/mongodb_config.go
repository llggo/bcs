package database

import (
	"fmt"
	"os"
	"qrcode-bulk/qrcode-bulk-generator/config/cons"
	"qrcode-bulk/qrcode-bulk-generator/x/db/mgo"
)

type DatabaseConfig struct {
	DBHost   string
	DBName   string
	Account  string
	Password string
}

func (o DatabaseConfig) String() string {
	return fmt.Sprintf("db:host=%s;name=%s", o.DBHost, o.DBName)
}

func (o *DatabaseConfig) Check() {
	var _, err = mgo.NewDB(o.DBHost, o.DBName)
	if err != nil {
		logger.Fatalf("Cannot connect to db at [%v]", o.DBHost)
	}
	os.Setenv(cons.ENV_OBJECT_DB, o.DBName)
}
