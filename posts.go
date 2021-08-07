package main

import "time"

type Post struct {
	Time   time.Time
	Text   string
	Author string
	Id     int64
}

var Posts = make(map[int64]*Post)
var lastId int64 = 0

func addPost(text, author string) {
	Posts[lastId] = &Post{
		Time:   time.Now(),
		Text:   text,
		Author: author,
		Id:     lastId,
	}
	lastId++
}

func getPosts() map[int64]*Post {
	return Posts
}
