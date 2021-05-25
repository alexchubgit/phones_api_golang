package auth

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	//"time"
	//"crypto/md5"
	//"encoding/hex"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

/*
Структура прав доступа JWT
*/
type Token struct {
	UserId uint
	jwt.StandardClaims
}

//структура для учётной записи пользователя
type Account struct {
	IDPERSON int     `json:"idperson"`
	Name     string  `json:"name"`
	Role     *string `json:"role"`
	Hash     *string `json:"hash"`
}

type Auth struct {
	Login    string
	Password string
}

var db *sql.DB
var err error

func HashPassword(password string) (string, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hashedPassword), nil
}

func CheckPassword(password string, hashedPassword string) error {

	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

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

	if r.Method != "POST" {
		fmt.Println("Not Post")
		return
	}

	decoder := json.NewDecoder(r.Body)

	var auth Auth
	err := decoder.Decode(&auth)

	if err != nil {
		panic(err)
	}

	login := auth.Login
	password := auth.Password

	hashFromDatabase, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Hash to store:", string(hashFromDatabase))

	// Comparing the password with the hash
	err = bcrypt.CompareHashAndPassword(hashFromDatabase, []byte(password))

	if err == nil {
		fmt.Println("OK")
	}

	//fmt.Println(err)

	//"SELECT idperson, name, role FROM persons LEFT JOIN role USING(idrole) WHERE (`cellular` = ? AND `passwd` = ?) OR (`business` = ? AND `passwd` = ?) LIMIT 1"

	result, err := db.Query("SELECT hash FROM persons WHERE `cellular` = ?  OR `business` = ? LIMIT 1", login, login)

	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	var account Account

	for result.Next() {

		err := result.Scan(&account.Hash)

		if err != nil {
			panic(err.Error())
		}

		fmt.Println(err)

	}

}

func GetUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	db, err = sql.Open("mysql", os.Getenv("MYSQL_URL"))

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
}

// const SECRET_KEY = process.env.SECRET_KEY

// const Login = (login, password) => {
//     return new Promise((resolve, reject) => {
//         pool.query('SELECT * FROM persons LEFT JOIN role USING(idrole) WHERE `cellular` = "' + login + '" AND `passwd` = "' + password + '" OR `business` = "' + login + '" AND `passwd` = "' + password + '" LIMIT 1', (err, results) => {
//             if (err) {
//                 return reject(err);
//             }
//             return resolve(results);
//         });
//     });
// }

// //Авторизация
// auth.post('/login', async (req, res) => {

//     const login = req.body.login;
//     const password = md5(req.body.password);
//     let payload = {}

//     //console.log(login)

//     if ((login == undefined) && (password == undefined)) {
//         return res.sendStatus(500);
//     }

//     try {
//         const results = await Login(login, password);

//         //console.log(results)

//         //здесь проверяем авторизацию и создаем токен
//         if (results.length == 0) {
//             res.json({ success: false, message: 'Логин или пароль указаны неверно' });
//         } else {
//             payload = {
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

//             res.json({ success: true, message: 'Запрос выполнен. Токен получен', token: token });
//             //res.status(200).json({ results });
//         }
//     } catch (e) {
//         console.log(e);
//         res.sendStatus(500);
//     }
// })

// auth.get('/getuser', async (req, res) => {

//     //Проверка токена из видео урока
//     const authHeader = req.get('Authorization');

//     if (!authHeader) {
//         res.status(401).json({ success: false, message: "Token not provided!" });
//     }

//     const token = authHeader.replace('token ', '')

//     //console.log('middleware ' + token);

//     try {
//         const decoded = jwt.verify(token, SECRET_KEY);

//         //ошибка заголовков
//         //res.status(200).json({ success: true, message: 'Good to authenticate token.' });

//         //console.log('decode ' + decoded.role + ' ' + decoded.name);
//         res.send(decoded);

//     } catch (e) {
//         if (e instanceof jwt.JsonWebTokenError) {
//             res.status(401).json({ success: false, message: "Token invalid!" });
//         }
//     }
// });

//fmt.Println(t.Login)
//fmt.Println(t.Password)

//login := r.FormValue("login")
//password := r.FormValue("password")
//fmt.Println(password)

//json.NewEncoder(w).Encode(account)

// Глобальный секретный ключ
//var mySigningKey = []byte("secret")

// Создаем новый токен
//token := jwt.New(jwt.SigningMethodHS256)

// Устанавливаем набор параметров для токена
//token.Claims["admin"] = true
//token.Claims["name"] = "Ado Kukic"
//token.Claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

// Подписываем токен нашим секретным ключем
//tokenString, _ := token.SignedString(mySigningKey)

// Отдаем токен клиенту
//w.Write([]byte(tokenString))
//json.NewEncoder(w).Encode(tokenString)
