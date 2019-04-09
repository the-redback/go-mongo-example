package main

import (
	"context"
	"fmt"
	"log"

	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// You will be using this Trainer type later in the program
type Trainer struct {
	Name string
	Age  int
	City string
}

func main() {
	// Set client options
	fmt.Println("Client Options...")
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	fmt.Println("Connect mongodb...")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Mongo Ping...")
	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	//// Check the connection
	//fmt.Println("List database names")
	//listDB, err := client.ListDatabaseNames(context.TODO(), nil, )
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println("DBs:", listDB)

}