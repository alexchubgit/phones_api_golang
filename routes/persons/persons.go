package persons

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"
)

type Person struct {
	IDPERSON int     `json:"idperson"`
	Name     string  `json:"name"`
	Date     string  `json:"date"`
	File     string  `json:"file"`
	Cellular string  `json:"cellular"`
	Business string  `json:"business"`
	Passwd   string  `json:"passwd"`
	Iddep    int     `json:"iddep"`
	Idpos    int     `json:"idpos"`
	Idrank   int     `json:"idrank"`
	Idrole   int     `json:"idrole"`
	Pos      *string `json:"pos"`
	Rank     *string `json:"rank"`
	Depart   *string `json:"depart"`
	Sdep     *string `json:"sdep"`
	Place    *string `json:"place"`
	Work     *string `json:"work"`
	Internal *string `json:"internal"`
	Ipphone  *string `json:"ipphone"`
	ARM      *string `json:"arm"`
}

var db *sql.DB
var err error

func GetDatesWeek(w http.ResponseWriter, r *http.Request) {

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

	result, err := db.Query("SELECT idperson, name, date_format(date,'%Y-%m-%d') AS date, IF(file IS NULL or file = '', 'photo.png', file) as file, pos, rank FROM persons LEFT JOIN pos USING(idpos) LEFT JOIN ranks USING(idrank) WHERE date_format(now()+interval 7 day,'%m-%d')>date_format(date,'%m-%d') AND date_format(now(),'%m-%d')<date_format(date,'%m-%d') AND iddep != 0 ORDER BY `name`")

	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	for result.Next() {

		var person Person

		err := result.Scan(&person.IDPERSON, &person.Name, &person.Date, &person.File, &person.Pos, &person.Rank)

		if err != nil {
			panic(err.Error())
		}

		persons = append(persons, person)
	}

	json.NewEncoder(w).Encode(persons)

}

func GetDatesToday(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	db, err = sql.Open("mysql", os.Getenv("MYSQL_URL"))

	if err != nil {
		panic(err.Error())
	}

	//defer db.Close()

	now := time.Now()
	date := now.Format("01-02")
	//month := strconv.Itoa(int(now.Month()))
	//day := strconv.Itoa(now.Day())

	//fmt.Println(now.Format("01-02"))

	var persons []Person

	result, err := db.Query("SELECT idperson, name, date_format(date,'%Y-%m-%d') AS date, IF(file IS NULL or file = '', 'photo.png', file) as file, pos, rank FROM persons LEFT JOIN pos USING(idpos) LEFT JOIN ranks USING(idrank) WHERE DATE_FORMAT(date, '%m-%d') LIKE ? ORDER BY `name`", date)

	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	for result.Next() {

		var person Person

		err := result.Scan(&person.IDPERSON, &person.Name, &person.Date, &person.File, &person.Pos, &person.Rank)

		if err != nil {
			panic(err.Error())
		}

		persons = append(persons, person)
	}

	json.NewEncoder(w).Encode(persons)

}

