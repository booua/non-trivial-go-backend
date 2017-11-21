package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    // "gopkg.in/mgo.v2"
    "fmt"
)

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/test", TestRouting).Methods("GET")
    router.HandleFunc("/opinions", opinions).Methods("GET")
    fmt.Println("Server started on http://localhost:8000")
    log.Fatal(http.ListenAndServe(":8000", router))
}


func TestRouting(w http.ResponseWriter, r *http.Request) {

    testArray := [10]string{"TEST1","TEST2","TEST3","TEST4","TEST5","TEST6","TEST7","TEST8","TEST9","TEST10"}
    json.NewEncoder(w).Encode(testArray)
}

func opinions(w http.ResponseWriter, r *http.Request) {

  type opinion struct {
    Author string `json:"author"`
    Content string `json:"content"`
    Rating int `json:"rating"`
  }

  var opinionsList = []opinion{
    {"Sąsiad Ziutka", "Ziutek kupił Grażynie książkę zamiast sobie kratę browarów. Żałuje. A ja nie, bo i tak stukam jego żonę.", 10},
    {"Trzyletni Jaś", "Ghyy ghyy bybyyy ghuuuu", 3},
    {"Samotna matka", "Mój syn nie docenił prezentu. Może dlatego, że jest niewidomy.", 5},
    {"Grafik", "Kto projektował tę stronę? Ja pierdziu, nie mogę na to patrzeć.", 0},
  }

  w.Header().Set("Access-Control-Allow-Origin", "*")
  json.NewEncoder(w).Encode(opinionsList)
}
