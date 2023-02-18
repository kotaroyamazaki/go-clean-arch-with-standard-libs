package usecases

import (
	"context"
	"time"

	"github.com/kauche-interview/backend-KotaroYamazaki/src/internal/entities/repository"
	"github.com/kauche-interview/backend-KotaroYamazaki/src/internal/entities/user"
	"github.com/kauche-interview/backend-KotaroYamazaki/src/internal/infra/db"
	"github.com/kauche-interview/backend-KotaroYamazaki/src/pkg/orm"
)

type userUsecase struct {
	userRepo            repository.UserRepository
	favoriteProductRepo repository.UserFavoriteProductsRepository
}
type UserUsecase interface {
	Register(ctx context.Context, birthDate time.Time, nickname string, favoriteProductIDs []int) error
}

func NewUserUsecase(userRepo repository.UserRepository, favoriteProductRepo repository.UserFavoriteProductsRepository) UserUsecase {
	return &userUsecase{
		userRepo:            userRepo,
		favoriteProductRepo: favoriteProductRepo,
	}
}

func (uc *userUsecase) Register(ctx context.Context, birthDate time.Time, nickname string, favoriteProductIDs []int) error {
	u, err := user.New(nickname, birthDate)
	if err != nil {
		return err
	}

	dbUser := orm.NewUser(
		string(u.NickName),
		u.BirthDate.GetBirthDate(),
	)
	return db.RunTransaction(ctx, func(ctx context.Context) error {
		if err := uc.userRepo.Save(ctx, dbUser); err != nil {
			return err
		}
		_favoriteProductIDs := make(orm.UserFavoriteProductSlice, 0, len(favoriteProductIDs))
		for _, pID := range favoriteProductIDs {
			_favoriteProductIDs = append(_favoriteProductIDs, orm.NewUserFavoriteProduct(
				dbUser.ID, pID,
			))
		}

		return uc.favoriteProductRepo.Save(ctx, _favoriteProductIDs)
	})
}
