package main

import (
	"log"
	"net/http"
	"strconv"
)

type IndexPageData struct {
	Posts map[int64]*Post
	Error int
}

func getIndexPageData(error int) *IndexPageData {
	return &IndexPageData{
		Posts: getPosts(),
		Error: error,
	}
}

func indexPageResponse(w http.ResponseWriter, r *http.Request) {
	log.Print("[INFO] [helloWorldResponse] ", r.Method, " request ", r.RequestURI, " from ", r.RemoteAddr)
	n, _ := strconv.Atoi(r.URL.Query().Get("error"))

	err := Templates.ExecuteTemplate(w, "index.html", getIndexPageData(n))
	if err != nil {
		log.Print("[ERROR] [helloWorldResponse] ", err.Error())
		return
	}
}
