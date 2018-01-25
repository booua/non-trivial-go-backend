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
	canvas *mgo.Collection
}

func NewCanvasController() *CanvasController {
	return &CanvasController{
		db(dbName).C(canvasCollectionName),
	}
}

func (c CanvasController) AddCanvas(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	o := models.Canvas{}
	if err := parseRequest(r, &o); err != nil {
		http.Error(w, parseRequest(r, &o).Error(), http.StatusInternalServerError)
		return
	}
	o.Id = bson.NewObjectId()

	c.canvas.Insert(o)

	uj, _ := json.Marshal(o)

	allowOrigin(w)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", uj)
}

func (c CanvasController) GetCanvasById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	err, id := getId(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	o := models.Canvas{}

	if err := c.canvas.FindId(id).One(&o); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
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
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err := c.canvas.RemoveId(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	allowOrigin(w)
	w.WriteHeader(200)
}
