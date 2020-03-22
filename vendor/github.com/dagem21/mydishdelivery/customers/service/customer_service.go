package service

import (
	"github.com/dagem21/mydishdelivery/customers"
	"github.com/dagem21/mydishdelivery/entity"
)

type CustomerService struct {
	customerRepo customers.CustomerRepository
}

func NewCustomerService(custRepo customers.CustomerRepository) customers.CustomerService {
	return &CustomerService{customerRepo: custRepo}
}

func (cs *CustomerService) Customers() ([]entity.Customer, []error) {
	custs, errs := cs.customerRepo.Customers()
	if len(errs) > 0 {
		return nil, errs
	}
	return custs, errs
}

func (cs *CustomerService) Customer(id uint) (*entity.Customer, []error) {
	cust, errs := cs.customerRepo.Customer(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return cust, errs
}

func (cs *CustomerService) UpdateCustomer(customer *entity.Customer) (*entity.Customer, []error) {
	cust, errs := cs.customerRepo.UpdateCustomer(customer)
	if len(errs) > 0 {
		return nil, errs
	}
	return cust, errs
}

func (cs *CustomerService) DeleteCustomer(id uint) (*entity.Customer, []error) {
	cust, errs := cs.customerRepo.DeleteCustomer(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return cust, errs
}

func (cs *CustomerService) StoreCustomer(customer *entity.Customer) (*entity.Customer, []error) {
	cust, errs := cs.customerRepo.StoreCustomer(customer)
	if len(errs) > 0 {
		return nil, errs
	}
	return cust, errs
}
