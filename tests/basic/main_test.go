package basic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func AddOne(x int) int {
	return x + 1
}

func TestA(t *testing.T) {
	assert.Equal(t, AddOne(1), 2, "AddOne(1) should be 2")
}
