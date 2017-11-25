package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"non-trivial-go-backend/conf"
	"non-trivial-go-backend/controllers"
)

func main() {
	r := httprouter.New()

	// OPINIONS
	var oc = controllers.NewOpinionController()

	r.GET("/opinion/:id", oc.GetOpinion)
	r.POST("/opinion", oc.AddOpinion)
	r.DELETE("/opinion/:id", oc.DeleteOpinion)

	serve(r, conf.Server.Port)
}

func serve(r *httprouter.Router, port string) {
	fmt.Println("Server started on http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
