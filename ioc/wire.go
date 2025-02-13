//go:build wireinject
// +build wireinject

package ioc

import (
	"go-ddd-sample/domain/repository"
	"go-ddd-sample/domain/service"

	appService "go-ddd-sample/application/service"

	"go-ddd-sample/dao/mysql"

	"github.com/google/wire"
)

var memberRepositorySet = wire.NewSet(
	repository.NewMemberRepositoryMySQL,
	wire.Bind(new(repository.MemberRepository), new(*repository.MemberRepositoryMySQL)),
	mysql.GetDB,
)

var memberServiceSet = wire.NewSet(
	service.NewMemberDomainServiceImpl,
	wire.Bind(new(service.MemberDomainService), new(*service.MemberDomainServiceImpl)),
)

func GetMemberService() *appService.MemberService {
	wire.Build(appService.NewMemberService, memberRepositorySet, memberServiceSet)
	return &appService.MemberService{}
}
