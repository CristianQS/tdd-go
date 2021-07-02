package pkg

import "strings"

func ValidatePassword(password string) bool {
	if HasMoreThan8Characters(password) { return false }
	if password == strings.ToLower(password) { return false }
	if password == strings.ToUpper(password) { return false }
	return true
}

func HasMoreThan8Characters(password string) bool {
	return len(password) <= 8
}


