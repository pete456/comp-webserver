package main

import (
	"comp-webserver/config"
	"comp-webserver/internal/user"
	"comp-webserver/internal/web"

	"net/http"
	"log"
)

func main() {
	sqldb := web.CreateDB()

	// Web pages
	html := http.FileServer(http.Dir("web/html/"))
	http.Handle("/",html)

	js := http.FileServer(http.Dir("web/js/"))
	http.Handle("/js/",http.StripPrefix("/js/",js))

	// Post
	http.HandleFunc("/adduser",user.AddUser(sqldb))

	log.Fatal(http.ListenAndServe(config.Serverport,nil))
}
