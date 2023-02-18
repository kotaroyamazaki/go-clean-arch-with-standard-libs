package repository

import (
	"context"
	"log"

	"github.com/kauche-interview/backend-KotaroYamazaki/src/internal/entities"
	"github.com/kauche-interview/backend-KotaroYamazaki/src/internal/entities/repository"
	"github.com/kauche-interview/backend-KotaroYamazaki/src/pkg/orm"
)

type userFavoriteProductsRepository struct {
	m map[int]orm.UserFavoriteProductSlice
}

func NewUserFavoriteProductsRepository() repository.UserFavoriteProductsRepository {
	var m = map[int]orm.UserFavoriteProductSlice{}
	return &userFavoriteProductsRepository{
		m: m,
	}
}

func (r *userFavoriteProductsRepository) Get(ctx context.Context, userID int) (orm.UserFavoriteProductSlice, error) {
	if r.m[userID] == nil {
		return nil, entities.ErrNotFound
	}
	return r.m[userID], nil
}

func (r *userFavoriteProductsRepository) Save(ctx context.Context, favotieProductSlice orm.UserFavoriteProductSlice) error {
	if len(favotieProductSlice) == 0 {
		return nil
	}
	r.m[favotieProductSlice[0].UserID] = favotieProductSlice
	log.Printf("save %v\n", favotieProductSlice)
	return nil
}
