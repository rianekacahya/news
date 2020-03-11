package elasticsearch

import (
	"github.com/olivere/elastic/v7"
	"github.com/rianekacahya/news/pkg/config"
	"log"
	"sync"
)

var (
	connection *elastic.Client
	mutex      sync.Mutex
)

func GetConnection() *elastic.Client {
	if connection == nil {
		mutex.Lock()
		defer mutex.Unlock()
		connection = newConnection()
	}

	return connection
}

func newConnection() *elastic.Client {
	client, err :=  elastic.NewClient(elastic.SetURL(config.GetConfig().ElasticsearchURL),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false))

	if err != nil {
		log.Panicf("got an error while connecting elasticsearch server, error: %s", err)
	}

	return client
}