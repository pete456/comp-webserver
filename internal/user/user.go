package user

import (
	"webserver/config"
        "net/http"
        "syscall"
        "log"
	"strings"
)

type user struct {
	User	string,
	Pass	string,
	Pagepath string,
	Title	string
}

func sanatizetitle(title string) {
	return strings.ToLower(title)
}

func createunixaccount(user string, pass string) {
                args := []string{config.Adduserscript,user,pass}
                syscall.ForkExec(config.Adduserscript,args,&syscall.ProcAttr{Dir:""})
}


func AddUser(w http.ResponseWriter, r *http.Request) {
        if (r.Method == "POST") {
                log.Printf("post user: %s\n",r.PostFormValue("user"))
                user := r.PostFormValue("user")
                pass := r.PostFormValue("pass")
		path := sanatizetitle(title)
		title := r.PostFormValue("title")
		user := user{user,pass,path,title)
		createunixaccount(user,pass)
        }
}
