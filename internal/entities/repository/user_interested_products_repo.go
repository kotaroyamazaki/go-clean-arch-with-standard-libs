package repository

import (
	"context"

	"github.com/kotaroyamazaki/go-clean-arch-sample-with-standard/pkg/orm"
)

type UserFavoriteBooksRepository interface {
	Save(ctx context.Context, favotieItemsorms orm.UserFavoriteBookSlice) error
}
