package supcription

import (
	"qrcode-bulk/qrcode-bulk-generator/o/org/feature"
	"qrcode-bulk/qrcode-bulk-generator/x/db/mgo"
	"time"
)

type Supcription struct {
	mgo.BaseModel `bson:",inline"`
	PackageCode   string    `bson:"package_code" json:"package_code"`
	QrcodeCount   int       `bson:"qrcode_count" json:"qrcode_count"`
	CreateTime    time.Time `bson:"create_time" json:"create_time"`
	Expirydate    time.Time `bson:"expiry_date" json:"expiry_date"`
}

func (s *Supcription) CheckSupcription() (bool, *feature.Message) {
	return true, nil
}

func (s *Supcription) CheckQrCodeLimit(name feature.FeatureName, action feature.FeatureAction) (bool, *feature.Message) {
	if name == feature.Qrcode {
		if action == feature.Create {
			if s.QrcodeCount > 0 {
				s.QrcodeCount--
				s.Update(s)
			} else {
				return false, &feature.Message{
					Access:  false,
					Message: "Access is denied. Check your subcrition qrcode limit",
				}
			}
		}
	}
	return true, nil
}
