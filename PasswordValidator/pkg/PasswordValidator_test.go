package pkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_should_password_pass_the_rules(t *testing.T) {
	validatedPassword := "AsqWlaO8"

	var result = ValidatePassword(validatedPassword)

	assert.Equal(t, true,result)
}
