package two

import (
	"github.com/urfave/cli/v2"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:      "two",
		ArgsUsage: "path/to/data/file",
		Usage:     "day 2 of aoc-23",
		Action: func(cCtx *cli.Context) error {
			return nil
		},
	}
}
