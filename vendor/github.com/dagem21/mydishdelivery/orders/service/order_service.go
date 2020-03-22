package service

import (
	"github.com/dagem21/mydishdelivery/orders"
	"github.com/dagem21/mydishdelivery/entity"
)

type OrderService struct {
	orderRepo orders.OrderRepository
}

func NewOrderService(ordRepo orders.OrderRepository) orders.OrderService {
	return &OrderService{orderRepo: ordRepo}
}

func (os *OrderService) Orders() ([]entity.Order, []error) {
	ords, errs := os.orderRepo.Orders()
	if len(errs) > 0 {
		return nil, errs
	}
	return ords, errs
}

func (os *OrderService) Order(id uint) (*entity.Order, []error) {
	ord, errs := os.orderRepo.Order(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return ord, errs
}

func (os *OrderService) UpdateOrder(order *entity.Order) (*entity.Order, []error) {
	ord, errs := os.orderRepo.UpdateOrder(order)
	if len(errs) > 0 {
		return nil, errs
	}
	return ord, errs
}

func (os *OrderService) DeleteOrder(id uint) (*entity.Order, []error) {
	ord, errs := os.orderRepo.DeleteOrder(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return ord, errs
}

func (os *OrderService) StoreOrder(order *entity.Order) (*entity.Order, []error) {
	ord, errs := os.orderRepo.StoreOrder(order)
	if len(errs) > 0 {
		return nil, errs
	}
	return ord, errs
}
