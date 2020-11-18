package routers

import (
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jacobd39/crmdemo/bd"
	model "github.com/jacobd39/crmdemo/models"
)

//CodUser es el valor de codUser usado en todos los endpoints
var CodUser string

//IDUsuario es el ID devuelto por el modelo, que se usará en todos los endpoints
var IDUsuario string

//ProcessToken proceso que extrae los valores del Token
func ProcessToken(tk string) (*model.Claim, bool, string, error) {
	myKey := []byte("kbjnfqfsfy79")
	claims := &model.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("invalid token format")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})

	if err == nil {
		_, found, _ := bd.UserVerification(claims.CodUser)
		if found == true {
			CodUser = claims.CodUser
			IDUsuario = claims.ID.Hex()
		}
		return claims, found, IDUsuario, nil
	}

	if !tkn.Valid {
		return claims, false, string(""), errors.New("Invalid Token")
	}

	return claims, false, string(""), err
}
