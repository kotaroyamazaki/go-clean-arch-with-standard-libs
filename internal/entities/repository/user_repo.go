package repository

import (
	"context"

	"github.com/kotaroyamazaki/go-clean-arch-sample-with-standard/pkg/orm"
)

type UserRepository interface {
	Save(ctx context.Context, user *orm.User) error
}
