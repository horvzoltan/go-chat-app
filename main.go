package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// MongoDB Atlas connection string
	// WARNING: Keep your connection string confidential.
	uri := "mongodb+srv://zoltanhorvath0721:almarepa2@cluster1.ahj6jat.mongodb.net/?retryWrites=true&w=majority&appName=Cluster1"

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB Atlas!")

	// Get a handle for your collection in a specific database
	// Replace "yourDatabase" and "users" with your actual database and collection names
	collection := client.Database("chat-app").Collection("users")

	// Insert a document
	// Replace these fields with the ones you'll use, and make sure to hash the password
	document := bson.D{{"username", "exampleUser"}, {"email", "user@example.com"}, {"password", "hashedPassword"}}
	insertResult, err := collection.InsertOne(context.TODO(), document)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	// Always close the connection when done to avoid memory leaks
	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connection to MongoDB Atlas closed.")
}
