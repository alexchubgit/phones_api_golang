package persons

import (
	//"crypto/md5"
	"database/sql"
	"encoding/json"
	"fmt"
	//"html/template"
	"io"
	"net/http"
	"os"
	"strconv"
	//"time"
	//"github.com/gorilla/mux"
)

type Person struct {
	IDPERSON int    `json:"idperson"`
	Name     string `json:"name"`
	Date     string `json:"date"`
	File     string `json:"file"`
	Cellular string `json:"cellular"`
	Business string `json:"business"`
	Passwd   string `json:"passwd"`
	Iddep    int    `json:"iddep"`
	Idpos    int    `json:"idpos"`
	Idrank   int    `json:"idrank"`
	Idrole   int    `json:"idrole"`
}

var db *sql.DB
var err error

func GetPersons(w http.ResponseWriter, r *http.Request) {

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

	var persons []Person

	result, err := db.Query("SELECT idperson, name, date_format(date,'%Y-%m-%d') AS date, IF(file IS NULL or file = '', 'photo.png', file) as file, cellular, business, iddep, idpos, idrank, idrole FROM persons WHERE iddep like ? ORDER BY `name`", iddep)

	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	for result.Next() {

		var person Person

		err := result.Scan(&person.IDPERSON, &person.Name, &person.Date, &person.File, &person.Cellular, &person.Business, &person.Iddep, &person.Idpos, &person.Idrank, &person.Idrole)

		if err != nil {
			panic(err.Error())
		}

		persons = append(persons, person)
	}

	json.NewEncoder(w).Encode(persons)

}

