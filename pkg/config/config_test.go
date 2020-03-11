package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfig(t *testing.T) {
	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Fatalf("Error when loading config %v", r)
			}
		}()

		cfg := newConfig()
		assert.Equal(t, configurations{}, cfg)
	}()
}

func TestGetConfig(t *testing.T) {
	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Fatalf("Error when loading config %v", r)
			}
		}()

		cfg := GetConfig()
		assert.Equal(t, configurations{}, cfg)
	}()
}
