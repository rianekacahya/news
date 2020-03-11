package response

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/rianekacahya/news/pkg/errors"
	"net/http"
)

const MessageEmpty = ""

type response struct {
	Message interface{} `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func Error(c echo.Context, err error) error {
	var (
		status   int
		response = new(response)
	)

	if err != nil {
		// Mapping Status
		errorStatus := errors.GetStatus(err)
		errorMessage := errors.GetError(err)
		switch errorStatus {
		case errors.Generic:
			status = http.StatusInternalServerError
		case errors.Forbidden:
			status = http.StatusForbidden
		case errors.Badrequest:
			status = http.StatusBadRequest
		case errors.Notfound:
			status = http.StatusNotFound
		case errors.Unauthorize:
			status = http.StatusUnauthorized
		case errors.SessionExpire:
			status = http.StatusLocked
		default:
			status = http.StatusInternalServerError
			response.Message = err.Error()
		}

		if errorStatus != errors.Notype {
			switch fmt.Sprintf("%T", errorMessage) {
			case "*errors.errorString":
				response.Message = errorMessage.Error()
			default:
				response.Message = errorMessage
			}
		}
	}

	return c.JSON(status, response)
}

func Render(c echo.Context, status int, message string, data interface{}) error {
	var response = new(response)

	response.Message = "success"

	if message != "" {
		response.Message = message
	}

	response.Data = data

	return c.JSON(status, response)
}
