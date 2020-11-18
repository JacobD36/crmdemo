package bds

import (
	"context"
	"time"

	model "github.com/jacobd39/twittor/models"
)

//BorroRelacion borra la relación en la BD
func BorroRelacion(t model.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("gotutorial")
	col := db.Collection("relacion")

	_, err := col.DeleteOne(ctx, t)

	if err != nil {
		return false, err
	}

	return true, nil
}
