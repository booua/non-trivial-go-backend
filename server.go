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

	var canvasController = controllers.NewCanvasController()
	r.GET("/canvas/:id", canvasController.GetCanvasById)
	r.POST("/canvas", canvasController.AddCanvas)
	r.DELETE("/canvas/:id", canvasController.DeleteCanvas)

	var opinionController = controllers.NewOpinionController()
	r.GET("/opinions", opinionController.GetAllOpinions)
	r.GET("/opinions/:id", opinionController.GetOpinion)
	r.POST("/opinions", opinionController.AddOpinion)
	r.DELETE("/opinions/:id", opinionController.DeleteOpinion)

	var userController = controllers.NewUserController()
	r.GET("/users", userController.GetAllUsers)
	r.GET("/users/:id", userController.GetUserById)
	r.POST("/users", userController.AddUser)
	r.DELETE("/users/:id", userController.DeleteUser)

	serve(r, conf.Server.Port)
}

func serve(r *httprouter.Router, port string) {
	fmt.Println("Server started on http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
