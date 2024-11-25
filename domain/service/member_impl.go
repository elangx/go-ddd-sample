package service

import "go-ddd-sample/domain/entity"

var _ MemberDomainService = (*memberDomainServiceImpl)(nil)

type memberDomainServiceImpl struct{}

func NewMemberDomainServiceImpl() MemberDomainService {
	return &memberDomainServiceImpl{}
}

func (m *memberDomainServiceImpl) CheckMember(member *entity.Member) bool {
	return member.CheckNickname() && member.CheckEmail()
}
