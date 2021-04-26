package places

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"os"
)

type Place struct {
	IDPLACE  string `json:"idplace"`
	Place    string `json:"place"`
	Work     string `json:"work"`
	Internal string `json:"internal"`
	Ipphone  string `json:"ipphone"`
	Arm      string `json:"arm"`
	Idperson int    `json:"idperson"`
	Idaddr   int    `json:"idaddr"`
}

var db *sql.DB
var err error

func GetPlaces(w http.ResponseWriter, r *http.Request) {

	db, err = sql.Open("mysql", os.Getenv("MYSQL_URL"))
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	w.Header().Set("Content-Type", "application/json")
	var places []Place
	result, err := db.Query("SELECT idplace, place from places")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		var place Place
		err := result.Scan(&place.IDPLACE, &place.Place)
		if err != nil {
			panic(err.Error())
		}
		places = append(places, place)
	}
	json.NewEncoder(w).Encode(places)
}
