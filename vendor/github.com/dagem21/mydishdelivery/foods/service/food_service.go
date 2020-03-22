package service

import (
	"github.com/dagem21/mydishdelivery/foods"
	"github.com/dagem21/mydishdelivery/entity"
)

type FoodService struct {
	foodRepo foods.FoodRepository
}

func NewFoodService(fdRepo foods.FoodRepository) foods.FoodService {
	return &FoodService{foodRepo: fdRepo}
}

func (fs *FoodService) Foods() ([]entity.Food, []error) {
	fds, errs := fs.foodRepo.Foods()
	if len(errs) > 0 {
		return nil, errs
	}
	return fds, errs
}

func (fs *FoodService) Food(id uint) (*entity.Food, []error) {
	fd, errs := fs.foodRepo.Food(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return fd, errs
}

func (fs *FoodService) UpdateFood(food *entity.Food) (*entity.Food, []error) {
	fd, errs := fs.foodRepo.UpdateFood(food)
	if len(errs) > 0 {
		return nil, errs
	}
	return fd, errs
}

func (fs *FoodService) DeleteFood(id uint) (*entity.Food, []error) {
	fd, errs := fs.foodRepo.DeleteFood(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return fd, errs
}

func (fs *FoodService) StoreFood(food *entity.Food) (*entity.Food, []error) {
	fd, errs := fs.foodRepo.StoreFood(food)
	if len(errs) > 0 {
		return nil, errs
	}
	return fd, errs
}
