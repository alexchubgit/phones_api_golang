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

func GetDatesWeek(w http.ResponseWriter, r *http.Request) {

	// const getDatesWeek = () => {
	//     return new Promise((resolve, reject) => {
	//         pool.query("SELECT *, IF(file IS NULL OR file = '', 'photo.png', file) as file, date_format(date,'%Y-%m-%d') AS date from persons LEFT JOIN depart USING(iddep) LEFT JOIN places USING(idperson) LEFT JOIN pos USING(idpos) LEFT JOIN ranks USING(idrank) WHERE date_format(now()+interval 7 day,'%m-%d')>date_format(date,'%m-%d') AND date_format(now(),'%m-%d')<date_format(date,'%m-%d') AND iddep != 0 ORDER BY `name`", (err, results) => {
	//             if (err) {
	//                 return reject(err);
	//             }
	//             return resolve(results);
	//         });
	//     });
	// }
}

func GetDatesToday(w http.ResponseWriter, r *http.Request) {

	// const getDatesToday = (day) => {
	//     return new Promise((resolve, reject) => {
	//         pool.query("SELECT *, IF(file IS NULL OR file = '', 'photo.png', file) as file, date_format(date,'%Y-%m-%d') AS date FROM persons LEFT JOIN depart USING(iddep) LEFT JOIN pos USING(idpos) LEFT JOIN ranks USING(idrank) WHERE DATE_FORMAT(date, '%m-%d') like '" + day + "' ORDER BY `name`", (err, results) => {
	//             if (err) {
	//                 return reject(err);
	//             }
	//             return resolve(results);
	//         });
	//     });
	// }
}

func Search(w http.ResponseWriter, r *http.Request) {

	// const seachPersonByPhone = (query) => {
	//     return new Promise((resolve, reject) => {
	//         pool.query('SELECT *, date_format(date,"%Y-%m-%d") AS date FROM persons LEFT JOIN depart USING(iddep) LEFT JOIN places USING(idperson) LEFT JOIN pos USING(idpos) LEFT JOIN ranks USING(idrank) WHERE persons.cellular like "%' + query + '%" OR persons.business like "%' + query + '%" OR places.work like "%' + query + '%" AND iddep != 0 ORDER BY persons.idperson LIMIT 10', (err, results) => {
	//             if (err) {
	//                 return reject(err);
	//             }
	//             return resolve(results);
	//         });
	//     });
	// }

	// const seachPersonByName = (query) => {
	//     return new Promise((resolve, reject) => {
	//         pool.query('SELECT *, date_format(date,"%Y-%m-%d") AS date FROM persons LEFT JOIN depart USING(iddep) LEFT JOIN places USING(idperson) LEFT JOIN pos USING(idpos) LEFT JOIN ranks USING(idrank) WHERE persons.name like "%' + query + '%" AND iddep != 0 ORDER BY persons.idperson LIMIT 10', (err, results) => {
	//             if (err) {
	//                 return reject(err);
	//             }
	//             return resolve(results);
	//         });
	//     });
	// }

}

func Dismiss(w http.ResponseWriter, r *http.Request) {

	// const dismissPerson = (idperson) => {
	//     return new Promise((resolve, reject) => {
	//         pool.query('UPDATE persons SET iddep="0", idpos="0", idrole="0" WHERE idperson="' + idperson + '"', (err, results) => {
	//             if (err) {
	//                 return reject(err);
	//             }
	//             return resolve(results);
	//         });
	//     });
	// }

}

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

	result, err := db.Query("SELECT idperson, name, date_format(date,'%Y-%m-%d') AS date, IF(file IS NULL or file = '', 'photo.png', file) as file, cellular, business, iddep, idpos, idrank, idrole FROM persons LEFT JOIN depart USING(iddep) LEFT JOIN places USING(idperson) LEFT JOIN pos USING(idpos) LEFT JOIN ranks USING(idrank) WHERE iddep like ? ORDER BY `name`", iddep)

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

	result, err := db.Query("SELECT idperson, name, date_format(date,'%Y-%m-%d') AS date, IF(file IS NULL or file = '', 'photo.png', file) as file, cellular, business, iddep, idpos, idrank, idrole FROM persons LEFT JOIN depart USING(iddep) LEFT JOIN places USING(idperson) LEFT JOIN pos USING(idpos) LEFT JOIN ranks USING(idrank) WHERE idperson = ? LIMIT 1", idperson)

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

