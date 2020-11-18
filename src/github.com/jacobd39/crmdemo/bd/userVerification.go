package bd

import (
	"context"
	"time"

	model "github.com/jacobd39/crmdemo/models"
	"go.mongodb.org/mongo-driver/bson"
)

//UserVerification recibe un coduser de parámetro y verifica si ya existe en la BD
func UserVerification(coduser string) (model.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("crmbase")
	col := db.Collection("users")

	condition := bson.M{
		"codUser": coduser,
	}

	var result model.User

	err := col.FindOne(ctx, condition).Decode(&result)

	ID := result.ID.Hex()

	if err != nil {
		return result, false, ID
	}

	return result, true, ID
}
