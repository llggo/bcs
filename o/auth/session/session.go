package session

import (
	"encoding/json"
	"qrcode-bulk/qrcode-bulk-generator/o/org/feature"
	"qrcode-bulk/qrcode-bulk-generator/o/org/user"
	"qrcode-bulk/qrcode-bulk-generator/x/db/mgo"
)

type Session struct {
	mgo.BaseModel `bson:",inline"`
	Username      string `json:"username"`
	UserID        string `json:"userid"`
	BranchID      string `json:"branch_id"`
	CTime         int64  `json:"ctime"`
}

func (a *Session) MarshalBinary() ([]byte, error) {
	return json.Marshal(a)
}

func (a *Session) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, a)
}

func (a *Session) CheckAccess(name feature.FeatureName, action feature.FeatureAction) (bool, *feature.Message) {
	var u, err = user.GetByID(a.UserID)
	if err != nil {
		return false, nil
	}

	return u.CheckAccess(name, action)
}
