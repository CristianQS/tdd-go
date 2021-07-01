package pkg

func ValidatePassword(password string) bool {
	if len(password) < 9 { return false}
	return true
}


