package user

import (
	"bar-code/bcs/o/org/feature"
	"bar-code/bcs/o/org/supcription"
)

func (u *User) CheckAccess(name feature.FeatureName, action feature.FeatureAction) (bool, *feature.Message) {
	var mes = &feature.Message{}
	var sup, err = supcription.GetByID(u.SupcriptionID)
	if err != nil {
		return false, &feature.Message{
			Access:  false,
			Message: "Chưa đăng ký package \nSystem: " + err.Error(),
		}
	}
	ok, mes := sup.CheckSupcription()
	if !ok {
		return false, mes
	}
	var pkg = feature.GetPkg(sup.PackageCode)
	ok, mes = pkg.CheckAccess(name, action)
	if !ok {
		return false, mes
	}
	ok, mes = sup.CheckQrCodeLimit(name, action)
	if !ok {
		return false, mes
	}
	return true, nil
}
