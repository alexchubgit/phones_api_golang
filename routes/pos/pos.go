package pos

import (
	"database/sql"
	"encoding/json"
	"fmt"
	//"io/ioutil"
	"net/http"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	//"github.com/gorilla/mux"
)

type Pos struct {
	IDPOS string `json:"idpos"`
	Pos   string `json:"pos"`
}

var db *sql.DB
var err error

func GetPoses(w http.ResponseWriter, r *http.Request) {

	db, err = sql.Open("mysql", os.Getenv("MYSQL_URL"))

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	w.Header().Set("Content-Type", "application/json")

	var poses []Pos

	result, err := db.Query("SELECT idpos, pos from pos")

	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	for result.Next() {

		var pos Pos

		err := result.Scan(&pos.IDPOS, &pos.Pos)

		if err != nil {
			panic(err.Error())
		}

		poses = append(poses, pos)

	}
	json.NewEncoder(w).Encode(poses)
}

func GetOnePos(w http.ResponseWriter, r *http.Request) {

	db, err = sql.Open("mysql", os.Getenv("MYSQL_URL"))

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	w.Header().Set("Content-Type", "application/json")

	idpos, err := strconv.Atoi(r.URL.Query().Get("idpos"))

	if err != nil || idpos < 1 {
		http.NotFound(w, r)
		return
	}

	result, err := db.Query("SELECT idpos, pos FROM pos WHERE idpos = ?", idpos)

	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	var pos Pos

	for result.Next() {

		err := result.Scan(&pos.IDPOS, &pos.Pos)

		if err != nil {
			panic(err.Error())
		}
	}

	json.NewEncoder(w).Encode(pos)
}

func DeletePos(w http.ResponseWriter, r *http.Request) {

	db, err = sql.Open("mysql", os.Getenv("MYSQL_URL"))

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	//params := mux.Vars(r)

	idpos, err := strconv.Atoi(r.URL.Query().Get("idpos"))

	if err != nil || idpos < 1 {
		http.NotFound(w, r)
		return
	}

	stmt, err := db.Prepare("DELETE FROM pos WHERE idpos = ?")

	if err != nil {
		panic(err.Error())
	}

	//_, err = stmt.Exec(params["idpos"])

	_, err = stmt.Exec(idpos)

	if err != nil {
		panic(err.Error())
	}

	//fmt.Fprintf(w, "Pos with ID = %s was deleted", params["idpos"])

	fmt.Fprintf(w, "Pos with ID = %s was deleted", strconv.Itoa(idpos))
}

func CreatePos(w http.ResponseWriter, r *http.Request) {

	//w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Content-Type", "application/json")
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Access-Control-Allow-Methods", "POST")
	// w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	db, err = sql.Open("mysql", os.Getenv("MYSQL_URL"))

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	if r.Method != "POST" {
		fmt.Println("Not Post")
		return
	}

	pos := r.FormValue("pos")
	fmt.Println(pos)

	if pos == "" {
		fmt.Println("Feild is empty")

	}

	result, err := db.Exec("INSERT INTO pos(pos) VALUES(455555)")
	if err != nil {
		panic(err)

	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err)

	}

	fmt.Fprintf(w, "Book %s created successfully (%d row affected)\n", pos, rowsAffected)

	//x-www-form-urlencoded

	// r.ParseForm()

	// fmt.Printf("%+v\n", r.Form)

	// for key, value := range r.Form {
	// 	fmt.Printf("%s = %s\n", key, value)
	// }

	// params := r.PostFormValue("pos")
	// fmt.Println(params)

	//получение параметра form-data

	// body, err := ioutil.ReadAll(r.Body)

	// if err != nil {
	// 	panic(err.Error())
	// }

	// keyVal := make(map[string]string)
	// json.Unmarshal(body, &keyVal)
	// pos := keyVal["pos"]
	// fmt.Println(pos)

	// fmt.Printf("%s\n", string(body))

	// if r.Method != "POST" {
	// 	http.Error(w, http.StatusText(405), 405)
	// 	return
	// }

	// pos := r.FormValue("pos")

	// fmt.Println(pos)

	// if pos == "" {
	// 	http.Error(w, http.StatusText(400), 400)
	// 	return
	// }

	// result, err := db.Exec("INSERT INTO pos VALUES($1)", pos)
	// if err != nil {
	// 	http.Error(w, http.StatusText(500), 500)
	// 	return
	// }

	// rowsAffected, err := result.RowsAffected()
	// if err != nil {
	// 	http.Error(w, http.StatusText(500), 500)
	// 	return
	// }

	// fmt.Fprintf(w, "Book %s created successfully (%d row affected)\n", pos, rowsAffected)

}

func UpdatePos(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	// db, err = sql.Open("mysql", os.Getenv("MYSQL_URL"))

	// if err != nil {
	// 	panic(err.Error())
	// }

	// defer db.Close()

	//params := mux.Vars(r)
	//fmt.Println(params)

	// stmt, err := db.Prepare("UPDATE pos SET pos = ? WHERE idpos = ?")

	// if err != nil {
	// 	panic(err.Error())
	// }

	// body, err := ioutil.ReadAll(r.Body)

	// if err != nil {
	// 	panic(err.Error())
	// }

	// keyVal := make(map[string]string)

	// json.Unmarshal(body, &keyVal)

	// newPos := keyVal["pos"]

	// _, err = stmt.Exec(newPos, params["idpos"])

	// if err != nil {
	// 	panic(err.Error())
	// }

	// fmt.Fprintf(w, "Pos with ID = %s was updated", params["idpos"])
}

//fmt.Println(idpos)
//idpos := r.URL.Query().Get("idpos")
// params := mux.Vars(r)
// fmt.Println(params)
//result, err := db.Query("SELECT idpos, pos FROM pos WHERE idpos = ?", params["idpos"])
