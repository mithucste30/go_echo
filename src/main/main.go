package main

import (
  "fmt"
  "net/http"
  "github.com/labstack/echo"
)

func main() {
  fmt.Println("welcome to the web server")

  e := echo.New()

  e.GET("/", func (c echo.Context) error  {
    return c.String(http.StatusOK, "yallo from the web side!")
  })

  e.Start(":8080")
}
