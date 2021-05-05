package dep

import (
	"database/sql"
	"encoding/json"
	//"fmt"
	"net/http"
	"os"
	"strconv"
	//"github.com/gorilla/mux"
)

type Dep struct {
	IDDEP    int    `json:"iddep"`
	Depart   string `json:"depart"`
	Sdep     string `json:"sdep"`
	Email    string `json:"email"`
	Abbr     string `json:"abbr"`
	Idparent int    `json:"idparent"`
	Idaddr   int    `json:"idaddr"`
}

var db *sql.DB
var err error

func GetDeps(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	db, err = sql.Open("mysql", os.Getenv("MYSQL_URL"))

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

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

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	db, err = sql.Open("mysql", os.Getenv("MYSQL_URL"))

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	iddep, err := strconv.Atoi(r.URL.Query().Get("iddep"))

	if err != nil || iddep < 1 {
		http.NotFound(w, r)
		return
	}

	//params := mux.Vars(r)
	//vals := r.URL.Query()
	//fmt.Println(params)

	result, err := db.Query("SELECT iddep, depart, sdep, email, abbr, idparent FROM depart WHERE iddep = ?", iddep)

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

func CreateDep(w http.ResponseWriter, r *http.Request) {

	db, err = sql.Open("mysql", os.Getenv("MYSQL_URL"))

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
}

func UpdateDep(w http.ResponseWriter, r *http.Request) {

	db, err = sql.Open("mysql", os.Getenv("MYSQL_URL"))

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
}

func DeleteDep(w http.ResponseWriter, r *http.Request) {

	db, err = sql.Open("mysql", os.Getenv("MYSQL_URL"))

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
}

//"SELECT * FROM depart ORDER BY `sdep`"
//'SELECT depart.*, addr.*, parent.sdep AS parent, parent.iddep AS idparent, COUNT(idperson) AS count FROM depart LEFT JOIN addr USING(idaddr) LEFT JOIN persons USING(iddep) LEFT JOIN depart AS parent ON depart.idparent=parent.iddep WHERE depart.iddep like ' + iddep + ' LIMIT 1'
//'SELECT * FROM depart WHERE sdep like "%' + query + '%" LIMIT 5'
//'INSERT INTO depart (depart , sdep, email, idaddr, idparent) VALUES (?, ?, ?, ?, ?)', [dep, sdep, email, idaddr, idparent]
//'UPDATE depart SET depart="' + dep + '", sdep="' + sdep + '", email="' + email + '", idaddr="' + idaddr + '", idparent="' + idparent + '" WHERE iddep="' + iddep + '"'
//'DELETE FROM depart WHERE iddep = "' + iddep + '"'
