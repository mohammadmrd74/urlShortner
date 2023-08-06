package get

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mohammadmrd74/urlShortner/db"
)

func GetController(writer http.ResponseWriter, request *http.Request) {

	writer.Header().Set("Content-Type", "application/json")

	dbConn := db.DbConnection()

	query := "select url from shorturls where shorturl = $1"
	rows, errDb := dbConn.Query(query, request.URL.Query().Get("url"))

	defer rows.Close()
	if errDb != nil {
		fmt.Println(errDb)
		writer.WriteHeader(http.StatusBadRequest)
		// json.NewEncoder(writer).Encode(map[string]any{"Status": false, "Message": "Duplicate user"}) //this could be done also
	} else {
		var bigUrl string
		for rows.Next() {
			if err := rows.Scan(&bigUrl); err != nil {
				fmt.Println(err)
			}
		}
		writer.WriteHeader(http.StatusCreated)
		json.NewEncoder(writer).Encode(map[string]any{"Status": true, "url": bigUrl})
	}

}
