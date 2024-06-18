package controller

import (
	"app/config"
	"app/model"
	"app/utils"
	"app/utils/res"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"
)

func IndexEmployee(c echo.Context) error {
	var employees []model.Employee

	err := config.DB.Find(&employees).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve employee"))
	}

	if len(employees) == 0 {
		return c.JSON(http.StatusNotFound, utils.ErrorResponse("Empty data"))
	}

	response := res.ConvertIndexEmployee(employees)

	return c.JSON(http.StatusOK, utils.SuccessResponse("Employee data successfully retrieved", response))
}

func ShowEmployee(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid ID"))
	}

	var employee model.Employee

	if err := config.DB.First(&employee, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve employee"))
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse("Employee data successfully retrieved", employee))
}

func StoreEmployee(c echo.Context) error {
	var employee model.Employee
	if err := c.Bind(&employee); err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body"))
	}

	//randomize the NIP
	employee.NIP = res.RandomNIP()

	if err := config.DB.Create(&employee).Error; err != nil {
		log.Fatal(err)
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to store employee data"))
	}

	return c.JSON(http.StatusCreated, utils.SuccessResponse("Success Created Data", employee))
}

func UpdateEmployee(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid ID"))
	}
	var updatedEmployee model.Employee

	if err := c.Bind(&updatedEmployee); err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body"))
	}

	var existingEmployee model.Employee
	result := config.DB.First(&existingEmployee, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve employee"))
	}
	config.DB.Model(&existingEmployee).Updates(updatedEmployee)

	return c.JSON(http.StatusOK, utils.SuccessResponse("Employee data successfully updated", nil))
}

func DeleteEmployee(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid ID"))
	}
	var existingEmployee model.Employee
	result := config.DB.First(&existingEmployee, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve employee"))
	}
	config.DB.Delete(&existingEmployee)

	return c.JSON(http.StatusOK, utils.SuccessResponse("Employee data successfully deleted", nil))
}
