package models

import "github.com/kotaroyamazaki/go-clean-arch-sample-with-standard/internal/entities"

type NickName string

func newNicKName(s string) (*NickName, error) {
	err := NickName(s).validate()
	if err != nil {
		return nil, err
	}

	n := NickName(s)
	return &n, nil
}

func (s NickName) validate() error {
	if s == "" {
		return entities.ErrValidattion
	}
	// バリデーションのルールを記述
	// :
	// :
	// e.g.
	// if utf8.RuneCountInString(v) < 3 {
	// 	return nil, errors.New("名前は3文字以上")
	// }
	return nil
}

func (n NickName) String() string {
	return string(n)
}
