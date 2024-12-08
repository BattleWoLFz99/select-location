package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"select-location/internal/config"
	"select-location/internal/db"
	"select-location/internal/models"
)

func main() {
	data, err := os.ReadFile("states-data.json")
	if err != nil {
		panic(err)
	}

	var states []models.State
	err = json.Unmarshal(data, &states)
	if err != nil {
		panic(err)
	}

	client, err := db.ConnectDB(config.MongoURI)
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(context.TODO())

	collection := client.Database("us_states").Collection("states")

	err = collection.Drop(context.TODO())
	if err != nil {
		panic(err)
	}

	for _, state := range states {
		_, err := collection.InsertOne(context.TODO(), state)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("States loaded into MongoDB successfully!")
}
