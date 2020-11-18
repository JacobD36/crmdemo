package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//User --> Define la estructura de la tabla usuarios
type User struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	CodUser    string             `bson:"codUser" json:"codUser"`
	IDProfile  int                `bson:"idProfile" json:"idProfile"`
	Status     int                `bson:"status" json:"status"`
	Campaign   int                `bson:"campaign" json:"campaign"`
	Name1      string             `bson:"name1" json:"name1,omitempty"`
	Name2      string             `bson:"name2" json:"name2,omitempty"`
	LastName1  string             `bson:"lastName1" json:"lastName1,omitempty"`
	LastName2  string             `bson:"lastName2" json:"lastName2,omitempty"`
	Password   string             `bson:"password" json:"password,omitempty"`
	BirthDate  time.Time          `bson:"birthDate" json:"birthDate,omitempty"`
	CreateAt   time.Time          `bson:"createAt" json:"createAt,omitempty"`
	ModifiedAt time.Time          `bson:"modifiedAt" json:"modifiedAt,omitempty"`
}
