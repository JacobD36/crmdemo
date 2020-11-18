package bd

import (
	"context"
	"fmt"
	"time"

	model "github.com/jacobd39/crmdemo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//UpdateRecord actualiza en la base de datos el registro de un usuario con id = ID
func UpdateRecord(u model.User, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("crmbase")
	col := db.Collection("users")

	record := make(map[string]interface{})

	record["name1"] = u.Name1
	record["name2"] = u.Name2

	/*if len(u.Name2) > 0 {
		record["name2"] = u.Name2
	}*/

	record["lastName1"] = u.LastName1
	record["lastName2"] = u.LastName2
	record["idProfile"] = u.IDProfile
	record["campaign"] = u.Campaign
	record["status"] = u.Status

	u.Password, _ = PasswordEncript(u.Password)

	if len(u.Password) > 0 {
		record["password"] = u.Password
	}

	record["birthDate"] = u.BirthDate
	record["modifiedAt"] = time.Now()

	updtString := bson.M{
		"$set": record,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)

	filter := bson.M{
		"_id": bson.M{
			"$eq": objID,
		},
	}

	_, err := col.UpdateOne(ctx, filter, updtString)

	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}

	return true, nil
}
