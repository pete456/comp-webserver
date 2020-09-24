package main

import (
	"comp-webserver/config"
	"comp-webserver/internal/web"

	"net/http"
	"log"
)

func main() {
	sqldb := web.CreateDB()

	// Web pages
	html := http.FileServer(http.Dir("web/html/"))
	http.Handle("/",html)

	http.HandleFunc("/site/",web.HandleUserSite)

	js := http.FileServer(http.Dir("web/js/"))
	http.Handle("/js/",http.StripPrefix("/js/",js))

	// Post
	http.HandleFunc("/adduser",web.AddUser(sqldb))

	// Get
	http.HandleFunc("/getportaldata",web.GetPortalData(sqldb))


	log.Fatal(http.ListenAndServe(config.Serverport,nil))
}
