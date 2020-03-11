package transport

import (
	"github.com/labstack/echo"
	"github.com/rianekacahya/news/internal/v1/news"
	"github.com/rianekacahya/news/internal/v1/news/entity"
	"github.com/rianekacahya/news/pkg/echoserver"
	"github.com/rianekacahya/news/pkg/echoserver/response"
	"github.com/rianekacahya/news/pkg/errors"
	"net/http"
)

type Rest struct {
	usecase news.Usecase
}

func NewRest(usecase news.Usecase) {
	transport := Rest{usecase}

	news := echoserver.GetServer().Group("/news")
	news.GET("", transport.list)
	news.POST("", transport.create)
}

func (DI *Rest) list(c echo.Context) error {
	var (
		err error
		req = new(entity.Request)
	)

	if err = c.Bind(req); err != nil {
		return response.Error(c, errors.Message("Binding request data failed."))
	}

	if err = req.QueryNewsValidation(); err != nil {
		return response.Error(c, errors.New(errors.Badrequest, err))
	}

	result, err := DI.usecase.List(req)
	if err != nil {
		return response.Error(c, err)
	}

	return response.Render(c, http.StatusOK, response.MessageEmpty, result)
}

func (DI *Rest) create(c echo.Context) error {
	var (
		err       error
		req       = new(entity.News)
		requestID = c.Response().Header().Get(echo.HeaderXRequestID)
	)

	if err = c.Bind(req); err != nil {
		return response.Error(c, errors.Message("Binding request data failed."))
	}

	if err = req.CommandNewsValidation(); err != nil {
		return response.Error(c, errors.New(errors.Badrequest, err))
	}

	result, err := DI.usecase.Create(req, requestID)
	if err != nil {
		return response.Error(c, err)
	}

	return response.Render(c, http.StatusOK, response.MessageEmpty, result)
}
