package favorites

import "github.com/dagem21/mydishdelivery/entity"

type FavoriteService interface {
	Favorites() ([]entity.Favorite, []error)
	Favorite(id uint) (*entity.Favorite, []error)
	UpdateFavorite(favorite *entity.Favorite) (*entity.Favorite, []error)
	DeleteFavorite(id uint) (*entity.Favorite, []error)
	StoreFavorite(favorite *entity.Favorite) (*entity.Favorite, []error)
}
