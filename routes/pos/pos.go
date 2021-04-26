package pos

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Pos struct {
	IDPOS string `json:"idpos"`
	Pos   string `json:"pos"`
}

var db *sql.DB
var err error

func GetPoses(w http.ResponseWriter, r *http.Request) {

	db, err = sql.Open("mysql", os.Getenv("MYSQL_URL"))
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	w.Header().Set("Content-Type", "application/json")
	var poses []Pos
	result, err := db.Query("SELECT idpos, pos from pos")
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

func GetPos(w http.ResponseWriter, r *http.Request) {

	db, err = sql.Open("mysql", os.Getenv("MYSQL_URL"))
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	fmt.Println(params)
	result, err := db.Query("SELECT idpos, pos FROM pos WHERE idpos = ?", params["idpos"])
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

func CreatePos(w http.ResponseWriter, r *http.Request) {

	db, err = sql.Open("mysql", os.Getenv("MYSQL_URL"))
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO pos(pos) VALUES(?)")
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	pos := keyVal["pos"]
	_, err = stmt.Exec(pos)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "New pos was created")
}

func UpdatePos(w http.ResponseWriter, r *http.Request) {

	db, err = sql.Open("mysql", os.Getenv("MYSQL_URL"))
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	params := mux.Vars(r)
	stmt, err := db.Prepare("UPDATE pos SET pos = ? WHERE idpos = ?")
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	newPos := keyVal["pos"]
	_, err = stmt.Exec(newPos, params["idpos"])
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Pos with ID = %s was updated", params["idpos"])
}

func DeletePos(w http.ResponseWriter, r *http.Request) {

	db, err = sql.Open("mysql", os.Getenv("MYSQL_URL"))
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	params := mux.Vars(r)
	stmt, err := db.Prepare("DELETE FROM pos WHERE idpos = ?")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(params["idpos"])
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Pos with ID = %s was deleted", params["idpos"])
}
