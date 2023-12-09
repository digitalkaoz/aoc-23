package four

import (
	"aoc-23/lib"
	"fmt"
	"github.com/urfave/cli/v2"
	"slices"
	"strconv"
	"strings"
)

type Game struct {
	name           string
	winningNumbers []int
	currentNumbers []int
}

func (g Game) Sum() int {
	var sum int
	for _, n := range g.currentNumbers {
		if slices.Contains(g.winningNumbers, n) {
			if sum == 0 {
				sum = 1
			} else {
				sum = sum * 2
			}
		}
	}
	return sum
}

func (g Game) SumMatches() int {
	var sum int
	for _, n := range g.currentNumbers {
		if slices.Contains(g.winningNumbers, n) {
			sum += 1
		}
	}
	return sum
}

func Command() *cli.Command {
	return &cli.Command{
		Name:      "four",
		ArgsUsage: "path/to/data/file",
		Usage:     "day 4 of aoc-23",
		Action: func(cCtx *cli.Context) error {
			sum, err := processFile(cCtx.Args().First(), defaultSummer)
			if err == nil {
				fmt.Printf("day 4: total sum is %d\n", sum)
			}
			return err
		},
	}
}

func processFile(file string, summer func([]*Game) int) (int, error) {
	lines, err := lib.ReadFile(file)
	if err != nil {
		return 0, err
	}
	games, err := parseGames(lines)
	if err != nil {
		return 0, err
	}
	return summer(games), nil
}

func defaultSummer(games []*Game) int {
	sums := sumGames(games)

	return lib.SumSlice(sums)
}

func parseGames(lines []string) ([]*Game, error) {
	var games []*Game

	for _, line := range lines {
		g, err := parseGame(line)
		if err != nil {
			return nil, err
		}
		games = append(games, g)
	}
	return games, nil
}

func multiplyGames(games []*Game) int {
	multiplied := make([][]*Game, len(games))
	for i, g := range games {
		multiplied[i] = []*Game{g}
	}

	for i, m := range multiplied {
		for _, g := range m {
			wins := g.SumMatches()
			if wins > 0 {
				clones := games[i+1 : i+1+wins]
				for j, clone := range clones {
					multiplied[i+1+j] = append(multiplied[i+1+j], clone)
				}
			}
		}
	}

	var sum int
	for _, m := range multiplied {
		sum += len(m)
	}

	return sum
}

func sumGames(games []*Game) []int {
	var sums []int
	for _, g := range games {
		sums = append(sums, g.Sum())
	}

	return sums
}

func parseGame(line string) (*Game, error) {
	g := new(Game)
	parts := strings.Split(line, ": ")
	g.name = parts[0]
	parts = strings.Split(parts[1], " | ")
	// winning numbers
	wn, err := parseNumbers(parts[0])
	if err != nil {
		return nil, err
	}
	g.winningNumbers = wn

	// current numbers
	cn, err := parseNumbers(parts[1])
	if err != nil {
		return nil, err
	}
	g.currentNumbers = cn

	return g, nil
}
func parseNumbers(line string) ([]int, error) {
	var numbers []int
	for _, w := range strings.Split(line, " ") {
		if w == "" {
			continue
		}
		n, err := strconv.Atoi(w)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, n)
	}

	return numbers, nil
}
