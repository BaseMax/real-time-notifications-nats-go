package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	conn *amqp.Connection
	ch   *amqp.Channel
)

const (
	QUEUE_NAME_ORDERS  = "orders"
	QUEUE_NAME_REFUNDS = "refunds"
)

func Connect() (err error) {
	conn, err = amqp.Dial(GetRabbitUrl())
	if err != nil {
		return err
	}
	ch, err = conn.Channel()
	return
}

func IsClosed() bool {
	if conn == nil || ch == nil {
		return true
	}
	return conn.IsClosed()
}

func RestartChannel() (err error) {
	if IsClosed() {
		if err = Connect(); err != nil {
			return
		}
	}

	if ch != nil {
		ch.Close()
		ch, err = conn.Channel()
		if err != nil {
			return err
		}
	}
	return err
}
