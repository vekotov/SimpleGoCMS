package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"time"
)

func testDB() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS posts(id bigint, time timestamp, text text, author text)")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("INSERT INTO posts values({}, {}, {}, {})", 1, time.Now(), "Hello world!", "vekotov")
	if err != nil {
		return
	}
}
