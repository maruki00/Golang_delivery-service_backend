package pkgRabbitmq

import (
	"errors"
	"log/slog"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	MAX_TRIES      = 3
	BACKOFFSECONDS = 3
)

type RabbitMQConnStr string

var ErrCannotConnectRabbitMQ = errors.New("cannot connect to rabbit")

func NewRabbitMQConn(rabbitMqURL RabbitMQConnStr) (*amqp.Connection, error) {
	var (
		amqpConn *amqp.Connection
		counts   int64
	)

	for {
		connection, err := amqp.Dial(string(rabbitMqURL))
		if err != nil {
			slog.Error("failed to connect to RabbitMq...", err.Error(), rabbitMqURL)
			counts++
		} else {
			amqpConn = connection
			break
		}

		if counts > MAX_TRIES {
			slog.Error("failed to retry", err)
			return nil, ErrCannotConnectRabbitMQ
		}

		slog.Info("Backing off for 2 seconds...")
		time.Sleep(BACKOFFSECONDS * time.Second)

		continue
	}

	slog.Info("📫 connected to rabbitmq 🎉")

	return amqpConn, nil
}
