package pkg

import (
	"strings"
	"unicode"
)

func ValidatePassword(password string) bool {
	if HasMoreThan8Characters(password) { return false }
	if NotContainsUpperCase(password) { return false }
	if NotContainsLowerCase(password) {	return false }
	if !HasDigit(password) { return false }
	return true
}


func HasMoreThan8Characters(password string) bool {
	return len(password) <= 8
}

func NotContainsUpperCase(password string) bool {
	return password == strings.ToLower(password)
}

func NotContainsLowerCase(password string) bool {
	return password == strings.ToUpper(password)
}

func HasDigit(password string) bool {
	for _, character := range password {
		if unicode.IsDigit(character) {
			return true
		}
	}
	return false
}




