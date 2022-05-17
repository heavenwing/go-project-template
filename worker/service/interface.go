package service

import "github.com/streadway/amqp"

type RabbitServiceInterface interface {
	ReceiveMessages() (<-chan amqp.Delivery, error)
	PublishMessage(key string, msg string) error
}
