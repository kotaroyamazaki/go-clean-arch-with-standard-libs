package orm

import (
	"time"
)

type User struct {
	ID        int
	NickeName string
	BitrhDate time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	R         *Relation
}

type Relation struct {
	FavoriteBooks []*Book
}

func NewUser(nickname string, birthDate time.Time) *User {
	return &User{
		ID:        NewUserID(),
		NickeName: nickname,
		BitrhDate: birthDate,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		R: &Relation{
			FavoriteBooks: []*Book{},
		},
	}
}
