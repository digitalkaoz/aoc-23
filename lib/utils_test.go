package lib

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadFile(t *testing.T) {
	file, err := ReadFile("../data/1.txt")

	assert.Nil(t, err)
	assert.Greater(t, len(file), 0)
}

func TestSumSlice(t *testing.T) {
	inp := []int{1, 2, 3, 4, 5}

	sum := SumSlice(inp)
	assert.Equal(t, 15, sum)
}

func TestReverseString(t *testing.T) {
	inp := "abc"

	assert.Equal(t, "cba", ReverseString(inp))
}
