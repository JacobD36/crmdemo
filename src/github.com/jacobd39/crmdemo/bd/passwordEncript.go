package bd

import "golang.org/x/crypto/bcrypt"

//PasswordEncript encripta el password que se le pasa como parámetro
func PasswordEncript(pass string) (string, error) {
	cost := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), cost)
	return string(bytes), err
}
