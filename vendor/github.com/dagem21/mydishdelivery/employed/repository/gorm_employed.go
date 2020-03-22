package repository

import (
	"github.com/dagem21/mydishdelivery/employed"
	"github.com/dagem21/mydishdelivery/entity"
	"github.com/jinzhu/gorm"
)

type EmployedGormRepo struct {
	conn *gorm.DB
}

func NewEmployedGormRepo(db *gorm.DB) employed.EmployedRepository {
	return &EmployedGormRepo{conn: db}
}

func (emdRepo *EmployedGormRepo) Employeds() ([]entity.Employed, []error) {
	empds := []entity.Employed{}
	errs := emdRepo.conn.Find(&empds).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return empds, errs
}

func (emdRepo *EmployedGormRepo) Employed(id uint) (*entity.Employed, []error) {
	empd := entity.Employed{}
	errs := emdRepo.conn.First(&empd, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &empd, errs
}

func (emdRepo *EmployedGormRepo) UpdateEmployed(employed *entity.Employed) (*entity.Employed, []error) {
	empd := employed
	errs := emdRepo.conn.Save(empd).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return empd, errs
}

func (emdRepo *EmployedGormRepo) DeleteEmployed(id uint) (*entity.Employed, []error) {
	empd, errs := emdRepo.Employed(id)

	if len(errs) > 0 {
		return nil, errs
	}

	errs = emdRepo.conn.Delete(empd, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return empd, errs
}

func (emdRepo *EmployedGormRepo) StoreEmployed(employed *entity.Employed) (*entity.Employed, []error) {
	empd := employed
	errs := emdRepo.conn.Create(empd).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return empd, errs
}
