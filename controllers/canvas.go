// AddCanvas
// GetCanvasById
// DeleteCanvas

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

const canvasCollectionName = "canvasData"

type CanvasController struct {
	canvasData *mgo.Collection
}

func NewCanvasController() *CanvasController {
	return &CanvasController{
		db(dbName).C(canvasCollectionName),
	}
}

func (c CanvasController) AddCanvas(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	o := models.Canvas{}
	parseRequest(r, &o)
	o.Id = bson.NewObjectId()

	c.canvasData.Insert(o)

	uj, _ := json.Marshal(o)

	allowOrigin(w)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", uj)
}

func (c CanvasController) GetCanvasById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	err, id := getId(p)
	if err != nil {
		w.WriteHeader(404)
		return
	}

	o := models.CanvasData{}

	if err := c.canvas.FindId(id).One(&o); err != nil {
		w.WriteHeader(404)
		return
	}

	j, _ := json.Marshal(o)

	allowOrigin(w)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", j)
}

func (c CanvasController) DeleteCanvas(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	err, id := getId(p)
	if err != nil {
		w.WriteHeader(404)
		return
	}

	if err := c.canvas.RemoveId(id); err != nil {
		w.WriteHeader(404)
		return
	}

	allowOrigin(w)
	w.WriteHeader(200)
}
