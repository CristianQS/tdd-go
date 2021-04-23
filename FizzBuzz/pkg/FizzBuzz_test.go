package pkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"should_return_100_elements"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result:= FizzBuzz()

			assert.Equal(t, 100, len(result))
		})
	}
}

