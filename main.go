package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	"alexchubgit/api/routes/addr"
	"alexchubgit/api/routes/dep"
	"alexchubgit/api/routes/persons"
	"alexchubgit/api/routes/places"
	"alexchubgit/api/routes/pos"
	"alexchubgit/api/routes/ranks"

	"alexchubgit/api/routes/auth"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

// init is invoked before main()
// loads values from .env into the system

func init() {

	// os.Setenv("MYSQL_URL", "phones:ZPwg4wHh@tcp(localhost:3306)/phones")

	// if err := godotenv.Load(".env"); err != nil {
	// 	log.Print("File .env not found")
	// }

	//логирование
	file, err := os.OpenFile("logs", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
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
	router.HandleFunc("/add_addr", auth.CheckSecurityRoute("", addr.CreateAddr)).Methods("POST")
	router.HandleFunc("/upd_addr", auth.CheckSecurityRoute("", addr.UpdateAddr)).Methods("PUT")
	router.HandleFunc("/del_addr", auth.CheckSecurityRoute("", addr.DeleteAddr)).Methods("DELETE")

	router.HandleFunc("/pos", pos.GetPoses).Methods("GET")
	router.HandleFunc("/one_pos", pos.GetOnePos).Methods("GET")
	router.HandleFunc("/list_pos", pos.GetListPos).Methods("GET")
	router.HandleFunc("/add_pos", auth.CheckSecurityRoute("", pos.CreatePos)).Methods("POST")
	router.HandleFunc("/upd_pos", auth.CheckSecurityRoute("", pos.UpdatePos)).Methods("PUT")
	router.HandleFunc("/del_pos", auth.CheckSecurityRoute("", pos.DeletePos)).Methods("DELETE")

	router.HandleFunc("/ranks", ranks.GetRanks).Methods("GET")
	router.HandleFunc("/one_rank", ranks.GetOneRank).Methods("GET")
	router.HandleFunc("/list_rank", ranks.GetListRank).Methods("GET")
	router.HandleFunc("/add_rank", auth.CheckSecurityRoute("", ranks.CreateRank)).Methods("POST")
	router.HandleFunc("/upd_rank", auth.CheckSecurityRoute("", ranks.UpdateRank)).Methods("PUT")
	router.HandleFunc("/del_rank", auth.CheckSecurityRoute("", ranks.DeleteRank)).Methods("DELETE")

	router.HandleFunc("/deps", dep.GetDeps).Methods("GET")
	router.HandleFunc("/one_dep", dep.GetOneDep).Methods("GET")
	router.HandleFunc("/list_dep", dep.GetListDep).Methods("GET")
	router.HandleFunc("/add_dep", auth.CheckSecurityRoute("", dep.CreateDep)).Methods("POST")
	router.HandleFunc("/upd_dep", auth.CheckSecurityRoute("", dep.UpdateDep)).Methods("PUT")
	router.HandleFunc("/del_dep", auth.CheckSecurityRoute("", dep.DeleteDep)).Methods("DELETE")

	router.HandleFunc("/places", places.GetPlaces).Methods("GET")
	router.HandleFunc("/one_place", places.GetOnePlace).Methods("GET")
	router.HandleFunc("/add_place", auth.CheckSecurityRoute("", places.CreatePlace)).Methods("POST")
	router.HandleFunc("/upd_place", auth.CheckSecurityRoute("", places.UpdatePlace)).Methods("PUT")
	router.HandleFunc("/del_place", auth.CheckSecurityRoute("", places.DeletePlace)).Methods("DELETE")
	router.HandleFunc("/del_person_place", auth.CheckSecurityRoute("", places.DeletePersonFromPlace)).Methods("PUT")

	router.HandleFunc("/persons", persons.GetPersons).Methods("GET")
	router.HandleFunc("/one_person", persons.GetOnePerson).Methods("GET")
	router.HandleFunc("/list_persons", persons.GetListPersons).Methods("GET")
	router.HandleFunc("/dismissed", persons.GetDismissed).Methods("GET")
	router.HandleFunc("/dates", persons.GetDatesWeek).Methods("GET")
	router.HandleFunc("/dates_today", persons.GetDatesToday).Methods("GET")
	router.HandleFunc("/search", persons.Search).Methods("GET")
	router.HandleFunc("/add_person", auth.CheckSecurityRoute("", persons.CreatePerson)).Methods("POST")
	router.HandleFunc("/upd_person", auth.CheckSecurityRoute("", persons.UpdatePerson)).Methods("PUT")
	router.HandleFunc("/del_person", auth.CheckSecurityRoute("", persons.DeletePerson)).Methods("DELETE")
	router.HandleFunc("/dismiss", auth.CheckSecurityRoute("", persons.Dismiss)).Methods("PUT")

	router.HandleFunc("/login", auth.Login).Methods("POST")
	router.HandleFunc("/checkauth", auth.CheckSecurityPage).Methods("GET")
	router.HandleFunc("/refresh", auth.Refresh).Methods("GET")

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))

	fmt.Println("Server running on port 8000")
	http.ListenAndServe(":8000", router)

}

// Here we are implementing the NotImplemented handler. Whenever an API endpoint is hit
// we will simply return the message "Not Implemented"
// var NotImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Not Implemented"))
// })

// router.HandleFunc("/status", NotImplemented).Methods("GET")
