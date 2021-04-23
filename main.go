package main

import (
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	"alexchubgit/api/routes/dep"
	"alexchubgit/api/routes/places"
	"alexchubgit/api/routes/ranks"

	"alexchubgit/api/routes/auth"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/ranks", ranks.GetRanks).Methods("GET")
	router.HandleFunc("/deps", dep.GetDeps).Methods("GET")
	router.HandleFunc("/places", places.GetPlaces).Methods("GET")

	router.HandleFunc("/login", auth.Login).Methods("POST")

	fmt.Println("Server running on port 8000")
	http.ListenAndServe(":8000", router)

}

// type Rank struct {
// 	IDRANK string `json:"idrank"`
// 	Rank   string `json:"rank"`
// }

// type Dep struct {
// 	IDDEP string `json:"iddep"`
// 	Sdep  string `json:"sdep"`
// }

// var db *sql.DB
// var err error

// db, err = sql.Open("mysql", "root:idEt38@tcp(127.0.0.1:3306)/phones")
// if err != nil {
// 	panic(err.Error())
// }
// defer db.Close()

//router.HandleFunc("/posts", createPost).Methods("POST")
//router.HandleFunc("/posts/{id}", getPost).Methods("GET")
//router.HandleFunc("/posts/{id}", updatePost).Methods("PUT")
//router.HandleFunc("/posts/{id}", deletePost).Methods("DELETE")

//fmt.Println("Hello, World!")

// func GreetingFor(name string) string {
// 	return fmt.Sprintf("Hello, %s!", name)
// }

//fmt.Println(pos.GreetingFor("Alex"))

// func getRanks(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var ranks []Rank
// 	result, err := db.Query("SELECT idrank, rank from ranks")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer result.Close()
// 	for result.Next() {
// 		var rank Rank
// 		err := result.Scan(&rank.IDRANK, &rank.Rank)
// 		if err != nil {
// 			panic(err.Error())
// 		}
// 		ranks = append(ranks, rank)
// 	}
// 	json.NewEncoder(w).Encode(ranks)
// }

// func getDeps(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var deps []Dep
// 	result, err := db.Query("SELECT iddep, sdep from depart")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer result.Close()
// 	for result.Next() {
// 		var dep Dep
// 		err := result.Scan(&dep.IDDEP, &dep.Sdep)
// 		if err != nil {
// 			panic(err.Error())
// 		}
// 		deps = append(deps, dep)
// 	}
// 	json.NewEncoder(w).Encode(deps)
// }
