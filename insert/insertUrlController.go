package insert

import (
	"encoding/json"
	"fmt"
	"hash/fnv"
	"log"
	"net/http"

	"github.com/mohammadmrd74/urlShortner/db"
)

type Url struct {
	Email string `json:"email"`
	Url   string `json:"url"`
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func InsertUrlController(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var userRequest Url

	dbConn := db.DbConnection()

	err := json.NewDecoder(request.Body).Decode(&userRequest)

	if err != nil {
		log.Fatalln("There is an error", err)
	}

	shortenUrl := fmt.Sprintf("%s%d", "https://nal.co/", hash(userRequest.Url))

	query := "INSERT INTO shortUrls (id, url, shorturl) VALUES ((select id from users where email = $1), $2, $3)"
	_, errDb := dbConn.Exec(query, userRequest.Email, userRequest.Url, shortenUrl)

	if errDb != nil {
		fmt.Println(errDb)
		writer.WriteHeader(http.StatusBadRequest)
		res := Response{false, "Duplicate"}
		json.NewEncoder(writer).Encode(&res)
		// json.NewEncoder(writer).Encode(map[string]any{"Status": false, "Message": "Duplicate user"}) //this could be done also
	} else {
		writer.WriteHeader(http.StatusCreated)
		json.NewEncoder(writer).Encode(map[string]any{"Status": true, "shortenUrl": shortenUrl})
	}

}
