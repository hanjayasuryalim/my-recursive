package my_recursive

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFactorial(t *testing.T) {
	expected := 120
	result := FactorialTailRecursive(5, 1)
	assert.Equal(t, expected, result, "Expected result to be 120")
}
