package models

import "gopkg.in/mgo.v2/bson"

// Represents a movie, we uses bson keyword to tell the mgo driver how to name
// the properties in mongodb document
type User struct {
	ID       bson.ObjectId `json:"_id" bson:"_id"`
	Name     string        `json:"Name" bson:"Name"`
	Email    string        `json:"Email" bson:"Email"`
	Password string        `json:"Password" bson:"Password" `
}
