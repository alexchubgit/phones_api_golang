package tokens

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"os"
)

type Token struct {
	IDTOKEN int    `json:"idtoken"`
	Number  string `json:"number"`
	Status  string `json:"status"`
	Idowner int    `json:"idowner"`
}

var db *sql.DB
var err error

func GetTokens(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	db, err = sql.Open("mysql", os.Getenv("MYSQL_URL"))

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	var tokens []Token

	result, err := db.Query("SELECT idtoken, number, idowner, status from tokens ORDER BY `number`")

	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	for result.Next() {

		var token Token

		err := result.Scan(&token.IDTOKEN, &token.Number, &token.Status, &token.Idowner)

		if err != nil {
			panic(err.Error())
		}

		tokens = append(tokens, token)
	}

	json.NewEncoder(w).Encode(tokens)
}

func CreateToken(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	db, err = sql.Open("mysql", os.Getenv("MYSQL_URL"))

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
}

func UpdateToken(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	db, err = sql.Open("mysql", os.Getenv("MYSQL_URL"))

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
}

func DeleteToken(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	db, err = sql.Open("mysql", os.Getenv("MYSQL_URL"))

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
}
