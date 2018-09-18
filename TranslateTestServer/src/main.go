package main

import (
	"log"
	"net/http"

	"./server"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	var server server.Server
	server.Init()

	router.HandleFunc("/translate", server.PostTranslate).Methods("POST")
	log.Fatal(http.ListenAndServe(":6000", router))
}
