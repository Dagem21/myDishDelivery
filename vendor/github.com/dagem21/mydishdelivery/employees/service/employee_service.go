package service

import (
	"github.com/dagem21/mydishdelivery/employees"
	"github.com/dagem21/mydishdelivery/entity"
)

type EmployeeService struct {
	employeeRepo employees.EmployeeRepository
}

func NewEmployeeService(emplRepo employees.EmployeeRepository) employees.EmployeeService {
	return &EmployeeService{employeeRepo: emplRepo}
}

func (es *EmployeeService) Employees() ([]entity.Employee, []error) {
	empls, errs := es.employeeRepo.Employees()
	if len(errs) > 0 {
		return nil, errs
	}
	return empls, errs
}

func (es *EmployeeService) Employee(id uint) (*entity.Employee, []error) {
	empl, errs := es.employeeRepo.Employee(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return empl, errs
}

func (es *EmployeeService) UpdateEmployee(employee *entity.Employee) (*entity.Employee, []error) {
	empl, errs := es.employeeRepo.UpdateEmployee(employee)
	if len(errs) > 0 {
		return nil, errs
	}
	return empl, errs
}

func (es *EmployeeService) DeleteEmployee(id uint) (*entity.Employee, []error) {
	empl, errs := es.employeeRepo.DeleteEmployee(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return empl, errs
}

func (es *EmployeeService) StoreEmployee(employee *entity.Employee) (*entity.Employee, []error) {
	empl, errs := es.employeeRepo.StoreEmployee(employee)
	if len(errs) > 0 {
		return nil, errs
	}
	return empl, errs
}
