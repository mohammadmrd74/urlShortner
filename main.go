package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mohammadmrd74/urlShortner/insert"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/user", insert.InsertController).Methods("POST")

	err := http.ListenAndServe(":8080", router)
	fmt.Println(err)
	if err != nil {
		log.Fatalln("There's an error with the server,", err)
	} else {
		fmt.Println("Listening on port 8080")
	}
}
