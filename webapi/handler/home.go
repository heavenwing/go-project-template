package handler

import (
	"fmt"
	"heavenwing/go-project-template/webapi/service"
	"net/http"
	"strings"
)

// HelloHandler is hello handler
type HomeHandler struct {
	authService service.AuthInterface
}

func NewHomeHandler(authService service.AuthInterface) *HomeHandler {
	return &HomeHandler{authService: authService}
}

const (
	// The route path of Index action of Hello handler (controller)
	RoutePath_Home_Index = "/"
)

func (s *HomeHandler) IndexAction(w http.ResponseWriter, r *http.Request) {
	if !s.authService.Login("fake", "fake") {
		return
	}

	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello world from home!")
}
