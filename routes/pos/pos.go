package pos

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Pos struct {
	IDPOS string `json:"idpos"`
	Pos   string `json:"pos"`
}

var db *sql.DB
var err error

func GetPoses(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	db, err = sql.Open("mysql", os.Getenv("MYSQL_URL"))

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	var poses []Pos

	result, err := db.Query("SELECT idpos, pos FROM pos ORDER BY `pos`")

	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	for result.Next() {

		var pos Pos

		err := result.Scan(&pos.IDPOS, &pos.Pos)

		if err != nil {
			panic(err.Error())
		}

		poses = append(poses, pos)

	}
	json.NewEncoder(w).Encode(poses)
}

func GetOnePos(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	db, err = sql.Open("mysql", os.Getenv("MYSQL_URL"))

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	idpos, err := strconv.Atoi(r.URL.Query().Get("idpos"))

	if err != nil || idpos < 1 {
		http.NotFound(w, r)
		return
	}

	result, err := db.Query("SELECT idpos, pos FROM pos WHERE idpos LIKE ? LIMIT 1", idpos)

	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	var pos Pos

	for result.Next() {

		err := result.Scan(&pos.IDPOS, &pos.Pos)

		if err != nil {
			panic(err.Error())
		}
	}

	json.NewEncoder(w).Encode(pos)
}

func GetListPos(w http.ResponseWriter, r *http.Request) {

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

	var poses []Pos

	result, err := db.Query("SELECT idpos, pos FROM pos WHERE pos LIKE concat('%', ?, '%') LIMIT 5", query)

	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	for result.Next() {

		var pos Pos

		err := result.Scan(&pos.IDPOS, &pos.Pos)

		if err != nil {
			panic(err.Error())
		}

		poses = append(poses, pos)

	}
	json.NewEncoder(w).Encode(poses)
}

func CreatePos(w http.ResponseWriter, r *http.Request) {

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

	pos := r.FormValue("pos")
	fmt.Println(pos)

	if pos == "" {
		fmt.Println("Feild is empty")

	}

	res, err := db.Exec("INSERT INTO pos(pos) VALUES(?)", pos)
	if err != nil {
		panic(err)
	}

	lastId, err := res.LastInsertId()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The last inserted row id: %d\n", lastId)

}

func UpdatePos(w http.ResponseWriter, r *http.Request) {

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

	pos := r.FormValue("pos")
	idpos := r.FormValue("idpos")

	_, err := db.Exec("UPDATE pos SET pos = ? WHERE idpos = ?", pos, idpos)

	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "Pos with ID = %s was updated", idpos)
}

func DeletePos(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	db, err = sql.Open("mysql", os.Getenv("MYSQL_URL"))

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	idpos, err := strconv.Atoi(r.URL.Query().Get("idpos"))

	if err != nil || idpos < 1 {
		http.NotFound(w, r)
		return
	}

	stmt, err := db.Prepare("DELETE FROM pos WHERE idpos = ?")

	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(idpos)

	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "Pos with ID = %s was deleted", strconv.Itoa(idpos))
}
