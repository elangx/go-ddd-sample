package service

import "go-ddd-sample/domain/entity"

type MemberDomainService interface {
	CheckMember(member *entity.Member) bool
}
