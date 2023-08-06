package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mohammadmrd74/urlShortner/get"
	"github.com/mohammadmrd74/urlShortner/insert"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/api/user", insert.InsertController).Methods("POST")
	router.HandleFunc("/api/addurl", insert.InsertUrlController).Methods("POST")
	router.HandleFunc("/api/url", get.GetController).Methods("GET")

	err1 := http.ListenAndServe(":8080", router)

	if err1 != nil {
		log.Fatalln("There's an error with the server,", err1)
	}
}
