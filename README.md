
### Getting Started
**Download ZIP from Github**
```go
go run main.go
```

### Build app image
```bash
docker build -t golang-app .
```

```bash
docker build --no-cache -t alexchub/golang-app:latest .
docker push alexchub/golang-app:latest
docker image rmi alexchub/golang-app:latest
```

### Environment variables
<!-- MYSQL_HOST
MYSQL_USER
MYSQL_PASSWORD -->
MYSQL_URL

**MariaDB image into docker hub**
```bash
docker pull mariadb:10.4
docker tag mariadb:10.4 alexchub/mariadb:10.4
docker push alexchub/mariadb:10.4
```

**Run MariaDB docker container**
```bash
docker run -d -p 3306:3306 --name mariadb -e MYSQL_ROOT_PASSWORD=ZPwg4wHh -e MYSQL_DATABASE=phones -e MYSQL_USER=phones -e MYSQL_PASSWORD=ZPwg4wHh -v /mysql:/var/lib/mysql mariadb:10.4
```

**Check ip address**
```bash
docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' mariadb
```

**Connect to MariaDB server**
```bash
mysql -h 172.17.0.2 -u phones -p
```

**Create database structure**
```bash
mysql -h 172.17.0.2 -u phones -p phones < /home/user/Downloads/Github/phones_api_golang/struct.sql
```



### Licensing
App is [MIT licensed](./LICENSE).



**Запуск go mod init github.com/jonpchin/gochess(создание модуля) и go get(определение и получение зависимостей) - это действительно все, что нужно (в результате получаются два файла go.mod и go.sum; добавление их в проект - все, что необходимо). -  Британцы, 21 фев, в 4:43**