func GetListPersons(w http.ResponseWriter, r *http.Request) {

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

	var persons []Person

	result, err := db.Query("SELECT idperson, name, cellular, business, iddep, idpos, idrank, idrole, IF(file IS NULL or file = '', 'photo.png', file) as file, date_format(date,'%Y-%m-%d') AS date FROM persons LEFT JOIN depart USING(iddep) WHERE name LIKE concat('%', ?, '%') LIMIT 5", query)

	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	for result.Next() {

		var person Person

		err := result.Scan(&person.IDPERSON, &person.Name, &person.Cellular, &person.Business, &person.Iddep, &person.Idpos, &person.Idrank, &person.Idrole, &person.File, &person.Date)

		if err != nil {
			panic(err.Error())
		}

		persons = append(persons, person)
	}

	json.NewEncoder(w).Encode(persons)

}

func GetDismissed(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	db, err = sql.Open("mysql", os.Getenv("MYSQL_URL"))

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	var persons []Person

	result, err := db.Query("SELECT idperson, name, cellular, business, iddep, idpos, idrank, idrole, IF(file IS NULL or file = '', 'photo.png', file) as file, date_format(date,'%Y-%m-%d') AS date FROM persons LEFT JOIN ranks USING(idrank) WHERE iddep = '0' ORDER BY name")

	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	for result.Next() {

		var person Person

		err := result.Scan(&person.IDPERSON, &person.Name, &person.Cellular, &person.Business, &person.Iddep, &person.Idpos, &person.Idrank, &person.Idrole, &person.File, &person.Date)

		if err != nil {
			panic(err.Error())
		}

		persons = append(persons, person)
	}

	json.NewEncoder(w).Encode(persons)

}

func DeletePerson(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
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

	stmt, err := db.Prepare("DELETE FROM persons WHERE idperson = ?")

	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(idperson)

	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "Pos with ID = %s was deleted", strconv.Itoa(idperson))
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {

	// const addPersonWithoutFile = (name, date, cellular, business, iddep, idpos, idrank) => {
	//     return new Promise((resolve, reject) => {
	//         pool.query('INSERT INTO persons (name, date, cellular, business, iddep, idpos, idrank, file) VALUES (?, ?, ?, ?, ?, ?, ?, ?)', [name, date, cellular, business, iddep, idpos, idrank, newname], (err, results) => {
	//             if (err) {
	//                 return reject(err);
	//             }
	//             return resolve(results);
	//         });
	//     });
	// }

	// const addPerson = (name, date, cellular, business, iddep, idpos, idrank, file) => {
	//     return new Promise((resolve, reject) => {
	//         pool.query('INSERT INTO persons (name, date, cellular, business, iddep, idpos, idrank, file) VALUES (?, ?, ?, ?, ?, ?, ?, ?)', [name, date, cellular, business, iddep, idpos, idrank, file], (err, results) => {
	//             if (err) {
	//                 return reject(err);
	//             }
	//             return resolve(results);
	//         });
	//     });
	// }

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

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

	// const updPersonWithoutFile = (name, date, cellular, business, iddep, idpos, idrank) => {
	//     return new Promise((resolve, reject) => {
	//         pool.query('UPDATE persons SET name="' + name + '", date="' + date + '", cellular="' + cellular + '", business="' + business + '", iddep="' + iddep + '", idpos="' + idpos + '", idrank="' + idrank + '" WHERE idperson="' + idperson + '"', (err, results) => {
	//             if (err) {
	//                 return reject(err);
	//             }
	//             return resolve(results);
	//         });
	//     });
	// }

	// const updPerson = (name, date, cellular, business, iddep, idpos, idrank, file) => {
	//     return new Promise((resolve, reject) => {
	//         pool.query('UPDATE persons SET name="' + name + '", date="' + date + '", cellular="' + cellular + '", business="' + business + '", iddep="' + iddep + '", idpos="' + idpos + '", idrank="' + idrank + '" WHERE idperson="' + idperson + '"', (err, results) => {
	//             if (err) {
	//                 return reject(err);
	//             }
	//             return resolve(results);
	//         });
	//     });
	// }

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
