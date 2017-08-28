package verify_code

import (
	"qrcode-bulk/qrcode-bulk-generator/x/db/mgo"
	"qrcode-bulk/qrcode-bulk-generator/x/mlog"
)

var objVerifyCodeLogging = mlog.NewTagLog("obj_Verify_Code")

type VerifyCode struct {
	mgo.BaseModel `bson:",inline"`
	VName         string `bson:"name" json:"name"`
	VCode         string `bson:"code" json:"code"`
	QrcodeID      string `bson:"qrcode_id" json:"qrcode_id"`
}

func (b *VerifyCode) GetName() string {
	return b.VName
}

func (b *VerifyCode) GetCode() string {
	return b.VCode
}

func (b *VerifyCode) GetQrcodeID() string {
	return b.QrcodeID
}
