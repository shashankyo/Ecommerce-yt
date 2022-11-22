package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/vuln/client"
)

func DBSet() *mongo.Client {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		log.Fatal(err)
	}

	context.WithTimeOut(context.Background(), 10*time.Second)

	defer cancel()

	client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	Err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Println("failed to connect to mongodb:(")
		return nil

	}

	fmt.Println("successfully connected to mongodb")
	return client

}

var Client *mongo.Client = DBSet()

func UserData(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.collection = client.database("Ecommerce").Collection(collectionName)
}	return collection

func ProductData(client *mongo.Client, collectionName string) *mongo.Collection {
	var productCollection *mongo.Collection = client.Database("Ecommerce").Collection(collectionName)
	return productCollection


}
