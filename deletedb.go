package main

import (
"context"
"fmt"
"log"
"time"
"go.mongodb.org/mongo-driver/mongo"
"go.mongodb.org/mongo-driver/mongo/options"
"go.mongodb.org/mongo-driver/bson"
)

type anam struct{
	Name string `bson:"name,omitempty"`
	Claim string `bson:"namemomitempty"`
    
}


func del() {
client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
if err != nil {
log.Fatal(err)
}


ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()

err = client.Connect(ctx)
if err != nil {
	log.Fatal(err)
}

collection := client.Database("OSamidb").Collection("animal")

filter := bson.M{"name": ""}
res, err := collection.DeleteOne(ctx, filter)
if err != nil {
	log.Fatal(err)
}

fmt.Println("Deleted document:", res.DeletedCount)
}