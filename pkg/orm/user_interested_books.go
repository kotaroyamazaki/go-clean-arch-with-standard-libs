package orm

import (
	"time"
)

type UserFavoriteBookSlice []*UserFavoriteBook

type UserFavoriteBook struct {
	UserID    int
	BookID    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUserFavoriteBook(userID, bookID int) *UserFavoriteBook {
	return &UserFavoriteBook{
		UserID:    userID,
		BookID:    bookID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
