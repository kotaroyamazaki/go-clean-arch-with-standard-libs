package usecases

import (
	"context"
	"time"

	"github.com/kotaroyamazaki/go-clean-arch-sample-with-standard/internal/entities/repository"
	"github.com/kotaroyamazaki/go-clean-arch-sample-with-standard/internal/entities/user"
	"github.com/kotaroyamazaki/go-clean-arch-sample-with-standard/internal/infra/db"
	"github.com/kotaroyamazaki/go-clean-arch-sample-with-standard/pkg/orm"
)

type userUsecase struct {
	userRepo         repository.UserRepository
	favoriteBookRepo repository.UserFavoriteBooksRepository
}
type UserUsecase interface {
	Register(ctx context.Context, birthDate time.Time, nickname string, favoriteBookIDs []int) error
}

func NewUserUsecase(userRepo repository.UserRepository, favoriteBookRepo repository.UserFavoriteBooksRepository) UserUsecase {
	return &userUsecase{
		userRepo:         userRepo,
		favoriteBookRepo: favoriteBookRepo,
	}
}

func (uc *userUsecase) Register(ctx context.Context, birthDate time.Time, nickname string, favoriteBookIDs []int) error {
	u, err := user.New(nil, nickname, birthDate, favoriteBookIDs)
	if err != nil {
		return err
	}

	dbUser := orm.NewUser(
		u.GetNickName().String(),
		u.GetBirthDate().Time(),
	)
	return db.RunTransaction(ctx, func(ctx context.Context) error {
		if err := uc.userRepo.Save(ctx, dbUser); err != nil {
			return err
		}
		_favoriteBookIDs := make(orm.UserFavoriteBookSlice, 0, len(favoriteBookIDs))
		for _, pID := range favoriteBookIDs {
			_favoriteBookIDs = append(_favoriteBookIDs, orm.NewUserFavoriteBook(
				dbUser.ID, pID,
			))
		}

		return uc.favoriteBookRepo.Save(ctx, _favoriteBookIDs)
	})
}
