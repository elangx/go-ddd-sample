package service

import "go-ddd-sample/domain/entity"

type MemberDomainService interface {
	CheckMember(member *entity.Member) bool
}

var memberDomainService MemberDomainService

func GetMemberDomainService() MemberDomainService {
	return memberDomainService
}
