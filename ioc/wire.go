//go:build wireinject
// +build wireinject

package ioc

import (
	"go-ddd-sample/domain/repository"
	"go-ddd-sample/domain/service"

	appService "go-ddd-sample/application/service"

	"github.com/google/wire"
)

var memberRepositorySet = wire.NewSet(
	repository.NewMemberRepositoryMySQL,
	wire.Bind(new(repository.MemberRepository), new(*repository.MemberRepositoryMySQL)),
)

var memberServiceSet = wire.NewSet(
	service.NewMemberDomainServiceImpl,
	wire.Bind(new(service.MemberDomainService), new(*service.MemberDomainServiceImpl)),
)

func GetMemberService() *appService.MemberService {
	wire.Build(appService.NewMemberService, memberRepositorySet, memberServiceSet)
	return &appService.MemberService{}
}
