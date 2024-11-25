package service

import (
	"context"
	"errors"
	"go-ddd-sample/application/dto"
	"go-ddd-sample/domain/entity"
	"go-ddd-sample/domain/repository"
	"go-ddd-sample/domain/service"
)

type MemberService struct {
	memberRepo    repository.MemberRepository
	memberService service.MemberDomainService
}

func NewMemberService(repo repository.MemberRepository, service service.MemberDomainService) *MemberService {
	return &MemberService{
		memberRepo:    repo,
		memberService: service,
	}
}

func (m *MemberService) Login(ctx context.Context, rq *dto.LoginRequest) (*dto.LoginResponse, error) {
	rs := &dto.LoginResponse{}
	memberInfo, err := m.memberRepo.GetMemberByEmail(rq.Email)
	if err != nil {
		return nil, err
	}
	if memberInfo == nil {
		return nil, errors.New("empty member")
	}
	if !memberInfo.VerifyPassword(rq.Password) {
		return nil, errors.New("invalid password")
	}
	rs.MemberId = memberInfo.MemberId
	return rs, nil
}

func (m *MemberService) Register(ctx context.Context, rq *dto.RegisterRequest) (*dto.RegisterResponse, error) {
	memberInfo := &entity.Member{
		Nickname: rq.Nickname,
		Email:    rq.Email,
	}
	if !m.memberService.CheckMember(memberInfo) {
		return nil, errors.New("member verify failed")
	}
	err := memberInfo.EncryptPassword(rq.Password)
	if err != nil {
		return nil, err
	}
	err = m.memberRepo.CreateMember(memberInfo)
	if err != nil {
		return nil, err
	}
	return &dto.RegisterResponse{
		MemberId: memberInfo.MemberId,
	}, nil
}
