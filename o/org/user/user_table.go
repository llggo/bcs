package user

import (
	"qrcode-bulk/qrcode-bulk-generator/o/model"

	"gopkg.in/mgo.v2/bson"
)

var TableUser = model.NewTable("users", "usr")

func NewUserID() string {
	return TableUser.Next()
}

func AllUserOnBranch(branchid string, role string) ([]User, error) {
	var res = make([]User, 0)
	var query = bson.M{"branch_id": branchid}
	if role != "" {
		query["role"] = role
	}
	return res, TableUser.C().Find(query).All(&res)
}

func (b *User) Create() error {
	if err := b.ensureUniqueUsername(); err != nil {
		return err
	}
	var p = password(b.Password)
	// replace
	if err := p.HashTo(&b.Password); err != nil {
		return err
	}
	return TableUser.Create(b)
}

func MarkDelete(id string) error {
	return TableUser.MarkDelete(id)
}

func (v *User) Update(newValue *User) error {
	var values = map[string]interface{}{
		"firstname": newValue.Firstname,
	}

	// if newValue.Username != v.Username {
	// 	if err := newValue.ensureUniqueUsername(); err != nil {
	// 		return err
	// 	}
	// 	values["username"] = newValue.Username
	// }

	if len(newValue.Password) > 0 {
		if newValue.Password != v.Password {
			var p = password(newValue.Password)
			if err := p.HashTo(&newValue.Password); err != nil {
				return err
			}
		}
		values["password"] = newValue.Password
	}

	if newValue.GetPhone() != v.GetPhone() {
		values["phone"] = newValue.GetPhone()
	}
	if newValue.GetOrigin() != v.GetOrigin() {
		values["origin"] = newValue.GetOrigin()
	}
	if newValue.GetFirstname() != v.GetFirstname() {
		values["firstname"] = newValue.GetFirstname()
	}
	if newValue.GetLastname() != v.GetLastname() {
		values["lastname"] = newValue.GetLastname()
	}
	if newValue.GetCompany() != v.GetCompany() {
		values["company"] = newValue.GetCompany()
	}
	if newValue.GetEmail() != v.GetEmail() {
		values["email"] = newValue.GetEmail()
	}

	values["supcription_id"] = newValue.SupcriptionID

	values["role"] = newValue.Role

	return TableUser.UnsafeUpdateByID(v.ID, values)
}
