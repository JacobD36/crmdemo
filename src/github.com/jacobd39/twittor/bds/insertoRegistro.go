package bds

import (
	"context"
	"time"

	model "github.com/jacobd39/twittor/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InsertoRegistro es la parada final con la BD para insertar los datos del usuario
func InsertoRegistro(u model.Usuario) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("gotutorial")
	col := db.Collection("usuarios")

	u.Password, _ = EncriptarPassword(u.Password)

	result, err := col.InsertOne(ctx, u)

	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)

	return ObjID.String(), true, nil
}
