package main

import (
	"fmt"
	"log"
	"net/http"
)

func addPostResponse(w http.ResponseWriter, r *http.Request) {
	log.Print("[INFO] [addPostResponse] ", r.Method, " request ", r.RequestURI, " from ", r.RemoteAddr)
	err := 0
	if r.Method == "POST" {
		text := r.PostFormValue("post_text")
		if len(text) < 10 {
			err = 1
		}
		if len(text) > 500 {
			err = 2
		}
		if err == 0 {
			id := addPost(r.PostFormValue("post_text"), r.RemoteAddr)
			dbAddPost(getPosts()[id])
		}
	}
	var redirectTo string
	if err != 0 {
		redirectTo = fmt.Sprintf("/?error=%d", err)
	} else {
		redirectTo = "/"
	}
	http.Redirect(w, r, redirectTo, 303)
}
