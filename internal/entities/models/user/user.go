package models

import (
	"math/rand"
	"time"
)

type User struct {
	ID            int
	nickName      NickName
	birthDate     BitrhDate
	favoriteItems []int
}

func NewUser(nickName string, birthDate time.Time, favoriteItems []int) (*User, error) {
	nn, err := newNicKName(nickName)
	if err != nil {
		return nil, err
	}
	bd, err := newBirthDate(birthDate)
	if err != nil {
		return nil, err
	}

	// TODO: UUIDやULIDを用いる
	id := rand.Int()

	return &User{
		ID:            id,
		nickName:      *nn,
		birthDate:     *bd,
		favoriteItems: favoriteItems,
	}, nil
}

func (u *User) GetNickName() NickName {
	return u.nickName
}

func (u *User) GetBirthDate() BitrhDate {
	return u.birthDate
}
