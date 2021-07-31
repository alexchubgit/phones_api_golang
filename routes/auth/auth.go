package auth

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	//"github.com/go-redis/redis/v7"
	//"github.com/twinj/uuid"
	//"log"
)

// Secret key to uniquely sign the token
var key []byte

//структура для учётной записи пользователя
type Account struct {
	Name string  `json:"name"`
	Role *string `json:"idrole"`
	Hash string  `json:"hash"`
	jwt.StandardClaims
}

type Auth struct {
	Login    string
	Password string
}

var db *sql.DB
var err error

//Генерация хэша на основе пароля
func HashPassword(password string) (string, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}

	fmt.Println("hashedPassword: " + string(hashedPassword))

	return string(hashedPassword), nil
}

//Проверка bcrypt хэша с помощью пароля
func CheckPassword(hashedPassword string, password string) bool {

	if (hashedPassword == "") || (password == "") {
		fmt.Println("Hash || password empty")

		fmt.Println("hashedPassword: " + hashedPassword)
		fmt.Println("password: " + password)

		return false
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	if err != nil {
		fmt.Println(err)
		return false
	}

	return true

}

//Проверка JWT хэша в маршрутах
func CheckSecurityRoute(password string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		tokenString := r.Header.Get("Authorization")
		fmt.Println(tokenString)

		key = []byte(os.Getenv("JWT_KEY"))

		// Initialize a new instance of `Claims`
		claims := &Account{}

		// Parse the JWT string and store the result in `claims`.
		// Note that we are passing the key in this method as well. This method will return an error
		// if the token is invalid (if it has expired according to the expiry time we set on sign in),
		// or if the signature does not match
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return key, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				fmt.Println("Токен не действителен")
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			fmt.Println("Ошибка")
			//проверить ошибку на клиенте и выкинуть в авторизацию
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if !token.Valid {
			fmt.Println("Токен не действителен")
			w.WriteHeader(http.StatusUnauthorized)
			return
		} else {
			fmt.Println("Токен действителен")
			//fmt.Println(claims.Name)
			//вывести должность и звание

			//fmt.Println(claims.StandardClaims.ExpiresAt)

			// Finally, return the welcome message to the user, along with their
			// username given in the token

			w.Write([]byte(fmt.Sprintf("Welcome %s!", claims.Name)))

			next(w, r)
		}

	}
}

func CheckSecurityPage(w http.ResponseWriter, r *http.Request) {

	tokenString := r.Header.Get("Authorization")
	fmt.Println(tokenString)

	key = []byte(os.Getenv("JWT_KEY"))

	// Initialize a new instance of `Claims`
	claims := &Account{}

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			fmt.Println("Токен не действителен")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		fmt.Println("Ошибка")
		//проверить ошибку на клиенте и выкинуть в авторизацию
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !token.Valid {
		fmt.Println("Токен не действителен")
		w.WriteHeader(http.StatusUnauthorized)
		return
	} else {
		fmt.Println("Токен действителен")
		//fmt.Println(claims.Name)
		//вывести должность и звание

		//fmt.Println(claims.StandardClaims.ExpiresAt)

		// Finally, return the welcome message to the user, along with their
		// username given in the token

		w.Write([]byte(fmt.Sprintf("Welcome %s!", claims.Name)))

		//next(w, r)
	}
}

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

	//Check Post query
	if r.Method != "POST" {
		fmt.Println("Not Post")
		return
	}

	//Get body params
	decoder := json.NewDecoder(r.Body)

	var auth Auth
	err := decoder.Decode(&auth)

	if err != nil {
		//panic(err)
		w.WriteHeader(http.StatusBadRequest)
	}

	//логин идет в проверку пользователя
	login := auth.Login
	//пароль идет в проверку пароля на основе полученного у пользователя хэша
	password := auth.Password

	//Get hash from database

	//Ищем человека по номеру телефона личному или служебному и вытаскиваем хэш brcypt

	result, err := db.Query("SELECT name, hash FROM persons WHERE cellular = ?  OR business = ? LIMIT 1", login, login)

	if err != nil {
		//panic(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}

	defer result.Close()

	var account Account

	for result.Next() {

		err := result.Scan(&account.Name, &account.Hash)

		if err != nil {
			//panic(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
		}
	}

	//fmt.Println(account.Name)
	//fmt.Println(account.Hash)

	//var hashedPassword string

	var hashedPassword string = account.Hash

	//Генерация хэша на основе пароля введенного пользователем
	//HashPassword(password)

	//Проверяем полученный из базы хэш с помощью введенного пароля
	if CheckPassword(hashedPassword, password) {
		//if CheckPassword(hashedPassword, password) {

		//fmt.Println("OK")

		//Создать токен JWT
		var claims = Account{
			Name: account.Name,
			StandardClaims: jwt.StandardClaims{
				// Enter expiration in milisecond
				ExpiresAt: time.Now().Add(5 * time.Minute).Unix(),
			},
		}

		key = []byte(os.Getenv("JWT_KEY"))

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, _ := token.SignedString(key)
		//fmt.Println(tokenString)

		//вывод результата
		values := map[string]string{"token": tokenString, "success": "true", "message": "Запрос выполнен. Токен получен"}
		json.NewEncoder(w).Encode(values)

	} else {

		//fmt.Println("Does not OK")

		//w.WriteHeader(http.StatusInternalServerError) //500 code
		//w.WriteHeader(http.StatusForbidden) //403 code

		//w.Write([]byte("500 - Something bad happened!"))
		// res.json({ success: false, message: 'Логин или пароль указаны неверно' });

		//вывод результата
		values := map[string]string{"success": "false", "message": "Логин или пароль указаны неверно"}
		json.NewEncoder(w).Encode(values)

	}

}

