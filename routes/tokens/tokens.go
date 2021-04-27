package tokens

import (
	"database/sql"
	"net/http"
	"os"
)

type Token struct {
	IDTOKEN int    `json:""`
	Number  string `json:"number"`
	Idowner int    `json:"idowner"`
	Status  string `json:"status"`
	Comment string `json:"comment"`
}

var db *sql.DB
var err error

func GetTokens(w http.ResponseWriter, r *http.Request) {

	db, err = sql.Open("mysql", os.Getenv("MYSQL_URL"))
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	w.Header().Set("Content-Type", "application/json")

}
