package service

import (
	"github.com/dagem21/mydishdelivery/employed"
	"github.com/dagem21/mydishdelivery/entity"
)

type EmployedService struct {
	employedRepo employed.EmployedRepository
}

func NewEmployedService(empdRepo employed.EmployedRepository) employed.EmployedService {
	return &EmployedService{employedRepo: empdRepo}
}

func (es *EmployedService) Employeds() ([]entity.Employed, []error) {
	empds, errs := es.employedRepo.Employeds()
	if len(errs) > 0 {
		return nil, errs
	}
	return empds, errs
}

func (es *EmployedService) Employed(id uint) (*entity.Employed, []error) {
	empd, errs := es.employedRepo.Employed(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return empd, errs
}

func (es *EmployedService) UpdateEmployed(employed *entity.Employed) (*entity.Employed, []error) {
	empd, errs := es.employedRepo.UpdateEmployed(employed)
	if len(errs) > 0 {
		return nil, errs
	}
	return empd, errs
}

func (es *EmployedService) DeleteEmployed(id uint) (*entity.Employed, []error) {
	empd, errs := es.employedRepo.DeleteEmployed(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return empd, errs
}

func (es *EmployedService) StoreEmployed(employed *entity.Employed) (*entity.Employed, []error) {
	empd, errs := es.employedRepo.StoreEmployed(employed)
	if len(errs) > 0 {
		return nil, errs
	}
	return empd, errs
}
