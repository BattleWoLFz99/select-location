package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const uri = "mongodb://localhost:27017"

type State struct {
	Name string `json:"name" bson:"name"`
	Code string `json:"code" bson:"code"`
}

func main() {
	// Read the states-data.json file
	data, err := os.ReadFile("states-data.json")
	if err != nil {
		panic(err)
	}

	// Parse JSON into states
	var states []State
	err = json.Unmarshal(data, &states)
	if err != nil {
		panic(err)
	}

	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// Send a ping to confirm a successful connection
	var result bson.M
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	collection := client.Database("us_states").Collection("states")

	// Clear existing data
	err = collection.Drop(context.TODO())
	if err != nil {
		panic(err)
	}

	// Insert the data
	for _, state := range states {
		_, err := collection.InsertOne(context.TODO(), state)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("States loaded into MongoDB successfully!")
}
