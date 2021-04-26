package pkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_should_return_100_elements(t *testing.T) {
	result:= FizzBuzz()

	assert.Equal(t, 100, len(result))
}

