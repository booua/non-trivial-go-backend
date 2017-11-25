package models

import "gopkg.in/mgo.v2/bson"

type Opinion struct {
	Id     bson.ObjectId `json:"id" bson:"_id"`
	Author string        `json:"author" bson:"author"`
	Text   string        `json:"text" bson:"text"`
	Rating int           `json:"rating" bson:"rating"`
}
