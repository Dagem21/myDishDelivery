package repository

import (
	"github.com/dagem21/mydishdelivery/employees"
	"github.com/dagem21/mydishdelivery/entity"
	"github.com/jinzhu/gorm"
)

type EmployeeGormRepo struct {
	conn *gorm.DB
}

func NewEmployeeGormRepo(db *gorm.DB) employees.EmployeeRepository {
	return &EmployeeGormRepo{conn: db}
}

func (emplRepo *EmployeeGormRepo) Employees() ([]entity.Employee, []error) {
	empls := []entity.Employee{}
	errs := emplRepo.conn.Find(&empls).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return empls, errs
}

func (emplRepo *EmployeeGormRepo) Employee(id uint) (*entity.Employee, []error) {
	empl := entity.Employee{}
	errs := emplRepo.conn.First(&empl, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &empl, errs
}

func (emplRepo *EmployeeGormRepo) UpdateEmployee(employee *entity.Employee) (*entity.Employee, []error) {
	empl := employee
	errs := emplRepo.conn.Save(empl).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return empl, errs
}

func (emplRepo *EmployeeGormRepo) DeleteEmployee(id uint) (*entity.Employee, []error) {
	empl, errs := emplRepo.Employee(id)

	if len(errs) > 0 {
		return nil, errs
	}

	errs = emplRepo.conn.Delete(empl, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return empl, errs
}

func (emplRepo *EmployeeGormRepo) StoreEmployee(employee *entity.Employee) (*entity.Employee, []error) {
	empl := employee
	errs := emplRepo.conn.Create(empl).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return empl, errs
}
