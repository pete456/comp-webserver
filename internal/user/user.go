package user

import (
	"comp-webserver/config"
	"syscall"
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

func CreateUnixAccount(user string, pass string) {
	args := []string{config.Adduserscript,user,pass,config.Usersites}
	syscall.ForkExec(config.Adduserscript,args,&syscall.ProcAttr{Dir:""})
}

