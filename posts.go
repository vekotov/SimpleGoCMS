package main

import "time"

type Post struct {
	Time     time.Time
	TimeText string
	Text     string
	Author   string
	Id       int64
}

var Posts = make(map[int64]*Post)
var lastId int64 = 0

func addPost(text, author string) (id int64) {
	Posts[lastId] = &Post{
		Time:     time.Now(),
		TimeText: time.Now().Format("2006-01-02 15:04"),
		Text:     text,
		Author:   author,
		Id:       lastId,
	}
	lastId++
	return lastId - 1
}

func getPosts() map[int64]*Post {
	return Posts
}

func dbLoadPosts() {
	count := dbCountPosts()
	for i := count; i > 0; i-- {
		var post *Post = dbGetPost(i)
		Posts[i] = post
	}
}
