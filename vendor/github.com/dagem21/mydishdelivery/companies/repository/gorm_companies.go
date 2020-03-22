package repository

import (
	"github.com/dagem21/mydishdelivery/companies"
	"github.com/dagem21/mydishdelivery/entity"
	"github.com/jinzhu/gorm"
)

type CompanyGormRepo struct {
	conn *gorm.DB
}

func NewCompanyGormRepo(db *gorm.DB) companies.CompanyRepository {
	return &CompanyGormRepo{conn: db}
}

func (cmpRepo *CompanyGormRepo) Companies() ([]entity.Company, []error) {
	cmps := []entity.Company{}
	errs := cmpRepo.conn.Find(&cmps).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return cmps, errs
}

func (cmpRepo *CompanyGormRepo) Company(id uint) (*entity.Company, []error) {
	cmp := entity.Company{}
	errs := cmpRepo.conn.First(&cmp, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &cmp, errs
}

func (cmpRepo *CompanyGormRepo) UpdateCompany(company *entity.Company) (*entity.Company, []error) {
	cmp := company
	errs := cmpRepo.conn.Save(cmp).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return cmp, errs
}

// DeleteComment deletes a given customer comment from the database
func (cmpRepo *CompanyGormRepo) DeleteCompany(id uint) (*entity.Company, []error) {
	cmp, errs := cmpRepo.Company(id)

	if len(errs) > 0 {
		return nil, errs
	}

	errs = cmpRepo.conn.Delete(cmp, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return cmp, errs
}

// StoreComment stores a given customer comment in the database
func (cmpRepo *CompanyGormRepo) StoreCompany(company *entity.Company) (*entity.Company, []error) {
	cmp := company
	errs := cmpRepo.conn.Create(cmp).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return cmp, errs
}
