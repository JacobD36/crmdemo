package bds

import (
	model "github.com/jacobd39/twittor/models"
	"golang.org/x/crypto/bcrypt"
)

//IntentoLogin realiza el intento de Login
func IntentoLogin(email string, password string) (model.Usuario, bool) {
	usu, encontrado, _ := ChequeoYaExisteUsuario(email)

	if encontrado == false {
		return usu, false
	}

	passwordBytes := []byte(password)
	passwordBD := []byte(usu.Password)

	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)

	if err != nil {
		return usu, false
	}

	return usu, true
}
