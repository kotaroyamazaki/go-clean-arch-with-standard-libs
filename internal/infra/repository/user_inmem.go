package repository

import (
	"context"
	"log"

	"github.com/kotaroyamazaki/go-clean-arch-sample-with-standard/internal/entities"
	models "github.com/kotaroyamazaki/go-clean-arch-sample-with-standard/internal/entities/models/user"
	"github.com/kotaroyamazaki/go-clean-arch-sample-with-standard/internal/entities/repository"
)

type userRepository struct {
	m map[int]*models.User
}

func NewUserRepository() repository.UserRepository {
	var m = map[int]*models.User{}
	return &userRepository{
		m: m,
	}
}

func (r *userRepository) Save(ctx context.Context, user *models.User) error {
	if r.m[user.ID] != nil {
		return entities.ErrDuplicatedPrimaryKey
	}
	r.m[user.ID] = user

	log.Printf("save %v\n", user)
	return nil
}
