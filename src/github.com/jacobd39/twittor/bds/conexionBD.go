package bds

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//MongoCN = variable exportable de conexion
var MongoCN = conectarBD()

var clientOptions = options.Client().ApplyURI("mongodb://superAdmin:pass1234@192.168.1.13:27017/?authSource=admin")

func conectarBD() *mongo.Client {
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

	log.Println("Conexión Exitosa con la BD")
	return client
}

//ChequeoConnection es el Ping a la BD
func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