func Search(w http.ResponseWriter, r *http.Request) {

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
	//fmt.Println(query)

	regexp_alph := `^[ЁёА-я]{2,}\s?[ЁёА-я]{2,}?\s?[ЁёА-я]{2,}?$`
	regexp_num := `^[[:digit:]]{2,11}$`

	var IsLetter = regexp.MustCompile(regexp_alph).MatchString
	var IsNumber = regexp.MustCompile(regexp_num).MatchString

	if IsLetter(query) {

		var persons []Person

		result, err := db.Query("SELECT idperson, name, sdep FROM persons LEFT JOIN depart USING(iddep) WHERE name LIKE concat('%', ?, '%') AND iddep != 0 LIMIT 10", query)

		if err != nil {
			panic(err.Error())
		}

		defer result.Close()

		for result.Next() {

			var person Person

			err := result.Scan(&person.IDPERSON, &person.Name, &person.Sdep)

			if err != nil {
				panic(err.Error())
			}

			persons = append(persons, person)
		}

		json.NewEncoder(w).Encode(persons)

	} else if IsNumber(query) {

		var persons []Person

		result, err := db.Query("SELECT idperson, name, sdep FROM persons LEFT JOIN depart USING(iddep) LEFT JOIN places USING(idperson) WHERE cellular LIKE concat('%', ?, '%')  OR business LIKE concat('%', ?, '%') OR work LIKE concat('%', ?, '%') AND iddep != 0 LIMIT 10", query, query, query)

		if err != nil {
			panic(err.Error())
		}

		defer result.Close()

		for result.Next() {

			var person Person

			err := result.Scan(&person.IDPERSON, &person.Name, &person.Sdep)

			if err != nil {
				panic(err.Error())
			}

			persons = append(persons, person)
		}

		json.NewEncoder(w).Encode(persons)

	}
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

	result, err := db.Query("SELECT idperson, name, date_format(date,'%Y-%m-%d') AS date, IF(file IS NULL or file = '', 'photo.png', file) as file, pos, rank FROM persons LEFT JOIN pos USING(idpos) LEFT JOIN ranks USING(idrank) WHERE iddep like ? ORDER BY `name`", iddep)

	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	for result.Next() {

		var person Person

		err := result.Scan(&person.IDPERSON, &person.Name, &person.Date, &person.File, &person.Pos, &person.Rank)

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

	result, err := db.Query("SELECT idperson, name, date_format(date,'%Y-%m-%d') AS date, IF(file IS NULL or file = '', 'photo.png', file) as file, pos, rank, cellular, business, depart, sdep, place, work, internal, ipphone, arm, iddep, idpos, idrank FROM persons LEFT JOIN depart USING(iddep) LEFT JOIN places USING(idperson) LEFT JOIN pos USING(idpos) LEFT JOIN ranks USING(idrank) WHERE idperson = ? LIMIT 1", idperson)

	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	var person Person

	for result.Next() {

		err := result.Scan(&person.IDPERSON, &person.Name, &person.Date, &person.File, &person.Pos, &person.Rank, &person.Cellular, &person.Business, &person.Depart, &person.Sdep, &person.Place, &person.Work, &person.Internal, &person.Ipphone, &person.ARM, &person.Iddep, &person.Idpos, &person.Idrank)

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

	result, err := db.Query("SELECT idperson, name FROM persons LEFT JOIN depart USING(iddep) WHERE name LIKE concat('%', ?, '%') LIMIT 5", query)

	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	for result.Next() {

		var person Person

		err := result.Scan(&person.IDPERSON, &person.Name)

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

	var person Person

	err = db.QueryRow("SELECT file FROM persons WHERE idperson = ?", idperson).Scan(&person.File)
	if err != nil {
		panic(err.Error())
	}

	if person.File != "" {

		path := "./static/photo/" + person.File
		err = os.Remove(path)

		if err != nil {
			fmt.Println(err)
			//return
		}
	}

	stmt, err := db.Prepare("DELETE FROM persons WHERE idperson = ?")

	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(idperson)

	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "Person with ID = %s was deleted", strconv.Itoa(idperson))

	//fmt.Println("File" + person.File + "successfully deleted")

	//res.json({ success: true, message: 'Запрос выполнен' });
}

func Dismiss(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	db, err = sql.Open("mysql", os.Getenv("MYSQL_URL"))

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	var dp Person

	err := json.NewDecoder(r.Body).Decode(&dp)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	idperson := dp.IDPERSON
	fmt.Println(idperson)

	tx, err := db.Begin()
	if err != nil {
		panic(err.Error())
	}

	// insert a record into table1
	_, err = tx.Exec("UPDATE persons SET iddep='0', idpos='0', idrole='0' WHERE idperson=?", idperson)
	if err != nil {
		tx.Rollback()
		panic(err.Error())
	}

	// insert record into table2, referencing the first record from table1
	_, err = tx.Exec("UPDATE places SET idperson='0' WHERE idperson=?", idperson)
	if err != nil {
		tx.Rollback()
		panic(err.Error())
	}

	// commit the transaction
	tx.Commit()

	fmt.Println("Done.")

	//({ success: true, message: 'Запрос выполнен' });
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {

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

	fmt.Println("method:", r.Method)

	if err := r.ParseMultipartForm(32 << 20); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("file")

	name := r.FormValue("name")
	date := r.FormValue("date")
	cellular := r.FormValue("cellular")
	business := r.FormValue("business")
	iddep := r.FormValue("iddep")
	idpos := r.FormValue("idpos")
	idrank := r.FormValue("idrank")

	if err != nil {
		//fmt.Println("Error Retrieving the File")
		//fmt.Println(err)
		//return

		res, err := db.Exec("INSERT INTO persons (name, date, cellular, business, hash, iddep, idpos, idrank, file) VALUES (?, ?, ?, ?, '', ?, ?, ?, '')", name, date, cellular, business, iddep, idpos, idrank)

		if err != nil {
			panic(err)
		}

		lastId, err := res.LastInsertId()

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("The last inserted row id: %d\n", lastId)

	} else {

		//fmt.Fprintf(w, "%v", handler.Header)
		//fmt.Printf("Uploaded File: %+v\n", handler.Filename)
		//fmt.Printf("File Size: %+v\n", handler.Size)
		//fmt.Printf("MIME Header: %+v\n", handler.Header)

		f, err := os.OpenFile("./static/photo/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)

		if err != nil {
			fmt.Println(err)
			//return
		}

		defer f.Close()
		io.Copy(f, file)

		res, err := db.Exec("INSERT INTO persons (name, date, cellular, business, hash, iddep, idpos, idrank, file) VALUES (?, ?, ?, ?, '', ?, ?, ?, ?)", name, date, cellular, business, iddep, idpos, idrank, handler.Filename)

		if err != nil {
			panic(err)
		}

		lastId, err := res.LastInsertId()

		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()
		fmt.Printf("The last inserted row id: %d\n", lastId)

	}

}

func UpdatePerson(w http.ResponseWriter, r *http.Request) {

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

	fmt.Println("method:", r.Method)

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	idperson := r.FormValue("idperson")
	// name := r.FormValue("name")
	// date := r.FormValue("date")
	// cellular := r.FormValue("cellular")
	// business := r.FormValue("business")
	// iddep := r.FormValue("iddep")
	// idpos := r.FormValue("idpos")
	// idrank := r.FormValue("idrank")

	//удаляем старый файл

	var person Person

	err = db.QueryRow("SELECT file FROM persons WHERE idperson = ?", idperson).Scan(&person.File)
	if err != nil {
		panic(err.Error())
	}

	if person.File != "" {

		path := "./static/photo/" + person.File
		err = os.Remove(path)

		if err != nil {
			fmt.Println(err)
			//return
		}
	}

	file, handler, err := r.FormFile("file")

	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()

	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	//fmt.Printf("MIME Header: %+v\n", handler.Header)

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	tempFile, err := os.CreateTemp("./static/photo/", "upload-*.png")
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	fmt.Println("Created File: " + tempFile.Name())

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)

	// Copy the uploaded file to the created file on the filesystem
	// if _, err := io.Copy(tempFile, file); err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	//fmt.Println(tempFile.Name())
	defer os.Remove(tempFile.Name()) // clean up

	// return that we have successfully uploaded our file!
	fmt.Fprintf(w, "Successfully Uploaded File\n")

	// if err != nil {
	// 	//fmt.Println("Error Retrieving the File")
	// 	//fmt.Println(err)
	// 	//return

	// 	_, err = db.Exec("UPDATE persons SET name=?, date=?, cellular=?, business=?, iddep=?, idpos=?, idrank=? WHERE idperson=?", name, date, cellular, business, iddep, idpos, idrank, idperson)

	// 	if err != nil {
	// 		panic(err.Error())
	// 	}

	// 	fmt.Fprintf(w, "Person with ID = %s was updated", idperson)

	// } else {

	// 	//fmt.Fprintf(w, "%v", handler.Header)
	// 	//fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	// 	//fmt.Printf("File Size: %+v\n", handler.Size)
	// 	//fmt.Printf("MIME Header: %+v\n", handler.Header)

	// 	//Загружаем новый файл

	// 	f, err := os.OpenFile("./static/photo/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)

	// 	if err != nil {
	// 		fmt.Println(err)
	// 		//return
	// 	}

	// 	defer f.Close()
	// 	io.Copy(f, file)

	// 	_, err = db.Exec("UPDATE persons SET name=?, date=?, cellular=?, business=?, iddep=?, idpos=?, idrank=?, file=? WHERE idperson=?", name, date, cellular, business, iddep, idpos, idrank, handler.Filename, idperson)

	// 	if err != nil {
	// 		panic(err.Error())
	// 	}

	// 	defer file.Close()
	// 	fmt.Fprintf(w, "Person with ID = %s was updated", idperson)

	// }

}

// var buff bytes.Buffer
// fileSize, err := buff.ReadFrom(file)
// if err != nil {
// 	fmt.Println(err)
// }
// fmt.Println(fileSize)

// fmt.Println(r.FormValue("name"))
// fmt.Println(r.FormValue("date"))
// fmt.Println(r.FormValue("cellular"))
// fmt.Println(r.FormValue("business"))
// fmt.Println(r.FormValue("iddep"))
// fmt.Println(r.FormValue("idpos"))
// fmt.Println(r.FormValue("idrank"))

// if person.File == "" {
// 	fmt.Println("file is empty")
// }
