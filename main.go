package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"alexchubgit/api/routes/addr"
	"alexchubgit/api/routes/certs"
	"alexchubgit/api/routes/dep"
	"alexchubgit/api/routes/docs"
	"alexchubgit/api/routes/persons"
	"alexchubgit/api/routes/places"
	"alexchubgit/api/routes/pos"
	"alexchubgit/api/routes/ranks"

	"alexchubgit/api/routes/auth"
)

// init is invoked before main()
// loads values from .env into the system

func init() {

	if err := godotenv.Load(".env"); err != nil {
		log.Print("File .env not found")
	}

	// var db *sql.DB
	// var err error

	// db, err = sql.Open("mysql", os.Getenv("MYSQL_URL"))

	// if err != nil {
	// 	panic(err.Error())
	// }

	// defer db.Close()
}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/addr", addr.GetAddr).Methods("GET")
	router.HandleFunc("/ranks", ranks.GetRanks).Methods("GET")
	router.HandleFunc("/pos", pos.GetPoses).Methods("GET")
	router.HandleFunc("/deps", dep.GetDeps).Methods("GET")
	router.HandleFunc("/places", places.GetPlaces).Methods("GET")
	router.HandleFunc("/certs", certs.GetCert).Methods("GET")
	router.HandleFunc("/docs", docs.GetDocs).Methods("GET")

	router.HandleFunc("/one_dep/{iddep}", dep.GetOneDep).Methods("GET")
	router.HandleFunc("/persons/{iddep}", persons.GetPersons).Methods("GET")
	router.HandleFunc("/one_person/{idperson}", persons.GetOnePerson).Methods("GET")

	router.HandleFunc("/login", auth.Login).Methods("POST")

	router.HandleFunc("/pos", pos.CreatePos).Methods("POST")

	router.HandleFunc("/pos/{idpos}", pos.UpdatePos).Methods("PUT")

	router.HandleFunc("/pos/{idpos}", pos.DeletePos).Methods("DELETE")

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	fmt.Println("Server running on port 8000")
	http.ListenAndServe(":8000", router)

}
