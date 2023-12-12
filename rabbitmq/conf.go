package rabbitmq

import (
	"fmt"
	"os"
)

func GetRabbitUrl() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%s",
		os.Getenv("RABBIT_USER"),
		os.Getenv("RABBIT_PASSWORD"),
		os.Getenv("RABBIT_HOSTNAME"),
		os.Getenv("RABBIT_PORT"),
	)
}
