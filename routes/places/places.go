package places

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Place struct {
	IDPLACE  string  `json:"idplace"`
	Place    string  `json:"place"`
	Work     string  `json:"work"`
	Internal string  `json:"internal"`
	Ipphone  string  `json:"ipphone"`
	Arm      string  `json:"arm"`
	Idperson int     `json:"idperson"`
	Idaddr   int     `json:"idaddr"`
	Name     *string `json:"name"`
	Addr     *string `json:"addr"`
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

	idaddr, err := strconv.Atoi(r.URL.Query().Get("idaddr"))

	if err != nil || idaddr < 1 {
		http.NotFound(w, r)
		return
	}

	var places []Place

	result, err := db.Query("SELECT idplace, place, work, internal, ipphone, arm, idperson, idaddr, name FROM places LEFT JOIN persons USING(idperson) WHERE idaddr = ? ORDER BY place;", idaddr)

	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	for result.Next() {

		var place Place

		err := result.Scan(&place.IDPLACE, &place.Place, &place.Work, &place.Internal, &place.Ipphone, &place.Arm, &place.Idperson, &place.Idaddr, &place.Name)

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

	result, err := db.Query("SELECT idplace, place, work, internal, ipphone, arm, idperson, idaddr, name, addr FROM places LEFT JOIN persons USING(idperson) LEFT JOIN addr USING(idaddr) WHERE idplace LIKE ? LIMIT 1", idplace)

	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	var place Place

	for result.Next() {

		err := result.Scan(&place.IDPLACE, &place.Place, &place.Work, &place.Internal, &place.Ipphone, &place.Arm, &place.Idperson, &place.Idaddr, &place.Name, &place.Addr)

		if err != nil {
			panic(err.Error())
		}
	}

	json.NewEncoder(w).Encode(place)
}

func CreatePlace(w http.ResponseWriter, r *http.Request) {

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

	place := r.FormValue("place")
	work := r.FormValue("work")
	internal := r.FormValue("internal")
	ipphone := r.FormValue("ipphone")
	arm := r.FormValue("arm")
	idaddr := r.FormValue("idaddr")
	idperson := r.FormValue("idperson")

	if place == "" {
		fmt.Println("Feild is empty")
	}

	res, err := db.Exec("INSERT INTO places (place, work, internal, ipphone, arm, idaddr, idperson) VALUES (?, ?, ?, ?, ?, ?, ?)", place, work, internal, ipphone, arm, idaddr, idperson)
	if err != nil {
		panic(err)
	}

	lastId, err := res.LastInsertId()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The last inserted row id: %d\n", lastId)

}

func UpdatePlace(w http.ResponseWriter, r *http.Request) {

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

	idplace := r.FormValue("idplace")
	place := r.FormValue("place")
	work := r.FormValue("work")
	internal := r.FormValue("internal")
	ipphone := r.FormValue("ipphone")
	arm := r.FormValue("arm")
	idaddr := r.FormValue("idaddr")
	idperson := r.FormValue("idperson")

	if place == "" {
		fmt.Println("Feild is empty")
	}

	_, err := db.Exec("UPDATE places SET place = ?, work = ?, internal = ?, ipphone = ?, arm = ?, idaddr = ?, idperson = ? WHERE idplace = ?", place, work, internal, ipphone, arm, idaddr, idperson, idplace)

	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "Rank with ID = %s was updated", idplace)

}

func DeletePlace(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
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

	stmt, err := db.Prepare("DELETE FROM places WHERE idplace = ?")

	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(idplace)

	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "Place with ID = %s was deleted", strconv.Itoa(idplace))

}

func DeletePersonFromPlace(w http.ResponseWriter, r *http.Request) {

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

	idplace := r.FormValue("idplace")

	if idplace == "" {
		fmt.Println("Feild is empty")
	}

	_, err := db.Exec("UPDATE places SET idperson='0' WHERE idplace = ?", idplace)

	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "Rank with ID = %s was updated", idplace)

}
