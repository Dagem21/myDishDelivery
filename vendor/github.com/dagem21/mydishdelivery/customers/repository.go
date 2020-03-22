package customers

import "github.com/dagem21/mydishdelivery/entity"

type CustomerRepository interface {
	Customers() ([]entity.Customer, []error)
	Customer(id uint) (*entity.Customer, []error)
	UpdateCustomer(customer *entity.Customer) (*entity.Customer, []error)
	DeleteCustomer(id uint) (*entity.Customer, []error)
	StoreCustomer(customer *entity.Customer) (*entity.Customer, []error)
}
