package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:   "zion-test",
		Usage:  "zion test framework",
		Action: start,
		Before: Init,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name: "config",
				Value: "config.json",
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal("Start error", "err", err)
	}
}

func start(ctx *cli.Context) (err error) {
	err = Run()
	return
}

func Init(ctx *cli.Context) (err error) {
	err = NewConfig(ctx.String("config"))
	if err != nil { return }
	err = Setup()
	return
}