package main

import (
	//  "bytes"
	"database/sql"
	"fmt"
	"github.com/go-faker/faker/v4"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"strconv"
	"time"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	} else {
		fmt.Println("successful...")
	}
}

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@127.0.0.1:5432/book?sslmode=disable")
	checkErr(err)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)
	err = db.Ping()
	checkErr(err)
	db.SetMaxIdleConns(100)
	db.SetMaxOpenConns(500)
	semaphore := make(chan int, 200)
	done := make(chan int)
	beginTime := time.Now()
	for i := 0; i < 1000000; i++ {
		go insertData(db, semaphore, done)
	}
	for respCount := 0; respCount < 1000000; respCount++ {
		<-done
	}
	print("time passed: " + time.Now().Sub(beginTime).String())
}

func insertData(db *sql.DB, semaphore chan int, done chan int) {
	semaphore <- 1
	year := faker.YearString()
	intYear, _ := strconv.ParseInt(year, 10, 64)
	_, _ = db.Exec("insert into books (title, author, year) values ($1 , $2 , $3);",
		faker.Name(), faker.Sentence(), intYear,
	)
	<-semaphore
	done <- 1
}
