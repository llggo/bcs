package customize

import (
	"qrcode-bulk/qrcode-bulk-generator/x/db/mgo"
	"qrcode-bulk/qrcode-bulk-generator/x/mlog"
)

var objCustomLoging = mlog.NewTagLog("obj_Customize")

type Customize struct {
	mgo.BaseModel `bson:",inline"`
	QrCodeID      string      `bson:"qrcode_id" json:"qrcode_id"`
	Name          string      `bson:"name" json:"name"`
	Logo          string      `bson:"logo" json:"logo"`
	T_Eye         Eye         `bson:"t_eye" json:"t_eye"`
	R_Eye         Eye         `bson:"r_eye" json:"r_eye"`
	B_Eye         Eye         `bson:"b_eye" json:"b_eye"`
	Frame         string      `bson:"frame" json:"frame"`
	Pattern       DataPattern `bson:"pattern" json:"pattern"`
	Background    string      `bson:"background" json:"background"`
	Img_size      int         `bson:"img_size" json:"img_size"`
	PositionOfEye int         `bson:"position_of_eye" json:"position_of_eye"`
	IsTemplate    bool        `bson:"is_template" json:"is_template"`
	Temp_Thumb    bool        `bson:"temp_thumb" json:"temp_thumb"`
}

type Eye struct {
	Inter  string `bson:"inter" json:"inter"`
	Outner string `bson:"outner" json:"outner"`
	Stype  string `bson:"type" json:"type"`
	Size   int    `bson:"size" json:"size"`
}

type DataPattern struct {
	Color  string `bson:"pt_color" json:"pt_color"`
	Matrix string `bson:"pt_matrix" json:"pt_matrix"`
	Shape  string `bson:"pt_shape" json:"pt_shape"`
}

func (cus *Customize) GetQrcodeID() string {
	return cus.QrCodeID
}

func (cus *Customize) GetLogo() string {
	return cus.Logo
}
func (cus *Customize) GetFrame() string {
	return cus.Frame
}

func (cus *Customize) GetBackground() string {
	return cus.Background
}

func (cus *Customize) GetTEyes() *Eye {
	var e = &Eye{
		cus.T_Eye.Inter,
		cus.T_Eye.Outner,
		cus.T_Eye.Stype,
		cus.T_Eye.Size,
	}
	return e
}

func (cus *Customize) GetREyes() *Eye {
	var e = &Eye{
		cus.R_Eye.Inter,
		cus.R_Eye.Outner,
		cus.R_Eye.Stype,
		cus.R_Eye.Size,
	}
	return e
}

func (cus *Customize) GetBEyes() *Eye {
	var e = &Eye{
		cus.B_Eye.Inter,
		cus.B_Eye.Outner,
		cus.B_Eye.Stype,
		cus.B_Eye.Size,
	}
	return e
}

func (cus *Customize) GetPattern() *DataPattern {
	var p = &DataPattern{
		cus.Pattern.Color,
		cus.Pattern.Matrix,
		cus.Pattern.Shape,
	}
	return p
}
func (cus *Customize) GetImgSize() int {
	return cus.Img_size
}
func (cus *Customize) GetPositionOfEye() int {
	return cus.PositionOfEye
}

func (cus *Customize) GetIsTemplate() bool {
	return cus.IsTemplate
}

func (cus *Customize) GetTempThumb() bool {
	return cus.Temp_Thumb
}
