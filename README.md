# API phones
API phones

### Getting Started

Download ZIP from Github  
`go run main.go`

### Licensing

API phones is [MIT licensed](./LICENSE).



Запуск go mod init github.com/jonpchin/gochess(создание модуля) и go get(определение и получение зависимостей) - это действительно все, что нужно (в результате получаются два файла go.mod и go.sum; добавление их в проект - все, что необходимо). -  Британцы, 21 фев, в 4:43



# Build image
docker build -t golang-app .

```bash
docker build --no-cache -t alexchub/golang-app:latest .
docker push alexchub/golang-app:latest
docker image rmi alexchub/golang-app:latest
```
