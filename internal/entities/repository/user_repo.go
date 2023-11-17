package repository

import (
	"context"

	models "github.com/kotaroyamazaki/go-clean-arch-sample-with-standard/internal/entities/models/user"
)

type UserRepository interface {
	Save(ctx context.Context, user *models.User) error
}
