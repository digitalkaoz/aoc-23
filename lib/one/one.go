package one

import (
	"aoc-23/lib"
	"fmt"
	"github.com/urfave/cli/v2"
	"strconv"
	"strings"
	"unicode"
)

var replacements = map[string]int{ //slice bc we need the replacements ordered
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func Command() *cli.Command {
	return &cli.Command{
		Name:      "one",
		ArgsUsage: "path/to/data/file",
		Usage:     "day 1 of aoc-23",
		Action: func(cCtx *cli.Context) error {
			sum, err := processFile(cCtx.Args().First(), sumLineDigitsAndLetters)
			if err == nil {
				fmt.Printf("day 1: total sum is %d\n", sum)
			}
			return err
		},
	}
}

func processFile(file string, summer func(string) (int, error)) (int, error) {
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

func sumLines(lines []string, summer func(string) (int, error)) ([]int, error) {
	filtered := make([]int, len(lines))

	for i, line := range lines {

		lineSum, err := summer(line)
		if err != nil {
			return nil, err
		}
		filtered[i] = lineSum
	}

	return filtered, nil
}

func sumLinePureDigits(line string) (int, error) {
	var first int
	var last int

	// ltr
	for _, c := range line {
		if unicode.IsDigit(c) {
			digit, _ := strconv.Atoi(string(c))
			first = digit
			break
		}
	}

	// rtl
	for _, c := range lib.ReverseString(line) {
		if unicode.IsDigit(c) {
			digit, _ := strconv.Atoi(string(c))
			last = digit
			break
		}
	}

	return strconv.Atoi(fmt.Sprintf("%d%d", first, last))
}

func sumLineDigitsAndLetters(line string) (int, error) {
	var first int
	var last int

	// ltr
	for i, c := range line {
		if unicode.IsDigit(c) {
			digit, _ := strconv.Atoi(string(c))
			first = digit
			break
		}
		letter := StartsWithLetteredNumber(line[i:], false)
		if letter != "" {
			first = replacements[letter]
			break
		}
	}

	// rtl
	line = lib.ReverseString(line)
	for i, c := range line {
		if unicode.IsDigit(c) {
			digit, _ := strconv.Atoi(string(c))
			last = digit
			break
		}
		letter := StartsWithLetteredNumber(line[i:], true)
		if letter != "" {
			last = replacements[lib.ReverseString(letter)]
			break
		}
	}

	return strconv.Atoi(fmt.Sprintf("%d%d", first, last))
}

func StartsWithLetteredNumber(line string, reverse bool) string {
	for search := range replacements {
		if reverse {
			search = lib.ReverseString(search)
		}
		if strings.HasPrefix(line, search) {
			return search
		}
	}

	return ""
}
