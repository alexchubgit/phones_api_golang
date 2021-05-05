package auth

import (
	"database/sql"
	"net/http"
	"os"
	//"encoding/json"
	// "fmt"
	//"log"
	//jwt "github.com/dgrijalva/jwt-go"
)

type Auth struct {
	Login  string `json:"login"`
	Passwd string `json:"passwd"`
}

var db *sql.DB
var err error

func Login(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	db, err = sql.Open("mysql", os.Getenv("MYSQL_URL"))

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
}
