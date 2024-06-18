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

func IndexSalary(c echo.Context) error {
	var salaries []model.Salary

	err := config.DB.Find(&salaries).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve salary"))
	}

	if len(salaries) == 0 {
		return c.JSON(http.StatusNotFound, utils.ErrorResponse("Empty data"))
	}

	response := res.ConvertIndexSalary(salaries)

	return c.JSON(http.StatusOK, utils.SuccessResponse("Salary data successfully retrieved", response))
}

func ShowSalary(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid ID"))
	}

	var salary model.Salary

	if err := config.DB.First(&salary, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve salary"))
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse("Salary data successfully retrieved", salary))
}

// StoreSalary Perhitungan gaji karyawan dilakukan
// berdasarkan gaji pokok, bonus dan PPH 5%. Bonus di dapat berdasarkan jabatan
// karyawan, dengan ketentuan Jabatan Manager 50% dari gaji pokok, Supervisor 40%
// dari gaji pokok, dan jabatan staff 30% dari gaji pokok./*
func StoreSalary(c echo.Context) error {
	var salary model.Salary

	employeeId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid ID"))
	}

	month, err := strconv.Atoi(c.Param("month"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid ID"))
	}

	year, err := strconv.Atoi(c.Param("year"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid ID"))
	}

	salary.EmployeeID = employeeId
	salary.Month = month
	salary.Year = year

	// Check if month is valid
	if salary.Month < 1 || salary.Month > 12 {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid month"))
	}

	// Check if salary of the employee in the same month already exists
	var existingSalary model.Salary
	if err := config.DB.Where("employee_id = ? AND month = ? AND year = ?", salary.EmployeeID, salary.Month, salary.Year).First(&existingSalary).Error; err == nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Salary data for the employee in the same month already exists"))
	}

	// Get employee data
	var employee model.Employee
	if err := config.DB.First(&employee, salary.EmployeeID).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve employee"))
	}

	// Get position data by level
	var position model.Position
	if err := config.DB.Where("level = ?", employee.Position).First(&position).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve position"))
	}

	// Calculate salary
	salary.BasicSalary = position.BasicSalary
	salary.Bonus = position.Bonus * position.BasicSalary / 100
	log.Println("position.BasicSalary : " + strconv.Itoa(position.BasicSalary))
	log.Println("position.Bonus : " + strconv.Itoa(position.Bonus))
	salary.Fee = 5 * (salary.BasicSalary + salary.Bonus) / 100
	salary.FinalSalary = salary.BasicSalary + salary.Bonus - salary.Fee
	if err := config.DB.Create(&salary).Error; err != nil {
		log.Fatal(err)
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to store salary data"))
	}

	return c.JSON(http.StatusCreated, utils.SuccessResponse("Success Created Data", salary))

}

func UpdateSalary(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid ID"))
	}

	var salary model.Salary
	if err := config.DB.First(&salary, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve salary"))
	}

	if err := c.Bind(&salary); err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body"))
	}

	if err := config.DB.Save(&salary).Error; err != nil {
		log.Fatal(err)
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to update salary data"))
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse("Success Update Data", salary))
}

func DeleteSalary(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid ID"))
	}

	var salary model.Salary
	if err := config.DB.First(&salary, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve salary"))
	}

	if err := config.DB.Delete(&salary).Error; err != nil {
		log.Fatal(err)
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to delete salary data"))
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse("Success Delete Data", nil))
}

func IndexSalaryByEmployee(c echo.Context) error {
	employeeId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid ID"))
	}

	var salaries []model.Salary

	err = config.DB.Where("employee_id = ?", employeeId).Find(&salaries).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve salary"))
	}

	if len(salaries) == 0 {
		return c.JSON(http.StatusNotFound, utils.ErrorResponse("Empty data"))
	}

	response := res.ConvertIndexSalarySortByMonth(salaries)

	return c.JSON(http.StatusOK, utils.SuccessResponse("Salary data successfully retrieved", response))

}

func IndexSalaryByMonth(c echo.Context) error {
	month := c.Param("month")

	var salaries []model.Salary

	err := config.DB.Where("month = ?", month).Find(&salaries).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve salary"))
	}

	if len(salaries) == 0 {
		return c.JSON(http.StatusNotFound, utils.ErrorResponse("Empty data"))
	}

	response := res.ConvertIndexSalary(salaries)

	return c.JSON(http.StatusOK, utils.SuccessResponse("Salary data successfully retrieved", response))
}

func IndexSalaryByEmployeeAndMonth(c echo.Context) error {
	employeeId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid ID"))
	}
	month := c.Param("month")

	var salaries []model.Salary

	err = config.DB.Where("employee_id = ? AND month = ?", employeeId, month).Find(&salaries).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve salary"))
	}

	if len(salaries) == 0 {
		return c.JSON(http.StatusNotFound, utils.ErrorResponse("Empty data"))
	}

	response := res.ConvertIndexSalary(salaries)

	return c.JSON(http.StatusOK, utils.SuccessResponse("Salary data successfully retrieved", response))
}

func IndexSalaryByEmployeeXLatestMonth(c echo.Context) error {
	employeeId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid ID"))
	}
	count, err := strconv.Atoi(c.Param("count"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid count"))
	}

	var salaries []model.Salary

	err = config.DB.Where("employee_id = ?", employeeId).Order("month desc").Limit(count).Find(&salaries).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve salary"))
	}

	if len(salaries) == 0 {
		return c.JSON(http.StatusNotFound, utils.ErrorResponse("Empty data"))
	}

	response := res.ConvertIndexSalary(salaries)

	return c.JSON(http.StatusOK, utils.SuccessResponse("Salary data successfully retrieved", response))
}
