package user

import "github.com/kauche-interview/backend-KotaroYamazaki/src/internal/entities"

type NickName string

func newNicKName(s string) (*NickName, error) {
	err := validateName(s)
	if err != nil {
		return nil, err
	}

	n := NickName(s)
	return &n, nil
}

func validateName(s string) error {
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
