package controllers

import (
	"net/http"
	"prakerja/config"
	"prakerja/models"

	"github.com/labstack/echo/v4"
)

func AddBookController(c echo.Context) error {
	var bookRequest models.Book

	c.Bind(&bookRequest)

	result := config.DB.Create(&bookRequest)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status:  false,
			Message: "Failed insert data to database",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusCreated, models.BaseResponse{
		Status:  true,
		Message: "Success insert data to database",
		Data:    bookRequest,
	})
}

func GetDetailBookController(c echo.Context) error {
	var id = c.Param("id")
	var data = models.Book{}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Status:  true,
		Message: id,
		Data:    data,
	})
}

func GetBookController(c echo.Context) error {
	var data []models.Book

	result := config.DB.Find(&data)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status:  false,
			Message: "Failed get data from book",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Status:  true,
		Message: "Success get Data",
		Data:    data,
	})
}
