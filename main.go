package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", helloWorldResponse)
	http.HandleFunc("/add_post", addPostResponse)
	log.Print("[INFO] Starting web-server at ", os.Getenv("PORT"), " port")
	http.Handle(
		"/static/",
		http.StripPrefix(
			"/static/",
			http.FileServer(http.Dir("static")),
		),
	)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+os.Getenv("PORT"), nil))
}

func helloWorldResponse(w http.ResponseWriter, r *http.Request) {
	log.Print("[INFO] [helloWorldResponse] Got request ", r.RequestURI, " from ", r.RemoteAddr)

	tmpl, err := template.ParseGlob("templates/*")
	if err != nil {
		log.Print("[WARNING] [helloWorldResponse] ", err.Error())
		return
	}

	data := struct {
		Posts map[int64]*Post
	}{
		Posts: getPosts(),
	}

	err = tmpl.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		log.Print("[WARNING] [helloWorldResponse] ", err.Error())
		return
	}
}

func addPostResponse(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		addPost(r.PostFormValue("post_text"), r.RemoteAddr)
	}
	http.Redirect(w, r, "/", 303)
}
