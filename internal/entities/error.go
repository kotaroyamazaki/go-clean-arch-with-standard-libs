package entities

import (
	"errors"
)

var (
	ErrNotFound             = errors.New("not found")
	ErrDuplicatedPrimaryKey = errors.New("duplicated primary key")
	ErrValidattion          = errors.New("invalid entity")
)
