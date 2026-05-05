package basic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestA(t *testing.T) {
	actual := 2
	expected := 2
	assert.Equal(t, expected, actual, "Expected %d, but got %d", expected, actual)
}
