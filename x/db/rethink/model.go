package rethink

import (
	"time"
)

// var idMaker = math.RandStringMaker{Prefix: "rad", Length: 20}

type IModel interface {
	GetID() string
	SetID(id string)
	IsDeleted() bool
	BeforeCreate()
	BeforeUpdate()
	BeforeDelete()
}

type BaseModel struct {
	ID    string `gorethink:"id,omitempty" json:"id"`
	MTime int64  `gorethink:"mtime" json:"mtime"`
	DTime int64  `gorethink:"dtime" json:"dtime"`
}

func (b *BaseModel) GetID() string {
	return b.ID
}

func (b *BaseModel) SetID(s string) {
	b.ID = s
}

func (b *BaseModel) IsDeleted() bool {
	return b.DTime > 0
}

func (b *BaseModel) BeforeCreate() {
	b.MTime = time.Now().Unix()
	b.DTime = 0
}

func (b *BaseModel) BeforeUpdate() {
	b.MTime = time.Now().Unix()
}

func (b *BaseModel) BeforeDelete() {

}

type CodeNameModel struct {
	BaseModel
	Name string `gorethink:"name" json:"name"`
	Code string `gorethink:"code" json:"code"`
}
