package notifications

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/nats-io/nats.go"
)

type NotifErr struct {
	Err       error
	HTTPError echo.HTTPError
}

func (e *NotifErr) Error() string { return e.Err.Error() }

func (e *NotifErr) Unwrap() error { return e.Err }

func GetNatsURL() string {
	url := os.Getenv("NATS_URL")
	if url == "" {
		return nats.DefaultURL
	}
	return url
}

func CreateSubject(id uint) string {
	return fmt.Sprintf("notify.%d", id)
}
