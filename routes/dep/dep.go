package dep

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type Dep struct {
	IDDEP string `json:"iddep"`
	Depart string `json:"depart"`
	Sdep  string `json:"sdep"`
	Email string `json:"email"`
	Abbr string `json:"abbr"`
}

var db *sql.DB
var err error

func GetDeps(w http.ResponseWriter, r *http.Request) {

	db, err = sql.Open("mysql", "root:idEt38@tcp(127.0.0.1:3306)/phones")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	w.Header().Set("Content-Type", "application/json")
	var deps []Dep
	result, err := db.Query("SELECT iddep, depart, sdep, email, abbr from depart")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		var dep Dep
		err := result.Scan(&dep.IDDEP, &dep.Depart, &dep.Sdep, &dep.Email, &dep.Abbr)
		if err != nil {
			panic(err.Error())
		}
		deps = append(deps, dep)
	}
	json.NewEncoder(w).Encode(deps)
}
