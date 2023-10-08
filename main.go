package main

import (
    "fmt"
    "net/http"
	"github.com/labstack/echo"
    h "github.com/DmitryDeveloper/event-service/handlers"
)

func main() {
    e := echo.New()

    e.GET("/health", func (c echo.Context) error {
        return c.String(http.StatusOK, `OK`)
    })

    e.GET("/events/:id", h.ShowEvent)
    e.GET("/events", h.AllEvents)
    e.POST("/events", h.CreateEvent)
    e.GET("/categories/:id/events", h.GetEventsForCategory)

    e.GET("/categories/:id", h.ShowCategory)
    e.GET("/categories", h.AllCategories)
    e.POST("/categories", h.CreateCategory)

    e.Logger.Fatal(e.Start(":8080"))

    fmt.Println("Server started!")
}
