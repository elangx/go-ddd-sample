package main

import (
	"github.com/urfave/cli/v2"
	"go-ddd-sample/dao/mysql"
	"go-ddd-sample/domain/repository"
	"go-ddd-sample/domain/service"
	"go-ddd-sample/httpserver"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name: "go-ddd-sample",
		Before: func(context *cli.Context) error {
			mysql.Init()
			repository.Init()
			service.Init()
			return nil
		},
		Action: func(context *cli.Context) error {
			httpserver.Run()
			return nil
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
