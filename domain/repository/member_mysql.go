package repository

import (
	"errors"
	"go-ddd-sample/dao/mysql"
	"go-ddd-sample/domain/entity"
	"gorm.io/gorm"
)

var _ MemberRepository = (*MemberRepositoryMySQL)(nil)

type MemberRepositoryMySQL struct {
	db *gorm.DB
}

func NewMemberRepositoryMySQL() *MemberRepositoryMySQL {
	return &MemberRepositoryMySQL{
		db: mysql.GetDB(),
	}
}

func (m *MemberRepositoryMySQL) GetMemberByEmail(email string) (*entity.Member, error) {
	mdb := &mysql.Member{}
	err := m.db.Where("email = ?", email).First(mdb).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &entity.Member{
		MemberId:  mdb.MemberId,
		Nickname:  mdb.Nickname,
		Email:     mdb.Email,
		Password:  mdb.Password,
		CreatedAt: mdb.CreatedAt,
		UpdatedAt: mdb.UpdatedAt,
	}, nil
}

func (m *MemberRepositoryMySQL) CreateMember(member *entity.Member) error {
	return m.db.Create(&mysql.Member{
		MemberId:  member.MemberId,
		Nickname:  member.Nickname,
		Email:     member.Email,
		Password:  member.Password,
		CreatedAt: member.CreatedAt,
		UpdatedAt: member.UpdatedAt,
	}).Error
}
