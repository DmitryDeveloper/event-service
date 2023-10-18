package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	m "github.com/DmitryDeveloper/event-service/models"
	"github.com/labstack/echo"
)

// Should I use RequestData or Binding?
type RequestData struct {
	UserID int `json:"user_id"`
}

func Subscribe(c echo.Context) error {
	var requestData RequestData

	err := json.NewDecoder(c.Request().Body).Decode(&requestData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	var s m.Subscription
	s.EventId, _ = strconv.Atoi(c.Param("id"))
	s.UserId = requestData.UserID

	fmt.Println(s.String())

	if res := s.Create(); !res {
		return c.String(http.StatusInternalServerError, "Cannot subscribe")
	}

	return c.String(http.StatusOK, "User subscribed")
}

func Unsubscribe(c echo.Context) error {
	var requestData RequestData

	err := json.NewDecoder(c.Request().Body).Decode(&requestData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	userID := requestData.UserID
	id := c.Param("id")
	eid, _ := strconv.Atoi(id)

	if m.DeleteByEventIdAndUserId(eid, userID) {
		return c.String(http.StatusOK, "User unsubscribed")
	} else {
		return c.String(http.StatusInternalServerError, "Failed to unsubscribe user")
	}
}
