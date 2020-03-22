package repository

import (
	"github.com/dagem21/mydishdelivery/customers"
	"github.com/dagem21/mydishdelivery/entity"
	"github.com/jinzhu/gorm"
)

type CustomerGormRepo struct {
	conn *gorm.DB
}

func NewCustomerGormRepo(db *gorm.DB) customers.CustomerRepository {
	return &CustomerGormRepo{conn: db}
}

func (custRepo *CustomerGormRepo) Customers() ([]entity.Customer, []error) {
	custs := []entity.Customer{}
	errs := custRepo.conn.Find(&custs).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return custs, errs
}

func (custRepo *CustomerGormRepo) Customer(id uint) (*entity.Customer, []error) {
	cust := entity.Customer{}
	errs := custRepo.conn.First(&cust, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &cust, errs
}

func (custRepo *CustomerGormRepo) UpdateCustomer(customer *entity.Customer) (*entity.Customer, []error) {
	cust := customer
	errs := custRepo.conn.Save(cust).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return cust, errs
}

func (custRepo *CustomerGormRepo) DeleteCustomer(id uint) (*entity.Customer, []error) {
	cust, errs := custRepo.Customer(id)

	if len(errs) > 0 {
		return nil, errs
	}

	errs = custRepo.conn.Delete(cust, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return cust, errs
}

func (custRepo *CustomerGormRepo) StoreCustomer(customer *entity.Customer) (*entity.Customer, []error) {
	cust := customer
	errs := custRepo.conn.Create(cust).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return cust, errs
}
