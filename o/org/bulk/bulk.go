package bulk

import (
	"bar-code/bcs/x/db/mgo"
	"bar-code/bcs/x/mlog"
)

var objBulkLogging = mlog.NewTagLog("obj_Bulk")

// type infoProduct map[string]interface{}

type Bulk struct {
	mgo.BaseModel `bson:",inline"`
	Name          string                  `bson:"name" json:"name"`
	Image         string                  `bson:"image" json:"image"`
	Logo          string                  `bson:"logo" json:"logo"`
	Status        bool                    `bson:"status" json:"status"`
	Type          int                     `bson:"type" json:"type"`
	VerifyNumber  int64                   `bson:"verify_number" json:"verify_number"`
	QrCodeNumber  int64                   `bson:"qrcode_number" json:"qrcode_number"`
	UserID        string                  `bson:"user_id" json:"user_id"`
	Hello         string                  `bson:"hello" json:"hello"`
	Contact       Contact                 `bson:"contact" json:"contact"`
	ReferProduct  map[string]ReferProduct `bson:"refer_product" json:"refer_product"`
	VerifyLimit   int                     `bson:"verify_limit" json:"verify_limit"`
	InfoProduct   Product                 `bson:"info_product" json:"info_product"`
}

type ReferProduct struct {
	Name  string `bson:"name" json:"name"`
	Image string `bson:"image" json:"image"`
	Price string `bson:"price" json:"price"`
	Link  string `bson:"link" json:"link"`
}

type Product struct {
	Name        string          `bson:"name" json:"name"`
	ImageInfo   string          `bson:"image_info" json:"image_info"`
	ImageVerify string          `bson:"image_verify" json:"image_verify"`
	Price       string          `bson:"price" json:"price"`
	Info        map[string]Info `bson:"info" json:"info"`
}

type Info struct {
	Title   string `bson:"title" json:"title"`
	Content string `bson:"content" json:"content"`
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

func (b *Bulk) GetVerifyLimit() int {
	return b.VerifyLimit
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

// func (b *Bulk) GetProduct() *Product {
// 	var c = &Product{
// 		b.Product.Name,
// 		b.Product.Image,
// 		b.Product.Price,
// 		b.Product.Tutorial,
// 		b.Product.Result,
// 		b.Product.Element,
// 	}
// 	return c
// }

func (b *Bulk) GetRefer() map[string]ReferProduct {
	return b.ReferProduct
}

func (b *Bulk) GetInfoProduct() Product {
	return b.InfoProduct
}
