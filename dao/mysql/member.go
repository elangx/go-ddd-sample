package mysql

import "time"

type Member struct {
	Id        int64     `json:"id"`
	MemberId  int64     `json:"member_id"`
	Nickname  string    `json:"nickname"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (m *Member) TableName() string {
	return "member"
}
