package three

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSampleInput(t *testing.T) {
	lines := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}

	matrixes := collectMatrixes(lines)

	assert.Equal(t, 4361, sumMatrixes(matrixes))
}

func TestGearInput(t *testing.T) {
	lines := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}

	matrixes := collectGearMatrixes(lines)

	assert.Equal(t, 467835, sumMatrixes(matrixes))
}
func TestFile(t *testing.T) {
	sum, err := processFile("../../data/3.txt", collectMatrixes)

	assert.Nil(t, err)
	assert.Equal(t, 546_563, sum)
}

func TestGearFile(t *testing.T) {
	sum, err := processFile("../../data/3.txt", collectGearMatrixes)

	assert.Nil(t, err)
	assert.Equal(t, 91_031_374, sum)
}

func TestContainsSymbol(t *testing.T) {
	assert.True(t, containsSymbol("..*.."))
	assert.False(t, containsSymbol("..1234.."))
	assert.True(t, containsSymbol("..*..12334..."))
	assert.False(t, containsSymbol("......."))
}

func TestCreateMatrix(t *testing.T) {
	inp := []string{
		"...13.....",
		"...12.....",
		"...*7.....",
	}

	matrixes := collectMatrixes(inp)

	assert.Len(t, matrixes, 3)
	assert.Equal(t, 0, matrixes[0].Number())
	assert.Equal(t, 12, matrixes[1].Number())
	assert.Equal(t, 7, matrixes[2].Number())
	assert.Equal(t, 19, sumMatrixes(matrixes))
}

func TestCreateGearMatrix(t *testing.T) {
	lines := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}

	gm := new(GearMatrix)
	gm.start = 3

	gm.CollectLines(lines, 1)

	assert.Equal(t, []string{
		"467..",
		".*.",
		".35.",
	}, gm.lines)
	assert.Equal(t, 16_345, gm.Number())
}

func TestGearMatrixEol(t *testing.T) {
	lines := []string{
		"....388",
		"541*...",
		".......",
	}

	gm := new(GearMatrix)
	gm.start = 3

	gm.CollectLines(lines, 1)

	assert.Equal(t, []string{
		"..388",
		"541*.",
		"...",
	}, gm.lines)
	assert.Equal(t, 209_908, gm.Number())
}
