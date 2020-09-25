package web

import (
	"comp-webserver/internal/user"
	"encoding/json"
	"net/http"
	"database/sql"
	"log"
)

func AddUser(sql *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if (r.Method == "POST") {
			var u user.User
			err := json.NewDecoder(r.Body).Decode(&u)
			if err != nil {
				http.Error(w,err.Error(),http.StatusBadRequest)
			}

			log.Printf("Creating user: %s\n",u.User)
			InsertUser(sql,u.User,u.Pass,u.Title)
			user.CreateUnixAccount(u.User,u.Pass)
			//AddUserHandle(u.User)
		}
	}
}

type portaldata struct {
	Name string `json:name`
	Title string `json:title`
}

type portal struct {
	PD []portaldata
}

func GetPortalData(sql *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		data := titlesusersfromdb(sql)
		p := portal{data}
		w.Header().Set("Content-Type","application/json")
		err := json.NewEncoder(w).Encode(p)
		if err != nil {
			log.Fatal(err)
		}
	}
}
