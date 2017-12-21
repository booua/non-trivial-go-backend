// GetCanvasDataById
// AddCanvasData
// DeleteCanvasData

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

const canvasDataCollectionName = "canvasData"

type CanvasDataController struct {
	canvasData *mgo.Collection
}

func NewCanvasDataController() *CanvasDataController {
	return &CanvasDataController{
		db(dbName).C(canvasDataCollectionName),
	}
}

func (c CanvasDataController) AddCanvasData(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	o := models.CanvasData{}
	parseRequest(r, &o)
	o.Id = bson.NewObjectId()

	c.canvasData.Insert(o)

	uj, _ := json.Marshal(o)

	allowOrigin(w)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", uj)
}



func (c CanvasDataController) GetCanvasDataById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	err, id := getId(p)
	if err != nil {
		w.WriteHeader(404)
		return
	}

	o := models.CanvasData{}

	if err := c.canvasData.FindId(id).One(&o); err != nil {
		w.WriteHeader(404)
		return
	}

	j, _ := json.Marshal(o)

	allowOrigin(w)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", j)
}

func (c CanvasDataController) DeleteCanvasData(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	err, id := getId(p)
	if err != nil {
		w.WriteHeader(404)
		return
	}

	if err := c.canvasData.RemoveId(id); err != nil {
		w.WriteHeader(404)
		return
	}

	allowOrigin(w)
	w.WriteHeader(200)
}
