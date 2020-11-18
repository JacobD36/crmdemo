package routers

import (
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	db "github.com/jacobd39/twittor/bds"
	model "github.com/jacobd39/twittor/models"
)

//Email valor de Email usado en todos los endpoints
var Email string

//IDUsuario es el ID devuelto del modelo, que se usará en todos los endpoints
var IDUsuario string

//ProcesoToken proceso que extrae los valores del Token
func ProcesoToken(tk string) (*model.Claim, bool, string, error) {
	miClave := []byte("MastersdelDesarrollo_grupodeFacebook")
	claims := &model.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token inválido")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})

	if err == nil {
		_, encontrado, _ := db.ChequeoYaExisteUsuario(claims.Email)
		if encontrado == true {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, encontrado, IDUsuario, nil
	}

	if !tkn.Valid {
		return claims, false, string(""), errors.New("token inválido")
	}

	return claims, false, string(""), err
}
