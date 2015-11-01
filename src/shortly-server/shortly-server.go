package main

import (
  "net/http"
  "github.com/labstack/echo"
  mw "github.com/labstack/echo/middleware"
)

func ping(c *echo.Context) error {
  return c.String(http.StatusOK, "Ok")
}

func main() {
  e := echo.New()

  e.Use(mw.Logger())
  e.Use(mw.Recover())

  e.Get("/ping", ping)

  e.Run(":1323")
}
