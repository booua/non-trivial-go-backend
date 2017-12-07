package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"non-trivial-go-backend/models"
)

const opinionCName = "opinions"

type OpinionController struct {
	opinions *mgo.Collection
}

func NewOpinionController() *OpinionController {
	return &OpinionController{
		db(dbName).C(opinionCName),
	}
}

func (c OpinionController) AddOpinion(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	o := models.Opinion{}
	parseRequest(r, &o)
	o.Id = bson.NewObjectId()

	c.opinions.Insert(o)

	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(o)

	// Write content-type, statuscode, payload
	allowOrigin(w)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", uj)
}


func (c OpinionController) GetAllOpinions(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	o := []models.Opinion{}

	if err := c.opinions.Find(nil).All(&o); err != nil {
		w.WriteHeader(404)
		return
	}

	// Marshal provided interface into JSON structure
	j, _ := json.Marshal(o)

	// Write content-type, statuscode, payload
	allowOrigin(w)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", j)
}

func (c OpinionController) GetOpinion(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	err, id := getId(p)
	if err != nil {
		w.WriteHeader(404)
		return
	}

	o := models.Opinion{}

	if err := c.opinions.FindId(id).One(&o); err != nil {
		w.WriteHeader(404)
		return
	}

	// Marshal provided interface into JSON structure
	j, _ := json.Marshal(o)

	// Write content-type, statuscode, payload
	allowOrigin(w)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", j)
}

func (c OpinionController) DeleteOpinion(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	err, id := getId(p)
	if err != nil {
		w.WriteHeader(404)
		return
	}

	if err := c.opinions.RemoveId(id); err != nil {
		w.WriteHeader(404)
		return
	}

	allowOrigin(w)
	w.WriteHeader(200)
}