func Refresh(w http.ResponseWriter, r *http.Request) {

	tokenString := r.Header.Get("Authorization")
	fmt.Println(tokenString)

	key = []byte(os.Getenv("JWT_KEY"))

	// Initialize a new instance of `Claims`
	claims := &Account{}

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			fmt.Println("Токен не действителен")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		fmt.Println("Ошибка")
		//проверить ошибку на клиенте и выкинуть в авторизацию
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !token.Valid {
		fmt.Println("Токен не действителен")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// (BEGIN) The code uptil this point is the same as the first part of the `Welcome` route
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tknStr := c.Value
	//claims := &Account{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	// (END) The code up-till this point is the same as the first part of the `Welcome` route

	// We ensure that a new token is not issued until enough time has elapsed
	// In this case, a new token will only be issued if the old token is within
	// 30 seconds of expiry. Otherwise, return a bad request status
	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Now, create a new token for the current use, with a renewed expiration time
	expirationTime := time.Now().Add(5 * time.Minute)
	claims.ExpiresAt = expirationTime.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(key)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Set the new token as the users `token` cookie
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	//"SELECT idperson, name, role FROM persons LEFT JOIN role USING(idrole) WHERE (`cellular` = ? AND `passwd` = ?) OR (`business` = ? AND `passwd` = ?) LIMIT 1"

	//Creating Refresh Token
	// rtClaims := jwt.MapClaims{}
	// rtClaims["refresh_uuid"] = td.RefreshUuid
	// rtClaims["user_id"] = "userid"
	// rtClaims["exp"] = td.RtExpires

	//rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	//refreshToken, _ := rt.SignedString([]byte(os.Getenv("refreshKey")))
	//fmt.Println(refreshToken)

	//payload = {
	//                 idperson: results[0].idperson,
	//                 name: results[0].name,
	//                 phone: results[0].cellular,
	//                 role: results[0].role
	//             };

	//             //console.log(payload)
	//             console.log(SECRET_KEY)
	//             //здесь создается JWT
	//             const token = jwt.sign(payload, SECRET_KEY, {
	//                 expiresIn: 60 * 60 * 24 // истекает через 24 часа
	//             });

}

// func Refresh(w http.ResponseWriter, r *http.Request) {
// 	// (BEGIN) The code uptil this point is the same as the first part of the `Welcome` route
// 	tokenString := r.Header.Get("Authorization")
// 	fmt.Println(tokenString)

// 	//key = []byte(os.Getenv("JWT_KEY"))

// 	// Initialize a new instance of `Claims`
// 	claims := &Token{}

// 	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
// 		return key, nil
// 	})
// 	if err != nil {
// 		if err == jwt.ErrSignatureInvalid {
// 			w.WriteHeader(http.StatusUnauthorized)
// 			return
// 		}
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// 	if !token.Valid {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		return
// 	}
// 	// (END) The code up-till this point is the same as the first part of the `Welcome` route

// 	// We ensure that a new token is not issued until enough time has elapsed
// 	// In this case, a new token will only be issued if the old token is within
// 	// 30 seconds of expiry. Otherwise, return a bad request status
// 	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	// Now, create a new token for the current use, with a renewed expiration time
// 	expirationTime := time.Now().Add(5 * time.Minute)
// 	claims.ExpiresAt = expirationTime.Unix()
// 	token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	tokenString, err = token.SignedString(key)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	// Set the new token as the users `token` cookie
// 	http.SetCookie(w, &http.Cookie{
// 		Name:    "token",
// 		Value:   tokenString,
// 		Expires: expirationTime,
// 	})
// }

// type Token struct {
// 	Name string `json:"name"`
// 	jwt.StandardClaims
// }

//fmt.Println("middleware")
//fmt.Println(password)

// if req.Header["Authorization"] == nil {
// 	fmt.Println("No Token Found")
// }

// authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
// if len(tokenString) != 2 {
// 	fmt.Println("Malformed token")
// 	w.WriteHeader(http.StatusUnauthorized)
// 	w.Write([]byte("Malformed Token"))
// }

// for result.Next() {

// 	err := result.Scan(&hashedPassword)

// 	if err != nil {
// 		panic(err.Error())
// 	}
// }

// http.SetCookie(w, &http.Cookie{
// 	Name:     "token",
// 	Value:    tokenString,
// 	HttpOnly: true,
// 	//Expires:  time.Now().Add(60 * time.Minute).Unix(),
// })
