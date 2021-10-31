package ranks

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Rank struct {
	IDRANK int    `json:"idrank"`
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

	result, err := db.Query("SELECT idrank, rank FROM ranks ORDER BY `rank`")

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

func GetListRank(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	db, err = sql.Open("mysql", os.Getenv("MYSQL_URL"))

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	query := r.URL.Query().Get("query")

	var ranks []Rank

	result, err := db.Query("SELECT idrank, rank FROM ranks WHERE rank LIKE concat('%', ?, '%') LIMIT 5", query)

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

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	db, err = sql.Open("mysql", os.Getenv("MYSQL_URL"))

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	if r.Method != "POST" {
		fmt.Println("Not Post")
		return
	}

	var cr Rank

	err := json.NewDecoder(r.Body).Decode(&cr)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	rank := cr.Rank

	if rank == "" {
		fmt.Println("Feild is empty")
	}

	res, err := db.Exec("INSERT INTO ranks (rank) VALUES (?)", rank)
	if err != nil {
		panic(err)
	}

	lastId, err := res.LastInsertId()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The last inserted row id: %d\n", lastId)

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

	if r.Method != "PUT" {
		fmt.Println("Not PUT")
		return
	}

	var er Rank

	err := json.NewDecoder(r.Body).Decode(&er)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	rank := er.Rank
	idrank := er.IDRANK

	_, err = db.Exec("UPDATE ranks SET rank = ? WHERE idrank = ?", rank, idrank)

	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "Rank with ID = %s was updated", strconv.Itoa(idrank))
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

	idrank, err := strconv.Atoi(r.URL.Query().Get("idrank"))

	if err != nil || idrank < 1 {
		http.NotFound(w, r)
		return
	}

	stmt, err := db.Prepare("DELETE FROM ranks WHERE idrank = ?")

	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(idrank)

	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "Rank with ID = %s was deleted", strconv.Itoa(idrank))
}
