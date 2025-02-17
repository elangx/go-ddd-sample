package service

import "go-ddd-sample/domain/entity"

var _ MemberDomainService = (*MemberDomainServiceImpl)(nil)

type MemberDomainServiceImpl struct{}

func NewMemberDomainServiceImpl() *MemberDomainServiceImpl {
	return &MemberDomainServiceImpl{}
}

func (m *MemberDomainServiceImpl) CheckMember(member *entity.Member) bool {
	return member.CheckNickname() && member.CheckEmail()
}
