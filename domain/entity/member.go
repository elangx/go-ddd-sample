package entity

import (
	"go-ddd-sample/util"
	"time"
)

type Member struct {
	MemberId  int64
	Nickname  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

const salt = "abcd"

func (m *Member) VerifyPassword(password string) bool {
	pass, err := util.ScryptWithSalt(password, salt)
	if err != nil {
		return false
	}
	return m.Password == pass
}

func (m *Member) SetPassword(password string) error {
	pass, err := util.ScryptWithSalt(password, salt)
	if err != nil {
		return err
	}
	m.Password = pass
	return nil
}
