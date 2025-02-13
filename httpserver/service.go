package httpserver

import (
	ioc "go-ddd-sample/ioc"
	"log"

	"github.com/gin-gonic/gin"
)

func NewServerHandler() *gin.Engine {
	e := gin.New()
	e.Use(gin.Recovery())
	return e
}

func Run() {
	s := ioc.GetMemberService()
	e := NewServerHandler()
	e.Handle("POST", "/login", HandlerFunc(s.Login))
	e.Handle("POST", "/register", HandlerFunc(s.Register))
	if err := e.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
