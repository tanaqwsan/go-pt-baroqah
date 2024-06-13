package routes

import (
	"app/controller"
	"app/middleware"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Init() *echo.Echo {

	e := echo.New()

	e.Use(middleware.NotFoundHandler)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to RESTful API Services test")
	})

	//Manage Employee
	e.POST("/employees", controller.StoreEmployee)
	e.GET("/employees", controller.IndexEmployee)
	e.GET("/employees/:id", controller.ShowEmployee)
	e.PUT("/employees/:id", controller.UpdateEmployee)
	e.DELETE("/employees/:id", controller.DeleteEmployee)

	//Manage Position
	e.POST("/positions", controller.StorePosition)
	e.GET("/positions", controller.IndexPosition)
	e.GET("/positions/:id", controller.ShowPosition)
	e.PUT("/positions/:level", controller.UpdatePositionByLevel)
	e.DELETE("/positions/:level", controller.DeletePositionByLevel)

	//Manage Salary
	e.POST("/salaries", controller.StoreSalary)
	e.GET("/salaries", controller.IndexSalary)
	e.GET("/salaries/:id", controller.ShowSalary)
	e.PUT("/salaries/:id", controller.UpdateSalary)
	e.DELETE("/salaries/:id", controller.DeleteSalary)
	e.GET("/salaries/:id/employee", controller.IndexSalaryByEmployee)
	e.GET("/salaries/:month/month", controller.IndexSalaryByMonth)
	e.GET("/salaries/:id/employee/:month/month", controller.IndexSalaryByEmployeeAndMonth)
	e.GET("/salaries/:id/employee/:count/count-month", controller.IndexSalaryByEmployeeXLatestMonth)

	return e

}
