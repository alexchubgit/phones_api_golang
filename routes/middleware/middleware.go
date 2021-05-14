package middleware

import (
	"fmt"
	"net/http"
)

func CheckSecurity(password string, next http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		fmt.Println("middleware")
		fmt.Println(password)
		// header := req.Header.Get("Super-Duper-Safe-Security")
		// if header != "password" {
		// 	fmt.Fprint(res, "Invalid password")
		// 	res.WriteHeader(http.StatusUnauthorized)
		// 	return
		// }
		next(res, req)
	}
}
