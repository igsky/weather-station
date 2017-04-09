package main

import (
	"io"
	"log"
	"net/http"
)

type webHandler struct{}

// HelloServer : "/" route function
func HelloServer(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

func (wh webHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

func main() {

	log.Fatal(http.ListenAndServe(":80", &webHandler{}))
}
