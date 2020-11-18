package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//MongoCN = variable exportable de conexion
var MongoCN = connectDatabase()

var clientOptions = options.Client().ApplyURI("mongodb://crmUser:kbjnfqfsfy79@127.0.0.1:27017/?authSource=admin")

func connectDatabase() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Successful DB Connection")
	return client
}

//CheckDBConnection verifica conexión a la base de datos
func CheckDBConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
