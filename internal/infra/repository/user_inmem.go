package repository

import (
	"context"
	"log"

	"github.com/kotaroyamazaki/go-clean-arch-sample-with-standard/internal/entities"
	"github.com/kotaroyamazaki/go-clean-arch-sample-with-standard/internal/entities/repository"
	"github.com/kotaroyamazaki/go-clean-arch-sample-with-standard/pkg/orm"
)

type userRepository struct {
	m map[int]*orm.User
}

func NewUserRepository() repository.UserRepository {
	var m = map[int]*orm.User{}
	return &userRepository{
		m: m,
	}
}

func (r *userRepository) Save(ctx context.Context, user *orm.User) error {
	if r.m[user.ID] != nil {
		return entities.ErrDuplicatedPrimaryKey
	}
	r.m[user.ID] = user

	log.Printf("save %v\n", user)
	return nil
}
