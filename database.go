package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"time"
)

var db *sql.DB

func initDB() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS posts(id bigint, time timestamp, text text, author text)")
	if err != nil {
		log.Fatal(err)
	}
}

func dbAddPost(post *Post) {
	_, err := db.Exec(
		"INSERT INTO posts value({},{},{},{})",
		post.Id,
		post.Time.Format("2006-01-02 15:04:05"),
		post.Text,
		post.Author,
	)
	if err != nil {
		log.Print("DB ERROR: ", err.Error())
	}
}

func dbCountPosts() (count int64) {
	row := db.QueryRow("SELECT COUNT(*) FROM posts")
	err := row.Scan(&count)
	if err != nil {
		log.Print("DB ERROR: ", err.Error())
	}
	return
}

func dbGetPost(id int64) *Post {
	row := db.QueryRow("SELECT * FROM posts WHERE id = {}", id)
	post := Post{}
	var timeText string
	err := row.Scan(&post.Id, &timeText, &post.Text, &post.Author)
	if err != nil {
		log.Print("DB ERROR: ", err.Error())
	}
	post.Time, err = time.Parse("2006-01-02 15:04:05", timeText)
	if err != nil {
		log.Print("DB ERROR: ", err.Error())
	}
	post.TimeText = post.Time.Format("2006-01-02 15:04:05")
	return &post
}
