package qrcode

import (
	"qrcode-bulk/qrcode-bulk-generator/x/db/mgo"
	"qrcode-bulk/qrcode-bulk-generator/x/mlog"
	qr "qrcode/qrcodelib"
)

var objQrCodeLoging = mlog.NewTagLog("obj_QrCoode")

type qrcodeData map[string]interface{}

type QrCode struct {
	mgo.BaseModel `bson:",inline"`
	Enable        bool             `bson:"enable" json:"enable"`
	Root          bool             `bson:"root" json:"root"`
	UserID        string           `bson:"user_id" json:"user_id"`
	BulkID        string           `bson:"bulk_id" json:"bulk_id"`
	QrData        qrcodeData       `bson:"data" json:"data"`
	QrName        string           `bson:"name" json:"name,omitempty"`
	QrType        string           `bson:"type" json:"type"`
	QrTemplate    string           `bson:"template" json:"template"`
	QrPathImg     string           `bson:"path_img" json:"path_img"`
	QrPathBase64  string           `bson:"path_base" json:"path_base"`
	QrMode        string           `bson:"mode" json:"mode"`
	QrSize        int              `bson:"size" json:"size"`
	QrLevel       qr.RecoveryLevel `bson:"level" json:"level"`
	TagetScan     string           `bson:"taget_scan" json:"taget_scan"`
	Host          string           ``
}

func (qr *QrCode) MakeStaticLink() (string, error) {
	switch qr.QrType {
	default:
		return "", nil
	case "text":
		return qr.HandleText()
	case "url":
		return qr.HandleUrl()
	case "urls":
		return qr.HandleUrl()
	case "sms":
		return qr.HandleSms()
	case "mms":
		return qr.HandleMms()
	case "phone":
		return qr.HandlePhone()
	case "email":
		return qr.HandleEmail()
	case "calendar":
		return qr.HandleCalendar()
	case "geo":
		return qr.HandleGeo()
	case "youtube":
		return qr.HandleYoutube()
	case "event":
		return qr.HandleEvent()
	// case "card":
	// return s.HandleCard(w, r, qr)
	case "wifi":
		return qr.HandleWifi()
	case "image":
		return qr.HandleImage()
	case "pdf":
		return qr.HandlePDF()
	case "audio":
		return qr.HandleAudio()
	}
}

func (qr *QrCode) MakeDynamicLink() string {
	return "http://mirascan.vn:31001/api/handle/welcome?qrcode_id=" + qr.GetID()
}

func (qr *QrCode) GetActiveData() map[string]interface{} {
	return qr.QrData
}

func (qr *QrCode) GetName() string {
	return qr.QrName
}

func (qr *QrCode) GetActiveTemplate() string {
	return qr.QrTemplate
}
func (qr *QrCode) IsDynamic() bool {
	return qr.QrMode == "dynamic"
}

func (qr *QrCode) GetQrPathImg() string {
	return qr.QrPathImg
}

func (qr *QrCode) GetType() string {
	return qr.QrType
}

func (qr *QrCode) GetSize() int {
	return qr.QrSize
}

func (qr *QrCode) GetLevel() qr.RecoveryLevel {
	return qr.QrLevel
}

func (qr *QrCode) GetPathBase64() string {
	return qr.QrPathBase64
}
func (qr *QrCode) GetTagetScan() string {
	return qr.TagetScan
}

func (qr *QrCode) GetBulkID() string {
	return qr.BulkID
}
