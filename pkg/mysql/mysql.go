package mysql

import (
	"database/sql"
	"fmt"
	"github.com/rianekacahya/news/pkg/errors"
	"github.com/rianekacahya/news/pkg/logger"
	"go.uber.org/zap"
	"log"
	"regexp"
	"strings"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rianekacahya/news/pkg/config"
)

var (
	connection *sql.DB
	mutex      sync.Mutex
)

func GetConnection() *sql.DB {
	if connection == nil {
		mutex.Lock()
		defer mutex.Unlock()
		connection = newConnection()
	}

	return connection
}

func newConnection() *sql.DB {
	conn, err := sql.Open(config.GetConfig().DatabaseDriver, config.GetConfig().DatabaseDSN)
	if err != nil {
		log.Panicf("got an error while connecting database server, error: %s", err)
	}

	conn.SetConnMaxLifetime(time.Duration(config.GetConfig().DatabaseTimeout) * time.Second)
	conn.SetMaxOpenConns(int(config.GetConfig().DatabaseMaxOpenConnection))
	conn.SetMaxIdleConns(int(config.GetConfig().DatabaseMaxIdleConnection))

	if err = conn.Ping(); err != nil {
		log.Panicf("got an error while connecting database server, error: %s", err)
	}

	return conn
}

func Debug(query string, args []interface{}, err error) {
	fields := []zap.Field{
		zap.String("start", time.Now().Format(time.RFC1123Z)),
		zap.String("statement", fmt.Sprintf(strings.ReplaceAll(regexp.MustCompile(`\s+`).ReplaceAllString(query, " "), "?", "'%v'"), args...)),
	}

	if err != nil {
		fields = append(fields, zap.String("error", err.Error()))
	}

	logger.Info("query-logger", fields...)
}

func Error(err error) error {
	switch {
	case err == sql.ErrNoRows:
		return errors.New(errors.Notfound, errors.Message(errors.ErrorNotFound))
	case err != nil:
		return errors.New(errors.Generic, errors.Message(errors.ErrorInternalServer))
	}

	return nil
}
