package user

import (
	"qrcode-bulk/qrcode-bulk-generator/x/db/mgo"
	"qrcode-bulk/qrcode-bulk-generator/x/mlog"
	// "qrcode/pba/x/web"
)

var objectUserLog = mlog.NewTagLog("object_user")

//User : Employee
type User struct {
	mgo.BaseModel `bson:",inline"`
	Username      string `bson:"username" json:"username,omitempty"`
	Role          string `bson:"role" json:"role"`
	Firstname     string `bson:"firstname" json:"firstname"`
	Lastname      string `bson:"lastname" json:"lastname"`
	Password      string `bson:"password" `
	Email         string `bson:"email" json:"email,omitempty"`
	Origin        string `bson:"origin" json:"origin"`
	Company       string `bson:"company" json:"company"`
	Phone         string `bson:"phone" json:"phone"`
	SupcriptionID string `bson:"supcription_id" json:"supcription_id"`
}

func (v *User) ensureUniqueUsername() error {
	// if len(v.Username) < 3 {
	// 	return web.BadRequest("Username must be at least 6 characters")
	// }
	if err := TableUser.NotExist(map[string]interface{}{
		"username": v.Username,
	}); err != nil {
		return err
	}
	return nil
}
func (v *User) GetFirstname() string {
	return v.Firstname
}

func (v *User) GetLastname() string {
	return v.Lastname
}

func (v *User) GetPhone() string {
	return v.Phone
}
func (v *User) GetOrigin() string {
	return v.Origin
}

func (v *User) GetCompany() string {
	return v.Company
}

func (v *User) GetEmail() string {
	return v.Email
}
