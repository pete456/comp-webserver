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
	err := web.InsertUser(sqldb,"test","tes pass","title for site")
	if err != nil {
		log.Println(err)
	}

	// Web pages
	fs := http.FileServer(http.Dir("web/"))
	http.Handle("/",fs)

	// Post
	http.HandleFunc("/adduser",user.AddUser(sqldb))
	log.Fatal(http.ListenAndServe(config.Serverport,nil))
}
