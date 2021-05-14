package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

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
	"alexchubgit/api/routes/tokens"

	"alexchubgit/api/routes/auth"
	"alexchubgit/api/routes/middleware"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

// init is invoked before main()
// loads values from .env into the system

func init() {

	if err := godotenv.Load(".env"); err != nil {
		log.Print("File .env not found")
	}

	//логирование
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {

	InfoLogger.Println("Starting the application...")
	//InfoLogger.Println("Something noteworthy happened")
	//WarningLogger.Println("There is something you should know about")
	//ErrorLogger.Println("Something went wrong")

	router := mux.NewRouter()

	router.HandleFunc("/addr", addr.GetAddr).Methods("GET")
	router.HandleFunc("/one_addr", addr.GetOneAddr).Methods("GET")
	router.HandleFunc("/list_addr", addr.GetListAddr).Methods("GET")
	router.HandleFunc("/add_addr", middleware.CheckSecurity("Password", addr.CreateAddr)).Methods("POST")
	router.HandleFunc("/upd_addr", middleware.CheckSecurity("Password", addr.UpdateAddr)).Methods("PUT")
	router.HandleFunc("/del_addr", middleware.CheckSecurity("Password", addr.DeleteAddr)).Methods("DELETE")

	router.HandleFunc("/pos", pos.GetPoses).Methods("GET")
	router.HandleFunc("/one_pos", pos.GetOnePos).Methods("GET")
	router.HandleFunc("/list_pos", pos.GetListPos).Methods("GET")
	router.HandleFunc("/add_pos", pos.CreatePos).Methods("POST")
	router.HandleFunc("/upd_pos", pos.UpdatePos).Methods("PUT")
	router.HandleFunc("/del_pos", pos.DeletePos).Methods("DELETE")

	router.HandleFunc("/ranks", ranks.GetRanks).Methods("GET")
	router.HandleFunc("/one_rank", ranks.GetOneRank).Methods("GET")
	router.HandleFunc("/list_rank", ranks.GetListRank).Methods("GET")
	router.HandleFunc("/add_rank", ranks.CreateRank).Methods("POST")
	router.HandleFunc("/upd_rank", ranks.UpdateRank).Methods("PUT")
	router.HandleFunc("/del_rank", ranks.DeleteRank).Methods("DELETE")

	router.HandleFunc("/deps", dep.GetDeps).Methods("GET")
	router.HandleFunc("/one_dep", dep.GetOneDep).Methods("GET")
	router.HandleFunc("/list_dep", dep.GetListDep).Methods("GET")
	router.HandleFunc("/add_dep", dep.CreateDep).Methods("POST")
	router.HandleFunc("/upd_dep", dep.UpdateDep).Methods("PUT")
	router.HandleFunc("/del_dep", dep.DeleteDep).Methods("DELETE")

	router.HandleFunc("/places", places.GetPlaces).Methods("GET")
	router.HandleFunc("/one_place", places.GetOnePlace).Methods("GET")
	router.HandleFunc("/add_place", places.CreatePlace).Methods("POST")
	router.HandleFunc("/upd_place", places.UpdatePlace).Methods("PUT")
	router.HandleFunc("/del_person_place", places.DeletePersonFromPlace).Methods("PUT")
	router.HandleFunc("/del_place", places.DeletePlace).Methods("DELETE")

	router.HandleFunc("/persons", persons.GetPersons).Methods("GET")
	router.HandleFunc("/one_person", persons.GetOnePerson).Methods("GET")
	router.HandleFunc("/add_person", persons.CreatePerson).Methods("POST")

	router.HandleFunc("/tokens", tokens.GetTokens).Methods("GET")
	router.HandleFunc("/certs", certs.GetCert).Methods("GET")
	router.HandleFunc("/docs", docs.GetDocs).Methods("GET")

	router.HandleFunc("/login", auth.Login).Methods("POST")

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	fmt.Println("Server running on port 8000")
	http.ListenAndServe(":8000", router)

}

//"gopkg.in/guregu/null.v4/zero"
//Name     zero.String `json:"name"`
//Name     NullString  `json:"name"`

// type NullString struct {
// 	sql.NullString
// }

// MarshalJSON for NullString
// func (ns *NullString) MarshalJSON() ([]byte, error) {
// 	if !ns.Valid {
// 		return []byte("null"), nil
// 	}
// 	return json.Marshal(ns.String)
// }
