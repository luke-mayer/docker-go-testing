package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// var port = os.Getenv("PORT")
var port = 1323

func helloHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func getUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func getQueries(c echo.Context) error {
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:"+team+", member:"+member)
}

func main() {
	e := echo.New()
	e.GET("/", helloHandler)
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
	}))
	e.GET("/users/:id", getUser)
	e.GET("/show", getQueries)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}
