package main

import (
	"database/sql"
	_ "errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Takao-Yamasaki/myapi_v2/api"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbUser     = "docker"
	dbPassword = "docker"
	dbDatabase = "sampledb"
	dbConn     = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
)

func main() {
	_, err1 := strconv.Atoi("a")

	err2 := strconv.ErrSyntax
	err3 := strconv.ErrRange

	fmt.Println(err1 == err2)
	fmt.Println(err1 == err3)

	// fmt.Println(errors.Is(err1, err2))
	// fmt.Println(errors.Is(err1, err3))
	// aは数値に直せないので、エラーが発生する
	// _, err0 := strconv.Atoi("a")
	// fmt.Printf("err0: [%T] %v\n", err0, err0)

	// err1 := errors.Unwrap(err0)
	// fmt.Printf("err1: [%T] %v\n", err1, err1)

	// err2 := errors.Unwrap(err1)
	// fmt.Printf("err2: [%T] %v\n", err2, err2)

	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		log.Println("fail to connect DB")
		return
	}

	r := api.NewRouter(db)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
