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

	db, err = sql.Open("mysql", os.Getenv("MYSQL_URL"))

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	w.Header().Set("Content-Type", "application/json")

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

}

func UpdateCert(w http.ResponseWriter, r *http.Request) {

}

func DeleteCert(w http.ResponseWriter, r *http.Request) {

}
