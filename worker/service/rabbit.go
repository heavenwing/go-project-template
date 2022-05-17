package service

import (
	"heavenwing/go-project-template/worker/global"

	"github.com/streadway/amqp"
)

const (
	queueName    = "test"
	exchangeName = "events"
)

type RabbitService struct {
	connection *amqp.Connection
	channel    *amqp.Channel
}

func NewRabbitService() *RabbitService {
	// Connect to the rabbitMQ instance
	connection, err := amqp.Dial(global.AMQP_URL)

	if err != nil {
		panic("could not establish connection with RabbitMQ:" + err.Error())
	}

	// Create a channel from the connection. We'll use channels to access the data in the queue rather than the connection itself.
	channel, err := connection.Channel()

	if err != nil {
		panic("could not open RabbitMQ channel:" + err.Error())
	}

	// We create an exchange that will bind to the queue to send and receive messages
	err = channel.ExchangeDeclare(exchangeName, amqp.ExchangeTopic, true, false, false, false, nil)

	if err != nil {
		panic(err)
	}

	// We create a queue named Test
	_, err = channel.QueueDeclare(queueName, true, false, false, false, nil)

	if err != nil {
		panic("error declaring the queue: " + err.Error())
	}

	// We bind the queue to the exchange to send and receive data from the queue
	err = channel.QueueBind(queueName, "#", exchangeName, false, nil)

	if err != nil {
		panic("error binding to the queue: " + err.Error())
	}

	return &RabbitService{connection: connection, channel: channel}
}

func (s *RabbitService) ReceiveMessages() (<-chan amqp.Delivery, error) {
	// We consume data in the queue named test using the channel we created in go.
	msgs, err := s.channel.Consume(queueName, "", false, false, false, false, nil)

	if err != nil {
		panic("error consuming the queue: " + err.Error())
	}

	return msgs, err
}

func (s *RabbitService) PublishMessage(key string, msg string) error {
	// We create a message to be sent to the queue.
	// It has to be an instance of the aqmp publishing struct
	message := amqp.Publishing{
		Body: []byte(msg),
	}

	// We publish the message to the exahange we created earlier
	err := s.channel.Publish(exchangeName, key, false, false, message)

	if err != nil {
		panic("error publishing a message to the queue:" + err.Error())
	}

	return err
}

func (s *RabbitService) Close() {
	s.connection.Close()
}
