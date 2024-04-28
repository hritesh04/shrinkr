package helper

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func EncodePassword(password string) string {
	hashedPassword,err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return string(hashedPassword)
}