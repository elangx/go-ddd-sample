package httpserver

import (
	"github.com/gin-gonic/gin"
	"go-ddd-sample/application/service"
	"go-ddd-sample/domain/repository"
	service2 "go-ddd-sample/domain/service"
	"log"
)

func NewServerHandler() *gin.Engine {
	e := gin.New()
	e.Use(gin.Recovery())
	return e
}

func Run() {
	s := service.NewMemberService(repository.NewMemberRepositoryMySQL(), service2.NewMemberDomainServiceImpl())
	e := NewServerHandler()
	e.Handle("POST", "/login", HandlerFunc(s.Login))
	e.Handle("POST", "/register", HandlerFunc(s.Register))
	if err := NewServerHandler().Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
