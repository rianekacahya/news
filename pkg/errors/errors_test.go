package errors

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	init := New(Notype, Message("error"))
	if init == nil {
		t.Errorf("Server should not be nil")
	}
}

func TestMessage(t *testing.T) {
	init := Message("error")
	assert.Equal(t, errors.New("error"), init)
}

func TestGetStatus(t *testing.T) {
	init := New(Badrequest, Message("error"))
	assert.Equal(t, Badrequest, GetStatus(init))
}

func TestGetError(t *testing.T) {
	init := New(Notype, Message("error"))
	assert.Equal(t, errors.New("error"), GetError(init))
}
