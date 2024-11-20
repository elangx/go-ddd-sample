package member

import (
	"go-ddd-sample/domain/entity"
	"gorm.io/gorm"
)

var _ MemberRepository = (*MemberRepositoryMySQL)(nil)

type MemberRepositoryMySQL struct {
	db *gorm.DB
}

func NewMemberRepositoryMySQL() *MemberRepositoryMySQL {
	return &MemberRepositoryMySQL{}
}

func (m *MemberRepositoryMySQL) GetMemberByEmail(email string) (*entity.Member, error) {
	return nil, nil
}

func (m *MemberRepositoryMySQL) UpdateMember(em *entity.Member) error {
	return nil
}
