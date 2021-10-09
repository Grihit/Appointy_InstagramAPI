package models

import "gopkg.in/mgo.v2/bson"
type User struct {
	ID       bson.ObjectId `json:"_id" bson:"_id"`
	Name     string        `json:"Name" bson:"Name"`
	Email    string        `json:"Email" bson:"Email"`
	Password string        `json:"Password" bson:"Password" `
}
