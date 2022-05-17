package main

import (
	"heavenwing/go-project-template/worker/global"
	"heavenwing/go-project-template/worker/service"
	"heavenwing/go-project-template/worker/worker"
)

//ref: https://dev.to/olushola_k/working-with-rabbitmq-in-golang-1kmj

func main() {

	global.LoadSetting()

	rabbit := service.NewRabbitService()
	queue := worker.NewQueueWorker(rabbit)

	queue.Process()

	defer rabbit.Close()
}
