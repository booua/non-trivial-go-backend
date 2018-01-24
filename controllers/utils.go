package controllers

import (
	"encoding/json"
	"errors"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"log"
)

var errIncorrectId = errors.New("Incorrect format of id")

func parseRequest(r *http.Request, o interface{}) {
	err := json.NewDecoder(r.Body).Decode(&o)
	if (err != nil) {
		log.Fatal(err)
	}
}

func getId(p httprouter.Params) (error, bson.ObjectId) {
	id := p.ByName("id")
	if !bson.IsObjectIdHex(id) {
		return errIncorrectId, ""
	}

	return nil, bson.ObjectIdHex(id)
}

func allowOrigin(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
}
