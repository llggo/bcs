package session

import (
	"gopkg.in/mgo.v2/bson"
)

func GetByUserId(id string) (*Session, error) {
	var sessions *Session
	return sessions, TableSession.ReadOne(bson.M{"userid": id}, &sessions)
}

func GetByID(id string) (*Session, error) {
	var s Session
	return &s, TableSession.ReadOne(bson.M{"_id": id}, &s)
}
func GetAll() ([]*Session, error) {
	var sessions = []*Session{}
	return sessions, TableSession.ReadAll(&sessions)
}
