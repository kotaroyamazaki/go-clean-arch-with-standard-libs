package user

import (
	"time"

	"github.com/kotaroyamazaki/go-clean-arch-sample-with-standard/internal/entities"
)

type BitrhDate time.Time

func newBirthDate(t time.Time) (*BitrhDate, error) {
	err := BitrhDate(t).validate()
	if err != nil {
		return nil, err
	}
	bd := BitrhDate(t)
	return &bd, nil
}

func (t BitrhDate) validate() error {
	now := time.Now()
	if t.Time().After(now) {
		return entities.ErrValidattion
	}
	return nil
}

func (bd BitrhDate) Time() time.Time {
	return time.Time(bd)
}
