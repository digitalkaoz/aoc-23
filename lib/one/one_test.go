package one

import (
	"aoc-23/lib"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSampleInput(t *testing.T) {
	lines := []string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	}

	sums, err := sumLines(lines, sumLinePureDigits)
	assert.Nil(t, err)

	sum := lib.SumSlice(sums)
	assert.Equal(t, 142, sum)
}

func TestSampleInputLetter(t *testing.T) {
	lines := []string{
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
	}

	sums, err := sumLines(lines, sumLineDigitsAndLetters)
	assert.Nil(t, err)

	sum := lib.SumSlice(sums)
	assert.Equal(t, 281, sum)
}

func TestFileDigitsOnly(t *testing.T) {
	sum, err := processFile("../../data/1.txt", sumLinePureDigits)

	assert.Nil(t, err)
	assert.Equal(t, 54388, sum)
}

func TestFileDigitsAndLetters(t *testing.T) {
	sum, err := processFile("../../data/1.txt", sumLineDigitsAndLetters)

	assert.Nil(t, err)
	assert.Equal(t, 53515, sum)
}
