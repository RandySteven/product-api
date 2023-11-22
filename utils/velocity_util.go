package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

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

func Validate(obj interface{}) []string {
	validate := validator.New()
	err := validate.Struct(obj)
	if err != nil {
		var errs = make([]string, 0)
		for _, currErr := range err.(validator.ValidationErrors) {
			errMsg := fmt.Sprintf("%s field is %s", currErr.Field(), currErr.ActualTag())
			errs = append(errs, errMsg)
		}
		return errs
	}
	return nil
}
