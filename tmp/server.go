package main

import (
	"database/sql"
	"encoding/json"
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

var ChunkReader func(...interface{}) (*sql.Rows, error)

func startServer() {
	router := mux.NewRouter()
	fs := http.FileServer(http.Dir("static"))

	var closeReader func() error
	closeReader, ChunkReader = CreateChunkReader()
	defer closeReader()

	router.Handle("/", fs)
	router.HandleFunc("/api/{id}", ApiHandler)

	http.Handle("/", router)
	http.ListenAndServe(":3000", nil)
}

func ApiHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	a := []byte{}

	if id == "all" {
		timeline := Timeline{}
		rows, err := ChunkReader()
		defer rows.Close()
		logErr(err)

		for rows.Next() {
			datapoint := Datapoint{}

			err := rows.Scan(
				&datapoint.ID,
				&datapoint.Temperature,
				&datapoint.Humidity,
				&datapoint.Pressure,
				&datapoint.Timestamp)

			logErr(err)
			timeline.Data = append(timeline.Data, datapoint)
		}

		j, err := json.Marshal(timeline)
		logErr(err)

		w.Write(j)
	} else if _, err := strconv.Atoi(id); err != nil {
		w.Write(a)
	} else {
		w.Write(a)
	}
}
