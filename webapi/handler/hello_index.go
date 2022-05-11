package handler

import (
	"fmt"
	"net/http"
	"strings"
)

const (
	// The route path of Index action of Hello handler (controller)
	RoutePath_Hello_Index = "/hello"
)

func (s *HelloHandler) IndexAction(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello world from hello!")
}
