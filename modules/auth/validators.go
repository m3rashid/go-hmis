package auth

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func LoginValidator(body *UserModel) error {
	if body.Email == "" || body.Password == "" || body.Name == "" {
		return fmt.Errorf("name, email and password are required")
	}

	if len(body.Password) < 8 {
		return fmt.Errorf("password must be at least 8 characters long")
	} else {
		if password, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10); err != nil {
			return err
		} else {
			body.Password = string(password)
		}
	}

	body.EmailVerified = false
	return nil
}
