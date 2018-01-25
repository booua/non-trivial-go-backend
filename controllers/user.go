// AddUser
// GetAllUsers
// GetUserById
// DeleteUser

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

const userCollection = "users"

type UserController struct {
	users *mgo.Collection
}

func NewUserController() *UserController {
	return &UserController{
		db(dbName).C(userCollection),
	}
}

func (c UserController) AddUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	o := models.User{}
	if err := parseRequest(r, &o); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	o.Id = bson.NewObjectId()

	c.users.Insert(o)

	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(o)

	// Write content-type, statuscode, payload
	allowOrigin(w)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", uj)
}

func (c UserController) GetAllUsers(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	o := []models.User{}

	if err := c.users.Find(nil).All(&o); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
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

func (c UserController) GetUserById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	err, id := getId(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	o := models.User{}

	if err := c.users.FindId(id).One(&o); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
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

func (c UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	err, id := getId(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err := c.users.RemoveId(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	allowOrigin(w)
	w.WriteHeader(200)
}
