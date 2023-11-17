package repository

import (
	"context"
)

type UserFavoriteBooksRepository interface {
	Save(ctx context.Context, favoriteBookIDs []int) error
}
