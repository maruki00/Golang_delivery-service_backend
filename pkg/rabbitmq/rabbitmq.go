package pkg_rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	Conn      *amqp.Connection
	Channel   *amqp.Channel
	QueneName string
	// RabbitMQDSN string
}

func NewRAbbitMQ(rabbitMQDSN string, queneName string) *RabbitMQ {
	conn, err := amqp.Connection(rabbitMQDSN)
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	return &RabbitMQ{
		Conn:      conn,
		Channel:   ch,
		QueneName: queneName,
	}
}
