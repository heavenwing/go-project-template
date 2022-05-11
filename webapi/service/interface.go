package service

type AuthInterface interface {
	Login(username string, password string) bool
}
