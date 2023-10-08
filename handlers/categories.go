package handlers

import (
    "github.com/labstack/echo"
	"strconv"
    "net/http"
	m "github.com/DmitryDeveloper/event-service/models"
)

func AllCategories(c echo.Context) error {
	var category m.Category
	categories, err := category.GetAll(50)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Cannot get categories"})
	}

    return c.JSON(http.StatusOK, categories)
}

func ShowCategory(c echo.Context) error {
	var category m.Category
	sid := c.Param("id")
	id, err := strconv.Atoi(sid)

	err = category.GetById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Cannot get category with id " + sid})
	}

	return c.JSON(http.StatusOK, category)
}

func CreateCategory(c echo.Context) error {
	category := new(m.Category)
	if err := c.Bind(category); err != nil {
        return err
    }

	res := category.Create()

	if res == false {
        return c.String(http.StatusInternalServerError, "Cannot create category")
    }

    return c.String(http.StatusOK, "Category created")
}
