package v1

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/rianekacahya/news/internal/v1/news"
	nws "github.com/rianekacahya/news/internal/v1/news/transport"
)

func StartEvent() {
	// Init Service
	nws.NewEvent(news.Initialize())

	fmt.Println("Event server has been started")

	// Quit with graceful shutdown
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
}
