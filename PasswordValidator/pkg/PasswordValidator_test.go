package pkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_should_password_pass_the_rules(t *testing.T) {
	validatedPassword := "AsqWla_q8"

	var result = ValidatePassword(validatedPassword)

	assert.Equal(t, true,result)
}

func Test_should_password_not_pass_when_his_length_is_lower_than_8(t *testing.T) {
	validatedPassword := "AsqWa_q8"

	var result = ValidatePassword(validatedPassword)

	assert.Equal(t, false,result)
}

func Test_should_password_not_pass_when_there_is_not_capital_letter(t *testing.T) {
	validatedPassword := "asqwla_q8"

	var result = ValidatePassword(validatedPassword)

	assert.Equal(t, false,result)
}

func Test_should_password_not_pass_when_there_is_not_lower_letter(t *testing.T) {
	validatedPassword := "ASQWLA_Q8"

	var result = ValidatePassword(validatedPassword)

	assert.Equal(t, false,result)
}

func Test_should_password_not_pass_when_there_is_not_number(t *testing.T) {
	validatedPassword := "AsqWla_qe"

	var result = ValidatePassword(validatedPassword)

	assert.Equal(t, false,result)
}

func Test_should_password_not_pass_when_there_is_not_underscore(t *testing.T) {
	validatedPassword := "AsqWla2q8"

	var result = ValidatePassword(validatedPassword)

	assert.Equal(t, false,result)
}
