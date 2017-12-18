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

	var oc = controllers.NewOpinionController()
	var canvasDataController = controllers.NewCanvasDataController()

	r.GET("/canvas_data/:id", canvasDataController.GetCanvasDataById)
	r.POST("/canvas_data", canvasDataController.AddCanvasData)
	r.DELETE("/canvas_data/:id", canvasDataController.DeleteCanvasData)

	r.GET("/opinions", oc.GetAllOpinions)
	r.GET("/opinions/:id", oc.GetOpinion)
	r.POST("/opinions", oc.AddOpinion)
	r.DELETE("/opinions/:id", oc.DeleteOpinion)

	serve(r, conf.Server.Port)
}

func serve(r *httprouter.Router, port string) {
	fmt.Println("Server started on http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
