package main

import (
	"mainpackage/printpackage"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	printpackage.Print()
	e := echo.New()

	e.GET("/users/:id", getUser)
	e.Logger.Fatal(e.Start(":1323"))
}
func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
