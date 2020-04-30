package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"templates/go/models"
)

func GetIndex(c echo.Context) error {
	return c.JSON(http.StatusOK, &models.IndexResponse{
		Data:    "data",
		Version: os.Getenv("APP_VERSION"),
	})
}
