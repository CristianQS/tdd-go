package pkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
const FIZZ = "Fizz"
const BUZZ = "Buzz"
const FIZZBUZZ = "FizzBuzz"


func Test_should_return_100_elements(t *testing.T) {
	result:= Transform()

	assert.Equal(t, 100, len(result))
}

func Test_should_return_number_when_do_not_fullfill_business_rules(t *testing.T) {
	_= Transform()

	assert.Equal(t, "1", GetNumber(1))
	assert.Equal(t, "47", GetNumber(47))
	assert.Equal(t, "97", GetNumber(97))
}

func Test_should_return_Fizz_when_is_multiple_of_three(t *testing.T) {
	_ = Transform()

	assert.Equal(t, FIZZ, GetNumber(3))
	assert.Equal(t, FIZZ, GetNumber(33))
	assert.Equal(t, FIZZ, GetNumber(99))
}

func Test_should_return_Buzz_when_is_multiple_of_five(t *testing.T) {
	_ = Transform()

	assert.Equal(t, BUZZ, GetNumber(5))
	assert.Equal(t, BUZZ, GetNumber(50))
	assert.Equal(t, BUZZ, GetNumber(95))
}

func Test_should_return_Buzz_when_is_multiple_of_five_and_three(t *testing.T) {
	_ = Transform()

	assert.Equal(t, FIZZBUZZ, GetNumber(15))
	assert.Equal(t, FIZZBUZZ, GetNumber(45))
	assert.Equal(t, FIZZBUZZ, GetNumber(90))
}

func Test_should_return_Fizz_when_contains_three(t *testing.T) {
	_ = Transform()

	assert.Equal(t, FIZZ, GetNumber(31))
	assert.Equal(t, FIZZ, GetNumber(13))
}

func Test_should_return_Buzz_when_contains_five(t *testing.T) {
	_ = Transform()

	assert.Equal(t, BUZZ, GetNumber(52))
	assert.Equal(t, BUZZ, GetNumber(59))
}
func Test_should_return_FizzBuzz_when_contains_three_and_five(t *testing.T) {
	_ = Transform()

	assert.Equal(t, FIZZBUZZ, GetNumber(35))
	assert.Equal(t, FIZZBUZZ, GetNumber(53))
}

