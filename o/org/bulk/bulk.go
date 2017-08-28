package bulk

import (
	"qrcode-bulk/qrcode-bulk-generator/x/db/mgo"
	"qrcode-bulk/qrcode-bulk-generator/x/mlog"
)

var objBulkLogging = mlog.NewTagLog("obj_Bulk")

type Info struct {
	Topic   string `bson:"topic" json:"topic"`
	Content string `bson:"content" json:"content"`
}
type iBulkInfo map[string]interface {
	// Data : Info
}

type Bulk struct {
	mgo.BaseModel `bson:",inline"`
	Name          string                 `bson:"name" json:"name"`
	Status        bool                   `bson:"status" json:"status"`
	Type          int                    `bson:"type" json:"type"`
	VerifyNumber  int64                  `bson:"verify_number" json:"verify_number"`
	QrCodeNumber  int64                  `bson:"qrcode_number" json:"qrcode_number"`
	UserID        string                 `bson:"user_id" json:"user_id"`
	Hello         string                 `bson:"hello" json:"hello"`
	Contact       Contact                `bson:"contact" json:"contact"`
	Product       Product                `bson:"product" json:"product"`
	ReferProduct  map[string]interface{} `bson:"refer_product" json:"refer_product"`
}

type Product struct {
	Name     string `bson:"name" json:"name"`
	Image    string `bson:"image" json:"image"`
	Price    string `bson:"price" json:"price"`
	Tutorial string `bson:"tutorial" json:"tutorial"`
	Result   string `bson:"result" json:"result"`
	Element  string `bson:"element" json:"element"`
}

type Contact struct {
	Logo    string `bson:"logo" json:"logo"`
	Company string `bson:"company" json:"company"`
	Email   string `bson:"email" json:"email"`
	Address string `bson:"address" json:"address"`
	Phone   string `bson:"phone" json:"phone"`
}

func (b *Bulk) GetName() string {
	return b.Name
}

func (b *Bulk) GetStatus() bool {
	return b.Status
}

func (b *Bulk) GetType() int {
	return b.Type
}

func (b *Bulk) GetVerifyNumber() int64 {
	return b.VerifyNumber
}

func (b *Bulk) GetUserID() string {
	return b.UserID
}

func (b *Bulk) GetHello() string {
	return b.Hello
}

func (b *Bulk) GetContact() *Contact {
	var c = &Contact{
		b.Contact.Logo,
		b.Contact.Company,
		b.Contact.Email,
		b.Contact.Address,
		b.Contact.Phone,
	}
	return c
}

func (b *Bulk) GetProduct() *Product {
	var c = &Product{
		b.Product.Name,
		b.Product.Image,
		b.Product.Price,
		b.Product.Tutorial,
		b.Product.Result,
		b.Product.Element,
	}
	return c
}

func (b *Bulk) GetRefer() map[string]interface{} {
	return b.ReferProduct
}
