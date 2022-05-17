package worker

import "fmt"

func processMessage(msg string) {
	fmt.Println("message received: " + msg)
}
