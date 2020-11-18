package models

import (
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Claim es la estructura usada para procesar el JWT
type Claim struct {
	CodUser string             `json:"codUser"`
	ID      primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	jwt.StandardClaims
}
