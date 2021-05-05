package places

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"os"
	"strconv"
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

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	db, err = sql.Open("mysql", os.Getenv("MYSQL_URL"))

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	var places []Place

	result, err := db.Query("SELECT idplace, place, work, internal, ipphone, arm, idperson, idaddr from places")

	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	for result.Next() {

		var place Place

		err := result.Scan(&place.IDPLACE, &place.Place, &place.Work, &place.Internal, &place.Ipphone, &place.Arm, &place.Idperson, &place.Idaddr)

		if err != nil {
			panic(err.Error())
		}

		places = append(places, place)
	}

	json.NewEncoder(w).Encode(places)
}

func GetOnePlace(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	db, err = sql.Open("mysql", os.Getenv("MYSQL_URL"))

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	idplace, err := strconv.Atoi(r.URL.Query().Get("idplace"))

	if err != nil || idplace < 1 {
		http.NotFound(w, r)
		return
	}

	result, err := db.Query("SELECT idplace, place, work, internal, ipphone, arm, idperson, idaddr from places WHERE idplace = ?", idplace)

	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	var place Place

	for result.Next() {

		err := result.Scan(&place.IDPLACE, &place.Place, &place.Work, &place.Internal, &place.Ipphone, &place.Arm, &place.Idperson, &place.Idaddr)

		if err != nil {
			panic(err.Error())
		}
	}

	json.NewEncoder(w).Encode(place)
}

func CreatePlace(w http.ResponseWriter, r *http.Request) {

	db, err = sql.Open("mysql", os.Getenv("MYSQL_URL"))

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
}

func UpdatePlace(w http.ResponseWriter, r *http.Request) {

	db, err = sql.Open("mysql", os.Getenv("MYSQL_URL"))

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
}

func DeletePlace(w http.ResponseWriter, r *http.Request) {

	db, err = sql.Open("mysql", os.Getenv("MYSQL_URL"))

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
}
