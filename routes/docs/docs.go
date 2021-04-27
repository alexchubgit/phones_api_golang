package docs

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"os"
)

type Doc struct {
	IDDOC    int    `json:"iddoc"`
	Filename string `json:"filename"`
}

var db *sql.DB
var err error

func GetDocs(w http.ResponseWriter, r *http.Request) {

	db, err = sql.Open("mysql", os.Getenv("MYSQL_URL"))

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	w.Header().Set("Content-Type", "application/json")

	var docs []Doc

	result, err := db.Query("SELECT iddoc, filename from docs ORDER BY `filename`")

	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	for result.Next() {

		var doc Doc

		err := result.Scan(&doc.IDDOC, &doc.Filename)

		if err != nil {
			panic(err.Error())
		}

		docs = append(docs, doc)
	}

	json.NewEncoder(w).Encode(docs)
}

func CreateDoc(w http.ResponseWriter, r *http.Request) {

	db, err = sql.Open("mysql", os.Getenv("MYSQL_URL"))

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
}

func UpdateDoc(w http.ResponseWriter, r *http.Request) {

	db, err = sql.Open("mysql", os.Getenv("MYSQL_URL"))

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
}

func DeleteDoc(w http.ResponseWriter, r *http.Request) {

	db, err = sql.Open("mysql", os.Getenv("MYSQL_URL"))

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
}
