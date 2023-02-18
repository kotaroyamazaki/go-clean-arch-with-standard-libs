package user

import "time"

type User struct {
	id            *int
	nickName      NickName
	birthDate     BitrhDate
	favoriteItems []int
}

func New(id *int, nickName string, birthDate time.Time, favoriteItems []int) (*User, error) {
	nn, err := newNicKName(nickName)
	if err != nil {
		return nil, err
	}
	bd, err := newBirthDate(birthDate)
	if err != nil {
		return nil, err
	}

	return &User{
		id:            id,
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
