package repository

import (
	"context"
	"log"

	"github.com/kauche-interview/backend-KotaroYamazaki/src/internal/entities"
	"github.com/kauche-interview/backend-KotaroYamazaki/src/internal/entities/repository"
	"github.com/kauche-interview/backend-KotaroYamazaki/src/internal/entities/user"
	"github.com/kauche-interview/backend-KotaroYamazaki/src/pkg/orm"
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

func (r *userRepository) Get(ctx context.Context, id int) (*user.User, error) {
	if r.m[id] == nil {
		return nil, entities.ErrNotFound
	}
	_u := r.m[id]

	return user.New(_u.NickeName, _u.BitrhDate)
}

func (r *userRepository) Save(ctx context.Context, user *orm.User) error {
	if r.m[user.ID] != nil {
		return entities.ErrDuplicatedPrimaryKey
	}
	r.m[user.ID] = user

	log.Printf("save %v\n", user)
	return nil
}
