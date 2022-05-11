package handler

import "heavenwing/go-project-template/webapi/manager"

// HelloHandler is hello handler
type HelloHandler struct {
	manager manager.SayInterface
}

func NewHelloHandler(manager manager.SayInterface) *HelloHandler {
	return &HelloHandler{manager: manager}
}
