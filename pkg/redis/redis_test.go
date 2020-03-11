package redis

import (
	"github.com/alicebob/miniredis"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestRedisConnect(t *testing.T) {
	mr, err := miniredis.Run()
	if err != nil {
		t.Fatalf("could not start miniredis, %s", err)
	}

	// set require environment variable
	os.Setenv("REDIS_ADDRESS", mr.Addr())

	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Fatalf("Error establishing connection %v", r)
			}
		}()

		client := newConnection()
		assert.False(t, client.Ping().Err() != nil)
	}()
}
