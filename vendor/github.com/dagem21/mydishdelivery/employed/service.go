package employed

import "github.com/dagem21/mydishdelivery/entity"

type EmployedService interface {
	Employeds() ([]entity.Employed, []error)
	Employed(id uint) (*entity.Employed, []error)
	UpdateEmployed(employed *entity.Employed) (*entity.Employed, []error)
	DeleteEmployed(id uint) (*entity.Employed, []error)
	StoreEmployed(employed *entity.Employed) (*entity.Employed, []error)
}
