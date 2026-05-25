package userdm

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type Email struct {
	value string
}

var ErrInvalidEmailFormat = errors.New("invalid email format")

func NewEmail(v string) (*Email, error) {
	match, err := regexp.MatchString(
		`^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$`,
		v,
	)
	if err != nil {
		return nil, fmt.Errorf("メールアドレスがフォーマット検証に失敗しました: %w", err)
	}
	if !match {
		return nil, fmt.Errorf(
			"無効なメールアドレスフォーマット:%w :%s",
			ErrInvalidEmailFormat,
			v)
	}

	return &Email{
		value: v,
	}, nil
}

func (v *Email) Value() string {
	return v.value
}

func (v *Email) Domain() string {
	parts := strings.Split(v.value, "@")

	return parts[1]
}

func (v *Email) Equal(other Email) bool {
	return v.value == other.value
}
