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

func sanatize(value string) string {
	return strings.ToLower(value)
}

func CreateUnixAccount(user string, pass string) {
	args := []string{config.Adduserscript,user,pass,config.Usersites}
	syscall.ForkExec(config.Adduserscript,args,&syscall.ProcAttr{Dir:""})
}

