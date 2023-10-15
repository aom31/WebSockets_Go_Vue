package main

import (
	"gowebsocketvue/handler"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	//init router mux
	router := mux.NewRouter()
	router.HandleFunc("/socket", handler.WsEndpoint)

	//start serve
	log.Println("starting server running")
	if err := http.ListenAndServe(":9100", router); err != nil {
		log.Fatal(err)
	}
}
