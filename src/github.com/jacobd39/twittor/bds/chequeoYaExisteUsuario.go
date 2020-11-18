package bds

import (
	"context"
	"time"

	"github.com/jacobd39/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
)

//ChequeoYaExisteUsuario recibe un email de parámetro y verifica si ya está en la BD
func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("gotutorial")
	col := db.Collection("usuarios")

	condicion := bson.M{"email": email}

	var resultado models.Usuario

	err := col.FindOne(ctx, condicion).Decode(&resultado)

	ID := resultado.ID.Hex()

	if err != nil {
		return resultado, false, ID
	}

	return resultado, true, ID
}
