package insert

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/mohammadmrd74/urlShortner/db"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func InsertController(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var userRequest User

	dbConn := db.DbConnection()

	err := json.NewDecoder(request.Body).Decode(&userRequest)
	fmt.Println(userRequest)
	if err != nil {
		log.Fatalln("There is an error", err)
	}

	query := "INSERT INTO users (name, email) VALUES ($1, $2)"
	_, errDb := dbConn.Exec(query, userRequest.Name, userRequest.Email)

	if errDb != nil {
		fmt.Println(errDb)
		panic(errDb)
	}

	writer.WriteHeader(http.StatusCreated)
}
