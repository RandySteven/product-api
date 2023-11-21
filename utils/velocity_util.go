package utils

import "golang.org/x/crypto/bcrypt"

func CheckOldAndNew(old, new string) string {
	if new == "" {
		return old
	}
	return new
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
