package certs

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"os"
)

type Cert struct {
	IDCERT    int    `json:"idcert"`
	Filename  string `json:"filename"`
	Startdate string `json:"startdate"`
	Enddate   string `json:"enddate"`
}

var db *sql.DB
var err error

func GetCert(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	db, err = sql.Open("mysql", os.Getenv("MYSQL_URL"))

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	var certs []Cert

	result, err := db.Query("SELECT idcert, filename, startdate, enddate from certs ORDER BY `startdate`")

	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	for result.Next() {

		var cert Cert

		err := result.Scan(&cert.IDCERT, &cert.Filename, &cert.Startdate, &cert.Enddate)

		if err != nil {
			panic(err.Error())
		}

		certs = append(certs, cert)
	}

	json.NewEncoder(w).Encode(certs)
}

func CreateCert(w http.ResponseWriter, r *http.Request) {

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

func UpdateCert(w http.ResponseWriter, r *http.Request) {

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

func DeleteCert(w http.ResponseWriter, r *http.Request) {

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
