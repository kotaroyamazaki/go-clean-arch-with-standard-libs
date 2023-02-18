package repository

import (
	"context"

	"github.com/kauche-interview/backend-KotaroYamazaki/src/pkg/orm"
)

type UserFavoriteProductsRepository interface {
	Get(ctx context.Context, userID int) (orm.UserFavoriteProductSlice, error)
	Save(ctx context.Context, favotieItemsorms orm.UserFavoriteProductSlice) error
}
