package usecases

import (
	"context"
	"time"

	models "github.com/kotaroyamazaki/go-clean-arch-sample-with-standard/internal/entities/models/user"
	"github.com/kotaroyamazaki/go-clean-arch-sample-with-standard/internal/entities/repository"

	"github.com/kotaroyamazaki/go-clean-arch-sample-with-standard/internal/infra/db"
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
	u, err := models.NewUser(nickname, birthDate, favoriteBookIDs)
	if err != nil {
		return err
	}

	return db.RunTransaction(ctx, func(ctx context.Context) error {
		if err := uc.userRepo.Save(ctx, u); err != nil {
			return err
		}

		return uc.favoriteBookRepo.Save(ctx, favoriteBookIDs)
	})
}
