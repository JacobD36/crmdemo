package bd

import (
	model "github.com/jacobd39/crmdemo/models"
	"golang.org/x/crypto/bcrypt"
)

//LoginIntent realiza el intento de Login
func LoginIntent(codUser string, password string) (model.User, bool, int) {
	user, found, _ := UserVerification(codUser)

	if found == false {
		return user, false, 1
	}

	passwordBytes := []byte(password)
	passwordBD := []byte(user.Password)

	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)

	if err != nil {
		return user, false, 2
	}

	return user, true, 0
}
