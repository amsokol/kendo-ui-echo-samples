package main

import (
	"github.com/amsokol/kendo-ui-echo-samples/Godeps/_workspace/src/github.com/labstack/echo"
	mw "github.com/amsokol/kendo-ui-echo-samples/Godeps/_workspace/src/github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(mw.Logger())
	e.Use(mw.Recover())
	e.Use(mw.Gzip())

	e.Static("/", "public")

	e.Run(":5091")
}