func GetOnePerson(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	db, err = sql.Open("mysql", os.Getenv("MYSQL_URL"))

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	idperson, err := strconv.Atoi(r.URL.Query().Get("idperson"))

	if err != nil || idperson < 1 {
		http.NotFound(w, r)
		return
	}

	result, err := db.Query("SELECT idperson, name, date, IF(file IS NULL or file = '', 'photo.png', file) as file, cellular, business, iddep, idpos, idrank, idrole FROM persons WHERE idperson = ?", idperson)

	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	var person Person

	for result.Next() {
		err := result.Scan(&person.IDPERSON, &person.Name, &person.Date, &person.File, &person.Cellular, &person.Business, &person.Iddep, &person.Idpos, &person.Idrank, &person.Idrole)
		if err != nil {
			panic(err.Error())
		}
	}

	json.NewEncoder(w).Encode(person)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {

	db, err = sql.Open("mysql", os.Getenv("MYSQL_URL"))

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {

	// db, err = sql.Open("mysql", os.Getenv("MYSQL_URL"))

	// if err != nil {
	// 	panic(err.Error())
	// }

	// defer db.Close()

	fmt.Println("method:", r.Method)

	if r.Method == "GET" {

		// crutime := time.Now().Unix()

		// h := md5.New()

		// io.WriteString(h, strconv.FormatInt(crutime, 10))

		// token := fmt.Sprintf("%x", h.Sum(nil))

		// t, _ := template.ParseFiles("upload.gtpl")

		// t.Execute(w, token)

	} else {

		if err := r.ParseMultipartForm(32 << 20); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		file, handler, err := r.FormFile("file")

		fmt.Println(r.FormValue("name"))
		fmt.Println(r.FormValue("date"))

		if err != nil {
			fmt.Println(err)
			return
		}

		defer file.Close()

		fmt.Fprintf(w, "%v", handler.Header)
		fmt.Printf("Uploaded File: %+v\n", handler.Filename)
		fmt.Printf("File Size: %+v\n", handler.Size)
		fmt.Printf("MIME Header: %+v\n", handler.Header)

		f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)

		if err != nil {
			fmt.Println(err)
			return
		}

		defer f.Close()

		io.Copy(f, file)
	}

}

func UpdatePerson(w http.ResponseWriter, r *http.Request) {

	db, err = sql.Open("mysql", os.Getenv("MYSQL_URL"))

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
}

//"SELECT *, IF(file IS NULL or file = '', 'photo.png', file) as file, date_format(date,'%Y-%m-%d') AS date FROM persons LEFT JOIN depart USING(iddep) LEFT JOIN places USING(idperson) LEFT JOIN pos USING(idpos) LEFT JOIN ranks USING(idrank) WHERE iddep like " + iddep + " ORDER BY name"
//"SELECT *, IF(file IS NULL or file = '', 'photo.png', file) as file, date_format(date,'%Y-%m-%d') AS date FROM persons LEFT JOIN ranks USING(idrank) WHERE iddep = '0' ORDER BY name"
//"SELECT *, IF(file IS NULL or file = '', 'photo.png', file) as file, date_format(date,'%Y-%m-%d') AS date FROM persons LEFT JOIN depart USING(iddep) LEFT JOIN places USING(idperson) LEFT JOIN pos USING(idpos) LEFT JOIN ranks USING(idrank) WHERE idperson like " + idperson + " LIMIT 1"
//'SELECT * FROM persons LEFT JOIN depart USING(iddep) WHERE name like "%' + query + '%" LIMIT 5'
//"SELECT *, IF(file IS NULL OR file = '', 'photo.png', file) as file, date_format(date,'%Y-%m-%d') AS date from persons LEFT JOIN depart USING(iddep) LEFT JOIN places USING(idperson) LEFT JOIN pos USING(idpos) LEFT JOIN ranks USING(idrank) WHERE date_format(now()+interval 7 day,'%m-%d')>date_format(date,'%m-%d') AND date_format(now(),'%m-%d')<date_format(date,'%m-%d') AND iddep != 0 ORDER BY `name`"
//"SELECT *, IF(file IS NULL OR file = '', 'photo.png', file) as file, date_format(date,'%Y-%m-%d') AS date FROM persons LEFT JOIN depart USING(iddep) LEFT JOIN pos USING(idpos) LEFT JOIN ranks USING(idrank) WHERE DATE_FORMAT(date, '%m-%d') like '" + day + "' ORDER BY `name`"
//'SELECT *, date_format(date,"%Y-%m-%d") AS date FROM persons LEFT JOIN depart USING(iddep) LEFT JOIN places USING(idperson) LEFT JOIN pos USING(idpos) LEFT JOIN ranks USING(idrank) WHERE persons.cellular like "%' + query + '%" OR persons.business like "%' + query + '%" OR places.work like "%' + query + '%" AND iddep != 0 ORDER BY persons.idperson LIMIT 10'
//'SELECT *, date_format(date,"%Y-%m-%d") AS date FROM persons LEFT JOIN depart USING(iddep) LEFT JOIN places USING(idperson) LEFT JOIN pos USING(idpos) LEFT JOIN ranks USING(idrank) WHERE persons.name like "%' + query + '%" AND iddep != 0 ORDER BY persons.idperson LIMIT 10'
//'INSERT INTO persons (name, date, cellular, business, iddep, idpos, idrank, file) VALUES (?, ?, ?, ?, ?, ?, ?, ?)', [name, date, cellular, business, iddep, idpos, idrank, newname]
//'INSERT INTO persons (name, date, cellular, business, iddep, idpos, idrank, file) VALUES (?, ?, ?, ?, ?, ?, ?, ?)', [name, date, cellular, business, iddep, idpos, idrank, file]
//'UPDATE persons SET name="' + name + '", date="' + date + '", cellular="' + cellular + '", business="' + business + '", iddep="' + iddep + '", idpos="' + idpos + '", idrank="' + idrank + '" WHERE idperson="' + idperson + '"'
//'UPDATE persons SET name="' + name + '", date="' + date + '", cellular="' + cellular + '", business="' + business + '", iddep="' + iddep + '", idpos="' + idpos + '", idrank="' + idrank + '" WHERE idperson="' + idperson + '"'
//'DELETE FROM persons WHERE idperson = "' + idperson + '"'
//'UPDATE persons SET iddep="0", idpos="0", idrole="0" WHERE idperson="' + idperson + '"'
