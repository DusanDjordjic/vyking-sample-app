package utils

import "regexp"

var (
	emailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
)

func IsEmailValid(email string) bool {
	if !emailRegex.MatchString(email) {
		return false
	}
	return true
}
