package rabbitmq

import (
	"encoding/json"
	"github.com/rianekacahya/news/pkg/config"
	"github.com/rianekacahya/news/pkg/logger"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"log"
	"sync"
	"time"
)

var (
	connection   *amqp.Connection
	mutex, queue sync.Mutex
)

func GetConnection() *amqp.Connection {
	if connection == nil {
		mutex.Lock()
		defer mutex.Unlock()
		connection = newConnection()
	}

	return connection
}

func newConnection() *amqp.Connection {
	client, err := amqp.Dial(config.GetConfig().RabbitmqDSN)
	if err != nil {
		log.Panicf("got an error while connecting rabbitMQ server, error: %s", err)
	}

	return client
}

func Publish(topic string, msg []byte, trackingID string) error {
	channel, err := GetConnection().Channel()
	if err != nil {
		return err
	}
	defer channel.Close()

	payload := amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "application/json",
		Body:         msg,
		Headers: amqp.Table{
			"tracking_id": trackingID,
		},
	}

	if err := channel.Publish("", topic, false, false, payload); err != nil {
		return err
	}

	// loging publish
	Log(topic, "publish", trackingID, msg)

	return nil
}

func Subscribe(topic string, handler func(amqp.Delivery)) error {
	channel, err := GetConnection().Channel()
	if err != nil {
		return err
	}
	defer channel.Close()

	queue, err := channel.QueueDeclare(topic, false, false, false, false, nil)
	if err != nil {
		return err
	}

	consumer, err := channel.Consume(queue.Name, "", false, false, false, false, nil)
	if err != nil {
		return err
	}

	for delivery := range consumer {
		handler(delivery)
	}

	return nil
}

func Log(topic, status, trackingID string, msg []byte) {
	logger.Info(
		"event-logger",
		zap.String("request_id", trackingID),
		zap.String("time", time.Now().Format(time.RFC1123Z)),
		zap.String("topic", topic),
		zap.String("status", status),
		zap.Any("payload", json.RawMessage(msg)),
	)
}
