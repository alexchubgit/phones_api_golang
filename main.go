package main

import (
	"fmt"
	"net/http"

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

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/addr", addr.GetAddr).Methods("GET")
	router.HandleFunc("/ranks", ranks.GetRanks).Methods("GET")
	router.HandleFunc("/pos", pos.GetPoses).Methods("GET")
	router.HandleFunc("/deps", dep.GetDeps).Methods("GET")
	router.HandleFunc("/places", places.GetPlaces).Methods("GET")

	router.HandleFunc("/one_dep/{iddep}", dep.GetOneDep).Methods("GET")
	router.HandleFunc("/persons/{iddep}", persons.GetPersons).Methods("GET")
	router.HandleFunc("/one_person/{idperson}", persons.GetOnePerson).Methods("GET")

	router.HandleFunc("/login", auth.Login).Methods("POST")

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	fmt.Println("Server running on port 8000")
	http.ListenAndServe(":8000", router)

}

//router.HandleFunc("/posts", createPost).Methods("POST")
//router.HandleFunc("/posts/{id}", getPost).Methods("GET")
//router.HandleFunc("/posts/{id}", updatePost).Methods("PUT")
//router.HandleFunc("/posts/{id}", deletePost).Methods("DELETE")
