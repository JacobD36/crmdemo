package bd

import (
	"context"
	"time"

	model "github.com/jacobd39/crmdemo/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//InsertNewUser inserta un nuevo usuario en la base de datos
func InsertNewUser(u model.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("crmbase")
	col := db.Collection("users")

	u.Password, _ = PasswordEncript(u.Password)

	result, err := col.InsertOne(ctx, u)

	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)

	return ObjID.String(), true, nil
}
