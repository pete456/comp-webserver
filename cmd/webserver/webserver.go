package main

import (
	"comp-webserver/config"
	"comp-webserver/internal/user"
	"net/http"
	"log"
)

func main() {
	// Web pages
	fs := http.FileServer(http.Dir("web/"))
	http.Handle("/",fs)

	// Post
	http.HandleFunc("/adduser",user.AddUser)
	log.Fatal(http.ListenAndServe(config.Serverport,nil))
}
