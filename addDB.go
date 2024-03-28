package main

import (
"context"
"log"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
	"fmt"
)

type Animal struct {
Name     string `bson:"name"`
Symptoms string `bson:"symptoms"`
id int `bson:"_id"`
}

func addDB(Name string, Claim string, id int) {
// Set client options

clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")


// Connect to MongoDB
client, err := mongo.Connect(context.Background(), clientOptions)
if err != nil {
	log.Fatal(err)
}

// Check the connection
err = client.Ping(context.Background(), nil)
if err != nil {
	log.Fatal(err)
}

// Connect to the "animals" collection
collection := client.Database("OSamidb").Collection("animal")

// Insert documents into the collection



doc := Animal{Name: Name, Symptoms: Claim, id:id,}



result, err := collection.InsertOne(context.TODO(), doc)


fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)

}
