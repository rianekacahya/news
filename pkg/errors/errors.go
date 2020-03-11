package errors

import (
	"errors"
)

// General Error Messages
var (
	ErrorNotFound       = "Data not found."
	ErrorForbidden      = "Access dont have permission."
	ErrorBadRequest     = "Invalid data request."
	ErrorUnauthorize    = "Access unauthorized."
	ErrorInternalServer = "Something went wrong in system."
)

type custom struct {
	status int
	error  error
}

const (
	Notype        = 1
	Generic       = 2
	Forbidden     = 3
	Badrequest    = 4
	Notfound      = 5
	Unauthorize   = 6
	SessionExpire = 7
)

// New generates an application error
func New(status int, err error) *custom {
	return &custom{status: status, error: err}
}

// Error returns the error message.
func (e custom) Error() string {
	return e.error.Error()
}

func Message(msg string) error {
	return errors.New(msg)
}

// Get status error
func GetStatus(err error) int {
	if custom, ok := err.(*custom); ok {
		return custom.status
	}

	return Notype
}

// Get error value
func GetError(err error) error {
	if custom, ok := err.(*custom); ok {
		return custom.error
	}

	return err
}
