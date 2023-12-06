package two

import (
	"aoc-23/lib"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSampleInput(t *testing.T) {
	lines := []string{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
	}

	sums, err := sumLines(lines, func(game *Game) int {
		return game.Number()
	})
	assert.Nil(t, err)
	assert.Len(t, sums, 5)
	sum := lib.SumSlice(sums)
	assert.Equal(t, 8, sum)
}

func TestSampleInputPower(t *testing.T) {
	lines := []string{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
	}

	sums, err := sumLines(lines, func(game *Game) int {
		return game.MinimumSet().Power()
	})
	assert.Nil(t, err)
	assert.Len(t, sums, 5)
	sum := lib.SumSlice(sums)
	assert.Equal(t, 2286, sum)
}

func TestFile(t *testing.T) {
	sum, err := processFile("../../data/2.txt", func(game *Game) int {
		return game.Number()
	})

	assert.Nil(t, err)
	assert.Equal(t, 2204, sum)
}

func TestFilePower(t *testing.T) {
	sum, err := processFile("../../data/2.txt", func(game *Game) int {
		return game.MinimumSet().Power()
	})

	assert.Nil(t, err)
	assert.Equal(t, 71036, sum)
}

func TestParseGame(t *testing.T) {
	inp := "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"

	game := parseGame(inp)

	assert.NotNil(t, game)
	assert.Equal(t, game.Name, "Game 1")
	assert.Len(t, game.Sets, 3)
	assert.Equal(t, 4, game.Sets[0].Red)
	assert.Equal(t, 3, game.Sets[0].Blue)
}

func TestMinimumSet(t *testing.T) {
	inp := "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"

	game := parseGame(inp)
	assert.NotNil(t, game)

	minSet := game.MinimumSet()
	assert.NotNil(t, minSet)

	assert.Equal(t, 4, minSet.Red)
	assert.Equal(t, 6, minSet.Blue)
	assert.Equal(t, 2, minSet.Green)

	assert.Equal(t, 48, minSet.Power())
}

func TestValidateGameSet(t *testing.T) {
	assert.True(t, GameSet{Red: 12}.IsValid())
	assert.True(t, GameSet{Green: 13}.IsValid())
	assert.True(t, GameSet{Blue: 14}.IsValid())

	assert.False(t, GameSet{Red: 13}.IsValid())
	assert.False(t, GameSet{Green: 14}.IsValid())
	assert.False(t, GameSet{Blue: 15}.IsValid())
}
