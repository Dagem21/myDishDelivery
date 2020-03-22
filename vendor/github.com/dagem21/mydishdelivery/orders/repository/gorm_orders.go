package repository

import (
	"github.com/dagem21/mydishdelivery/orders"
	"github.com/dagem21/mydishdelivery/entity"
	"github.com/jinzhu/gorm"
)

type OrderGormRepo struct {
	conn *gorm.DB
}

func NewOrderGormRepo(db *gorm.DB) orders.OrderRepository {
	return &OrderGormRepo{conn: db}
}

func (ordRepo *OrderGormRepo) Orders() ([]entity.Order, []error) {
	ords := []entity.Order{}
	errs := ordRepo.conn.Find(&ords).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return ords, errs
}

func (ordRepo *OrderGormRepo) Order(id uint) (*entity.Order, []error) {
	ord := entity.Order{}
	errs := ordRepo.conn.First(&ord, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &ord, errs
}

func (ordRepo *OrderGormRepo) UpdateOrder(order *entity.Order) (*entity.Order, []error) {
	ord := order
	errs := ordRepo.conn.Save(ord).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return ord, errs
}

func (ordRepo *OrderGormRepo) DeleteOrder(id uint) (*entity.Order, []error) {
	ord, errs := ordRepo.Order(id)

	if len(errs) > 0 {
		return nil, errs
	}

	errs = ordRepo.conn.Delete(ord, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return ord, errs
}

func (ordRepo *OrderGormRepo) StoreOrder(order *entity.Order) (*entity.Order, []error) {
	ord := order
	errs := ordRepo.conn.Create(ord).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return ord, errs
}
