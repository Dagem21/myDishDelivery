package repository

import (
	"github.com/dagem21/mydishdelivery/favorites"
	"github.com/dagem21/mydishdelivery/entity"
	"github.com/jinzhu/gorm"
)

type FavoriteGormRepo struct {
	conn *gorm.DB
}

func NewFavoriteGormRepo(db *gorm.DB) favorites.FavoriteRepository {
	return &FavoriteGormRepo{conn: db}
}

func (fvrtRepo *FavoriteGormRepo) Favorites() ([]entity.Favorite, []error) {
	fvrt := []entity.Favorite{}
	errs := fvrtRepo.conn.Find(&fvrt).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return fvrt, errs
}

func (fvrtRepo *FavoriteGormRepo) Favorite(id uint) (*entity.Favorite, []error) {
	fvrt := entity.Favorite{}
	errs := fvrtRepo.conn.First(&fvrt, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &fvrt, errs
}

func (fvrtRepo *FavoriteGormRepo) UpdateFavorite(favorite *entity.Favorite) (*entity.Favorite, []error) {
	fvrt := favorite
	errs := fvrtRepo.conn.Save(fvrt).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return fvrt, errs
}

func (fvrtRepo *FavoriteGormRepo) DeleteFavorite(id uint) (*entity.Favorite, []error) {
	fvrt, errs := fvrtRepo.Favorite(id)

	if len(errs) > 0 {
		return nil, errs
	}

	errs = fvrtRepo.conn.Delete(fvrt, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return fvrt, errs
}

func (fvrtRepo *FavoriteGormRepo) StoreFavorite(favorite *entity.Favorite) (*entity.Favorite, []error) {
	fvrt := favorite
	errs := fvrtRepo.conn.Create(fvrt).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return fvrt, errs
}
