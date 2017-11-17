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
    fmt.Println("Server started")
    log.Fatal(http.ListenAndServe(":8000", router))
}


func TestRouting(w http.ResponseWriter, r *http.Request) {

    testArray := [10]string{"TEST1","TEST2","TEST3","TEST4","TEST5","TEST6","TEST7","TEST8","TEST9","TEST10"}
    json.NewEncoder(w).Encode(testArray)
}
