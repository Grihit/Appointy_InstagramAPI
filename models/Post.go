package models

import "gopkg.in/mgo.v2/bson"

type Post struct {
	ID       bson.ObjectId       `json:"_id" bson:"_id"`
	UserID   string              `json:"Userid" bson:"Userid"`
	Caption  string              `json:"Caption" bson:"Caption"`
	ImageURL string              `json:"ImageURL" bson:"ImageURL"`
	Time     bson.MongoTimestamp `json:"Time" bson:"Time" `
}
