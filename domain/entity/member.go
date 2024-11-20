package entity

import (
	"errors"
	"go-ddd-sample/util"
	"regexp"
	"time"
)

type Member struct {
	MemberId  int64
	Nickname  string
	Email     string
	password  string
	createdAt time.Time
	updatedAt time.Time
}

const salt = "abcd"

func (m *Member) VerifyPassword(password string) bool {
	pass, err := util.ScryptWithSalt(password, salt)
	if err != nil {
		return false
	}
	return m.password == pass
}

func (m *Member) EncryptPassword(password string) error {
	pass, err := util.ScryptWithSalt(password, salt)
	if err != nil {
		return err
	}
	m.password = pass
	return nil
}

func (m *Member) VerifyInfo() error {
	if m == nil {
		return errors.New("member is nil")
	}
	if !regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`).MatchString(m.Email) {
		return errors.New("invalid email")
	}
	if m.Nickname == "" || len(m.Nickname) > 50 {
		return errors.New("invalid nickname")
	}
	return nil
}
