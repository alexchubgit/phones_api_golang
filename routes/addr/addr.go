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

	result, err := db.Query("SELECT idaddr, addr, lat, lng, postcode from addr  WHERE idaddr like ? LIMIT 1", idaddr)

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

	// query, err := strconv.Atoi(r.URL.Query().Get("query"))

	// if err != nil {
	// 	http.NotFound(w, r)
	// 	return
	// }

	var addrs []Addr

	result, err := db.Query("SELECT idaddr, addr, lat, lng, postcode from addr WHERE addr like concat('%', ?, '%') LIMIT 5", query)

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

	addr := r.FormValue("addr")
	lat := r.FormValue("lat")
	lng := r.FormValue("lng")
	postcode := r.FormValue("postcode")
	//fmt.Println(addr)

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

	// stmt, err := db.Prepare("INSERT INTO addr(addr, lat, lng, postcode) VALUES(?, ?, ?, ?)")

	// if err != nil {
	// 	panic(err.Error())
	// }

	// body, err := ioutil.ReadAll(r.Body)

	// if err != nil {
	// 	panic(err.Error())
	// }

	// keyVal := make(map[string]string)

	// json.Unmarshal(body, &keyVal)

	// addr := keyVal["addr"]

	// _, err = stmt.Exec(addr)

	// if err != nil {
	// 	panic(err.Error())
	// }

	// fmt.Fprintf(w, "New address was created")

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

	//'UPDATE addr SET addr="' + addr + '", postcode="' + postcode + '", lat="' + lat + '", lng="' + lng + '" WHERE idaddr="' + idaddr + '"'

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
