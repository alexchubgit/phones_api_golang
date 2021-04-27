package ranks

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"os"
)

type Rank struct {
	IDRANK string `json:"idrank"`
	Rank   string `json:"rank"`
}

var db *sql.DB
var err error

func GetRanks(w http.ResponseWriter, r *http.Request) {

	db, err = sql.Open("mysql", os.Getenv("MYSQL_URL"))

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	w.Header().Set("Content-Type", "application/json")

	var ranks []Rank

	result, err := db.Query("SELECT idrank, rank from ranks")

	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	for result.Next() {

		var rank Rank

		err := result.Scan(&rank.IDRANK, &rank.Rank)

		if err != nil {
			panic(err.Error())
		}

		ranks = append(ranks, rank)
	}

	json.NewEncoder(w).Encode(ranks)
}

func CreateRank(w http.ResponseWriter, r *http.Request) {

}

func UpdateRank(w http.ResponseWriter, r *http.Request) {

}

func DeleteRank(w http.ResponseWriter, r *http.Request) {

}
