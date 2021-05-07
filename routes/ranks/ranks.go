package ranks

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"os"
	"strconv"
)

type Rank struct {
	IDRANK string `json:"idrank"`
	Rank   string `json:"rank"`
}

var db *sql.DB
var err error

func GetRanks(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	db, err = sql.Open("mysql", os.Getenv("MYSQL_URL"))

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

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

func GetOneRank(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	db, err = sql.Open("mysql", os.Getenv("MYSQL_URL"))

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	idrank, err := strconv.Atoi(r.URL.Query().Get("idrank"))

	if err != nil || idrank < 1 {
		http.NotFound(w, r)
		return
	}

	result, err := db.Query("SELECT idrank, rank FROM ranks WHERE idrank = ?", idrank)

	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	var rank Rank

	for result.Next() {

		err := result.Scan(&rank.IDRANK, &rank.Rank)

		if err != nil {
			panic(err.Error())
		}
	}

	json.NewEncoder(w).Encode(rank)
}

func DeleteRank(w http.ResponseWriter, r *http.Request) {

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

func CreateRank(w http.ResponseWriter, r *http.Request) {

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

func UpdateRank(w http.ResponseWriter, r *http.Request) {

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
