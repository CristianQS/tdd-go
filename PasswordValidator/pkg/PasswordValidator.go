package pkg

import "strings"

func ValidatePassword(password string) bool {
	if len(password) < 9 { return false}
	if password == strings.ToLower(password) { return false }
	return true
}


