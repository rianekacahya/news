package echoserver

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/rianekacahya/news/pkg/config"
	custom_middleware "github.com/rianekacahya/news/pkg/echoserver/middleware"
)

var (
	server *echo.Echo
	mutex  sync.Once
)

func GetServer() *echo.Echo {
	mutex.Do(func() {
		server = newServer()
	})

	return server
}

func newServer() *echo.Echo {
	return echo.New()
}

func InitServer() {

	// Hide banner
	GetServer().HideBanner = true

	// Set debug status parameter
	GetServer().Debug = config.GetConfig().AppDebug

	// init default middleware
	GetServer().Use(
		middleware.RequestID(),
		middleware.RecoverWithConfig(middleware.RecoverConfig{
			DisableStackAll:   true,
			DisablePrintStack: true,
		}),
		custom_middleware.CORS(),
		custom_middleware.Headers(),
		middleware.BodyDump(custom_middleware.Logger),
	)

	// healthCheck endpoint
	GetServer().GET("/infrastructure/healthcheck", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "OK")
	})

	// custom error handler
	GetServer().HTTPErrorHandler = Handler

	// Custom binder
	GetServer().Binder = &CustomBinder{bind: &echo.DefaultBinder{}}
}

func StartServer() {
	if err := GetServer().StartServer(&http.Server{
		Addr:         fmt.Sprintf(":%v", config.GetConfig().ServerRestPort),
		ReadTimeout:  time.Duration(config.GetConfig().ServerRestReadTimeout) * time.Second,
		WriteTimeout: time.Duration(config.GetConfig().ServerRestWriteTimeout) * time.Second,
		IdleTimeout:  time.Duration(config.GetConfig().ServerRestIdleTimeout) * time.Second,
	}); err != nil {
		log.Fatal(err.Error())
	}
}

func Shutdown() {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := GetServer().Shutdown(ctx); err != nil {
		log.Fatal(err.Error())
	}
}
