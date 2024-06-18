package res

import (
	"app/model"
	"app/model/web"
	"math/rand"
	"strconv"
)

func ConvertIndexEmployee(employees []model.Employee) []web.GetEmployeeResponse {
	var results []web.GetEmployeeResponse
	for _, employee := range employees {
		employeeResponse := web.GetEmployeeResponse{
			Id:            int(employee.ID),
			Name:          employee.Name,
			NIP:           employee.NIP,
			Address:       employee.Address,
			Position:      employee.Position,
			BirthDate:     employee.BirthDate,
			FirstWorkDate: employee.FirstWorkDate,
		}
		results = append(results, employeeResponse)
	}

	return results
}

func ConvertIndexPosition(positions []model.Position) []web.GetPositionResponse {
	var results []web.GetPositionResponse
	for _, position := range positions {
		positionResponse := web.GetPositionResponse{
			Level:       position.Level,
			Name:        position.Name,
			BasicSalary: position.BasicSalary,
			Bonus:       position.Bonus,
			Id:          int(position.ID),
		}
		results = append(results, positionResponse)
	}

	return results
}

func ConvertIndexSalary(salaries []model.Salary) []web.GetSalaryResponse {
	var results []web.GetSalaryResponse
	for _, salary := range salaries {
		salaryResponse := web.GetSalaryResponse{
			Id:          int(salary.ID),
			EmployeeID:  salary.EmployeeID,
			Month:       salary.Month,
			Year:        salary.Year,
			BasicSalary: salary.BasicSalary,
			Bonus:       salary.Bonus,
			Fee:         salary.Fee,
			FinalSalary: salary.FinalSalary,
		}
		results = append(results, salaryResponse)
	}

	return results
}

func ConvertIndexSalarySortByMonth(salaries []model.Salary) []web.GetSalaryResponse {
	var results []web.GetSalaryResponse
	for _, salary := range salaries {
		salaryResponse := web.GetSalaryResponse{
			Id:          int(salary.ID),
			EmployeeID:  salary.EmployeeID,
			Month:       salary.Month,
			Year:        salary.Year,
			BasicSalary: salary.BasicSalary,
			Bonus:       salary.Bonus,
			Fee:         salary.Fee,
			FinalSalary: salary.FinalSalary,
		}
		results = append(results, salaryResponse)
	}

	//Sort by month
	for i := 0; i < len(results); i++ {
		for j := i + 1; j < len(results); j++ {
			if results[i].Month > results[j].Month {
				temp := results[i]
				results[i] = results[j]
				results[j] = temp
			}
		}
	}

	//Sort by year
	for i := 0; i < len(results); i++ {
		for j := i + 1; j < len(results); j++ {
			if results[i].Year > results[j].Year {
				temp := results[i]
				results[i] = results[j]
				results[j] = temp
			}
		}
	}

	return results
}

func RandomNIP() string {
	//return 10 random number
	return strconv.Itoa(rand.Intn(1000000000))
}
