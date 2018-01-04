package models

import "gopkg.in/mgo.v2/bson"

type Canvas struct {
	Id     bson.ObjectId `json:"id" bson:"_id"`
	Author string        `json:"author" bson:"author"`
	Canvas string        `json:"content" bson:"canvas"`
}
