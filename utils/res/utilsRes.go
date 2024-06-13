package res

import (
	"app/model"
	"app/model/web"
)

func ConvertIndexEmployee(employees []model.Employee) []web.GetEmployeeResponse {
	var results []web.GetEmployeeResponse
	for _, employee := range employees {
		employeeResponse := web.GetEmployeeResponse{
			Id:            int(employee.ID),
			Name:          employee.Name,
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
			BasicSalary: salary.BasicSalary,
			Bonus:       salary.Bonus,
			Fee:         salary.Fee,
			FinalSalary: salary.FinalSalary,
		}
		results = append(results, salaryResponse)
	}

	return results
}
