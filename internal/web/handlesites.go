package web

import (
	"comp-webserver/config"
	"net/http"
	"html"
	"strings"
	"path"
)

func removeholes(lst []string) []string {
	var newlst []string
	for _,v := range lst {
		if v != "" {
			newlst = append(newlst,v)
		}
	}
	return newlst
}

func shiftpath(p string) (head, tail string) {
	p = path.Clean("/"+p)
	i := strings.Index(p[1:],"/") + 1
	if i <= 0 {
		return p[1:],"/"
	}
	return p[1:i], p[i:]
}

func HandleUserSite(w http.ResponseWriter, r *http.Request) {
	url := html.EscapeString(r.URL.Path)
	_, userquery := shiftpath(url)
	if len(userquery) == 1 {
		http.Error(w,"Adding a user g, and stop breaking stuff",404)
	}
	http.ServeFile(w,r,config.Usersites+userquery)
}
