package transport

import (
	"encoding/json"
	"fmt"
	"github.com/rianekacahya/news/internal/v1/news"
	"github.com/rianekacahya/news/internal/v1/news/entity"
	"github.com/rianekacahya/news/pkg/rabbitmq"
	"github.com/streadway/amqp"
)

type Event struct {
	usecase news.Usecase
}

func NewEvent(usecase news.Usecase) {
	transport := Event{usecase}

	// Run subscriber
	go rabbitmq.Subscribe(entity.TopicInsertNews, transport.insert)
}

func (DI *Event) insert(msg amqp.Delivery)  {
	var (
		request = new(entity.News)
		trackingID = msg.Headers["tracking_id"]
	)

	// log receive data
	rabbitmq.Log(entity.TopicInsertNews, "receive", fmt.Sprint(trackingID), msg.Body)

	err := json.Unmarshal(msg.Body, request)
	if err != nil {
		rabbitmq.Error(entity.TopicInsertNews, fmt.Sprint(trackingID), msg, err)
	}

	// Insert data news
	if err = DI.usecase.Insert(request); err != nil {
		rabbitmq.Error(entity.TopicInsertNews, fmt.Sprint(trackingID), msg, err)
	}

	// Ack event
	msg.Ack(true)
}