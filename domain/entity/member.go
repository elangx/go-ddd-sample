package entity

import (
	"go-ddd-sample/util"
	"regexp"
	"time"
)

type Member struct {
	MemberId  int64     `json:"member_id"`
	Nickname  string    `json:"nickname"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

const salt = "abcd"

func (m *Member) VerifyPassword(password string) bool {
	pass, err := util.ScryptWithSalt(password, salt)
	if err != nil {
		return false
	}
	return m.Password == pass
}

func (m *Member) EncryptPassword(password string) error {
	pass, err := util.ScryptWithSalt(password, salt)
	if err != nil {
		return err
	}
	m.Password = pass
	return nil
}

func (m *Member) CheckEmail() bool {
	return m != nil && regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`).MatchString(m.Email)
}

func (m *Member) CheckNickname() bool {
	return m != nil && m.Nickname != "" && len(m.Nickname) <= 50
}
