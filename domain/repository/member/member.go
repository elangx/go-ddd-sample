package member

import "go-ddd-sample/domain/entity"

type MemberRepository interface {
	GetMemberByEmail(email string) (*entity.Member, error)
	UpdateMember(member *entity.Member) error
}

var memberIns MemberRepository

func SetMemberRepository(ins MemberRepository) {
	memberIns = ins
}

func GetMemberRepository() MemberRepository {
	return memberIns
}
