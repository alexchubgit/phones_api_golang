package auth

import (
	"database/sql"
	"net/http"
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

}
