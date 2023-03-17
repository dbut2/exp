package math

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargestN(t *testing.T) {
	s := randomInt(1e6)
	n := 100

	s1 := make([]int, len(s))
	copy(s1, s)
	s2 := make([]int, len(s))
	copy(s2, s)

	expected := Order(s1, true)[:n]
	actual := LargestN(s2, n)
	assert.Equal(t, expected, actual)
}

func TestSmallestN(t *testing.T) {
	s := randomInt(1e6)
	n := 100

	s1 := make([]int, len(s))
	copy(s1, s)
	s2 := make([]int, len(s))
	copy(s2, s)

	expected := Order(s1, false)[:n]
	actual := SmallestN(s2, n)
	assert.Equal(t, expected, actual)
}

func randomInt(n int) []int {
	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = rand.Int()
	}
	return s
}
