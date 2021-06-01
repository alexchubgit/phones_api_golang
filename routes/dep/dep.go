package dep

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Dep struct {
	IDDEP    int     `json:"iddep"`
	Depart   string  `json:"depart"`
	Sdep     string  `json:"sdep"`
	Email    string  `json:"email"`
	Abbr     string  `json:"abbr"`
	Postcode *string `json:"postcode"`
	Addr     *string `json:"addr"`
	Idparent int     `json:"idparent"`
	Idaddr   int     `json:"idaddr"`
	Count    int     `json:"count"`
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

	result, err := db.Query("SELECT iddep, depart, sdep, addr, email, postcode, abbr, idparent, idaddr, COUNT(idperson) AS count FROM depart LEFT JOIN addr USING(idaddr) LEFT JOIN persons USING(iddep) WHERE iddep = ?", iddep)

	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	var dep Dep

	for result.Next() {

		err := result.Scan(&dep.IDDEP, &dep.Depart, &dep.Sdep, &dep.Addr, &dep.Email, &dep.Postcode, &dep.Abbr, &dep.Idparent, &dep.Idaddr, &dep.Count)

		if err != nil {
			panic(err.Error())
		}

	}

	json.NewEncoder(w).Encode(dep)

	//'SELECT depart.*, addr.*, parent.sdep AS parent, parent.iddep AS idparent, COUNT(idperson) AS count FROM depart LEFT JOIN addr USING(idaddr) LEFT JOIN persons USING(iddep) LEFT JOIN depart AS parent ON depart.idparent=parent.iddep WHERE depart.iddep like ' + iddep + ' LIMIT 1'

}

func GetListDep(w http.ResponseWriter, r *http.Request) {

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

	var deps []Dep

	result, err := db.Query("SELECT iddep, depart, sdep, email, abbr, idparent, idaddr from depart WHERE sdep like concat('%', ?, '%') LIMIT 5", query)

	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	for result.Next() {

		var dep Dep

		err := result.Scan(&dep.IDDEP, &dep.Depart, &dep.Sdep, &dep.Email, &dep.Abbr, &dep.Idparent, &dep.Idaddr)

		if err != nil {
			panic(err.Error())
		}

		deps = append(deps, dep)
	}

	json.NewEncoder(w).Encode(deps)

}

func CreateDep(w http.ResponseWriter, r *http.Request) {

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

	var cd Dep

	err := json.NewDecoder(r.Body).Decode(&cd)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	depart := cd.Depart
	sdep := cd.Sdep
	email := cd.Email
	abbr := cd.Abbr
	idaddr := cd.Idaddr
	idparent := cd.Idparent

	fmt.Println(depart)
	fmt.Println(sdep)
	fmt.Println(email)
	fmt.Println(abbr)
	fmt.Println(idaddr)
	fmt.Println(idparent)

	if depart == "" {
		fmt.Println("Feild is empty")
	}

	res, err := db.Exec("INSERT INTO depart (depart , sdep, email, abbr, idaddr, idparent) VALUES (?, ?, ?, ?, ?, ?)", depart, sdep, email, abbr, idaddr, idparent)
	if err != nil {
		panic(err)
	}

	lastId, err := res.LastInsertId()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The last inserted row id: %d\n", lastId)
}

func UpdateDep(w http.ResponseWriter, r *http.Request) {

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

	var ed Dep

	err := json.NewDecoder(r.Body).Decode(&ed)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	iddep := ed.IDDEP
	depart := ed.Depart
	sdep := ed.Sdep
	email := ed.Email
	abbr := ed.Abbr
	idaddr := ed.Idaddr
	idparent := ed.Idparent

	fmt.Println(iddep)
	fmt.Println(depart)
	fmt.Println(sdep)
	fmt.Println(email)
	fmt.Println(abbr)
	fmt.Println(idaddr)
	fmt.Println(idparent)

	if depart == "" {
		fmt.Println("Feild is empty")
	}

	_, err = db.Exec("UPDATE depart SET depart = ?, sdep = ?, email = ?, abbr = ?, idaddr = ?, idparent = ? WHERE iddep = ?", depart, sdep, email, abbr, idaddr, idparent, iddep)

	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "Dep with ID = %s was updated", strconv.Itoa(iddep))

}

func DeleteDep(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
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

	stmt, err := db.Prepare("DELETE FROM depart WHERE iddep = ?")

	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(iddep)

	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "Dep with ID = %s was deleted", strconv.Itoa(iddep))
}
