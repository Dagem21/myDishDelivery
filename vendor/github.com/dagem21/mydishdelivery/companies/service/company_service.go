package service

import (
	"github.com/dagem21/mydishdelivery/companies"
	"github.com/dagem21/mydishdelivery/entity"
)

type CompanyService struct {
	companyRepo companies.CompanyRepository
}

func NewCompanyService(cmpRepo companies.CompanyRepository) companies.CompanyService {
	return &CompanyService{companyRepo: cmpRepo}
}

func (cs *CompanyService) Companies() ([]entity.Company, []error) {
	cmps, errs := cs.companyRepo.Companies()
	if len(errs) > 0 {
		return nil, errs
	}
	return cmps, errs
}

func (cs *CompanyService) Company(id uint) (*entity.Company, []error) {
	cmp, errs := cs.companyRepo.Company(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return cmp, errs
}

func (cs *CompanyService) UpdateCompany(company *entity.Company) (*entity.Company, []error) {
	cmp, errs := cs.companyRepo.UpdateCompany(company)
	if len(errs) > 0 {
		return nil, errs
	}
	return cmp, errs
}

func (cs *CompanyService) DeleteCompany(id uint) (*entity.Company, []error) {
	cmp, errs := cs.companyRepo.DeleteCompany(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return cmp, errs
}

func (cs *CompanyService) StoreCompany(company *entity.Company) (*entity.Company, []error) {
	cmp, errs := cs.companyRepo.StoreCompany(company)
	if len(errs) > 0 {
		return nil, errs
	}
	return cmp, errs
}
