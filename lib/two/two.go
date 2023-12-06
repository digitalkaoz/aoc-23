package two

import (
	"aoc-23/lib"
	"fmt"
	"github.com/urfave/cli/v2"
	"strconv"
	"strings"
)

type Game struct {
	Name string
	Sets []*GameSet
}

func (s Game) Number() int {
	for _, set := range s.Sets {
		if !set.IsValid() {
			return 0
		}
	}

	number, _ := strconv.Atoi(strings.TrimPrefix(s.Name, "Game "))

	return number
}

func (s Game) MinimumSet() *GameSet {
	set := new(GameSet)

	for _, gameSet := range s.Sets {
		if gameSet.Red > set.Red {
			set.Red = gameSet.Red
		}
		if gameSet.Green > set.Green {
			set.Green = gameSet.Green
		}
		if gameSet.Blue > set.Blue {
			set.Blue = gameSet.Blue
		}
	}

	return set
}

type GameSet struct {
	Blue  int
	Red   int
	Green int
}

func (s GameSet) IsValid() bool {
	return s.Green <= 13 && s.Red <= 12 && s.Blue <= 14
}

func (s GameSet) Power() int {
	return s.Green * s.Red * s.Blue
}

func Command() *cli.Command {
	return &cli.Command{
		Name:      "two",
		ArgsUsage: "path/to/data/file",
		Usage:     "day 2 of aoc-23",
		Action: func(cCtx *cli.Context) error {
			sum, err := processFile(cCtx.Args().First(), func(game *Game) int {
				return game.Number()
			})
			if err == nil {
				fmt.Printf("day 2: total sum is %d\n", sum)
			}
			return err
		},
	}
}

func processFile(file string, summer func(*Game) int) (int, error) {
	lines, err := lib.ReadFile(file)
	if err != nil {
		return 0, err
	}
	sums, err := sumLines(lines, summer)
	if err != nil {
		return 0, err
	}

	return lib.SumSlice(sums), nil
}

func sumLines(lines []string, summer func(*Game) int) ([]int, error) {
	games := make([]int, len(lines))

	for i, line := range lines {
		game := parseGame(line)
		if game != nil {
			games[i] = summer(game)
		}
	}

	return games, nil
}

func parseGame(line string) *Game {
	game := new(Game)
	parts := strings.Split(line, ": ")
	game.Name = parts[0]

	for _, set := range strings.Split(parts[1], ";") {
		gameSet := new(GameSet)
		for _, draw := range strings.Split(set, ",") {
			draw = strings.TrimSpace(draw)
			if strings.HasSuffix(draw, "red") {
				number, _ := strconv.Atoi(strings.TrimSuffix(draw, " red"))
				gameSet.Red = number
			} else if strings.HasSuffix(draw, "green") {
				number, _ := strconv.Atoi(strings.TrimSuffix(draw, " green"))
				gameSet.Green = number
			} else if strings.HasSuffix(draw, "blue") {
				number, _ := strconv.Atoi(strings.TrimSuffix(draw, " blue"))
				gameSet.Blue = number
			}
		}

		game.Sets = append(game.Sets, gameSet)
	}

	return game
}
