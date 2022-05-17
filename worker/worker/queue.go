package worker

import (
	"heavenwing/go-project-template/worker/service"
)

type QueueWorker struct {
	service service.RabbitServiceInterface
}

func NewQueueWorker(service service.RabbitServiceInterface) *QueueWorker {
	return &QueueWorker{service: service}
}

func (s *QueueWorker) Process() {

	msgs, err := s.service.ReceiveMessages()

	if err != nil {
		panic("error consuming the queue: " + err.Error())
	}

	// We loop through the messages in the queue and print them in the console.
	// The msgs will be a go channel, not an amqp channel
	for msg := range msgs {
		processMessage(string(msg.Body))
		msg.Ack(false)
	}
}
