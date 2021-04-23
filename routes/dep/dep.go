package dep

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Dep struct {
	IDDEP    int    `json:"iddep"`
	Depart   string `json:"depart"`
	Sdep     string `json:"sdep"`
	Email    string `json:"email"`
	Abbr     string `json:"abbr"`
	Idparent int    `json:"idparent"`
}

var db *sql.DB
var err error

func GetDeps(w http.ResponseWriter, r *http.Request) {

	db, err = sql.Open("mysql", "root:ju0jiL@tcp(127.0.0.1:3306)/phones")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	w.Header().Set("Content-Type", "application/json")
	var deps []Dep
	result, err := db.Query("SELECT iddep, depart, sdep, email, abbr, idparent from depart ORDER BY `sdep`")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		var dep Dep
		err := result.Scan(&dep.IDDEP, &dep.Depart, &dep.Sdep, &dep.Email, &dep.Abbr, &dep.Idparent)
		if err != nil {
			panic(err.Error())
		}
		deps = append(deps, dep)
	}
	json.NewEncoder(w).Encode(deps)
}

func GetOneDep(w http.ResponseWriter, r *http.Request) {

	db, err = sql.Open("mysql", "root:ju0jiL@tcp(127.0.0.1:3306)/phones")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	fmt.Println(params)
	result, err := db.Query("SELECT iddep, depart, sdep, email, abbr, idparent FROM depart WHERE iddep = ?", params["iddep"])
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var dep Dep
	for result.Next() {
		err := result.Scan(&dep.IDDEP, &dep.Depart, &dep.Sdep, &dep.Email, &dep.Abbr, &dep.Idparent)
		if err != nil {
			panic(err.Error())
		}
	}
	json.NewEncoder(w).Encode(dep)
}
