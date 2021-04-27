package addr

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"os"
)

type Addr struct {
	IDADDR   int    `json:"idaddr"`
	Addr     string `json:"addr"`
	Lat      string `json:"lat"`
	Lng      string `json:"lng"`
	Postcode string `json:"postcode"`
}

var db *sql.DB
var err error

func GetAddr(w http.ResponseWriter, r *http.Request) {

	db, err = sql.Open("mysql", os.Getenv("MYSQL_URL"))

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	w.Header().Set("Content-Type", "application/json")

	var addrs []Addr

	result, err := db.Query("SELECT idaddr, addr, lat, lng, postcode from addr ORDER BY `addr`")

	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	for result.Next() {

		var addr Addr

		err := result.Scan(&addr.IDADDR, &addr.Addr, &addr.Lat, &addr.Lng, &addr.Postcode)

		if err != nil {
			panic(err.Error())
		}

		addrs = append(addrs, addr)
	}

	json.NewEncoder(w).Encode(addrs)
}

func GetOneAddr(w http.ResponseWriter, r *http.Request) {

}

func CreateAddr(w http.ResponseWriter, r *http.Request) {

}

func UpdateAddr(w http.ResponseWriter, r *http.Request) {

}

func DeleteAddr(w http.ResponseWriter, r *http.Request) {

}
