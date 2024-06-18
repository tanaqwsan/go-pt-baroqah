package controller

import (
	"app/config"
	"app/model"
	"app/utils"
	"app/utils/res"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func IndexPosition(c echo.Context) error {
	var positions []model.Position

	err := config.DB.Find(&positions).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve position"))
	}

	if len(positions) == 0 {
		return c.JSON(http.StatusNotFound, utils.ErrorResponse("Empty data"))
	}

	response := res.ConvertIndexPosition(positions)

	return c.JSON(http.StatusOK, utils.SuccessResponse("Position data successfully retrieved", response))
}

func ShowPosition(c echo.Context) error {
	//Get position by level
	level := c.Param("level")

	var position model.Position

	if err := config.DB.Where("level = ?", level).First(&position).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve position"))
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse("Position data successfully retrieved", position))
}

func StorePosition(c echo.Context) error {
	var position model.Position
	if err := c.Bind(&position); err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body"))
	}

	// Check if position already exists
	var checkPosition model.Position
	if err := config.DB.Where("level = ?", position.Level).First(&checkPosition).Error; err == nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Position already exists"))
	}

	if err := config.DB.Create(&position).Error; err != nil {
		log.Fatal(err)
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to store position data"))
	}

	return c.JSON(http.StatusCreated, utils.SuccessResponse("Success Created Data", position))
}

func UpdatePositionByLevel(c echo.Context) error {
	level := c.Param("level")
	var position model.Position

	if err := config.DB.Where("level = ?", level).First(&position).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve position"))
	}

	if err := c.Bind(&position); err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body"))
	}

	if err := config.DB.Save(&position).Error; err != nil {
		log.Fatal(err)
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to update position data"))
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse("Success Update Data", position))
}

func DeletePositionByLevel(c echo.Context) error {
	level := c.Param("level")
	var position model.Position

	if err := config.DB.Where("level = ?", level).First(&position).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve position"))
	}

	if err := config.DB.Delete(&position).Error; err != nil {
		log.Fatal(err)
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to delete position data"))
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse("Success Delete Data", nil))
}
