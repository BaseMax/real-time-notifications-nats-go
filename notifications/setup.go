package notifications

import "github.com/nats-io/nats.go"

var conn *nats.Conn

func InitNats() error {
	var err error
	conn, err = nats.Connect(GetNatsURL())
	if err != nil {
		return err
	}
	return nil
}
