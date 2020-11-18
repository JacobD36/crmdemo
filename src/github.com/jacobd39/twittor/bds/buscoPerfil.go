package bds

import (
	"context"
	"fmt"
	"time"

	model "github.com/jacobd39/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//BuscoPerfil busca un perfil en la BD
func BuscoPerfil(ID string) (model.Usuario, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("gotutorial")
	col := db.Collection("usuarios")

	var perfil model.Usuario

	objID, _ := primitive.ObjectIDFromHex(ID)

	condicion := bson.M{
		"_id": objID,
	}

	err := col.FindOne(ctx, condicion).Decode(&perfil)

	perfil.Password = ""

	if err != nil {
		fmt.Println("Registro no encontrado " + err.Error())
		return perfil, err
	}

	return perfil, nil
}
