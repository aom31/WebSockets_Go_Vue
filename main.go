package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	//init router mux
	router := mux.NewRouter()

	//start serve
	if err := http.ListenAndServe(":9100", router); err != nil {
		log.Fatal(err)
	}
}
