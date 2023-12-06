package lib

import (
	"os"
	"strings"
)

func SumSlice(digits []int) int {
	sum := 0
	for _, d := range digits {
		sum += d
	}
	return sum
}

func ReadFile(name string) ([]string, error) {
	file, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(file), "\n")
	return lines, nil
}

func ReverseString(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
