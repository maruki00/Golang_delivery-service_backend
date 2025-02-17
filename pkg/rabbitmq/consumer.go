package pkg_rabbitmq

import (
	"context"

	"github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	AMQPConn *amqp091.Connection
	Queues   []string
}

func NewConsumer(dsn string, queues []string) (*Consumer, error) {
	conn, err := amqp091.Dial(dsn)
	if err != nil {
		return nil, err
	}

	return &Consumer{
		AMQPConn: conn,
		Queues:   queues,
	}, nil
}

func (c *Consumer) Start(worder func(ctx context.Context, delivery <-chan amqp091.Delivery)) {
	ch, err := c.AMQPConn.Channel()
	if err != nil {
		panic("error: could not oppen the channel")
	}
	if err := ch.Qos(1, 1, false); err != nil {
		panic("error: error Qos")
	}
	for _, queueName := range c.Queues {
		queue, err := ch.QueueDeclare(
			queueName,
			true,
			false,
			false,
			false,
			nil,
		)
		if err != nil {
			panic(err)
		}
		deliveryMessages, err := ch.Consume(
			queue.Name,
			"",
			false,
			false,
			false,
			false,
			nil,
		)
		if err != nil {
			panic(err)
		}
		ctx := context.Background()
		for range 10 {
			go worder(ctx, deliveryMessages)
		}

	}

	// var forever chan struct{}

	// go func() {
	// 	for d := range msgs {
	// 		log.Printf("Received a message: %s", d.Body)
	// 	}
	// }()

	// log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	// <-forever
}
