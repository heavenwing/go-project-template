package main

import (
	"heavenwing/go-project-template/webapi/handler"
	"heavenwing/go-project-template/webapi/manager"
	"heavenwing/go-project-template/webapi/middleware"
	"heavenwing/go-project-template/webapi/service"
	"log"
	"net/http"
)

func init() {
	// just only init code
	// don't contains any IO access
	// don't contains any codes which will impact unit test running
}

func main() {

	authService := service.NewAuthService()

	logger := middleware.Logger{}

	home := handler.NewHomeHandler(authService)
	http.Handle(handler.RoutePath_Home_Index, logger.Middleware(http.HandlerFunc(home.IndexAction)))

	sayManager := manager.NewSayManager()
	hello := handler.NewHelloHandler(sayManager)
	http.Handle(handler.RoutePath_Hello_Index, logger.Middleware(http.HandlerFunc(hello.IndexAction)))
	http.Handle(handler.RoutePath_Hello_Say, logger.Middleware(http.HandlerFunc(hello.SayAction)))

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
