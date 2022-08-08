package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var MongoCN = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://bblanqui:brian312@cluster0.hac8ckg.mongodb.net/?retryWrites=true&w=majority")

func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return client
	}
	log.Println("conexion exitosa")

	return client
}

func ChequeoConnection() bool {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return false
	}
	return true
}

func DbColeccion(db string, coleccion string) (contexto context.Context, cancelar context.CancelFunc, client *mongo.Collection) {
	contextos, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	mongoDb := MongoCN.Database(db).Collection(coleccion)
	return contextos, cancel, mongoDb
}
