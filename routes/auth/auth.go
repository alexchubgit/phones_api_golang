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
