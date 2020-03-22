package employees

import "github.com/dagem21/mydishdelivery/entity"

type EmployeeService interface {
	Employees() ([]entity.Employee, []error)
	Employee(id uint) (*entity.Employee, []error)
	UpdateEmployee(employee *entity.Employee) (*entity.Employee, []error)
	DeleteEmployee(id uint) (*entity.Employee, []error)
	StoreEmployee(employee *entity.Employee) (*entity.Employee, []error)
}
