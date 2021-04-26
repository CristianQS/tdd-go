package pkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_should_return_100_elements(t *testing.T) {
	result:= Transform()

	assert.Equal(t, 100, len(result))
}

func Test_should_return_Fizz_when_is_multiple_of_three(t *testing.T) {
	_ = Transform()

	assert.Equal(t, "Fizz", GetNumber(3))
	assert.Equal(t, "Fizz", GetNumber(33))
	assert.Equal(t, "Fizz", GetNumber(99))
}

func Test_should_return_Buzz_when_is_multiple_of_five(t *testing.T) {
	_ = Transform()

	assert.Equal(t, "Buzz", GetNumber(5))
	assert.Equal(t, "Buzz", GetNumber(50))
	assert.Equal(t, "Buzz", GetNumber(95))
}

