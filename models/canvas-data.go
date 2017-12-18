package models

import "gopkg.in/mgo.v2/bson"

type CanvasData struct {
	Id     bson.ObjectId `json:"id" bson:"_id"`
	Author string        `json:"author" bson:"author"`
	CanvasJSONData   string        `json:"content" bson:"canvasJSONData"`
}
