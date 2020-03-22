package orders

import "github.com/dagem21/mydishdelivery/entity"

type OrderService interface {
	Orders() ([]entity.Order, []error)
	Order(id uint) (*entity.Order, []error)
	UpdateOrder(order *entity.Order) (*entity.Order, []error)
	DeleteOrder(id uint) (*entity.Order, []error)
	StoreOrder(order *entity.Order) (*entity.Order, []error)
}
