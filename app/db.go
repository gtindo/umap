package app

import (
	"database/sql"
	"log"
)

var db *sql.DB

func InitDb(dbStr string) {
	db, err := sql.Open("mysql", dbStr)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Connection established with MYSQL DB")
}
