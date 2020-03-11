package redis

import (
	"log"
	"sync"
	"time"

	"github.com/go-redis/redis"
	"github.com/rianekacahya/news/pkg/config"
)

var (
	connection *redis.Client
	mutex      sync.Mutex
)

func GetConnection() *redis.Client {
	if connection == nil {
		mutex.Lock()
		defer mutex.Unlock()
		connection = newConnection()
	}

	return connection
}

func newConnection() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:         config.GetConfig().RedisAddress,
		Password:     "",
		PoolTimeout:  2 * time.Second,
		IdleTimeout:  2 * time.Second,
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
	})

	_, err := rdb.Ping().Result()
	if err != nil {
		log.Panicf("got an error while connecting redis server, error: %s", err)
	}

	return rdb
}
