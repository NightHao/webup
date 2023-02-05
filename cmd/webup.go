package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	"webup/internal/gdoc"
)

func main() {
	// ensure we can instantiate a HttpClient
	gdoc.ClientMustFromFile("cred.json")

	e := echo.New()
	e.Logger.SetLevel(log.INFO)
	e.Logger.(*log.Logger).SetHeader("${time_rfc3339} ${level} ${short_file}:L${line} ${message}")
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} ${latency_human} ${status} ${method} ${uri} err=\"${error}\"\n",
	}))

	e.GET("/:id/", func(c echo.Context) error {
		id := c.Param("id")
		raw, err := gdoc.Request(id)
		if err != nil {
			c.Logger().Error(err)
			_ = c.String(http.StatusBadGateway, "error")
			return err
		}
		return c.HTML(http.StatusOK, gdoc.Parse(raw))
	})

	e.Logger.Fatal(e.Start(os.Getenv("LISTEN")))
}
