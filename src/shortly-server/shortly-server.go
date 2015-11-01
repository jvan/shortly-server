package main

import (
  "net/http"
  "github.com/labstack/echo"
  mw "github.com/labstack/echo/middleware"

  "strconv"

  "shortly"
)

func ping(c *echo.Context) error {
  return c.String(http.StatusOK, "Ok")
}

func encode_value(c *echo.Context) error {
  // Convert the `val` parameter to an integer.
  value := c.Param("val")
  i, _ := strconv.Atoi(value)

  // Return the base-62 encoded value.
  return c.String(http.StatusOK, shortly.Encode(i))
}

func decode_value(c *echo.Context) error {
  // Decode the base-62 `val` parameter.
  value := c.Param("val")
  i := shortly.Decode(value)

  // Convert the base-10 integer to a string and return.
  s := strconv.Itoa(i)
  return c.String(http.StatusOK, s)
}

func main() {
  e := echo.New()

  e.Use(mw.Logger())
  e.Use(mw.Recover())

  e.Get("/ping", ping)

  e.Get("/encode/:val", encode_value)
  e.Get("/decode/:val", decode_value)

  e.Run(":1323")
}
