
package main

import (
    "net/http"

    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
)

func main() {
    e := echo.New()

    e.Use(middleware.Recover())
    e.Use(middleware.Logger())
    e.Use(middleware.Gzip())

    e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
        AllowOrigins: []string{os.Getenv("CORS_ALLOW_ORIGIN")},
        AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
    }))

    e.GET("/health", func(c echo.Context) error {
        return c.NoContent(http.StatusOK)
    })

    e.HideBanner = true
    e.Logger.Fatal(e.Start(":4000"))
}


