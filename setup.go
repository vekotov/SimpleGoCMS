package main

import (
	"html/template"
	"log"
	"net/http"
)

func setupStatic() {
	http.Handle(
		"/static/",
		http.StripPrefix(
			"/static/",
			http.FileServer(http.Dir("static")),
		),
	)
}

func setupDynamic() {
	http.HandleFunc("/", indexPageResponse)
	http.HandleFunc("/add_post", addPostResponse)
}

func setupTemplates() {
	var err error
	Templates, err = template.ParseGlob("templates/*")
	if err != nil {
		log.Fatal("[ERROR] [setupTemplates] ", err.Error())
	}
}
