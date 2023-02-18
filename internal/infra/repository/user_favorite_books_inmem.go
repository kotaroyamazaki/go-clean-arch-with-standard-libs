package repository

import (
	"context"
	"log"

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

func (r *userFavoriteBooksRepository) Save(ctx context.Context, favotieBookSlice orm.UserFavoriteBookSlice) error {
	if len(favotieBookSlice) == 0 {
		return nil
	}
	r.m[favotieBookSlice[0].UserID] = favotieBookSlice
	log.Printf("save %v\n", favotieBookSlice)
	return nil
}
