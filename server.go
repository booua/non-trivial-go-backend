package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"non-trivial-go-backend/controllers"
)

func main() {
	r := httprouter.New()

	// OPINIONS
	var oc = controllers.NewOpinionController()

	r.GET("/opinion/:id", oc.GetOpinion)
	r.POST("/opinion", oc.AddOpinion)
	r.DELETE("/opinion/:id", oc.DeleteOpinion)

	fmt.Println("Server started on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
