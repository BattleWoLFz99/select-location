// Connects to MongoDB and sets a Stable API version
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql"
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
	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	// Since we repeatedly use this part, we can revise this part if time allows.
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
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}

	// Define GraphQL schema
	fields := graphql.Fields{
		"states": &graphql.Field{
			Type: graphql.NewList(graphql.NewObject(graphql.ObjectConfig{
				Name: "State",
				Fields: graphql.Fields{
					"name": &graphql.Field{Type: graphql.String},
					"code": &graphql.Field{Type: graphql.String},
				},
			})),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				collection := client.Database("us_states").Collection("states")
				var states []State
				cursor, err := collection.Find(context.TODO(), bson.D{})
				if err != nil {
					return nil, err
				}
				defer cursor.Close(context.TODO())
				if err = cursor.All(context.TODO(), &states); err != nil {
					return nil, err
				}
				return states, nil
			},
		},
	}

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name:   "Query",
			Fields: fields,
		}),
	})
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		result := graphql.Do(graphql.Params{
			Schema:        schema,
			RequestString: r.URL.Query().Get("query"),
		})
		json.NewEncoder(w).Encode(result)
	})

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
