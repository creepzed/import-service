package rest

import (
	"bitbucket.org/ripleyx/import-service/app/shared/infrastructure/log"
	"fmt"
	"github.com/labstack/echo"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"net/http"
	"time"
)

const (
	httpReadTimeout  = 3 * time.Minute
	httpWriteTimeout = 3 * time.Minute
)

func New() *echo.Echo {

	echo := echo.New()

	echo.Use(log.EchoLogger())
	echo.Use(echoMiddleware.Logger())
	echo.Use(echoMiddleware.Recover())
	echo.Use(echoMiddleware.CORS())

	echo.Validator = NewValidator()

	echo.HideBanner = true

	NewHealthHandler(echo)
	//echo.GET("/swagger/*", echoSwagger.WrapHandler)
	return echo
}

func Setup(host string, port string) *http.Server {
	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%s", host, port),
		ReadTimeout:  httpReadTimeout,
		WriteTimeout: httpWriteTimeout,
	}
	return server

}
