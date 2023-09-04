package rabbitmq

import (
	"context"
	"encoding/json"

	"github.com/rabbitmq/amqp091-go"
)

func EnqueueTask[T any](model T, queueName string) error {
	_, err := ch.QueueDeclare(queueName, true, false, false, false, amqp091.Table{})
	if err != nil {
		return err
	}

	body, _ := json.Marshal(model)
	err = ch.PublishWithContext(context.Background(), "", queueName, false, false, amqp091.Publishing{
		ContentType: "json/application",
		Body:        body,
	})
	return err
}
