package bds

import "golang.org/x/crypto/bcrypt"

//EncriptarPassword encripta el password que se le pasa como parámetro
func EncriptarPassword(pass string) (string, error) {
	costo := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	return string(bytes), err
}
