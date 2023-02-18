package orm

import (
	"time"
)

type Book struct {
	ID          int
	Name        string
	Description string
	Price       int
	ImageURL    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
