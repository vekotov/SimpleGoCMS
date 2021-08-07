package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloWorldResponse)
	log.Print("[INFO] Starting web-server at :8080 port")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}

func helloWorldResponse(w http.ResponseWriter, r *http.Request) {
	log.Print("[INFO] [helloWorldResponse] Got request ", r.RequestURI, " from ", r.RemoteAddr)
	if _, err := fmt.Fprintf(w, "Hello world! You are: %s", r.RemoteAddr); err != nil {
		log.Print("[WARNING] [helloWorldResponse] ", err.Error())
	}
}
