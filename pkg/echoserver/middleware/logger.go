package middleware

import (
	"encoding/json"
	"github.com/labstack/echo"
	"github.com/rianekacahya/news/pkg/logger"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func Logger(c echo.Context, reqBody, resBody []byte) {

	// init zap field
	fields := []zap.Field{
		zap.String("requestID", c.Response().Header().Get(echo.HeaderXRequestID)),
		zap.Int("status", c.Response().Status),
		zap.String("time", time.Now().Format(time.RFC1123Z)),
		zap.String("hostname", c.Request().Host),
		zap.String("user_agent", c.Request().UserAgent()),
		zap.String("method", c.Request().Method),
		zap.String("path", c.Request().URL.Path),
		zap.String("query", c.QueryString()),
	}

	// kondisi jika request body kosong
	if len(reqBody) == 0 {
		logger.Info("http-logger", append(fields, zap.Any("request", reqBody))...)
	}

	// kondisi normal logger
	if c.Response().Status >= http.StatusOK && c.Response().Status < http.StatusMultipleChoices {
		logger.Info("REST Logger", append(fields, zap.Any("response", json.RawMessage(resBody)))...)
	} else if c.Response().Status >= http.StatusBadRequest && c.Response().Status < http.StatusInternalServerError {
		logger.Error("REST Logger", append(fields, zap.Any("response", json.RawMessage(resBody)))...)
	}
}
