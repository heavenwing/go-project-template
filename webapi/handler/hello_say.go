package handler

import (
	"fmt"
	"net/http"
	"strings"
)

const (
	// The route path of Say action of Hello handler (controller)
	RoutePath_Hello_Say = "/hello/say"
)

func (s *HelloHandler) SayAction(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	var name string
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
		if k == "name" {
			name = v[0]
		}
	}
	fmt.Fprintf(w, "Hello %s!", name)
}
