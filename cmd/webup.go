package main

import (
	"fmt"
	"net/http"
	"os"
	"webup/internal/cms"
	"webup/internal/gdrive"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	"webup/internal/gdoc"
)

func main() {
	// ensure we can instantiate a HttpClient
	gdoc.ClientMustFromFile("cred.json")

	defaultFolder := "1GXeYQOvNvDvqhtSZW4mOhJCBgcG-d_6r"
	cmsId := "12CkaxfCn4RMs1gmt3tJxtYq-As0F22ssP6XndGhWDsY"

	e := echo.New()
	e.Logger.SetLevel(log.INFO)
	e.Logger.(*log.Logger).SetHeader("${time_rfc3339} ${level} ${short_file}:L${line} ${message}")
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} ${latency_human} ${status} ${method} ${uri} err=\"${error}\"\n",
	}))

	e.GET("/", func(c echo.Context) error {
		files, err := gdrive.List(defaultFolder)
		if err != nil {
			c.Logger().Error(err)
			_ = c.String(http.StatusBadGateway, "error")
			return err
		}

		resp := "<html><body><ul>"
		for _, file := range files {
			resp += fmt.Sprintf("<li><a href=\"%s/\">%s</a></li>", file.Id, file.Name)
		}
		resp += "</ul></body></html>"
		return c.HTML(http.StatusOK, resp)
	})

	e.GET("/:id/", func(c echo.Context) error {
		id := c.Param("id")
		raw, err := gdoc.Request(id)
		if err == nil {
			err = gdoc.CheckError(raw)
		}
		if err != nil {
			c.Logger().Error(err)
			_ = c.String(http.StatusBadGateway, "error")
			return err
		}
		result, _ := gdoc.Parse(raw)
		return c.HTML(http.StatusOK, result)
	})

	c := e.Group("/cms")
	c.GET("/menu/:lang", func(c echo.Context) error {
		lang := c.Param("lang")

		items, err := cms.GetMenu(cmsId)
		if err != nil {
			c.Logger().Error(err)
			_ = c.String(http.StatusBadGateway, "error")
			return err
		}

		return c.JSON(http.StatusOK, items[lang])
	})

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*://idea.cs.nthu.edu.tw"},
		//AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.Logger.Fatal(e.Start(os.Getenv("LISTEN")))
}
