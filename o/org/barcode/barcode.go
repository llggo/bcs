package barcode

import (
	"bar-code/bcs/x/db/mgo"
	"qrcode/pba/x/mlog"
)

var objBarCodeLoging = mlog.NewTagLog("obj_BarCoode")

type Status string

const Pending = Status("pending")

type BarCode struct {
	mgo.BaseModel `bson:",inline"`
	UserID        string `bson:"user_id" json:"user_id"`
	Status        Status `bson:"status" json:"status"`
	Name          string `bson:"name" json:"name,omitempty"`
	Type          string `bson:"type" json:"type"`
	Image         string `bson:"image" json:"image"`
	Code          string `bson:"code" json:"code"`
}

func (b *BarCode) GetName() string {
	return b.Name
}

func (b *BarCode) GetUserID() string {
	return b.UserID
}

func (b *BarCode) GetStatus() Status {
	return b.Status
}

func (b *BarCode) GetType() string {
	return b.Type
}

func (b *BarCode) GetImage() string {
	return b.Image
}

func (b *BarCode) GetCode() string {
	return b.Code
}
