package user

import (
	"comp-webserver/config"
	"comp-webserver/internal/web"
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
	args := []string{config.Adduserscript,config.Usersites,user,pass}
	syscall.ForkExec(config.Adduserscript,args,&syscall.ProcAttr{Dir:""})
}


func AddUser(sql *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if (r.Method == "POST") {
			log.Printf("post user: %s\n",r.PostFormValue("user"))
			uname := r.PostFormValue("user")
			pass := r.PostFormValue("pass")
			title := r.PostFormValue("title")
			//u := user{uname,pass,title}
			web.InsertUser(sql,uname,pass,title)
			createunixaccount(uname,pass)
		}
	}
}
