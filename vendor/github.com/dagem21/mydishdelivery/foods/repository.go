package foods

import "github.com/dagem21/mydishdelivery/entity"

type FoodRepository interface {
	Foods() ([]entity.Food, []error)
	Food(id uint) (*entity.Food, []error)
	UpdateFood(food *entity.Food) (*entity.Food, []error)
	DeleteFood(id uint) (*entity.Food, []error)
	StoreFood(food *entity.Food) (*entity.Food, []error)
}
