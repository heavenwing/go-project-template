package manager

import "fmt"

type SayManager struct {
}

func NewSayManager() *SayManager {
	return &SayManager{}
}

func (s *SayManager) Say(name string) string {
	return fmt.Sprintf("Hello %s!", name)
}
