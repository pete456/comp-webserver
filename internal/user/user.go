package user

import (
	"comp-webserver/config"
	"comp-webserver/internal/web"
	"encoding/json"
	"net/http"
	"database/sql"
	"syscall"
	"log"
	"strings"
)

type User struct {
	User	string
	Pass	string
	Title	string
}

func sanatizetitle(title string) string {
	return strings.ToLower(title)
}

func createunixaccount(user string, pass string) {
	args := []string{config.Adduserscript,user,pass,config.Usersites}
	syscall.ForkExec(config.Adduserscript,args,&syscall.ProcAttr{Dir:""})
}

func AddUser(sql *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if (r.Method == "POST") {
			var user User
			err := json.NewDecoder(r.Body).Decode(&user)
			if err != nil {
				http.Error(w,err.Error(),http.StatusBadRequest)
			}

			log.Printf("Creating user: %s\n",user.User)
			web.InsertUser(sql,user.User,user.Pass,user.Title)
			createunixaccount(user.User,user.Pass)
		}
	}
}
