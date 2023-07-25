package insert

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func InsertController(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	var userRequest User

	err := json.NewDecoder(request.Body).Decode(&userRequest)
	fmt.Println(userRequest)
	if err != nil {
		log.Fatalln("There is an error", err)
	}

	json.NewEncoder(writer).Encode(userRequest)

}
