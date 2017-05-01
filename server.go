package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Datapoint struct {
	ID          int       `json:"id"`
	Temperature float32   `json:"temperature"`
	Humidity    float32   `json:"humidity"`
	Pressure    float32   `json:"pressure"`
	Timestamp   time.Time `json:"timestamp"`
}

type Timeline struct {
	Data []Datapoint `json:"data"`
}

func ApiHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	a := []byte{}

	if id == "all" {
		w.Write(a)
	} else if _, err := strconv.Atoi(id); err != nil {
		w.Write(a)
	} else {
		w.Write(a)
	}
}

func startServer() {
	router := mux.NewRouter()
	fs := http.FileServer(http.Dir("static"))

	router.Handle("/", fs)
	router.HandleFunc("/api/{id}", ApiHandler)

	http.Handle("/", router)
	http.ListenAndServe(":3000", nil)
}
