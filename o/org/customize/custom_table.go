package customize

import (
	"qrcode-bulk/qrcode-bulk-generator/o/model"

	"gopkg.in/mgo.v2/bson"
)

var TableCustomize = model.NewTable("customize", "custom")

func NewCustomID() string {
	return TableCustomize.Next()
}

func AllCustomOnBranch(branchid string, role string) ([]Customize, error) {
	var res = make([]Customize, 0)
	var query = bson.M{"branch_id": branchid}
	if role != "" {
		query["role"] = role
	}
	return res, TableCustomize.C().Find(query).All(&res)
}

func (cus *Customize) Create() error {
	return TableCustomize.Create(cus)
}

func MarkDelete(id string) error {
	return TableCustomize.MarkDelete(id)
}

func (cus *Customize) Update(newValue *Customize) error {
	var values = map[string]interface{}{
		"name": newValue.Name,
	}
	if newValue.GetLogo() != cus.GetLogo() {
		values["logo"] = newValue.GetLogo()
	}
	if newValue.GetFrame() != cus.GetFrame() {
		values["frame"] = newValue.GetFrame()
	}
	if newValue.GetPattern() != cus.GetPattern() {
		values["pattern"] = newValue.GetPattern()
	}
	if newValue.GetTEyes() != cus.GetTEyes() {
		values["t_eye"] = newValue.GetTEyes()
	}
	if newValue.GetREyes() != cus.GetREyes() {
		values["r_eye"] = newValue.GetREyes()
	}
	if newValue.GetBEyes() != cus.GetBEyes() {
		values["b_eye"] = newValue.GetBEyes()
	}
	if newValue.GetBackground() != cus.GetBackground() {
		values["background"] = newValue.GetBackground()
	}
	if newValue.GetPattern() != cus.GetPattern() {
		values["pattern"] = newValue.GetPattern()
	}
	if newValue.GetImgSize() != cus.GetImgSize() {
		values["img_size"] = newValue.GetImgSize()
	}
	if newValue.GetPositionOfEye() != cus.GetPositionOfEye() {
		values["position_of_eye"] = newValue.GetPositionOfEye()
	}
	if newValue.GetIsTemplate() != cus.GetIsTemplate() {
		values["is_template"] = newValue.GetIsTemplate()
	}
	if newValue.GetTempThumb() != cus.GetTempThumb() {
		values["temp_thumb"] = newValue.GetTempThumb()
	}
	return TableCustomize.UnsafeUpdateWhere(map[string]interface{}{"qrcode_id": cus.QrCodeID}, values)
}
