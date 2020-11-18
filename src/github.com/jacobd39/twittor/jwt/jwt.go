package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	model "github.com/jacobd39/twittor/models"
)

//GeneroJWT nos proporciona un metodo para generar un JSON Web Token
func GeneroJWT(t model.Usuario) (string, error) {
	miClave := []byte("MastersdelDesarrollo_grupodeFacebook")

	payload := jwt.MapClaims{
		"email":            t.Email,
		"nombre":           t.Nombre,
		"apellidos":        t.Apellidos,
		"fecha_nacimiento": t.FechaNacimiento,
		"biografia":        t.Biografia,
		"ubicacion":        t.Ubicacion,
		"sitioweb":         t.SitioWeb,
		"_id":              t.ID.Hex(),
		"exp":              time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(miClave)

	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}