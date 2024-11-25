package httpserver

import (
	"github.com/gin-gonic/gin"
	"log"
)

func NewServerHandler() *gin.Engine {
	e := gin.New()
	e.Use(gin.Recovery())
	return e
}

func Run() {
	if err := NewServerHandler().Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
