package bd

import (
	"context"
	"fmt"
	"time"

	model "github.com/jacobd39/crmdemo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//GetUsers devuelve el listado de usuarios que cumplan con los criterios de búsqueda y paginación
func GetUsers(page int64, search string) ([]*model.User, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("crmbase")
	col := db.Collection("users")

	var results []*model.User

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 10)
	findOptions.SetLimit(10)

	query := bson.M{
		"name1": bson.M{"$regex": `(?i)` + search},
	}

	cur, err := col.Find(ctx, query, findOptions)

	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}

	for cur.Next(ctx) {
		var s model.User

		err := cur.Decode(&s)

		if err != nil {
			fmt.Println(err.Error())
			return results, false
		}

		results = append(results, &s)
	}

	err = cur.Err()

	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}

	cur.Close(ctx)
	return results, true
}
