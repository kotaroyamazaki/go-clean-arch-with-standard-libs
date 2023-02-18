package user

import "time"

type User struct {
	id            int
	NickName      NickName
	BirthDate     BitrhDate
	favoriteItems []int
}

func New(nickName string, birthDate time.Time) (*User, error) {
	nn, err := newNicKName(nickName)
	if err != nil {
		return nil, err
	}
	bd, err := newBirthDate(birthDate)
	if err != nil {
		return nil, err
	}

	return &User{
		NickName:  *nn,
		BirthDate: *bd,
	}, nil
}
