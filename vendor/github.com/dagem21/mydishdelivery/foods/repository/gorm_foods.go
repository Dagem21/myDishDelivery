package repository

import (
	"github.com/dagem21/mydishdelivery/foods"
	"github.com/dagem21/mydishdelivery/entity"
	"github.com/jinzhu/gorm"
)

type FoodGormRepo struct {
	conn *gorm.DB
}

func NewFoodGormRepo(db *gorm.DB) foods.FoodRepository {
	return &FoodGormRepo{conn: db}
}

func (fdRepo *FoodGormRepo) Foods() ([]entity.Food, []error) {
	fds := []entity.Food{}
	errs := fdRepo.conn.Find(&fds).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return fds, errs
}

func (fdRepo *FoodGormRepo) Food(id uint) (*entity.Food, []error) {
	fd := entity.Food{}
	errs := fdRepo.conn.First(&fd, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &fd, errs
}

func (fdRepo *FoodGormRepo) UpdateFood(food *entity.Food) (*entity.Food, []error) {
	fd := food
	errs := fdRepo.conn.Save(fd).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return fd, errs
}

func (fdRepo *FoodGormRepo) DeleteFood(id uint) (*entity.Food, []error) {
	fd, errs := fdRepo.Food(id)

	if len(errs) > 0 {
		return nil, errs
	}

	errs = fdRepo.conn.Delete(fd, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return fd, errs
}

func (fdRepo *FoodGormRepo) StoreFood(food *entity.Food) (*entity.Food, []error) {
	fd := food
	errs := fdRepo.conn.Create(fd).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return fd, errs
}
