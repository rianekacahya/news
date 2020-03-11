package echoserver

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/rianekacahya/news/pkg/logger"
	"go.uber.org/zap"
	"log"
	"net/http"
	"time"
)

type response struct {
	Message interface{} `json:"message"`
}

func Handler(err error, c echo.Context) {
	var (
		code        = http.StatusInternalServerError
		msg, logmsg interface{}
	)

	if GetServer().Debug {
		msg = err.Error()
		switch err.(type) {
		case *echo.HTTPError:
			code = err.(*echo.HTTPError).Code
		}
	} else {
		switch e := err.(type) {
		case *echo.HTTPError:
			code = e.Code
			msg = e.Message
			if e.Internal != nil {
				msg = fmt.Sprintf("%v, %v", err, e.Internal)
			}
		default:
			msg = http.StatusText(code)
			logmsg = e.Error()
		}

		if _, ok := msg.(string); ok {
			msg = response{Message: msg}
		}
	}

	if !c.Response().Committed {
		if c.Request().Method == "HEAD" {
			err = c.NoContent(code)
		} else {
			err = c.JSON(code, msg)
		}

		if code == http.StatusInternalServerError {
			// Log error message
			logger.Error(
				"http-logger",
				zap.String("requestID", c.Response().Header().Get(echo.HeaderXRequestID)),
				zap.Int("status", code),
				zap.String("time", time.Now().Format(time.RFC1123Z)),
				zap.String("hostname", c.Request().Host),
				zap.String("user_agent", c.Request().UserAgent()),
				zap.String("method", c.Request().Method),
				zap.String("path", c.Path()),
				zap.String("query", c.QueryString()),
				zap.Any("response", response{Message: logmsg}),
			)
		}

		if err != nil {
			log.Panicf("got an error while serve data, error: %s", err)
		}
	}
}
