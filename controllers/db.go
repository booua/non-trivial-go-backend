package controllers

import "gopkg.in/mgo.v2"

var url = "mongodb://localhost"

func session() *mgo.Session {
	// Connect to local mongo instance
	s, err := mgo.Dial(url)

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}
	return s
}

func db(name string) *mgo.Database {
	s := session()
	return s.DB(name)
}
