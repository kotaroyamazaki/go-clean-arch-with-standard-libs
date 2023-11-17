package repository

import (
	"context"

	"github.com/kotaroyamazaki/go-clean-arch-sample-with-standard/internal/entities/repository"
	"github.com/kotaroyamazaki/go-clean-arch-sample-with-standard/pkg/orm"
)

type userFavoriteBooksRepository struct {
	m map[int]orm.UserFavoriteBookSlice
}

func NewUserFavoriteBooksRepository() repository.UserFavoriteBooksRepository {
	var m = map[int]orm.UserFavoriteBookSlice{}
	return &userFavoriteBooksRepository{
		m: m,
	}
}

func (r *userFavoriteBooksRepository) Save(ctx context.Context, favotieBookIDs []int) error {
	if len(favotieBookIDs) == 0 {
		return nil
	}
	// TODO: save
	return nil
}
