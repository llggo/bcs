package mgo

import (
	"qrcode-bulk/qrcode-bulk-generator/x/mlog"
)

var mongoDBLog = mlog.NewTagLog("MongoDB")

type M map[string]interface{}
