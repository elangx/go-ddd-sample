package repository

import "go-ddd-sample/domain/entity"

type MemberRepository interface {
	GetMemberByEmail(email string) (*entity.Member, error)
	CreateMember(member *entity.Member) error
}
