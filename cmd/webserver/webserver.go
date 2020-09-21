package main

import (
	"webserver/config"
	"webserver/internal/user"
	"net/http"
	"log"
)

func main() {
	// Web pages
	http.HandleFunc("/",web.Index)

	// Post
	http.HandleFunc("/adduser",user.AddUser)
	log.Fatal(http.ListenAndServe(config.Serverport,nil))
}
