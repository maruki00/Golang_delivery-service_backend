package pkg_rabbitmq

import (
	"context"

	"github.com/rabbitmq/amqp091-go"
)

type Publisher struct {
	AMQPConn *amqp091.Connection
}

func NewPublisher(dsn string) *Publisher {
	con, err := amqp091.Dial(dsn)
	if err != nil {
		panic("could not ")
	}
	return &Publisher{
		AMQPConn: con,
	}
}

func (p *Publisher) Publish(ctx context.Context, data []byte, queueName, contentType, msgTypeName string) error {
	ch, err := p.AMQPConn.Channel()
	if err != nil {
		return err
	}
	return ch.PublishWithContext(
		ctx,
		"",
		queueName,
		false,
		false,
		amqp091.Publishing{
			ContentType: contentType,
			Type:        msgTypeName,
			Body:        data,
		},
	)
}
