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

type Response struct {
	Status     bool   `json:"status"`
	ErrMessage string `json:"errorMessage"`
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
		writer.WriteHeader(http.StatusBadRequest)
		res := Response{false, "Duplicate"}
		json.NewEncoder(writer).Encode(&res)
		// json.NewEncoder(writer).Encode(map[string]any{"Status": false, "Message": "Duplicate user"}) //this could be done also
	}

	writer.WriteHeader(http.StatusCreated)
}
