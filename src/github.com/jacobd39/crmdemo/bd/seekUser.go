package bd

import (
	"context"
	"fmt"
	"time"

	model "github.com/jacobd39/crmdemo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//SeekUser busca un registro en la base de datos
func SeekUser(ID string) (model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("crmbase")
	col := db.Collection("users")

	var usrData model.User

	objID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id": objID,
	}

	err := col.FindOne(ctx, condition).Decode(&usrData)

	usrData.Password = ""

	if err != nil {
		fmt.Println("Record not found " + err.Error())
		return usrData, err
	}

	return usrData, nil
}
