package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

var Templates *template.Template = nil

func main() {
	initDB()
	dbLoadPosts()
	setupDynamic()
	setupStatic()
	setupTemplates()
	log.Print("[INFO] Starting web-server at ", os.Getenv("PORT"), " port")
	log.Fatal(http.ListenAndServe("0.0.0.0:"+os.Getenv("PORT"), nil))
}
