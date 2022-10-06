package main

import (
	"log"
	"net/http"
	"fmt"
	"database/sql"

	_ "github.com/lib/pq"
)

const (
	host     = "172.17.0.2"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "postgres"
)

func OpenConnection() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/", GETHandler)
	http.HandleFunc("/inset", POSTHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

