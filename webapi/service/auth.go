package service

import "log"

type AuthService struct {
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) Login(username string, password string) bool {
	log.Printf("username: %s", username)
	return true
}
