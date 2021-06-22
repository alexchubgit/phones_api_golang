package addr

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
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

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	db, err = sql.Open("mysql", os.Getenv("MYSQL_URL"))

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	var addrs []Addr

	result, err := db.Query("SELECT idaddr, addr, lat, lng, postcode FROM addr ORDER BY `addr`")

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

	result, err := db.Query("SELECT idaddr, addr, lat, lng, postcode FROM addr WHERE idaddr like ? LIMIT 1", idaddr)

	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	var addr Addr

	for result.Next() {

		err := result.Scan(&addr.IDADDR, &addr.Addr, &addr.Lat, &addr.Lng, &addr.Postcode)

		if err != nil {
			panic(err.Error())
		}
	}

	json.NewEncoder(w).Encode(addr)
}

func GetListAddr(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	db, err = sql.Open("mysql", os.Getenv("MYSQL_URL"))

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	query := r.URL.Query().Get("query")

	var addrs []Addr

	result, err := db.Query("SELECT idaddr, addr, lat, lng, postcode FROM addr WHERE addr like concat('%', ?, '%') LIMIT 5", query)

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

func CreateAddr(w http.ResponseWriter, r *http.Request) {

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

	var ca Addr

	err := json.NewDecoder(r.Body).Decode(&ca)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	addr := ca.Addr
	lat := ca.Lat
	lng := ca.Lng
	postcode := ca.Postcode

	if addr == "" {
		fmt.Println("Feild is empty")
	}

	res, err := db.Exec("INSERT INTO addr (addr, lat, lng, postcode) VALUES (?, ?, ?, ?)", addr, lat, lng, postcode)
	if err != nil {
		panic(err)
	}

	lastId, err := res.LastInsertId()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The last inserted row id: %d\n", lastId)
}

func UpdateAddr(w http.ResponseWriter, r *http.Request) {

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

	var ea Addr

	err := json.NewDecoder(r.Body).Decode(&ea)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	idaddr := ea.IDADDR
	addr := ea.Addr
	lat := ea.Lat
	lng := ea.Lng
	postcode := ea.Postcode

	if addr == "" {
		fmt.Println("Feild is empty")
	}

	_, err = db.Exec("UPDATE addr SET addr = ?, postcode = ?, lat = ?, lng = ? WHERE idaddr = ?", addr, postcode, lat, lng, idaddr)

	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "Addr with ID = %s was updated", strconv.Itoa(idaddr))

}

func DeleteAddr(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
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

	stmt, err := db.Prepare("DELETE FROM addr WHERE idaddr = ?")

	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(idaddr)

	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "Addr with ID = %s was deleted", strconv.Itoa(idaddr))

}
