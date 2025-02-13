package main

import (
	"go-ddd-sample/httpserver"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name: "go-ddd-sample",
		Before: func(context *cli.Context) error {
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
