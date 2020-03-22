package service

import (
	"github.com/dagem21/mydishdelivery/favorites"
	"github.com/dagem21/mydishdelivery/entity"
)

type FavoriteService struct {
	favoriteRepo favorites.FavoriteRepository
}

func NewFavoriteService(fvrtRepo favorites.FavoriteRepository) favorites.FavoriteService {
	return &FavoriteService{favoriteRepo: fvrtRepo}
}

func (fs *FavoriteService) Favorites() ([]entity.Favorite, []error) {
	fvrts, errs := fs.favoriteRepo.Favorites()
	if len(errs) > 0 {
		return nil, errs
	}
	return fvrts, errs
}

func (fs *FavoriteService) Favorite(id uint) (*entity.Favorite, []error) {
	fvrt, errs := fs.favoriteRepo.Favorite(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return fvrt, errs
}

func (fs *FavoriteService) UpdateFavorite(favorite *entity.Favorite) (*entity.Favorite, []error) {
	fvrt, errs := fs.favoriteRepo.UpdateFavorite(favorite)
	if len(errs) > 0 {
		return nil, errs
	}
	return fvrt, errs
}

func (fs *FavoriteService) DeleteFavorite(id uint) (*entity.Favorite, []error) {
	fvrt, errs := fs.favoriteRepo.DeleteFavorite(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return fvrt, errs
}

func (fs *FavoriteService) StoreFavorite(favorite *entity.Favorite) (*entity.Favorite, []error) {
	fvrt, errs := fs.favoriteRepo.StoreFavorite(favorite)
	if len(errs) > 0 {
		return nil, errs
	}
	return fvrt, errs
}
