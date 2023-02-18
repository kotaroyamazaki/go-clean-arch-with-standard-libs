package user

import (
	"time"

	"github.com/kauche-interview/backend-KotaroYamazaki/src/internal/entities"
)

type BitrhDate struct {
	birthDate time.Time
}

func newBirthDate(t time.Time) (*BitrhDate, error) {
	err := validateBirthDate(t)
	if err != nil {
		return nil, err
	}

	return &BitrhDate{
		birthDate: t,
	}, nil
}

func validateBirthDate(t time.Time) error {
	now := time.Now()
	if t.After(now) {
		return entities.ErrValidattion
	}
	return nil
}

func (bd *BitrhDate) GetBirthDate() time.Time {
	return bd.birthDate
}

// anti-YAGNI
func (bd *BitrhDate) GetAge() int {
	now := time.Now()
	years := now.Year() - bd.birthDate.Year()
	if now.YearDay() < bd.birthDate.YearDay() {
		years--
	}
	return years
}
