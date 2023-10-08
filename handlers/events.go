package handlers

import (
	"fmt"
    "github.com/labstack/echo"
	"strconv"
    "net/http"
	m "github.com/DmitryDeveloper/event-service/models"
)

func AllEvents(c echo.Context) error {
	var e m.Event
	events, err := e.GetAll(100)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Cannot get events"})
	}

    return c.JSON(http.StatusOK, events)
}

func ShowEvent(c echo.Context) error {
	var event m.Event
	sid := c.Param("id")
	id, err := strconv.Atoi(sid)

	fmt.Println("Event ID = ", id)

	err = event.GetById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Cannot get event with id " + sid})
	}

	return c.JSON(http.StatusOK, event)
}

func GetEventsForCategory(c echo.Context) error {
	cid := c.Param("id")
	categoryId, err := strconv.Atoi(cid)

	var e m.Event
	events, err := e.GetByCategoryId(categoryId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Cannot get events"})
	}

    return c.JSON(http.StatusOK, events)
}

func CreateEvent(c echo.Context) error {
	data := make(map[string]interface{})

    // Пробуем прочитать и распарсить JSON-данные из тела запроса
    if err := c.Bind(&data); err != nil {
        return err
    }

	categoryIDs, ok := data["categories"].([]interface{})
	if !ok {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid categories format"})
    }

	// Проверяем существование категорий
	var categories []m.Category
	err := m.GetDB().Where("id IN (?)", categoryIDs).Find(&categories).Error
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid category IDs"})
	}

	data["categories"] = categories

	e := m.Event{
		Title:            data["title"].(string),
		ShortDescription: data["short_description"].(string),
		Description:      data["description"].(string),
		UserId:           int(data["user_id"].(float64)),
		IsApproved:       false,
		Categories:       data["categories"].([]m.Category),
	}

	res := e.Create()

	if res == false {
        return c.String(http.StatusInternalServerError, "Cannot create event")
    }

    return c.String(http.StatusOK, "Event Created")
}
