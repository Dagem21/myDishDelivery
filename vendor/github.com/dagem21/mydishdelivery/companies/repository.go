package companies

import "github.com/dagem21/mydishdelivery/entity"

// CompanyRepository specifies customer comment related database operations
type CompanyRepository interface {
	Companies() ([]entity.Company, []error)
	Company(id uint) (*entity.Company, []error)
	UpdateCompany(company *entity.Company) (*entity.Company, []error)
	DeleteCompany(id uint) (*entity.Company, []error)
	StoreCompany(company *entity.Company) (*entity.Company, []error)
}
