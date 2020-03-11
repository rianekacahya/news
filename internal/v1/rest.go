package v1

import (
	"github.com/rianekacahya/news/pkg/echoserver"

	"github.com/rianekacahya/news/internal/v1/news"
	nws "github.com/rianekacahya/news/internal/v1/news/transport"
)

func StartRest() {
	// run echo server
	echoserver.InitServer()
	go echoserver.StartServer()

	// Init Service
	nws.NewRest(news.Initialize())

	// Shutdown server gracefully
	echoserver.Shutdown()
}
