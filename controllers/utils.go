package controllers

import (
	"encoding/json"
	"errors"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

var errIncorrectId = errors.New("Incorrect format of id")

func parseRequest(r *http.Request, o interface{}) {
	json.NewDecoder(r.Body).Decode(&o)
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
