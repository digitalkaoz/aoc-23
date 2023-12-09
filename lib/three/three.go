package three

import (
	"aoc-23/lib"
	"fmt"
	"github.com/urfave/cli/v2"
	"math"
	"strconv"
	"strings"
	"unicode"
)

type Summable interface {
	Number() int
}

type Matrix struct {
	number    int
	start     int
	end       int
	lineIndex int
	lines     []string
}

func (m *Matrix) AddDigit(digit string) {
	newNumber, _ := strconv.Atoi(strconv.Itoa(m.number) + digit)
	m.number = newNumber
}

func (m Matrix) Number() int {
	lines := m.lines
	lines[m.lineIndex] = strings.Replace(lines[m.lineIndex], strconv.Itoa(m.number), "", 1)
	if containsSymbol(strings.Join(lines, "")) {
		return m.number
	}
	return 0
}

func (m *Matrix) CollectLines(lines []string, lineIndex int) {
	//prev line
	if lineIndex > 0 {
		m.lines = append(m.lines, lines[lineIndex-1][m.start:m.end])
		m.lineIndex = 1 // we have a previous line, so the line of interest is in the second row, otherwise first
	}

	// current line
	m.lines = append(m.lines, lines[lineIndex][m.start:m.end])

	// next line
	if lineIndex < len(lines)-1 {
		m.lines = append(m.lines, lines[lineIndex+1][m.start:m.end])
	}
}

type GearMatrix struct {
	start     int
	lineIndex int
	lines     []string
}

func (m GearMatrix) Number() int {
	// join all lines together, split by "." and collect the numbers
	var numbers []int
	fragments := strings.Split(strings.Replace(strings.Join(m.lines, "."), "*", ".", -1), ".")
	for _, fragment := range fragments {
		if len(fragment) == 0 {
			continue
		}
		number, _ := strconv.Atoi(fragment)
		numbers = append(numbers, number)
	}
	if len(numbers) != 2 {
		//skipped gear with more or less than 2 numbers
		return 0
	}

	//log.Printf("m: %q, a: %d * %d", m.lines, numbers[0], numbers[1])
	return numbers[0] * numbers[1]
}

func (m *GearMatrix) CollectLines(lines []string, lineIndex int) {
	if lineIndex > 0 {
		m.lines = append(m.lines, collectNumbersSlice(lines[lineIndex-1], m.start))
	}

	m.lines = append(m.lines, collectNumbersSlice(lines[lineIndex], m.start))

	if lineIndex < len(lines)-1 {
		m.lines = append(m.lines, collectNumbersSlice(lines[lineIndex+1], m.start))
	}
}

func Command() *cli.Command {
	return &cli.Command{
		Name:      "three",
		ArgsUsage: "path/to/data/file",
		Usage:     "day 3 of aoc-23",
		Action: func(cCtx *cli.Context) error {
			//sum, err := processFile[Matrix](cCtx.Args().First(), collectMatrixes)
			sum, err := processFile[GearMatrix](cCtx.Args().First(), collectGearMatrixes)
			if err == nil {
				fmt.Printf("day 3: total sum is %d\n", sum)
			}
			return err
		},
	}
}

func processFile[T Summable](file string, collector func([]string) []T) (int, error) {
	lines, err := lib.ReadFile(file)
	if err != nil {
		return 0, err
	}

	matrixes := collector(lines)

	return sumMatrixes(matrixes), nil
}

func sumMatrixes[T Summable](matrixes []T) int {
	var sum int
	for _, matrix := range matrixes {
		sum += matrix.Number()
	}
	return sum
}

func collectMatrixes(lines []string) []Matrix {
	matrixes := make([]Matrix, 0)

	for lineIndex, line := range lines {
		var matrix *Matrix
		for numberIndex, char := range line {
			// new digit found
			if matrix == nil && unicode.IsDigit(char) {
				matrix = new(Matrix)
				matrix.start = int(math.Max(float64(numberIndex-1), float64(0)))
			}
			// a digit so add to current matrix number
			if unicode.IsDigit(char) {
				matrix.AddDigit(string(char))
			}

			// end of digit or EOL
			if matrix != nil && (!unicode.IsDigit(char) || numberIndex+1 == len(line)) {
				matrix.end = int(math.Min(float64(numberIndex+1), float64(len(line))))
				matrix.CollectLines(lines, lineIndex)
				matrixes = append(matrixes, *matrix)
				matrix = nil
			}
		}
	}
	return matrixes
}

func collectGearMatrixes(lines []string) []GearMatrix {
	matrixes := make([]GearMatrix, 0)

	for lineIndex, line := range lines {
		var matrix *GearMatrix
		for numberIndex, char := range line {
			// new digit found
			if string(char) == "*" {
				matrix = new(GearMatrix)
				matrix.start = int(math.Max(float64(numberIndex), float64(0)))
				matrix.CollectLines(lines, lineIndex)
				matrixes = append(matrixes, *matrix)
				matrix = nil
			}
		}
	}
	return matrixes
}

func containsSymbol(line string) bool {
	for _, r := range line {
		if string(r) != "." && !unicode.IsDigit(r) {
			return true
		}
	}
	return false
}

func collectNumbersSlice(line string, start int) string {
	var realStart int
	var realEnd int

	for i := start; i >= 0; i-- {
		if i <= start-1 && !unicode.IsDigit(rune(line[i])) && string(line[i]) != "*" {
			realStart = i
			break
		}
	}

	for i := start; i < len(line); i++ {
		if i+1 == len(line) || (i >= start+1 && !unicode.IsDigit(rune(line[i])) && string(line[i]) != "*") {
			realEnd = i
			break
		}
	}

	start = int(math.Min(float64(realStart), float64(start-1)))
	end := int(math.Max(float64(realEnd), float64(start)))

	return line[start : end+1]
}
