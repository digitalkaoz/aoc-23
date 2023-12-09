package four

import (
	"aoc-23/lib"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSampleInput(t *testing.T) {
	lines := []string{
		"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
		"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
		"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
		"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
		"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
		"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
	}

	games, err := parseGames(lines)

	assert.Nil(t, err)
	assert.Len(t, games, 6)
	assert.Equal(t, []int{8, 2, 2, 1, 0, 0}, sumGames(games))
	assert.Equal(t, 13, lib.SumSlice(sumGames(games)))
}

func TestFile(t *testing.T) {
	sum, err := processFile("../../data/4.txt", defaultSummer)

	assert.Nil(t, err)
	assert.Equal(t, 20855, sum)
}

func TestSampleInputMultiplier(t *testing.T) {
	lines := []string{
		"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
		"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
		"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
		"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
		"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
		"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
	}

	games, err := parseGames(lines)

	assert.Nil(t, err)
	assert.Equal(t, 30, multiplyGames(games))
}

func TestMultiplierFile(t *testing.T) {
	sum, err := processFile("../../data/4.txt", multiplyGames)

	assert.Nil(t, err)
	assert.Equal(t, 5489600, sum)
}
