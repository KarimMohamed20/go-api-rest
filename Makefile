build:
	go build -o bin/main src/main.go

get:
	go get github.com/gin-gonic/gin
	go get github.com/jinzhu/gorm/dialects/mysql
	go get github.com/jinzhu/gorm
run:
	go run src/main.go

