package main

import (
	"aoc-23/lib/one"
	"aoc-23/lib/two"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := buildCli()

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func buildCli() *cli.App {
	return &cli.App{
		Commands: []*cli.Command{
			one.Command(),
			two.Command(),
		},
	}
}
