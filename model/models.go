package model

import (
	"gorm.io/gorm"
)

type Employee struct {
	Name          string `json:"name" form:"name"`
	NIP           string `json:"nip" form:"nip"`
	Address       string `json:"address" form:"address"`
	Position      int    `json:"position" form:"position"`
	BirthDate     int    `json:"birth_date" form:"birth_date"`
	FirstWorkDate int    `json:"first_work_date" form:"first_work_date"`
	*gorm.Model
}

// Position id = 1 => Manager
// id = 2 => Supervisor
// id = 3 => Staff
type Position struct {
	Level       int    `json:"level" form:"level"`
	Name        string `json:"name" form:"name"`
	BasicSalary int    `json:"basic_salary" form:"basic_salary"`
	Bonus       int    `json:"bonus" form:"bonus"`
	*gorm.Model
}

type Salary struct {
	EmployeeID  int `json:"employee_id" form:"employee_id"`
	Month       int `json:"month" form:"month"`
	Year        int `json:"year" form:"year"`
	BasicSalary int `json:"basic_salary" form:"basic_salary"`
	Bonus       int `json:"bonus" form:"bonus"`
	Fee         int `json:"fee" form:"fee"`
	FinalSalary int `json:"final_salary" form:"final_salary"`
	*gorm.Model
}
