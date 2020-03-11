package rabbitmq

import (
	"encoding/json"
	"fmt"
	"github.com/rianekacahya/news/pkg/errors"
	"github.com/rianekacahya/news/pkg/logger"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"time"
)

const (
	InternalServerError = "Internal server error"
	Unauthorize         = "Unauthorize"
	NotFound            = "Not found"
	BadRequest          = "Bad Request"
	Forbidden           = "Forbidden"
)


func Error(topic, trackingID string, msg amqp.Delivery, err error) {
	// init require variable
	var (
		status  string
		Errors  interface{}
	)

	if err != nil {
		// Mapping error
		errorStatus := errors.GetStatus(err)
		errorMessage := errors.GetError(err)
		switch errorStatus {
		case errors.Generic:
			status = InternalServerError
		case errors.Forbidden:
			status = Forbidden
		case errors.Badrequest:
			status = BadRequest
		case errors.Notfound:
			status = NotFound
		case errors.Unauthorize:
			status = Unauthorize
		default:
			status = InternalServerError
			Errors = err.Error()
		}

		if errorStatus != errors.Notype {
			switch fmt.Sprintf("%T", errorMessage) {
			case "*errors.errorString":
				Errors = errorMessage.Error()
			default:
				Errors = errorMessage
			}
		}

		// logging error
		logger.Error(
			"event-logger",
			zap.String("request_id", trackingID),
			zap.String("time", time.Now().Format(time.RFC1123Z)),
			zap.String("topic", topic),
			zap.String("status", status),
			zap.Any("error", Errors),
			zap.Any("payload", json.RawMessage(msg.Body)),
		)
	}
}
