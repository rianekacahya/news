package echoserver

import (
	"github.com/labstack/echo"
	"github.com/rianekacahya/news/pkg/logger"
	"go.uber.org/zap"
	"time"
)

// CustomBinder struct
type CustomBinder struct {
	bind echo.Binder
}

func NewBinder() *CustomBinder {
	return &CustomBinder{bind: &echo.DefaultBinder{}}
}

// Bind tries to bind request into interface, and if it does then validate it
func (cb *CustomBinder) Bind(i interface{}, c echo.Context) error {
	if err := cb.bind.Bind(i, c); err != nil && err != echo.ErrUnsupportedMediaType {
		return err
	}

	// log bind data
	logger.Info(
		"http-logger",
		zap.String("requestID", c.Response().Header().Get(echo.HeaderXRequestID)),
		zap.Int("status", c.Response().Status),
		zap.String("time", time.Now().Format(time.RFC1123Z)),
		zap.String("hostname", c.Request().Host),
		zap.String("user_agent", c.Request().UserAgent()),
		zap.String("method", c.Request().Method),
		zap.String("path", c.Path()),
		zap.String("query", c.QueryString()),
		zap.Any("request", i),
	)

	return nil
}
